# ports

`ports` 端口是你的应用程序的输入，也是外部世界能够到达它的唯一途径。即入口

它可能是:
- 一个HTTP或gRPC服务器 
  - http: 提供 http 服务
  - grpc: 提供 gprc 服务
- 一个CLI命令，
- 或一个Pub/Sub消息订阅者


