package base

type UnitOfWork interface {
	// 注册跟踪 新增的, 变更|脏的 ，删除的 对象
	RegisterNew(obj interface{})
	RegisterDirty(obj interface{})
	// registerClean(obj interface{})
	RegisterDeleted(obj interface{})
	//
	Commit() // 遍历所有变更 逐个写到db去(cud-create update delete)  让后commit
	// Rollback()
}

// @see https://ders.github.io/post/2016-12-23-lean-clean-golang-machine/
// @see http://docs.sqlalchemy.org/en/latest/orm/session_basics.html
// @see http://martinfowler.com/eaaCatalog/unitOfWork.html
/***

 工作单位是PAEE设计模式中的一个
 主要配合Repository设计模式来完成 对象的事务持久化操作
 把需要持久到db的对象 全部注册到工作单位上 业务操作完成后 统一做事务的提交
 内部维护了三个集合（inserting  updating  deleting）
 事务提交时遍历这些集合 分别调用对象对应mapper的 insert update delete


 还需要配合 MapperRegistry 来实现cud操作

>  MapperRegistry.GetMapper(obj.class).insert(obj) // insert|update|delete

代码的意思就是有一个Mapper集中营 注册了各种实体对应的DbMapper 这里需要用到反射啦
在持久化某个领域对象时  我们先需要在注册表中得到其对应的Mapper  然后再调用对应的写方法

有一个这样的库 不妨看看： https://github.com/go-gorp/gorp


要实现这个功能 意味着我们需要**提前**注册所有需要被持久化的对象的DbMapper

## 关于提交的实现

- 由工作单位来自己完成所有操作 全权实现
  这种实现 模型对UnitOfWork是无感知的 但对于特殊情形无法应对

- 创建 更新 删除 这些操作会被委托给模型对象
  这种实现需要把相应的操作 委托给模型对象 意味着模型对UnitOfWork 承担一定的职责（即UnitOfWork
  对模型有要求 -- 有侵入性）



*/
