package main

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/xxbandy/presto-go/presto"
	"log"
	"os"
)

func fetchall(host, sql string, port int) {
	req, posterr := presto.NewQuery(host, port, "root", "", "hive", "default", sql)
	if posterr != nil {
		fmt.Println(posterr)
	} else {
		fmt.Println(req.GetQueryId())

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

/*
func getflow(host,queryid string,port int) {
  req := &presto.Conn{
    host:    host,
    port:    port,
    user:    user,
    source:  source,
    catalog: catalog,
    schema:  schema,
      }
}
*/

func main() {
	app := cli.NewApp()

	app.Name = "prestogo"
	app.Usage = "a tiny presto client for golang."
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:      "query",
			ShortName: "q",
			//Aliases: []string{"q"},     // like shortname1,shortname2
			Usage:       "submit a query to presto.",
			Description: "fetch the data from presto with sql.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "sql",
					//Value:    "select * from default.test",
					Usage:    "a query langurage with presto synax.",
					FilePath: "/export/server/presto.sql",
				},
				cli.StringFlag{
					Name:   "server, s",
					Value:  "localhost",
					Usage:  "the presto host.",
					EnvVar: "PRESTO_HOST,PRESTO",
				},
				cli.IntFlag{
					Name:   "port,p",
					Value:  8080,
					Usage:  "the presto instances port.",
					EnvVar: "PRESTO_PORT",
				},
			},
			Action: func(c *cli.Context) error {
        var errdata error
        if c.String("sql") == "" {
            errdata = fmt.Errorf("the sql is null,please exec:%s query [--help|-h]",os.Args[0])
        } else {
            fetchall(c.String("server"),c.String("sql"),c.Int("port"))
        }
				return errdata
			},
		},
		{
			Name:  "stats",
			Usage: "check the query task status.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "queryid",
					Usage: "get a query status.",
				},
				cli.StringFlag{
					Name:   "server, s",
					Value:  "localhost",
					Usage:  "the presto host.",
					EnvVar: "PRESTO_HOST,PRESTO",
				},
				cli.IntFlag{
					Name:   "port,p",
					Value:  8080,
					Usage:  "the presto instances port.",
					EnvVar: "PRESTO_PORT",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println(c.String("queryid"))
				fmt.Println(c.String("server"))
				fmt.Println(c.Int("port"))
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		var reqdata error
		if len(c.Args()) < 3 {
			reqdata = fmt.Errorf("Usage:%s -h|help|--help\n", os.Args[0])
		}
		return reqdata
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
