package demo

//"github.com/goodmall/goodmall/base"

// TodoRepo manager the entity Todo as a collection
// （Collection-oriented Repositories(add addAll remove removeAll ...)
//  |Persisted-Oriented Repositories()）
//
// 在ddd中 只有聚合根才有对应的Repo 并非所有实体 领域模型都有对应的Repo！多个实体也可以只有一个Repo
type TodoRepo interface {

	//Store(td *Todo) error
	//Save(td *Todo) error
	Create(td *Todo) error
	//
	Update(id int, td *Todo) error
	// Remove(td *Todo)
	Remove(id int) error

	// ## Query methods:

	// ##  finder methods (latestItems(since) )
	Load(id int) (*Todo, error)

	// Query(spec Specification)
	// 实现方法 可以参考 https://www.sohamkamani.com/blog/2017/10/18/golang-adding-database-to-web-application/
	// 有人用string来表示查询串  这个有点跟url中的query串类似 ：?page=0&per-page=10&name=someName&age=10&title=...
	Query(sm TodoSearch /*criteria base.Query*/) ([]Todo, error)

	// ## Extra Behavior
	// Size()
	Count() (int, error)
}

/**
				从DDD观点
关于仓储和dao 的根本区别：
-  Repository 是领域层的东西 讲的是领域语言（ubiquitous-language）
-  DAO 是面向存储介质的 比如db表 或者mongodb的集合 ...
-  DAO更底层 Repository抽象级别更高点  Repository可以使用DAO来完成其业务接口！

- Repository 作为聚合根 可以获取领域对象  可以被其他领域对象引用（通常Repository只被更上层的
service层 使用  但也是可以被领域对象使用的！）

>The main differences between the Repository and the DAO are that :

- The DAO is at a lower level of abstraction than the Repository and can contain
 plumbing codes to pull out data from the database. We have one DAO per database
 table, but one repository per domain type or aggregate.

- The contracts provided by the Repository are purely "domain centric" and
 speak the same domain language.



*/
/**
错误认知：   FindByXxx 是DAO式接口特征 Repository不建议用 一般只需要一个Query就够了
有人认为 Query(query string|SomeQueryInterface) []object 这种风格的接口不具有业务语义
是很宽泛的接口，领域层的接口应该具象化 应该具有业务语义 所以GetByXxx 是可以接受的！
> Hint: Minimize the complexity at the entry/exist seams of your domain by making
 the seams as explicit as possible.

在.net 中有GenericRepository实现  对于需要明确业务语义的仓库 可以组合它（不是继承哦  虽然有很多人
采用继承手段 但我们不需要我们用不上的接口 对于其他通用操作仍旧可以使用泛化的仓储实现）
*/
/**

查询风格跟DAO的不同  当然有人也设计成类似DAO的那样  主要看查询性质方法如：FindByXxx|GetByXxx 系列方法
因为出现在接口上 如果有很多实现类（比如针对不同的DB存储）在增加此类方法时 需要修改的类数量会很快增长
包括单元测试需要变动的地方也很多

~~~go

package user

import "time"

type Repository interface {
	Save(user User) error
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
	All() (CollectionRepository, error)
}

type CollectionRepository interface {
	AddActiveFilter() error
	AddCreatedAtFilter(time.Time, time.Time) error
	Get() ([]User, error)
	Slice(int, int) error
}

~~~

         Persistence-Oriented Repository
面向集合的Repo 有时候跟持久化机制不那么契合 此时就可以使用面向持久化的Repo了
比如

~~~php

interface PostRepository
{
	public function nextIdentity();
	public function postOfId(PostId $anId);
	public function save(Post $aPost);
	public function saveAll(array $posts);
	public function remove(Post $aPost);
	public function removeAll(array $posts);
}
~~~
特征：

add|addAll  --->   save|saveAll

使用风格也不一样
- 面向集合的仓储，add方法只调用一次
- 面向存储的仓储，save方法 在首次创建聚合对象时 和聚合被修改时都需要调用

*/
