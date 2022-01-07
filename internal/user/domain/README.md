# domain

该层每个包的包名都是一个聚合根
- Aggregate 聚合根定义
- service 领域服务实现
- repository  仓储接口定义
 <!-- VO 值对象定义 -->



## 领域服务是否能直接调用Repositoty层，这样的话感觉更内聚?
备注：以下是 <淘系技术> 公众号 作者回复

`Domain Service`不应该直接调用 `Repository` (以及其他的跨网络调用)，哪怕`Repository Interface`是在`Domain`层。我们需要看一下`DomainService`的核心目的:**封装业务逻辑(也就是各种规则)，而不是业务流程。**

也就是说`DomainService`天生是Stateless的纯内存操作。`DomainService`的所有入参都必须是上层调用方提前查出来给予的(也就是`ApplicationService`的职责)。

至于说“更内聚”，其实是有问题的，等于是在业务逻辑上加入了一个外部依赖，如果`Repo`有问题，你的`Domain`层都会出问题，连业务逻辑正确性都无法验证。如果说要问“内聚”应该关注啥，我认为应该关注的是`Application`层的边界。

对外部来说`Application`层就是一个领域`BoundedContext`的边界，在这里面的都是内聚的，而外部只需要关注`Application`接口的入参/出参即可
## 业务规则与业务流程怎么区分？

有个很简单的办法区分： 
- 业务规则是有if/else的，业务流程没有
  

参考：
- https://cloud.tencent.com/developer/article/1803939 DDD之Repository