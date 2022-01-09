# Dapr 实战
## DDD 架构

- DDD 端口和适配器模型 + CQRS 
- domain 层采用充血模型

## 路由 gorilla/mux
 go-sdk >= 1.3 不再支持 go原生 `net/http`的`ServerMux`，而是使用`gorilla/mux`来作为其路由
### 之前是go-chi（dapr/go-sdk <=1.2）
选用go-chi的理由：
- 100% 兼容 `net/http`  重点考虑，没有做过度包装 可以完美集成
- 只需要个路由功能，其他功能，如中间件等，使用 `dapr` 提供的
- 性能很好
- 参数解析和参数校验不打算使用 `go-playground/validate` 库，而是 [`bytedance/go-tagexpr`](https://github.com/bytedance/go-tagexpr)
  解析参数+校验，支持原生 `net/http`，亲测性能优于`validate`

## DDD 分层图
![](https://pic3.zhimg.com/80/v2-7b9f668218cee34d21d5a6a966e46602_1440w.jpg)

左边为新DDD分层，右边为传统DDD分层

domain层是业务规则的集合，application service编排业务，domain service编排领域；

domain体现在业务语义显现化，不仅仅是一堆代码，代码即文档、代码即业务；
要达到高内聚就得充分发挥domain层的优势，domain层不单单是domain service，还有entity、vo、aggregate

domain层是最最需要拥抱变化的一层，为什么？domain代表了业务规则，业务规则来自于需求，日常开发中，需求是经常变化的

我们需要逆向思维，以往我们去封装第三方服务，解耦外部依赖，大多数时候是考虑外部的变化不要影响自身，而现实中，更多的变化来自内部：需求变了，所以我们应该更多关注一个业务架构的目标：独立性，不因外部变化而变化，更要不因自身变化影响外部服务的适应性

在《DDD之Repository》中指出Domain Service是业务规则的集合，不是业务流程，所以Domain Service不应该有需要调用到Repo的地方。

如果需要从另一个地方拿数据，最好作为入参，而不是在内部调用。DomainService需要是无状态的，加了Repo就有状态了。domainService是规则引擎，appService才是流程引擎。Repo跟规则无关

也就是domain层应该是一个纯内存操作，不依赖外部任何服务，这样提高了domain层的可测试性，拥抱变化的底气也来自于完整的UT，而application层UT全部得mock
## 层与层之间的依赖关系
>旧版， 借鉴Java: starter -> controller -> application -> domain -> infrastructure

现在实现：
main -> port -> application  -> adapters(infrastructure) -> domain 

有些类似菱形架构

![菱形架构图](https://pic3.zhimg.com/80/v2-4c1692a5538ecf16779db7c4bb3979d2_1440w.jpg)

![菱形架构图2](https://pic2.zhimg.com/80/v2-7aabac19536d028a6f8cf30738bc60f1_1440w.jpg)
把六边形架构与分层架构整合时，发现六边形架构与领域驱动设计的分层架构存在设计概念上的冲突

出口端口用于抽象领域模型对外部环境的访问，位于领域六边形的边线之上。根据分层架构的定义，领域六边形的内部属于领域层，介于领域六边形与应用六边形的中间区域属于基础设施层，那么，位于六边形边线之上的出口端口就应该既不属于领域层，又不属于基础设施层。它的职责与属于应用层的入口端口也不同，因为应用层的应用服务是对外部请求的封装，相当于是一个业务用例的外观。

根据六边形架构的协作原则，领域模型若要访问外部设备，需要调用出口端口。依据整洁架构遵循的“稳定依赖原则”，领域层不能依赖于外层。因此，出口端口只能放在领域层。事实上，领域驱动设计也是如此要求的，它在领域模型中定义了资源库（Repository），用于管理聚合的生命周期，同时，它也将作为抽象的访问外部数据库的出口端口。

将资源库放在领域层确有论据佐证，毕竟，在抹掉数据库技术的实现细节后，资源库的接口方法就是对聚合领域模型对象的管理，包括查询、修改、增加与删除行为，这些行为也可视为领域逻辑的一部分。

然而，限界上下文可能不仅限于访问数据库，还可能访问同样属于外部设备的文件、网络与消息队列。

为了隔离领域模型与外部设备，同样需要为它们定义抽象的出口端口，这些出口端口该放在哪里呢？如果依然放在领域层，就很难自圆其说。

例如，出口端口EventPublisher支持将事件消息发布到消息队列，要将这样的接口放在领域层，就显得不伦不类了。倘若不放在位于内部核心的领域层，就只能放在领域层外部，这又违背了整洁架构思想。

如果我们将六边形架构看作是一个对称的架构，以领域为轴心，入口适配器和入口端口就应该与出口适配器和出口端口是对称的；同时，适配器又需和端口相对应，如此方可保证架构的松耦合。

![菱形架构图3](https://pic4.zhimg.com/80/v2-27dc8f6623871478b801a999c5c1f927_1440w.jpg)
```shell
<modules>
 <module>assist-ohs</module> <!-- ohs -->
 <module>assist-service</module> <!-- domain -->
 <module>assist-acl</module> <!-- acl -->
 <module>starter</module> <!-- 启动入口及test -->
</modules>
```
> 代码实战:[https://github.com/agiledon/diamond](https://github.com/agiledon/diamond)

这有点类似《DDD之形》中提到的端口模式，把资源库Repository从domain层转移到端口层和其它端口元素统一管理，原来的四层架构变成了三层架构，对repository的位置从物理与逻辑上一致，相当于扩大了ACL范围

这个架构结构清晰，算是六边形架构与分层架构的融合体，至于怎么选择看个人喜爱
## 入口、出口适配器作用
- 入口适配器负责处理系统外部发送的请求，也就是驱动应用程序运行的用户、程序、自动化测试或批处理脚本会向入口适配器发起请求，适配器将该请求适配为符合内部应用程序执行的输入格式，转交给端口，再由端口调用应用程序。

- 出口适配器负责接收内部应用程序通过出口端口传递的请求，对其进行适配后，向位于外部的运行时设备和数据库发起请求。

![示例](https://img-blog.csdnimg.cn/img_convert/4aeaa7ba2bb515771a24118f9b6655b7.png)

以预定机票场景为例，用户通过浏览器访问订票网站，向订票系统发起订票请求。根据六边形架构的规定，前端 UI 位于应用六边形之外，属于驱动应用程序运行的起因。

订票请求发送给以 RESTful 契约定义的资源服务ReservationResource，它作为入口适配器，介于应用六边形与领域六边形的边界之内。

ReservationResource在接收到以 JSON 格式传递的前端请求后，将其转换（反序列化）为入口端口ReservationAppService需要的请求对象。

入口端口为应用服务，位于领域六边形的边界之上。当它在接收到入口适配器转换后的请求对象后，调用位于领域六边形边界内的领域服务TicketReservation，执行领域逻辑。

在执行订票的领域逻辑时，需要向数据库添加一条订票记录。这时，位于领域六边形边界内的领域模型对象会调用出口端口ReservationRepository。

出口端口为资源库，位于领域六边形的边界之上，定义为接口，真正访问数据库的逻辑则由介于应用六边形与领域六边形边界内的出口适配器ReservationRepositoryAdapter实现。

该实现访问了数据库，将端口发送过来的插入订票记录的请求转换为数据库能够接收的消息，执行插入操作。



本项目调用链：
> main -> ports -> application -> domain (调用领域服务)

其中，领域服务中数据读写的实际逻辑是在 adapters 实现

## 开发思路
如果开始一个新的业务，在划分好领域模型后，考虑到最终的依赖关系：
- 我们可能先写Domain层的业务逻辑，
- 然后再写Application层的组件编排
- 最后才写每个外部依赖的具体实现
  
这种架构思路和代码组织结构就叫做`Domain-Driven Design`（领域驱动设计，或DDD）。所以DDD不是一个特殊的架构设计，而是所有Transction Script代码经过合理重构后一定会抵达的终点。


参考：
- https://zhuanlan.zhihu.com/p/401604739 DDD落地指南
- https://blog.csdn.net/u012921921/category_11421961.html
- [DDD系列 第五讲 - 如何避免写流水账代码](https://mp.weixin.qq.com/s/1rdnkROdcNw5ro4ct99SqQ)