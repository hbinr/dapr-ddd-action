## 项目介绍

## 特别说明: 数据CURD

由于[Dapr 状态管理](https://docs.dapr.io/reference/components-reference/supported-state-stores/)提供的 `key/value`API， 不支持关系数据存储或图形数据存储，但是组件却支持关系型数据库，其底层实现 **主键ID为 `key`，值为 `value`**，这样的设计是不支持根据条件查询和范围查询的，相当于将关系型数据库(MySQL)当做了NoSQL来使用

但是看到官方提供了 `Binding`组件，所以选择类型为 `bindings.mysql` 的数据库绑定组件，但是仍遇到了问题:
- SQL注入无法避免，需要手动判断每个入参的合法性
- 多次SQL操作无法保证事务

> 看了下微软官方 [Dapr实战项目](https://github.com/dotnet-architecture/eShopOnDapr) ，使用的是 SQL Server 来存储。所以自己也打算换掉数据存储方式，不再使用Dapr提供的Binding组件
## 状态管理
### 工作原理
应用程序与 Dapr sidecar 交互，以存储和检索键/值数据。 在后台，sidecar API 使用可配置的状态存储组件来持久保存数据。 开发人员可以从不断增长的支持状态存储集合中进行选择，其中包括 Redis、MongoDB 、MySQL等

可以通过 `HTTP` 或 `gRPC` 调用 API。 使用以下 URL 调用 HTTP API：

> http://localhost:<dapr-port>/v1.0/state/<store-name>/

- `<dapr-port>`： Dapr 侦听的 HTTP 端口。
- `<store-name>`：要使用的状态存储组件的名称。



### 本项目示例
`GetUserFromCache` 和 `SaveUserCache` 实现都是调 dapr.Client 的API，分别是 `GetState` 和 `SaveState`
```go
// GetUserFromCache 获取 user(查询缓存)
// Dapr 底层调用 GET 请求: http://127.0.0.1:3500/v1.0/state/ddd-action-statestore/user:info6
// key: user:info6
GetUserFromCache(context.Context, int64) (*do.User, error)

// SaveUserCache 保存 user(缓存)
// Dapr 底层调用 POST 请求 http://127.0.0.1:3500/v1.0/state/ddd-action-statestore
// key: user:info6, data: 为代码中的业务逻辑组成的数据，是个数组，示例如下:
// [{
// 	"key":"user:info6",
// 	"value": {
// 		 "id": 6,
// 		 "user_name": "redis-test333"
// 	}
// }]
SaveUserCache(context.Context, *do.User) error
```
#### 通过 http 调用
提供了http路由：
```go
	r.HandleFunc("/user/{id}", ctl.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/user", ctl.SaveUser).Methods(http.MethodPost)
```
所以可以通过http请求来调用，例如：
`GetUser`
```sh
curl http://127.0.0.1:8090/user/6
```
`SaveUser` 
```sh
curl -X POST http://127.0.0.1:8090/user \
  -H "Content-Type: application/json" \
  -d '[
        {
          "key": "user:info6",
          "user_name": "redis-test333"
        }
      ]'
```


#### 直接通过 Dapr API 调用
```sh
http://localhost:<dapr-port>/v1.0/state/<store-name>/
```
即`GetUser`:
```sh
curl http://127.0.0.1:3500/v1.0/state/ddd-action-statestore/user:info6
```

即`SaveUser`:
```sh
curl  http://127.0.0.1:3500/v1.0/state/ddd-action-statestore \ 
  -H "Content-Type: application/json" \
  -d '[
        {
          "key": "user:info6",
          "user_name": "redis-test333"
        }
      ]'
```

这两种方式都能调通，但是第一种是我们希望真正调用的，它会执行业务逻辑，第二种是Dapr 提供的标准API调用，不执行业务逻辑，直接操作状态存储组件

所以在第二种方式下，你甚至能直接发起DELETE请求，删除指定key的数据，例如：
```sh
curl -v -X DELETE http://127.0.0.1:3500/v1.0/state/ddd-action-statestore/user:info6

*   Trying 127.0.0.1:3500...
* Connected to 127.0.0.1 (127.0.0.1) port 3500 (#0)
> DELETE /v1.0/state/ddd-action-statestore/user:info6 HTTP/1.1
> Host: 127.0.0.1:3500
> User-Agent: curl/7.77.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 204 No Content
< Server: fasthttp
< Date: Sun, 02 Jan 2022 10:08:43 GMT
< Traceparent: 00-695e148b46b68f36333ee2a26527d1eb-d21a1adde62b82b7-00
< 
* Connection #0 to host 127.0.0.1 left intact
```

那这就带来一个问题，如果是团队开发，通过配置文件肯定知道 `<dapr-port>` 和 `<store-name>`，再从项目中知道业务Redis key，那么这些数据不是透明的了？

直接调 Dapr client API，可以对数据进行操作，如果和公司有矛盾，执行了恶操作，直接删除了呢？感觉会有企业极大的安全风险，有没有安全措施呢？

待去深入探索Dapr，我觉得设计Dapr的人可能考虑到这个问题了，不知道后面的官方文档有没有提如何解决，如果没有，那这个安全风险太大了，全靠人自觉
## 如何运行项目？
run:
```bash
cd app/user

dapr run \
    --app-id server-service-demo \
    --app-port 8080 \
    --app-protocol server \
    --dapr-server-port 3500 \
    --components-path ./config \
    go run main.go
```

or 
```bash
cd app/user
make run
```

## 分支介绍
- `main` 主分支
- `dapr-study`  是学习dapr所创建的分支,包含:
    - 使用dapr开发 http 服务
    - dapr 集成 gin 或 mux
    - 绑定MySQL组件
- `ddd` 使用 DDD 思想来开发一个完整的项目 
    


## VS Code 中Debug Dapr 应用

### 1. 安装 Dapr 插件
插件市场搜索 dapr 安装即可
### 2. debug 启动 dapr 应用
第一次点 debug 图标按钮启动应用，会自动生成相关文件：
- `extensions.json` 插件配置文件
- `launch.json` 调试配置文件
- `task.json` 调试任务配置文件，其中定义了启动应用的命令、参数、调试环境等
- `setting.json`  设置 `go.inferGopath` 为 `false`

`task.json` 示例定义如下
```json
{
	"version": "2.0.0",
	"tasks": [
		{
			"appId": "dapr-user-service",
			"appPort": 8091,
            "appProtocol": "http",
			"label": "daprd-debug",
			"type": "daprd",
            "logLevel": "debug",
			// "componentsPath": "${workspaceFolder}/dapr/components",
			"componentsPath": "/Users/xx/go/workspace/my-project/dapr-ddd-action/dapr/components",
			// "config": "${workspaceFolder}/internal/user/config/configs.yaml",
			"config": "/Users/xx/go/workspace/my-project/dapr-ddd-action/internal/user/configs/config.yaml",
		},
		{
			"appId": "dapr-user-service",
			"label": "daprd-down",
			"type": "daprd-down"
		}
	]
}
```

其中:
- `appId` 为应用的id
- `appPort` 为应用的端口
- `appProtocol` 为应用的协议，示例中指定了 http
- `type` 为 `daprd` 或 `daprd-down`
- `logLevel` 为日志级别，示例中指定了 debug
- `componentsPath` 为dapr使用的组件路径
- `${workspaceFolder}` 是项目的根目录，需要手动指定