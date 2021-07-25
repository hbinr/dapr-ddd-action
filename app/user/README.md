#

## 路由: go-chi
选用go-chi的理由：
- 100% 兼容 `net/http`  重点考虑，没有做过度包装 可以完美集成
- 只需要个路由功能，其他功能，如中间件等，使用 `dapr` 提供的
- 性能很好
- 参数解析和参数校验不打算使用 `go-playground/validate` 库，而是 [`bytedance/go-tagexpr`](https://github.com/bytedance/go-tagexpr)
  解析参数+校验，支持原生 `net/http`，亲测性能优于`validate`
  
