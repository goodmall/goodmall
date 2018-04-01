用户模块
==========

属于整个应用的一个Pod（多module结构中的一个模块  为了避免混淆 所以采用pod）

作为整个系统的一个子模块|子系统  考虑如何集成到系统整体去？
还有自身应该需要的一些配置参数 和依赖组件如何获取
~~~

InitPod(app App , config Config , container Container)
~~~

还有就是模块间通讯方式  app---module间的通信方式 包括设计期 运行期 如何进行信息传递

前端世界现在比较流行的 redux中的Store  是一个全局存储 各个组件自己拉取自己的信息 这种数据通讯的方式也比较有借鉴意义。

基于事件驱动的框架 一般借用event-bus 来进行控制 和信息传递的 这种方式也可以考虑集成进来
在分布式方式下 event-bus 也可以使用一些专业mq 比如rabbitmq 等来作为底层支持。 





## 目录结构说明


### 传统web架构
~~~

mvc
├── controllers
├── models
└── views
~~~

### hex|onion|clean-arch：

~~~

user
├── app         application-service|usercases  interactor
|
├── domain      domain-model entity valueobject|events|exception/error|domain-service|repository-interface
| 
├── infra       infrastructure 
|
##        ├── framework   各个框架的适配器  实现内层定义的接口|适配从外层到内层的请求
##        ├               adapting raw requests into our application

├── delivery    interfaces-layer
                delivery/web/|delivery/cli/|delivery/web/|delivery/desktop/                                         http://retromocha.com/obvious/                

~~~

~~~


.
├── delivery            // Serve content via HTTP? CLI? everything related to that should be here
目录名决定采用 adapters 
|
├── domain              // Where we have our domain logic
|
├── infrastructure      // Where we have our implementation details (Database connections, Queues, External services)
|
└── usecase             // The glue between our delivery layer and our domain layer. Different 
                        // Delivery mechanisms probably will have the same use cases or very similiar
                        // use cases, this allows you to use the same code for different mechanisms by
                        // using the same use case interactors in different mechanisms

~~~

## 主要参考：
- https://github.com/CaptainCodeman/clean-go
- https://github.com/moul/cleanarch
- https://github.com/ManuelKiessling/go-cleanarchitecture

### 墙
- [clean-architecture-using-golang](https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f)

布局：
http://idiomaticgo.com/post/best-practice/server-project-layout/

### 关于用例

