package main

import (
	"fmt"

	. "github.com/go-xorm/builder"
	b2 "github.com/goodmall/goodmall/base/xorm/builder"
)

func main() {
	// sql, args, _ := ToSQL(Like{"a", "c"})
	// sql, args, _ := ToSQL(And(Eq{"a": 1}, Like{"b", "c"}, Neq{"d", 2}))
	// sql, args, _ := ToSQL(Or(Eq{"a": 1}, Like{"b", "c"}, Neq{"d", 2}))
	// a=? OR b LIKE ? OR d<>? [1, %c%, 2]
	// sql, args, _ := ToSQL(Or(Eq{"a": 1}, And(Like{"b", "c"}, Neq{"d": 2})))
	//sql, args, _ := ToSQL(Eq{"d": 2})

	//sql, args, _ := ToSQL(Between{"a", 1, 2})

	// fmt.Println(sql, args)
	playFilterCond()

}

func playFilterCond() {
	// sql, args, _ := ToSQL(Or(Eq{"a": 1}, And(Like{"b", "c"}, Neq{"d": 2})))
	sql, args, _ := ToSQL(Eq{"a": 1})

	fmt.Println(sql, args, "\n\n")

	sql, args, _ = ToSQL(Like{"b", "c"})

	fmt.Println(sql, args, "\n\n")

	// -------------- ## 测试开始啦 我是分割线哦      ---- ---- --------- ++|

	sql, args, _ = ToSQL(b2.FilterCond(Eq{"a": 1}))

	fmt.Println(sql, args, "\n\n")

	sql, args, _ = ToSQL(b2.FilterCond(Eq{"a": 0}))

	fmt.Println(sql, args, "\n\n")

	sql, args, _ = ToSQL(b2.FilterCond(Like{"b", ""}))

	fmt.Println(sql, args, "\n\n")
}
