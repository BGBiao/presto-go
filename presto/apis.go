package presto

import (
	"time"
)

const (
	version             = "0.0.1"
	userHeader          = "X-Presto-User"
	sourceHeader        = "X-Presto-Source"
	catalogHeader       = "X-Presto-Catalog"
	schemaHeader        = "X-Presto-Schema"
	userAgent           = "go-presto/" + version
	prestoSessionHeader = "X-Presto-Session"

	initialRetry            = 50 * time.Millisecond
	maxRetry                = 800 * time.Millisecond
	ProgressUnknown float64 = -1.0
)

// define a presto connection and cache struct
type Conn struct {
	host    string // presto主机地址
	port    int    // presto主机端口(一般默认8080)
	user    string // presto主机用户可默认指定
	source  string // 请求来源
	catalog string // presto数据源(hive,kafka...)
	schema  string // 对应数据源的对象,比如hive的schema就是指定表

	query string // 指定相关查询语句

	id      string // 查询id
	nextUri string // 查询任务的下一个uri

	columns []string // 列名

	closed       bool            // 查询是否关闭
	bufferedRows [][]interface{} // 缓存的查询结果
	state        string          // 当前查询状态
	nodes        int             // 查询总共分布在多少节点
	progress     float64         // 查询进度(分割任务和已完成任务占比)

}

// define a result data struct
type PrestoDatas struct {
	Id               string `json:"id"`
	InfoUri          string `json:"infourl"`
	PartialCancelUri string `json:"partialCancelUri"`
	NextUri          string `json:"nextUri"`
	// 只取列名(name:appname,type:varchar,typeSignature:struct)
	Columns []struct {
		Name string `json:"name"`
	} `json:"columns"`

	//Data数据为每一条查询内部获取的数据，最终需要将全部的数据汇聚到一个列表中
	Data [][]interface{} `json:"data"`

	// 只获取关心的几个列(state,scheduled,)
	// 注意:此处应该每个任务的状态数据都是相同的
	Stats struct {
		State           string // 当前该任务的状态[QUEUED|RUNNING|FINISHED|FAILED]
		Nodes           int    // 当前任务节点
		Scheduled       bool   // 当前该任务是否需要调度[false|true|true|false] 任务类的肯定属于需要调度
		TotalSplits     int    // 当前整个查询总共分割成几个任务执行
		CompletedSplits int    // 当前已完成的任务数(当totalSplits=completedSplits肯定就全部成功)
	} `json:"stats"`

	// 获取错误信息
	Error struct {
		Message   string `json:"message"`   // 错误原因(需要返回给用户的数据)
		ErrorCode int    `json:"errorCode"` // 错误编码
		ErrorName string `json:"errorName"` // 错误名称(SYNTAX_ERROR)
		ErrorType string `json:"errorType"` // 错误类型(USER_ERROR)
		// 错误堆栈信息
		FailureInfo struct {
			Type    string   `json:"type"`    //错误类型[com.facebook.presto.sql.analyzer.SemanticException]
			Message string   `json:"message"` //错误详细信息(同上级的Message)
			Stack   []string `json:"stack"`   //错误的栈信息
		} `json:"failureInfo"`
	} `json:"error"`
}

// define a flow and task data struct
// for http:/presto-host/query.html?queryid

type FlowTaskDat struct {
	Task []struct {
		// 仅取几个相关信息
		TaskId string // 任务id
		State  string // 任务状态
		Host   string // 任务所在主机
		Uptime int64  // 运行时间
	} `json:"task"`
	Flows []struct {
		From     string // flow是从哪个任务
		To       string // flow是到那个任务
		Finished bool   // flow是否完成
	} `json:"flows"`
}
