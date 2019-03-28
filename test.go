package main
import (
    "presto-go/presto"
    "fmt"
)

func main() {
  //host,port,user,catalog,schema,query
  sql := "select * from default.testapp_vip_ip "
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
