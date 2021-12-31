# domain

该层每个包的包名都是一个聚合根
- Aggregate 聚合根定义、DO 领域对象定义、 VO 值对象定义
- 领域服务实现
- repository接口定义
- 领域对象创建 + 

注意：

本项目分层结构下， `domain` 不调 `repository`，而是 CQRS 去调用 


