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
			// "componentsPath": "${workspaceFolder}/app/user/components",
			"componentsPath": "/Users/duanhaobin/go/workspace/my-project/dapr-ddd-action/app/user/components",
			// "config": "${workspaceFolder}/app/user/config/configs.yaml",
			"config": "$/Users/duanhaobin/go/workspace/my-project/dapr-ddd-action/app/user/config/configs.yaml",
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
- `${workspaceFolder}` 是项目的根目录，需要手动指定，eg: `/Users/duanhaobin/go/workspace/my-project/dapr-ddd-action`
