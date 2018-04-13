BEEGO Swagger 文档生成分析
=============

实现思路：

- 静态扫描项目目录来生成 swagger 规范需要的两个文件 swagger.json 和 swagger.yml

- 整个api文档 分几大部分 （见： github.com/astaxie/beego/swagger）结构定义
    -  api根 Swagger::Information 描述所有api的基础共享信息  比如版本 描述 协议之类
	-  api定义集合 Swagger::Item 这个就是每个api-endpoint的收集 
	-  Definitions 表示的是一些复合类型的定义 就是api方法签名中出现的复合类型（in-param（方法参数类型） out-param（返回值类型））
	
- 结构体构造好后 会直接转换为文件的

文档生成的主要任务也就是构造这个结构体 等结构体构造完后 就可以写出为swagger.json 和 swagger.yml了

## 我们先假设我们需要完成这样的功能（黑盒思路）：

1. 黑盒入参 即为了完成功能所需要的原料 进口：
	项目所在的根目录
2. 黑盒出参 黑盒功能完成后的产出	出口：
	指定目录生成 swagger.json和swagger.yml

3. 实现：
	假设了beego的标准结构：
	~~~
	
	project-root        项目跟目录
	|
	├── vendor          第三方库
	|
	├── main.go
	|             
	├── routers     	集中的路由配置
	| 
	├── controllers     控制器目录
	|
	├── models     		模型目录 
	##            
	
	├── ...             其他目录（对此问题而言 无关紧要）
	
	~~~
	
	先找到路由器的配置 routers/router.go 将作为api文档生成的“根”  其他的api项会顺此根来顺藤摸瓜的
	-  解析出api-root的根信息 即Swagger::Information 
	-  根据router.go 的所有导入包 来遍历分析控制器
		有了控制器所在的包位置 然后计算出对应的目录 在利用go语言提供的解析器来解析这些目录 主要是提取分析评论部分：
	    parser.ParseDir(....., parser.ParseComments)
		来构造抽象语法树
		
	-  分析名空间 和导入的名空间
		
	