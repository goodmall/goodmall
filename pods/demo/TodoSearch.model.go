package demo

import (
	"github.com/go-ozzo/ozzo-validation"
)

type TodoSearch struct {
	Todo
}

// Validate validates the Todo fields.
func (m TodoSearch) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title /* validation.Required,*/, validation.Length(0, 120)),
	)
}
