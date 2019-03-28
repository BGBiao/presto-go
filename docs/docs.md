
```
# 1. 使用post请求提交一个新的查询，并且获取查询id和nexturl
## 此时会返回查询id,以及nextUri相关信息
## 需要注意的是这里的infoUri不是返回相关json串信息，因此需要查询每个任务详情，可以根据queryid进行查询该查询任务相关的任务

# 根据id查询这个查询任务相关执行状态(任务流关系，状态信息；每个任务详情信息state,taskId,host)
$$ curl -s  "http://localhost:8080/v1/query-execution/20190326_063230_00001_t6972"   | python -m json.tool
# 根据查出来的taskid信息查询每个task详情
curl "http://172.25.238.104:8080/v1/task/20190326_063230_00001_t6972.0.0?pretty"



$ curl -s -X POST  -H 'X-Presto-Catalog:hive' -H 'X-Presto-Schema:tds' -H 'X-Presto-Source:presto-cli' -H 'User-Agent:golang-client' -H 'X-Presto-User:root' --data "select * from ops_app limit 1"  localhost:8080/v1/statement

{
    "id": "20190326_063230_00001_t6972",
    "infoUri": "http://localhost:8080/query.html?20190326_063230_00001_t6972",
    "nextUri": "http://localhost:8080/v1/statement/20190326_063230_00001_t6972/1",
    "stats": {
        "state": "QUEUED",
        "scheduled": false,
        "nodes": 0,
        "totalSplits": 0,
        "queuedSplits": 0,
        "runningSplits": 0,
        "completedSplits": 0,
        "userTimeMillis": 0,
        "cpuTimeMillis": 0,
        "wallTimeMillis": 0,
        "processedRows": 0,
        "processedBytes": 0
    }
}

# 2. 使用获取到的查询id和nexturl进行相关事情(
queryid:20190326_063230_00001_t6972
nexturl: http://localhost:8080/v1/statement/20190326_063230_00001_t6972/1
)

## 注意,当任务完成之后该url将不会有相关数据了

$ curl -s "http://localhost:8080/v1/statement/20190326_063230_00001_t6972/1" | python -m json.tool
{
    "id": "20190326_063230_00001_t6972",
    "infoUri": "http://localhost:8080/query.html?20190326_063230_00001_t6972",
    "partialCancelUri": "http://localhost:8080/v1/stage/20190326_063230_00001_t6972.0",
    "nextUri": "http://localhost:8080/v1/statement/20190326_063230_00001_t6972/2",
    "columns": [
        {
            "name": "appname",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_cnname",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_status",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_level",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_type",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_online_platform",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_line",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_owner",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_backup_owner",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_developers",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_class",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_backup_log",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_sudo_privilege",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_dept_full",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "ip",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "ip_state",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "idc",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "grouptype",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "envtype",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        }
    ],
    "data": [
        [
            "cbe_service",
            "",
            "1",
            "P1",
            "web_tomcat",
            "SURE",
            "\u8de8\u5883\u652f\u4ed8",
            "litaigang",
            "wyliujiangbin",
            "yuhesong,zhangdong8,zhanggaofeng,wyliujiangbin",
            "\u7ba1\u63a7\u53f0",
            "yes",
            "no",
            "\u4eac\u4e1c\u96c6\u56e2-\u4eac\u4e1c\u6570\u5b57\u79d1\u6280-\u4e2a\u4eba\u670d\u52a1\u7fa4\u7ec4-\u652f\u4ed8\u4e8b\u4e1a\u90e8-\u652f\u4ed8\u7814\u53d1\u90e8-\u652f\u4ed8\u4e1a\u52a1\u7814\u53d1\u7ec4",
            "172.23.111.42",
            "\u5728\u7ebf",
            "BJM6",
            "M6",
            "\u751f\u4ea7\u73af\u5883"
        ]
    ],
    "stats": {
        "state": "RUNNING",
        "scheduled": true,
        "nodes": 2,
        "totalSplits": 2,
        "queuedSplits": 0,
        "runningSplits": 0,
        "completedSplits": 2,
        "userTimeMillis": 20,
        "cpuTimeMillis": 20,
        "wallTimeMillis": 123,
        "processedRows": 1024,
        "processedBytes": 608972,
        "rootStage": {
            "stageId": "0",
            "state": "RUNNING",
            "done": false,
            "nodes": 1,
            "totalSplits": 1,
            "queuedSplits": 0,
            "runningSplits": 0,
            "completedSplits": 1,
            "userTimeMillis": 0,
            "cpuTimeMillis": 1,
            "wallTimeMillis": 2,
            "processedRows": 1,
            "processedBytes": 3409,
            "subStages": [
                {
                    "stageId": "1",
                    "state": "FINISHED",
                    "done": true,
                    "nodes": 1,
                    "totalSplits": 1,
                    "queuedSplits": 0,
                    "runningSplits": 0,
                    "completedSplits": 1,
                    "userTimeMillis": 20,
                    "cpuTimeMillis": 19,
                    "wallTimeMillis": 121,
                    "processedRows": 1024,
                    "processedBytes": 608972,
                    "subStages": []
                }
            ]
        }
    }
}


# 3. 再次获取nexturl进行获取数据(如果没有nexturl就直接关闭连接)
$ curl -s  "http://localhost:8080/v1/statement/20190326_063230_00001_t6972/2" | python -m json.tool
{
    "id": "20190326_063230_00001_t6972",
    "infoUri": "http://localhost:8080/query.html?20190326_063230_00001_t6972",
    "columns": [
        {
            "name": "appname",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_cnname",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_status",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_level",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_type",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_online_platform",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_line",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_owner",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_backup_owner",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_developers",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_class",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_backup_log",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_sudo_privilege",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "app_dept_full",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "ip",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "ip_state",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "idc",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "grouptype",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        },
        {
            "name": "envtype",
            "type": "varchar",
            "typeSignature": {
                "rawType": "varchar",
                "typeArguments": [],
                "literalArguments": [],
                "arguments": [
                    {
                        "kind": "LONG_LITERAL",
                        "value": 2147483647
                    }
                ]
            }
        }
    ],
    "stats": {
        "state": "FINISHED",
        "scheduled": true,
        "nodes": 2,
        "totalSplits": 2,
        "queuedSplits": 0,
        "runningSplits": 0,
        "completedSplits": 2,
        "userTimeMillis": 20,
        "cpuTimeMillis": 20,
        "wallTimeMillis": 123,
        "processedRows": 1024,
        "processedBytes": 608972,
        "rootStage": {
            "stageId": "0",
            "state": "FINISHED",
            "done": true,
            "nodes": 1,
            "totalSplits": 1,
            "queuedSplits": 0,
            "runningSplits": 0,
            "completedSplits": 1,
            "userTimeMillis": 0,
            "cpuTimeMillis": 1,
            "wallTimeMillis": 2,
            "processedRows": 1,
            "processedBytes": 3409,
            "subStages": [
                {
                    "stageId": "1",
                    "state": "FINISHED",
                    "done": true,
                    "nodes": 1,
                    "totalSplits": 1,
                    "queuedSplits": 0,
                    "runningSplits": 0,
                    "completedSplits": 1,
                    "userTimeMillis": 20,
                    "cpuTimeMillis": 19,
                    "wallTimeMillis": 121,
                    "processedRows": 1024,
                    "processedBytes": 608972,
                    "subStages": []
                }
            ]
        }
    }
}



```