在[Organizing modules in a project](https://fsharpforfunandprofit.com/posts/recipe-part3/)
有提到用例
> And just above it is the code for the use cases in the application. The code in this file is where all the functions from all the other modules are “glued together” into a single function that represents a particular use case or service request. (The nearest equivalent of this in an OO design are the [“application services”](http://stackoverflow.com/questions/2268699/domain-driven-design-domain-service-application-service), which serve roughly the same purpose.)

用例等价于ddd中的应用服务application-service 角色



有的资料中 用例出现在领域层
但根据用例驱动开发过程  用例是比较接近用户的  所以出现在应用层或者框架层比较合适（为了隔离框架层--框架是可被替换的属于不稳定比较强的 所以用例最好出现在应用层啦！）

用例的变种名称： Command(见 设计模式中的 命令设计模式)
如： RegisterUserCommand ，UpdateBillingCommand
词法特征： 动-名 结构
在commandBus 寻找处理器的过程中 类名就扮演了key的角色 而某个command的其他属性则是所携带的信息 整个command看起来就是 {cmdType,{attribute-list}}  类型+DTO的感觉。命令类型等价路由key(感觉就像传统web路由path 因为要做到全局唯一 所以类名的抉择需要规划 不然容易冲突) 命令的属性是额外的payload请求负载数据。
这里有个 go语言实现的[命令总线](https://github.com/dadamssg/commandbus))

> Use Case/Command's main benefit is keeping code DRY - we can re-use the same use case code in multiple contexts (web, API, CLI, etc).

> Use Cases also serve to further decouple your application from the framework. This gives some protection from framework changes (upgrades, etc) and also makes testing easier.

用例可以根据关联性归为组  所以我们这里用app-service层来 管理|归组用例
如： UserService:registerUser  UserService:updatePassword

还有当今各大前端框架中的命令通信(比如VUE) 方式： command = {action-type:xxx , payload:{} }

有的资料 程序员  实现用例的手段是： Command-CommandBus-Handler 

[REST?RPC?是时候改变你对微服务的认知了！](https://mp.weixin.qq.com/s/HTeQNU-1P-hWloEdjl1QYg##)
该文中提到 可以把微服务的 请求/响应 方式改为 事件驱动
*Reciver Driven Flow Control，接收者驱动流程控制* 控制反转带来很多好处 

现在流行的Rxjs Rx-xxx 都是用的响应式 可以了解下  如果可以的话 可以打通跟UI上的发布订阅

###  关于仓储Repository

书籍见：《企业应用架构模式》 《ddd》 《实现ddd》 等

仓储的位置 在领域层 和用例层出现  
根据可测试性 领域模型是从仓储中获取的  
[初探领域驱动设计--Repository在DDD中的应用](http://www.uml.org.cn/zjjs/201412112.asp)

[repository-is-dead-long-live-repository](http://codebetter.com/gregyoung/2009/04/23/repository-is-dead-long-live-repository/)
>  but if you apply command and query separation you will have almost none (read: none) read methods on your repositories in your domain.

[ddd-the-generic-repository](http://codebetter.com/gregyoung/2009/01/16/ddd-the-generic-repository/)

~~~C#

public interface ICustomerRepository {
    IEnumerable<Customer> GetCustomersWithFirstNameOf(string _Name);
}

In the customer repository composition would be used.

Public class CustomerRepository {
    private Repository<Customer> internalGenericRepository;
    Public IEnumerable<Customer> GetCustomersWithFirstNameOf(string _Name) {
         internalGenericRepository.FetchByQueryObject(new CustomerFirstNameOfQuery(_Name)); //could be hql or whatever
     }
}
~~~

from[domain-driven-design-inject](http://debasishg.blogspot.com/2007/02/domain-driven-design-inject.html)
>    The main differences between the Repository and the DAO are that :

	The DAO is at a lower level of abstraction than the Repository and can contain plumbing codes to pull out data from the database. We have one DAO per database table, but one repository per domain type or aggregate.
	
	The contracts provided by the Repository are purely "domain centric" and speak the same domain language.

>	The generic layering works like the following :
a) The Struts Action is the Controller layer of your web application. It uses the Domain Service layer for all its action logic. 
b) The Domain Service layer is at a coarser level and uses domain entities and Repositories for implementation. The Domain Services have Repositories injected through some sort of DI.
c) The domain entities, again can have repositories injected to perform domain logic.
	

### 关于infrastructure 基础设施层

http://blog.csdn.net/sven_xu/article/details/46323929

基础设施层，Infrastructure为Interfaces、Application和Domain三层提供支撑。所有与具体平台、框架相关的实 现会在Infrastructure中提供，避免三层特别是Domain层掺杂进这些实现，从而“污染”领域模型。Infrastructure中最常见 的一类设施是对象持久化的具体实现。

[池建强](http://www.infoq.com/cn/articles/cjq-ddd)
> 领域驱动设计除了对系统架构进行了分层描述，还对对象（Object）做了明确的职责和策略划分：
- 实体（Entities）：具备唯一ID，能够被持久化，具备业务逻辑，对应现实世界业务对象。
- 值对象（Value objects）：不具有唯一ID，由对象的属性描述，一般为内存中的临时对象，可以用来传递参数或对实体进行补充描述。
- 工厂（Factories）：主要用来创建实体，目前架构实践中一般采用IOC容器来实现工厂的功能。
- 仓库（Repositories）：用来管理实体的集合，封装持久化框架。
- 服务（Services）：为上层建筑提供可操作的接口，负责对领域对象进行调度和封装，同时可以对外提供各种形式的服务。



## TDD

用例驱动开发的应用 用例的basic course 和 alternative course 是测试的主要方向  特别是分支流不能遗漏 确保全面覆盖

~~~ref
    onion/
    ├── cmd # cmd is for our binaries and artefacts. It is al
    │   ├── onionctl
    │   │   └── main.go #setup and dependency injection
    │   └── server
    │       └── main.go #setup and dependency injection
    └── pkg
        ├── database # database is an infrastructure dependency not part of the core buisiness logic
        │   └── orderRepository.go
        │   └── customerRepository.go
        │   └── ...
        ├── shop #shop is our core business domain
        │   ├── customer #customer is a distinct subdomain
        │   │   └── service.go #Domain Service
        │   │   └── ...    
        │   ├── catalog #catalog is a distinct subdomain
        │   │   └── search.go #Domain Service
        │   │   └── ...        
        │   ├── orders #orders is a distinct subdomain
        │   │   └── dispatch.go
        │   │   └── ...            
        │   └── placeOrder.go  #Application Services or use cases. Bringing suddomains together.      
        │   ├── interfaces.go #holds the interfaces defined as part of our buisness logic, these are implemented by the outer layers (repositories, filehandling, external dependencies and services.)
        │   └── types.go #types holds our domain model
        └── web # web like database is another external dependency
            ├── catalog.go
            ├── orders.go
            └── router.go
~~~            