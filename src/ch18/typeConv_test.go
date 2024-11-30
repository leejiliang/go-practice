package ch18

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAssert(t *testing.T) {
	var i interface{}
	i = "hello"
	if s, ok := i.(string); ok {
		t.Log("断言成功" + s)
	}
	OutPrint(i)
	i = 123
	OutPrint(i)
}

func OutPrint(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("我是string: " + i.(string))
	case int:
		fmt.Print("我是int: " + strconv.Itoa(i.(int)))
	}
}
