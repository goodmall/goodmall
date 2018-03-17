>	Martin Fowler defines a Repository as:
The mechanism between the domain and data mapping layers, acting like an inmemory
domain object collection. Client objects construct query specifications
declaratively and submit them to Repository for satisfaction. Objects can be added
to and removed from the Repository, as they can from a simple collection of
objects, and the mapping code encapsulated by the Repository will carry out the
appropriate operations behind the scenes. Conceptually, a Repository
encapsulates the set of objects persisted in a data store and the operations
performed over them, providing a more object-oriented view of the persistence
layer. Repository also supports the objective of achieving a clean separation and
one-way dependency between the domain and data mapping layers.

##  Repository VS  DAO

DAO 是持久化领域对象到DB的常用模式

跟仓储最大的区别是 仓储表示一个**集合** DAO更接近于DB 并且经常更像以表为中心的（table-centric）
典型地 DAO含有针对某个特定领域对象的CRUD方法

例子：
~~~php

interface UserDAO
{
	/**
	* @param string $username
	* @return User
	*/
	public function get($username);
	public function create(User $user);
	public function update(User $user);
	/**
	* @param string $username
	*/
	public function delete($username);
}

~~~

DAO的实现 可以是基于ORM 或者 常规SQL查询 
DAO最大问题是其职责没有被明确定义，经常像一个DB的网关 相对比较容易极大地降低内聚性 用许多特定的方法
为了查询db

~~~php

interface BloatUserDAO
{
	public function get($username);
	public function create(User $user);
	public function update(User $user);
	public function delete($username);
	
	public function getUserByLastName($lastName);
	public function getUserByEmail($email);
	public function updateEmailAddress($username, $email);
	public function updateLastName($username, $lastName);
}
~~~

如你所见 越添加更多的实现方法 单元测试DAO就越困难 并且其变得更加持续地耦合到User对象 这些问题会随
时间增加的。

## 仓储  Repository

两大类： 面向集合 和  面向持久化

### Collection-Oriented Repositories

~~~go

package main

import "fmt"

type MapFunc func(interface{}) interface{}
type FilterFunc func(interface{}) bool
type Collection []interface{}

type User string
type Host string

func main() {
	c := Collection{User("A"), Host("Z"), User("B"), User("A"), Host("Y")}
	fmt.Println(c.Filter(ByUser("A")).Map(RenameUser("A", "C")))
}

func (c Collection) Map(fn MapFunc) Collection {
	d := make(Collection, 0, len(c))
	for _, c := range c {
		d = append(d, fn(c))
	}
	return d
}

func (c Collection) Filter(fn FilterFunc) Collection {
	d := make(Collection, 0)
	for _, c := range c {
		if fn(c) {
			d = append(d, c)
		}
	}
	return d
}

func ByUser(u User) FilterFunc {
	return func(v interface{}) bool {
		if w, ok := v.(User); ok && w == u {
			return true
		}
		return false
	}
}

func RenameUser(oldu, newu User) MapFunc {
	return func(v interface{}) interface{} {
		if w, ok := v.(User); ok && w == oldu {
			return newu
		}
		return v
	}
}


~~~

### Persistence-Oriented Repository

## 仓储实现
位置：  namespace Infrastructure\Persistence\Redis;

例子如下
~~~php

composer require predis/predis:~1.0
namespace Infrastructure\Persistence\Redis;
use Domain\Model\Post;
use Domain\Model\PostId;
use Domain\Model\PostRepository;
use Predis\Client;

class RedisPostRepository implements PostRepository
{
	private $client;
	public function __construct(Client $client)
	{
		$this->client = $client;
	}
	
	public function save(Post $aPost)
	{
		$this->client->hset(
		'posts',
		(string) $aPost->id(), serialize($aPost)
		);
	}
	
	public function remove(Post $aPost)
	{
		$this->client->hdel('posts', (string) $aPost->id());
	}
	
	public function postOfId(PostId $anId)
	{
		if($data = $this->client->hget('posts', (string) $anId)) {
		return unserialize($data);
		}
		return null;
	}
	
	public function latestPosts(\DateTimeImmutable $sinceADate)
	{
		$latest = $this->filterPosts(
		function(Post $post) use ($sinceADate) {
		return $post->createdAt() > $sinceADate;
		}
		);
		$this->sortByCreatedAt($latest);
		return array_values($latest);
	}
	private function filterPosts(callable $fn)
	{
		return array_filter(array_map(function ($data) {
			return unserialize($data);
		},
		$this->client->hgetall('posts')), $fn);
	}
	
	private function sortByCreatedAt(&$posts)
	{
		usort($posts, function (Post $a, Post $b) {
			if ($a->createdAt() == $b->createdAt()) {
			return 0;
			}
			return ($a->createdAt() < $b->createdAt()) ? -1 : 1;
		});
	}
	
	public function nextIdentity()
	{
	return new PostId();
	}
}

~~~

### 查询仓储

使用 criterion  而不是一些QueryXxxByXxx  

criterion会被翻译为sql或者orm的queries 或者迭代在in-memory集合上的过滤器

#### Specification Pattern

~~~php

interface PostSpecification
{
	/**
	* @return boolean
	*/
	public function specifies(Post $aPost);
}

// We just need to add a query method to our Repository:
interface PostRepository
{
	// ...
	public function query($specification);
}

// in-memory 实现
namespace Infrastructure\Persistence\InMemory;
use Domain\Model\Post;

interface InMemoryPostSpecification
{
	/**
	* @return boolean
	*/
	public function specifies(Post $aPost);
}

namespace Infrastructure\Persistence\InMemory;
use Domain\Model\Post;
class InMemoryLatestPostSpecification

implements InMemoryPostSpecification
{
	private $since;
	public function __construct(\DateTimeImmutable $since)
	{
		$this->since = $since;
	}
	
	public function specifies(Post $aPost)
	{
		return $aPost->createdAt() > $this->since;
	}
}

class InMemoryPostRepository implements PostRepository
{
	// ...
	/**
	* @param InMemoryPostSpecification $specification
	*
	* @return Post[]
	*/
	public function query($specification)
	{
		return $this->filterPosts(
			function (Post $post) use($specification) {
				return $specification->specifies($post);
			}
		);
	}
}

~~~


## 处理事务

一般需要配合UnitOfWork 来实现事务的提交

基本思想是 仓储不负责对象的持久化（ 对于持久化仓储而言 忽略该方法 持久化仓储的实现是细粒度的写 每次变更都
写入到db了 这样会很频繁的调用db 而且多仓储协作完成一个事务时  事务必须包裹在这些仓储写操作的外围 ）
需要持久化的对象 注册自己到UnitOfWork 工作单位上 （新增的 更新过的-dirty 删除过的） 然后再一次性提交给底层db


~~~go

// 可以作为全局共享实现的

package domain

type UnitOfWork interface {
	// 注册跟踪 新增的, 变更|脏的 ，删除的 对象
	registerNew(obj interface{})
	registerDirty(obj interface{})
	registerClean(obj interface{})
	registerDeleted(obj interface{})
	//

	Commit() // 遍历所有变更 逐个写到db去(cud-create update delete)  让后commit
	Rollback()
}

~~~

## 参考

https://programmingwithmosh.com/entity-framework/common-mistakes-with-the-repository-pattern/
