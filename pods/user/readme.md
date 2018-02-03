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
├── domain      domain-model|events|exception/error|domain-service
| 
├── infra       infrastructure 
|
├── framework   各个框架的适配器  实现内层定义的接口|适配从外层到内层的请求
├               adapting raw requests into our application

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

> Use Case/Command's main benefit is keeping code DRY - we can re-use the same use case code in multiple contexts (web, API, CLI, etc).

> Use Cases also serve to further decouple your application from the framework. This gives some protection from framework changes (upgrades, etc) and also makes testing easier.

用例可以根据关联性归为组  所以我们这里用app-service层来 管理|归组用例
如： UserService:registerUser  UserService:updatePassword

还有当今各大前端框架中的命令通信(比如VUE) 方式： command = {action-type:xxx , payload:{} }




    >   
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