Migration
=========

多人开发中 队员对底层存储schema的修改 需要在其他成员的db环境中得到同步 此时 Migration工具是有
效手段

目前市面上的migrations 工具 基本都对迁移文件的存放路径有惯例假设（比如放在项目根目录的migrations
文件夹）

考察了下go社区目前的主要实现  对迁移文件的要求：

- 文件名中暗示出 迁移是向前还是向后 参考这个star最多的[migrate](https://github.com/mattes/migrate)

> {version}_{title}.up.{extension}
> {version}_{title}.down.{extension}

- 迁移文件里面用特殊的字符串暗示出是正向迁移还是反向迁移:[sql-migrate](https://github.com/rubenv/sql-migrate)

~~~

-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE people (id int);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE people;

~~~

- 可以使用编程语言来自己操控 迁移过程 [goose](https://github.com/pressly/goose)

允许使用go文件来自己编写迁移代码

此库克隆自bitbuckt [goose]( https://bitbucket.org/liamstask/goose)项目 实际做了很大幅
度的风格修改（个人感觉没原来的好 不过修改版本也可以 风格比较偏向过程化 原始的oo特征比较重些）并将原始
支持时间戳作为版本前缀的用法改为使用整数序列做为迁移文件的前缀

值得一提的是 原始版本 有个很新颖的做法 对待.go文件的迁移 执行方法是 动态生成一个临时项目到临时目录
然后利用 go run xxx.go  来执行迁移代码 执行完毕后清除目录 这种方式很值得思考！

## 我们的需求

多pod 可能使用不同的存储  这个跟微服务常见的一服务一存储需求类似 我们可以一个pod（隔离仓 模块）一种
存储 

这样要求迁移文件具有分散位置 不同存储的需求

符合这样需求的工具只有goose啦 不过要修改为使用时间戳作为版本前缀 不然多人开发时生迁移文件模板时
很容易出现版本冲突（版本区分上 只记录数字前缀 不管后面的字符串是啥 比如 1_init_user.go,
1_init_items.go 会被认为是相同的版本！ 只考察1 这个数字前缀）

为此稍加修改了原始项目 改为使用时间戳了 使用自己的版本[goose](github.com/goyes/goose)

安装命令工具
>  go get -u github.com/goyes/goose/cmd/goose

使用上同原始版本一样  先把当前目录切到想生成迁移文件的目标目录去 比如pods/user/migrations
然后使用命令：

~~~cmd

...github.com\goodmall\goodmall\pods\demo\migrations>goose mysql "root:@/test?parseTime=true" create init_todo go

~~~

然后想办法加载这个包就可以了

## 多db类型

注意 我们可以使用不同的db配置来执行迁移任务的 尽管函数参数有个事务引用；

~~~go

package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20180328193636, Down20180328193636)
}

func Up20180328193636(tx *sql.Tx) error {
	println("hiiiiii up")
	// This code is executed when the migration is applied.
	return nil
}

func Down20180328193636(tx *sql.Tx) error {
	println("hiii down!")
	// This code is executed when the migration is rolled back.
	return nil
}

~~~ 

我们仍旧可以不使用这个参数 从其他地方获取 或者 注入其他db组件 来执行迁移

这样命令行 db的配置参数 只是管理迁移执行进度的 并不代表我们的迁移表必须生成在对应的这个db库中！

由于migrations目录的 的特殊性 里面的go文件 会扰乱原先逻辑 因此对于需要其他db的场景
只能通过全局组件查询方式来进行了 不能使用依赖注入啦

比如 myDb := app.MysqlDb()




