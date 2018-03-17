package demo

// TodoRepo manager the entity Todo as a collection
// （Collection-oriented Repositories(add addAll remove removeAll ...)
//  |Persisted-Oriented Repositories()）
//
// 在ddd中 只有聚合根才有对应的Repo 并非所有实体 领域模型都有对应的Repo！多个实体也可以只有一个Repo
type TodoRepo interface {

	// ##  finder methods (latestItems(since) )
	Load(id int) (Todo, error)
	FindById(id int) Todo

	//
	Store(td *Todo) error
	Remove(td *Todo)

	// ## Extra Behavior
	// Size()

	// ## Query
	// Query(spec Specification)
	Query(criteria Query)
}

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
