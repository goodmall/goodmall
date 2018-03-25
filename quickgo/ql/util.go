package main

import (
	qlBase "github.com/cznic/ql"
)

func Schema(v interface{}, name string, opt *qlBase.SchemaOptions) string {
	schema := qlBase.MustSchema(v, name, opt)

	return schema.String()
}
