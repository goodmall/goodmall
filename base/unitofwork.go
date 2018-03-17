package base

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


要实现这个功能 意味着我们需要注册所有需要被持久化的对象的DbMapper
*/
