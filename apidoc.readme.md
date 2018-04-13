用法：
~~~cmd

src\github.com\goodmall\goodmall>apidoc -i ./pods/ -o ./docs/
~~~

go语言也有国人实现了一个apidoc工具 不要跟js版的同时装哦！两个同名的程序会被覆盖一个的
[apidoc](github.com/caixw/apidoc)

## swagger

https://github.com/go-swagger/go-swagger

从源码提取annotation 来生成doc
https://github.com/yvasiyarov/swagger 此项目灵感来自beego中的功能 不过进行了用途扩充 使之
使用所有go项目 而不似beego版本跟框架耦合较深
~~~shell

$GOPATH/bin/swagger -apiPackage="github.com/goodmall/goodmall/pods/demo/adapters/api/gin" -mainApiFile=github.com/goodmall/goodmall/cmd/api/gin/main.go -output=./API.md -format=markdown

~~~

- 不支持目录递归扫描 需要精确指定含有api的目录位置g