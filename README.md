## 如何运行项目？
run:
```bash
dapr run \
    --app-id http-service-demo \
    --app-port 8080 \
    --app-protocol http \
    --dapr-http-port 3500 \
    --components-path ./config \
    go run main.go
```

or run `make run`

## 分支介绍
- `main` 主分支
- `dapr-study`  是学习dapr所创建的分支,包含:
    - 使用dapr开发 http 服务
    - dapr 集成 gin 或 mux
    - 绑定MySQL组件
- `ddd` 使用 DDD 思想来开发一个完整的项目 
    

