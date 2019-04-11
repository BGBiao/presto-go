# presto-go

`注意:presto官方也提供了一个golang的presto客户端包，但是要求presto的版本大于0.16x`

一个简单版本的Golang语言的[presto](https://prestodb.github.io/)的客户端。



## Todos

- [X] HTTP Basic 认证
- [X] Presto基本查询
- [X] command line
- - [X] 支持catalog参数
- - [X] 支持--file指定sql文件执行
- [ ] Tasks任务状态[使用独立请求实时查看任务状态]
- [ ] 格式化输出
- [ ] Stages任务流

## Requirements

- go1.8.3+
- Presto 0.14X or newer 


## Installation and Usage

```
$ go get -v github.com/xxbandy/presto-go/presto

$ cat querydata.go
package main
import (
    "fmt"
    "github.com/xxbandy/presto-go/presto"
)

func main() {
  //host,port,user,catalog,schema,query
  sql := "select * from default.userinfo limit 10"
  req, posterr := presto.NewQuery("localhost", 8080, "root", "", "hive", "default", sql)
  if posterr != nil {
    fmt.Println(posterr)
  } else {
    fmt.Println(req.GetQueryId())
    //fmt.Println(req.GetNextUri())

    // rows为获取到的全量数据
    rows, fetcherr := req.Next()
    if fetcherr != nil {
      fmt.Println(fetcherr)
    } else {
      fmt.Println("进度:", req.GetProcess())
      fmt.Println("节点:", req.GetNodes())
      fmt.Println("是否关闭:", req.Getclose())
      fmt.Println("状态:", req.GetState())

      fmt.Println("列名:", req.GetColumns())
      fmt.Println("数据:")
      if rows != nil {
        for _, v := range rows {
          fmt.Println(v)
        }
      }
    }
    //fmt.Println("任务流状态:")
    //req.GetTasks()
  }

}

$ go run querydata.go
query id:20190328_102808_00012_t6972
进度: 1
节点: 1
是否关闭: true
状态: FINISHED
列名: [name age sex]
数据:
[bgops 18 male]
```
