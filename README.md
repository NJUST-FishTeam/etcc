# etcc

一个集中式的配置存储中心

## 编译

> go get github.com/codegangsta/cli

> go get github.com/gorilla/mux

> go build

## 运行

> ./etcc [-path /your/config/files/]

**-path** 表明存放配置文件的目录路径，没有的情况下，默认使用`./config`

## API

提供 HTTP+JSON 的 API

配置文件会隶属于一个服务 service

> GET /

获得所有服务的名称

返回：

    {
        "status": "OK",
        "data": [
            "cat",
            "judge",
            "web"
        ]
    }

> POST /

增加一个新的服务

`Form`参数：`service: service_name`

> GET /{service}

获得服务下所有配置

返回

    [
        {
            "name": "dev",
            "service": "web",
            "data": "{\"a\": 111111, \"b\": 222222222}"
        },
        {
            "name": "production",
            "service": "web",
            "data": "{\"b\": \"bbbbbb\", \"aaa\": 12314}"
        }
    ]

> GET /{service}/{config}

获得具体某一配置

返回

    {
        "name": "production",
        "service": "web",
        "data": "{\"b\": \"bbbbbb\", \"aaa\": 12314}"
    }

> POST /{service}/{config}

新建或者修改某一配置

HTTP Body 为配置的 JSON