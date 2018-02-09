目录结构说明
=========

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