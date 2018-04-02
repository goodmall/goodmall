package main

import (
	"fmt"

	"github.com/doug-martin/goqu"
)

func main() {
	e := goqu.ExOr{
		"col1": "a",
		"col2": 1,
		"col3": true,
		"col4": false,
		"col5": nil,
		"col6": []string{"a", "b", "c"},
	}

	fmt.Printf("%#v", e)
}
