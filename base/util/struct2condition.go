package util

// TODO  任务比较艰巨呀
// 做法：  遍历结构体的字段  然后针对每个字段 判断其类型  如果是字符串就构造like条件 如果是整数|浮点型 则构造Eq条件

import (
	"fmt"
	"reflect"
)

func Struct2Condition(s interface{}) string {
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		fmt.Printf("%#v \n", st.Field(i))
	}
	return ""
}
