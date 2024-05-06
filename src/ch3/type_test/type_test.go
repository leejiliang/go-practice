package type_test

import "testing"

type MyInt int64

func TestImplicint(t *testing.T) {
	var a int32 = 42
	var b int64
	//b=a  // 不支持隐示类型装换
	b = int64(a) // 显示类型转换
	var c MyInt
	//c = b // 同类型别名之间也无法直接进行隐式转换
	c = MyInt(b) // 显示转换
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr=aPtr+1 // 指针不支持运算
	t.Log(aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string // default "" 空字符串
	t.Log("#" + s + "*")
	t.Log(len(s))
}
