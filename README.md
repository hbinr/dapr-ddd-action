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
    

