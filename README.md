# etcc

一个集中式的配置存储中心

## 安装

> pip install -r requirements.txt

## 部署运行

修改`api/config.py`中`STORE_PATH`的值，表示存储配置文件的路径

> gunicorn -b 0.0.0.0:8009 api.app


## API

提供 HTTP + JSON 的 API

配置文件会隶属于一个服务 service

> GET /services

获得所有服务的名称

返回：

    {
        "status": "success",
        "data": [
            "web"
        ]
    }

> POST /services

创建新的 service

POST 的 body JSON 格式如下

    {
        "action": "create_service",
        "service": "mysql"
    }

> DELETE /services

删除所有的 service


> GET /services/{service_name}

获得服务 service_name 的信息

返回

    {
        "status": "success",
        "count": 1,
        "data": [
            "dev"
        ]
    }

> PUT /services/{service_name}

重命名 service_name

PUT 的 body JSON 格式如下：

    {
        "action": "rename_service",
        "new_name":"web2"
    }

> DELETE /services/{service_name}

删除某个服务

> GET /services/{service_name}/configures

获取某个service_name 服务下所有配置信息

    {
        "status": "success",
        "count": 1,
        "data": [
            {
                "data": {
                    "aaa": 1111,
                    "bb": 22
                },
                "name": "dev"
            }
        ]
    }

> POST /services/{service_name}/configures

增加一个配置

POST 的body JSON 格式如下

    {
        "action": "create_configure",
        "configure_name": "production",
        "configure": {"a":1, "b":11}
    }

> GET /services/{service_name}/configures/{conf_name}

获取一个配置的具体内容

    {
        "status": "success",
        "data": {
            "aaa": 1111,
            "bb": 22
        }
    }

> PUT /services/{service_name}/configures/{conf_name}

更新一个配置的内容

PUT 的 body 的 JSON 就是配置内容


> DELETE /services/{service_name}/configures/{conf_name}

删除某个配置
