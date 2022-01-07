# adapters
## 一般的 `adapters` 层作用
`adapters` 适配器是你的应用程序如何与外部世界对话。即出口

你必须使你的内部结构适应外部 API 所期望的东西，想象一下： 
- SQL 查询
  - repository 接口实现，比如MySQL的CURD
  - 数据模型 entity 定义？
- HTTP 或 gRPC 客户端
  - http: 直接调 http 接口即可
  - gRPC: 使用别人提供的 proto 文件，在本地生成对应的调用 client 
- 文件读写器
- Pub/Sub 消息发布器

## 使用 dapr 后，该层作用？

使用 Dapr 编写应用后，只要项目(简称A项目)启动，就会创建一个 dapr sidecar，项目的服务也就发布了

那别的项目如何调用 A 项目服务：
- 如果也是使用dapr写的，那么在调用时是通过 Dapr 提供的API来调用，每个 dapr sidecar之间是走的 grpc 协议
- 如果不是，那么http服务就通过访问 http 接口来调用； grpc服务需要使用别人提供的 proto 文件，在本地生成对应的调用 client ，然后再调用对应的方法即可

## gorm-gen SQL生成命令
```
gentool -dsn "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local" -tables "user" -outPath "./data/dao" -modelPkgName "entity" 
```