Request Scope
======

> RequestScope contains the application-specific information that are carried around in a request.

这个的详细用法见强哥：[request scope](https://github.com/qiangxue/golang-restful-starter-kit/blob/master/app/scope.go)


关于这个用法的介绍 可以参考：
[organising-database-access](http://www.alexedwards.net/blog/organising-database-access)

## 其他思考

在每个方法中传递requestScope 对象 这样很多方法签名都需要添加一个额外参数  如果是基于接口风格的编程
那么这个添加数量是很可观的 而且很多方法的实现并不需要它

但如果不是传递它 而是直接通过某种“空降”的方式获取（比如 全局依赖：settings.Request.Context...） 
这样又有点隐式依赖的意味  违反 迪米特法则

参考 [beego 控制器](https://github.com/astaxie/beego/blob/master/controller.go)
折中： 请求域信息 可以传递给控制器 然后控制器方法从控制器获取 这样貌似所有方法共享了 上下文信息 所需的信息只管从
当前session 当前Ctx 中获取即可



在考虑接口稳定性（方法签名 尽量保持长久不变---  不然 方法使用者就惨了 特别是第三方client ）和灵活性
上需要一个折中  这个需要仔细思考  区分必须参数 可选参数 如何影响接口签名！


总之 还是要根据情况使用特定技术 不要滥用 也不要为了用而用 使用场景尽量符合技术出现的初衷 比如context.Context
