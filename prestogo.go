package main

/*
需求列表:
- [X] command line 支持catalog的选择
- [X] 增加--file指定sql文件
- [X] 增加数据记录数量
- [ ] 增加任务详情查询
*/
import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/xxbandy/presto-go/presto"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type fileName struct {
	Name string
}

func (c *fileName) getContext() string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("read file %s with error :%s\n", c.Name, err)
		}
	}()
	fileobj, _ := os.Open(c.Name)
	defer fileobj.Close()
	contents, _ := ioutil.ReadAll(fileobj)

	return strings.Replace(string(contents), "\n", "", 1)
}

func fetchall(host, sql, catalog string, port int) {
	fmt.Println("本次查询sql:", sql)
	req, posterr := presto.NewQuery(host, port, "root", "", catalog, "default", sql)
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
			fmt.Printf("记录条数:%d\n", len(rows))
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

	// commands 仅支持query,stats
	app.Commands = []cli.Command{
		{
			Name:      "query",
			ShortName: "q",
			//Aliases: []string{"q"},     // like shortname1,shortname2
			Usage:       "submit a query to presto.",
			Description: "fetch the data from presto with sql.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "sql",
					//Value:    "select * from default.test",
					Usage:    "a query langurage with presto synax.",
					FilePath: "/export/server/presto.sql",
				},
				cli.StringFlag{
					Name:   "server, s",
					Value:  "10.221.196.92",
					Usage:  "the presto host.",
					EnvVar: "PRESTO_HOST,PRESTO",
				},
				cli.StringFlag{
					Name:   "catalog",
					Value:  "hive",
					Usage:  "set the presto catalog for query [hive|kafka|mysql...]",
					EnvVar: "PRESTO_CATALOG",
				},
				cli.StringFlag{
					// 指定sql文件进行查询，需要读取sql文件内容
					Name:  "file,f",
					Usage: "set the sqlfile with presto query language.eg:/export/server/presto.sql",
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
				sqlfile := fileName{Name: c.String("file")}

				// 使用switch case 比较合适一些
				switch {
				// 先读取--file内容
				case sqlfile.getContext() != "":
					fetchall(c.String("server"), sqlfile.getContext(), c.String("catalog"), c.Int("port"))
				// 再读取sql变量参数
				// 由于sql可能是从文件中获取，因此需要去除换行符
				case strings.Replace(c.String("sql"), "\n", "", 1) != "":
					fetchall(c.String("server"), c.String("sql"), c.String("catalog"), c.Int("port"))
				default:
					errdata = fmt.Errorf("the sql is null,please exec:%s query [--help|-h]", os.Args[0])
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
