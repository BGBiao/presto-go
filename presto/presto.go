package presto
import (
        "encoding/json"
        "fmt"
        requests "github.com/xxbandy/go-utils/requests"
        _ "io/ioutil"
        _ "net/http"
        "strings"
        "time"
)
// make a new presto connection and cache object
func NewQuery(host string, port int, user, source, catalog, schema, query string) (*Conn, error) {
  if user == "" {
    user = "anonymous"
  }

  if source == "" {
    source = userAgent
  }

  if catalog == "" {
    catalog = "hive"
  }

  if schema == "" {
    schema = "tds"
  }

  // 初始化一个Conn
  c := &Conn{
    host:    host,
    port:    port,
    user:    user,
    source:  source,
    catalog: catalog,
    schema:  schema,
    query:   query,
  }

  postErr := c.PostQuery()

  if postErr != nil {
    return nil, postErr
  }

  return c, nil

}

// make some methods
func (c *Conn) Geturl() string {
  return fmt.Sprintf("http://%s:%d", c.host, c.port)
}

func (c *Conn) GetQueryId() string {
  return fmt.Sprintf("query id:%s", c.id)
}

func (c *Conn) GetNextUri() string {
  return fmt.Sprintf("next uri:%s", c.nextUri)
}

func (c *Conn) GetColumns() []string {
  return c.columns
}

func (c *Conn) GetDataRows() int {
  return len(c.bufferedRows)
}

func (c *Conn) Getclose() bool {
  return c.closed
}

func (c *Conn) GetState() string {
  return c.state
}

// get nodes for task exec.
func (c *Conn) GetNodes() int {
  return c.nodes
}

// get the task process 
func (c *Conn) GetProcess() float64 {
  return c.progress
}

// submit the first query
func (c *Conn) PostQuery() error {
  baseurl := c.Geturl()
  uri := fmt.Sprintf("/v1/statement")
  newapi := requests.NewApi(baseurl)

  // 构造一个POST请求
  newreq, reqErr := newapi.GetRequest("POST", uri, strings.NewReader(c.query))
  if reqErr != nil {
    fmt.Printf("sorry,failed to get a request with:%s", reqErr)
  }
  // 添加相关的header
  newreq.Header.Add(catalogHeader, c.catalog)
  newreq.Header.Add(schemaHeader, c.schema)
  newreq.Header.Add(sourceHeader, c.source)
  newreq.Header.Add("User-Agent", userAgent)
  newreq.Header.Add(userHeader, c.user)

  // 执行上面构造的newreq
  respbody, _ := requests.NewClient(newreq, 50*time.Second)

  result := PrestoDatas{}
  // 获取完结果数据后直接赋值就丢弃该数据
  // 解析查询任务的响应，获取查询id和nexturi
  resErr := json.Unmarshal(respbody, &result)
  if resErr != nil {
    fmt.Printf("unmarshal %s to PrestoDatas error with %s.", string(respbody), resErr)
    return fmt.Errorf("error:%s", resErr)
  }

  if result.Error.Message != "" {
    return fmt.Errorf("parser the result data error:%s", result.Error.Message)
  }

  c.id = result.Id
  c.nextUri = result.NextUri
  return nil
}

// query next task 
func (c *Conn) queryNext() error {
  var errdata error
  nexturl := requests.NewApi(c.nextUri)
  data, _ := nexturl.Get("")
  result := PrestoDatas{}
  resErr := json.Unmarshal(data, &result)
  if resErr != nil {
    c.closed = true
    errdata = fmt.Errorf("failed to unmarshal the task %s with error :%s\n", c.nextUri, resErr)
  }

  if result.Error.Message != "" {
    c.closed = true
    errdata = fmt.Errorf("failed to exec the sub task with error :%s\n", result.Error.Message)
  } else if result.Stats.State == "FAILED" {
    c.closed = true
    errdata = fmt.Errorf("the sub task %s is failed,please contact the administors", c.nextUri)
  } else {
    c.bufferedRows = result.Data
    c.state = result.Stats.State
    c.nodes = result.Stats.Nodes
    // 构造相关列信息
    if c.columns == nil {
      c.columns = make([]string, len(result.Columns))
      for i, col := range result.Columns {
        c.columns[i] = col.Name
      }
    }

    // 判断调度调度次数2/2 通常为1.0的时候即完全成功调度
    if result.Stats.Scheduled {
      c.progress = float64(result.Stats.CompletedSplits) / float64(result.Stats.TotalSplits)
    } else {
      c.progress = ProgressUnknown
    }

    c.nextUri = result.NextUri
    // 当最后一个任务没有nexturi的时候，说明已经是最后一个结尾任务了
    if result.NextUri == "" {
      c.closed = true
    }
    errdata = nil
  }
  return errdata

}

// fetch all data
func (c *Conn) Next() ([][]interface{}, error) {
  var rowdatas [][]interface{}
  var errordata error

  retry := initialRetry
  // 判断链接没有关闭，并且没有获取相关数据
  for !c.closed && c.nextUri != "" {
    err := c.queryNext()
    if err != nil {
      errordata = err
    }

    if len(c.bufferedRows) == 0 {
      time.Sleep(retry)
      retry *= 2
      if retry > maxRetry {
        retry = maxRetry
      }
    } else if len(c.bufferedRows) > 0 {
      for _, v := range c.bufferedRows {
        rowdatas = append(rowdatas, v)
      }
    }
  }

  /*
     if len(rowdatas) > 0 {
         errordata = nil
     }else {
         errordata = fmt.Errorf("Sorry,No record matched!")
     }
  */
  return rowdatas, errordata
}

// get flow and task infos

func (c *Conn) GetTasks() {
  taskurl := fmt.Sprintf("%s/v1/query-execution/%s", c.Geturl(), c.id)
  newapi := requests.NewApi(taskurl)
  data, _ := newapi.Get("")
  resdata := FlowTaskDat{}
  json.Unmarshal(data, &resdata)
  fmt.Println(string(data))
  fmt.Println(resdata.Flows)
}
