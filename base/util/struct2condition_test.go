package util

import (
	"testing"
	//	"github.com/stretchr/testify/assert"
)

func TestStruct2Cond(t *testing.T) {
	t.Log("hiii ")
	Struct2Condition(struct{ Name string }{Name: "yiqing"})
}
