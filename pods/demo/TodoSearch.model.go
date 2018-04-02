package demo

import (
	"github.com/go-ozzo/ozzo-validation"
)

// FIXME 关于表名 比较疑惑  Todo最好实现TableName 方法吧！
type TodoSearch struct {
	Todo
}

// Validate validates the Todo fields.
func (m TodoSearch) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title /* validation.Required,*/, validation.Length(0, 120)),
	)
}
