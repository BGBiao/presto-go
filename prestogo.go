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
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"github.com/xxbandy/presto-go/presto"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// 定义一个文件名对象
type fileName struct {
	Name string
}

// 获取文件内容
// 应该首先查看文件是否为可读文件(暂时先使用异常恢复)
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

// 指定presto主机，端口，presto数据源(catalog)，待执行sql来获取数据
func fetchall(host, sql, catalog string, port int, outputfile string) {
	table := tablewriter.NewWriter(os.Stdout)
	fmt.Println("本次查询sql:", sql)
	req, posterr := presto.NewQuery(host, port, "root", "", catalog, "default", sql)
	if posterr != nil {
		fmt.Println(posterr)
	} else {
		fmt.Println(req.GetQueryId())

		// rows为获取到的全量数据
		// func (c *Conn) Next() ([][]interface{}, error)

		rows, fetcherr := req.Next()
		if fetcherr != nil {
			fmt.Println(fetcherr)
		} else {
			fmt.Println("进度:", req.GetProcess())
			fmt.Println("节点:", req.GetNodes())
			fmt.Println("是否关闭:", req.Getclose())
			fmt.Println("状态:", req.GetState())

			// 列名: []string{}
			columns := req.GetColumns()
			fmt.Println("列名:", columns)
			table.SetHeader(columns)

			fmt.Println("数据:")
			if rows != nil {
				// 判断输出文件不为空，创建文件并追加到列名
				if outputfile != "" {

					headers := strings.Join(columns, ",")

					// 判断文件是否存在，存在即删除
					_, fileerr := os.Stat(outputfile)
					// 文件存在即删除
					if fileerr == nil {
						os.Remove(outputfile)
					}

					// 以append模式创建文件
					fileObj, fileerr := os.OpenFile(outputfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
					defer func() {
						err := recover()
						if err != nil {
							fmt.Printf("数据文件:%s 导出失败:%v\n", outputfile, err)
						}
					}()
					defer fileObj.Close()
					// 写入文件头部
					if _, err := io.WriteString(fileObj, headers); err == nil {
					}
					if _, err := io.WriteString(fileObj, "\r\n"); err == nil {
					}
					// 写入数据内容
					for _, v := range rows {
						var paramSlice []string
						for _, param := range v {
							ptype := fmt.Sprintf("%T", param)
							switch {
							case ptype == "string":
								paramSlice = append(paramSlice, param.(string))
							case ptype == "float64":
								tmpparam := strconv.FormatFloat(param.(float64), 'f', 2, 64)
								paramSlice = append(paramSlice, tmpparam)
							default:
								paramSlice = append(paramSlice, "")
							}

						}
						table.Append(paramSlice)
						// 写入文件内容
						if _, err := io.WriteString(fileObj, strings.Join(paramSlice, ",")); err == nil {
						}
						if _, err := io.WriteString(fileObj, "\r\n"); err == nil {
						}
					}

				} else {

					for _, v := range rows {
						// v's type is []interface{}
						var paramSlice []string
						for _, param := range v {
							ptype := fmt.Sprintf("%T", param)
							// 每一列数据类型进行断言，并进行适当格式替换插入到table
							switch {
							case ptype == "string":
								paramSlice = append(paramSlice, param.(string))
							case ptype == "float64":
								// fmt.Printf("%T %T %v\n",ptype,strconv.FormatFloat(param.(float64),'f',2,64),param)
								tmpparam := strconv.FormatFloat(param.(float64), 'f', 2, 64)
								paramSlice = append(paramSlice, tmpparam)
							case ptype == "int64":
								paramSlice = append(paramSlice, param.(string))
							// 假设为<nil>
							default:
								paramSlice = append(paramSlice, "")
							}
						}
						table.Append(paramSlice)
					}
				}

				table.Render()
			}
			fmt.Printf("记录条数:%d\n", len(rows))
		}
		//fmt.Println("任务流状态:")
		//req.GetTasks()
	}
}

/*
// 获取每个流程的状态
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
	app.Version = "0.0.2"

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
				cli.StringFlag{
					// 使用-o对查询内容进行文件输出到csv格式
					Name:  "output,o",
					Usage: "set the records output file(csv)",
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
					fetchall(c.String("server"), sqlfile.getContext(), c.String("catalog"), c.Int("port"), c.String("output"))
				// 再读取sql变量参数
				// 由于sql可能是从文件中获取，因此需要去除换行符
				case strings.Replace(c.String("sql"), "\n", "", 1) != "":
					fetchall(c.String("server"), c.String("sql"), c.String("catalog"), c.Int("port"), c.String("output"))
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
