package demo

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Todo struct {
	Id          int    `json:"id"` // int32
	Created     int    `json:"created"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`

	// 依赖Repo
	repo TodoRepo `json:"-" form:",omitempty"` //
}

// Validate validates the Todo fields.
func (m Todo) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required, validation.Length(0, 120)),
	)
}

func (m Todo) checkTitleDup(value interface{}) error {
	/*
		s, _ := value.(string)

		models, err := m.Repo.Query(nil)

		if len(models) > 0 {
			return errors.New("title duplicated !")
		}
	*/
	return nil
}
