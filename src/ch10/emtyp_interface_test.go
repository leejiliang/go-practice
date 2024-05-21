package ch10

import (
	"fmt"
	"testing"
)

func TypeConvert(i interface{}) {

	// 空接口可以转换为任意类型, 格式:
	//if intVal, ok := i.(int); ok {
	//	fmt.Println(intVal)
	//}
	//if strVal, ok := i.(string); ok {
	//	fmt.Println(strVal)
	//}

	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknown type", v)
	}
}

func TestTypeConvert(t *testing.T) {
	TypeConvert(1)
	TypeConvert("hello")
	TypeConvert(1.2222)
}
