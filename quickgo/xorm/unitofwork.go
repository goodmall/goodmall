package xorm

import (
	"log"
	"reflect"

	baseorm "github.com/go-xorm/xorm"

	"github.com/goodmall/goodmall/base"
)

type unitOfWork struct {
	engine *baseorm.Engine
	//
	inserting []interface{}
	updating  []interface{}
	deleting  []interface{}
}

// 注册跟踪 新增的, 变更|脏的 ，删除的 对象
func (uow *unitOfWork) RegisterNew(obj interface{}) {
	// TODO 存在性未检测
	// TODO 必须是对象的引用么？ 还是可以是值？
	uow.inserting = append(uow.inserting, obj)
}
func (uow *unitOfWork) RegisterDirty(obj interface{}) {
	// TODO 必须是对象的引用么？ 还是可以是值？
	uow.updating = append(uow.updating, obj)
}
func (uow *unitOfWork) RegisterClean(obj interface{}) {}
func (uow *unitOfWork) RegisterDeleted(obj interface{}) {
	// TODO 必须是对象的引用么？ 还是可以是值？
	uow.deleting = append(uow.deleting, obj)
}

//
func (uow *unitOfWork) Commit() { // 遍历所有变更 逐个写到db去(cud-create update delete)  让后commit
	// TODO 开启事务

	for idx, obj := range uow.inserting {
		log.Println("inserting the ", idx+1, " object now ")
		affect, err := uow.engine.Insert(obj)
		if err != nil {
			panic(err)
		}
		log.Println("affected num :", affect)

		// log.Println(obj)
	}

	// -----------------------------------------------  +|
	for idx, obj := range uow.updating {
		log.Println("update the ", idx+1, " object now ")
		// TODO 修改对象的方法 需要传递ID？ 根据对象的id来修改的 为了应对更灵活的更新 还是用回调吧！obj.(Updater) 用接口来判断是否实现了这个回调类型
		// type Updater func(engin *baseorm.Engin)
		//
		// affect, err := uow.engine.Update(obj)
		affect, err := uow.engine.ID(uow.idOfStruct(obj)).Update(obj)
		if err != nil {
			panic(err)
		}
		log.Println("affected num :", affect)

		// log.Println(obj)
	}

	// -----------------------------------------------  +|
	for idx, obj := range uow.deleting {
		log.Println("delete the ", idx+1, " object now ")
		// 注意这里传递的对象的 属性会作为删除条件的！ 所以最好只赋予主键值
		// 删除操作除非有统一性  不然可能需要某种回调机制才行
		// https://stackoverflow.com/questions/7850140/how-do-you-create-a-new-instance-of-a-struct-from-its-type-at-run-time-in-go
		// https://blog.golang.org/laws-of-reflection
		// 创建对象的零值： reflect.Zero(reflect.TypeOf(obj)).Interface()
		/**
		*  参考http://xorm.io/docs/ 的delete record 章节
		*       err := engine.Id(1).Delete(&User{})
		*  可以看出我们需要对象的ID 跟其零值结构
		 */
		//		log.Println(uow.idOfStruct(obj))
		sess := uow.engine.ID(uow.idOfStruct(obj))
		if reflect.ValueOf(obj).Kind() == reflect.Ptr {
			// log.Panic("hiiii ")
			// obj = reflect.ValueOf(obj).Elem()
			// log.Panic(uow.idOfStruct(obj))

			affect, err := sess.Delete(uow.zeroValueOf(reflect.ValueOf(obj).Elem().Interface()))
			if err != nil {
				panic(err)
			}
			log.Println("affected num :", affect)

		} else {
			affect, err := sess.Delete(uow.zeroValueOf(obj))
			if err != nil {
				panic(err)
			}
			log.Println("affected num :", affect)
		}

		// log.Println(obj)
	}

	// -----------------------------------------------  +|

	// 清空
	uow.inserting = []interface{}{}
	// 清空
	uow.updating = []interface{}{}
	// 清空
	uow.deleting = []interface{}{}
	// TODO 提交事务
}

func (uow *unitOfWork) idOfStruct(obj interface{}) interface{} {
	tbl := uow.engine.TableInfo(obj)
	pks := tbl.PKColumns()
	if len(pks) == 0 {
		panic("struct have no pks ")
	}
	pkCol := pks[0] // 只取第一个
	v := reflect.ValueOf(obj)
	// 如果是引用
	if v.Kind() == reflect.Ptr {
		return v.Elem().FieldByName(pkCol.FieldName).Interface()
	}

	return v.FieldByName(pkCol.FieldName).Interface()

}

func (uow *unitOfWork) zeroValueOf(obj interface{}) interface{} {
	// TODO check obj type , when it is reflect.Ptr will need another special handling
	return reflect.Zero(reflect.TypeOf(obj)).Interface()
}

func (uow *unitOfWork) Rollback() {}

func NewUintOfWork(engine *baseorm.Engine) base.UnitOfWork {
	log.Println("create unitofwork instance")
	return &unitOfWork{
		engine: engine,
		// inserting: make(interface{}, 0),
	}
}
