样例模块
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


## 参考

- (wtf-dial)[https://medium.com/wtf-dial/wtf-dial-domain-model-9655cd523182]

- [REST Microservices in Go with Gin](http://txt.fliglio.com/2014/07/restful-microservices-in-go-with-gin/)

- [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)



## 目录结构说明

- 



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
          