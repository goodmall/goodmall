package builder

import (
	//	"fmt"
	"reflect"

	. "github.com/go-xorm/builder"
)

// FilterCond defines optional conditions when the inner Cond 's operand is empty-like
//
// see https://github.com/yiisoft/yii2/blob/master/framework/db/QueryTrait.php#L231
//
// 缘起: 客户端查询条件一股脑全部提交给db  后端根据查询条件来构造where部分的条件子句 如果对于值是空的字段
//      可以在条件构造中予以忽略
// 例子：对于GET 查询请求 localhost:8080/todos?title=hello&description=
//      对应的条件构造：

/*			// 构造条件子句

sql, args, _ := ToSQL(
And(
	FilterCond(Like{"title", q.Title}),
	FilterCond(Like{"description", q.Description})))

因为请求中description并没有赋值  最终生成的sql片段中 是不含description部分的

其中 q 是一个查询结构体 可以使用 "github.com/gorilla/schema" 库从url请求体来填充： Request.URL.Query()

	***/
type filterCond struct {
	cond Cond
}

var _ Cond = filterCond{}

func FilterCond(cond Cond) Cond {
	return filterCond{cond}
}

// WriteTo writes SQL to Writer
func (fc filterCond) WriteTo(w Writer) error {
	return fc.cond.WriteTo(w)
}

// And implements And with other conditions
func (fc filterCond) And(conds ...Cond) Cond {
	return And(fc, And(conds...))
}

// Or implements Or with other conditions
func (fc filterCond) Or(conds ...Cond) Cond {
	return Or(fc, Or(conds...))
}

// IsValid tests if this condition is valid
func (fc filterCond) IsValid() bool {
	innerCond := fc.cond

	// TODO 目前暂时只处理两个  后续可以加哦
	// switch t := innerCond.(type) {
	switch innerCond.(type) {
	case Eq:
		eqObj := fc.cond.(Eq)
		for k, v := range eqObj {
			// 如果值是空的那么从映射表中删除掉
			if isBlank(reflect.ValueOf(v)) {
				delete(eqObj, k)
			}
		}
		return eqObj.IsValid()
	case Like:
		likeObj := fc.cond.(Like)
		// 检查like 第二个元素的空性
		if isBlank(reflect.ValueOf(likeObj[1])) {
			return false
		}
		return likeObj.IsValid()
	default:
		return !isBlank(reflect.ValueOf(fc.cond)) && fc.cond.IsValid()
	}

}

// borrow from gorm/utils.go
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}

	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
