# DDD架构分层
DDD的代码模型是按照严格分层进行设计的，主要包括，interface(接口层)、application(应用服务层)、domain(领域层)、infrastructure(基础层)
进行划分，由于采用了严格的分层架构，不允许进行跨层调用，调用逻辑严格按照以下规则进行；
## interface(接口层)：
向上提供前端业务接口，向下只能调用application(应用服务层)，infrastructure(基础层)、不允许跨层调用domain(领域层)的相
关接口或服务
## application(应用服务层)：
向上提供interface(接口层)调用，组织前端的业务逻辑，向下对domain(领域层)相关服务进行服务编排和组合，同时可以
调用infrastructure(基础层)相关服务处理一些跟领域业务逻辑不强的前端业务逻辑。
## domain(领域层)：
为应用服务层提供领域内相关业务逻辑服务，同时向下调用infrastructure(基础层)对相关的DO对象进行数据持久化
## infrastructure(基础层)：
主要存放与基础资源服务相关的代码，为其他各层提供的通用技术能力
