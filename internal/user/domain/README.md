# domain
domain 在最底层，传统架构之前是数据库，而domain是Entity、Domain Primitive和Domain Service。

**这些对象不依赖任何外部服务和框架，而是纯内存中的数据和操作**。domain 层不再调用 repository ，而是 application 层去调用

> 之前有ddd解析中谈到，domain service为领域服务，编写具体业务逻辑，包括调repository来对数据进行操作。

这些对象我们打包为Domain Layer（领域层）。领域层没有任何外部依赖关系。 这些对象不依赖任何外部服务和框架，而是纯内存中的数据和操作。这些对象我们打包为Domain Layer（领域层）。领域层没有任何外部依赖关系。
## domain 包含的内容介绍

- Aggregate 聚合根定义，其中包括了domain object 的行为 + 业务规则，包括业务校验
- service 领域服务实现，主要处理业务规则。更倾向于封装**多对象**逻辑（所以属于跨domain obeject的业务逻辑）。
  - 如果业务很简单，那么该domain service 的封装可不写，domain object 已经足够了
- repository  仓储接口定义
 <!-- VO 值对象定义 -->

 该层每个包的包名都是一个聚合根
## 领域对象  domain object 

### 不可以强依赖其他聚合根domain object 或领域服务

一个领域对象的原则是高内聚、低耦合，即一个领域对象类不能直接在内部直接依赖一个外部的领域对象或服务。这个原则和绝大多数ORM框架都有比较严重的冲突，所以是一个在开发过程中需要特别注意的。这个原则的必要原因包括：对外部对象的依赖性会直接导致领域对象无法被单测；以及一个领域对象无法保证外部领域对象变更后不会影响本领域对象的一致性和正确性。


所以，正确的对外部依赖的方法有两种：

- 只保存外部领域对象的ID：这里我再次强烈建议使用强类型的ID对象，而不是Long型ID。强类型的ID对象不单单能自我包含验证代码，保证ID值的正确性，同时还能确保各种入参不会因为参数顺序变化而出bug。具体可以参考 [Domain Primitive文章](https://mp.weixin.qq.com/s/tTnj4XHy-Q0S_25VO9F7gQ)。
- 针对于“无副作用”的外部依赖，通过方法入参的方式传入。比如上文中的equip(Weapon，EquipmentService）方法。

如果方法对外部依赖有副作用，不能通过方法入参的方式，只能通过Domain Service解决

### 任何领域对象的行为只能直接影响到本领域对象（和其子领域对象）

这个原则更多是一个确保代码可读性、可理解的原则，即任何领域对象的行为不能有“直接”的”副作用“，即直接修改其他的领域对象类。这么做的好处是代码读下来不会产生意外。

另一个遵守的原因是可以降低未知的变更的风险。在一个系统里一个领域对象对象的所有变更操作应该都是预期内的，如果一个领域对象能随意被外部直接修改的话，会增加代码bug的风险。


## 领域服务 domain service

## 领域服务是否能直接调用Repositoty层，这样的话感觉更内聚?
> 备注：以下是 <淘系技术> 公众号 作者回复

`Domain Service`不应该直接调用 `Repository` (以及其他的跨网络调用)，哪怕`Repository Interface`是在`Domain`层。我们需要看一下`DomainService`的核心目的:**封装业务逻辑(也就是各种规则)，而不是业务流程。**

也就是说`DomainService`天生是Stateless的纯内存操作。`DomainService`的所有入参都必须是上层调用方提前查出来给予的(也就是`ApplicationService`的职责)。

至于说“更内聚”，其实是有问题的，等于是在业务逻辑上加入了一个外部依赖，如果`Repo`有问题，你的`Domain`层都会出问题，连业务逻辑正确性都无法验证。如果说要问“内聚”应该关注啥，我认为应该关注的是`Application`层的边界。

对外部来说`Application`层就是一个领域`BoundedContext`的边界，在这里面的都是内聚的，而外部只需要关注`Application`接口的入参/出参即可
## 业务规则与业务流程怎么区分？

有个很简单的办法区分： 
- 业务规则是有if/else的，业务流程没有
  
## Double Dispatch

在DDD里，一个domain object 不应该直接参考另一个domain object或领域服务，也就是说以下的代码是错误的：
```java

public class Player {
    @Autowired
    EquipmentService equipmentService; // BAD: 不可以直接依赖

    public void equip(Weapon weapon) {
       // ...
    }
}
```
这里的问题是domain object 只能保留自己的状态（或非聚合根的对象）。任何其他的对象，无论是否通过依赖注入的方式弄进来，都会破坏domain object 的Invariance（不变性），并且还难以单测。



正确的引用方式是通过方法参数引入（Double Dispatch）：
```java
public class Player {

    public void equip(Weapon weapon, EquipmentService equipmentService) {
        if (equipmentService.canEquip(this, weapon)) {
            this.weaponId = weapon.getId();
        } else {
            throw new IllegalArgumentException("Cannot Equip: " + weapon);
        }
    }
}
```

Double Dispatch是一个使用Domain Service经常会用到的方法，类似于调用反转。

在这里，无论是Weapon还是EquipmentService都是通过方法参数传入，确保不会污染Player的自有状态。



Double Dispatch是一个使用Domain Service经常会用到的方法，类似于调用反转。

参考：
- [DDD系列 第一弹 - Domain Primitive](https://mp.weixin.qq.com/s/tTnj4XHy-Q0S_25VO9F7gQ)
- [DDD系列 第二弹 - 应用架构](https://mp.weixin.qq.com/s/bhfnyhlKfrPpSh9-6of9xw)
- [DDD系列 第三讲 - Repository模式](https://mp.weixin.qq.com/s/1bcymUcjCkOdvVygunShmw)

