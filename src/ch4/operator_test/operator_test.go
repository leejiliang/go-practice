package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestOperatorArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	d := [...]int{1, 2, 3}
	b := [...]int{1, 3, 2}
	c := [...]int{3, 2, 1}
	//e := [...]int{1, 2, 5, 6}
	t.Log(a == b)
	t.Log(a == c)
	t.Log(a == d)
	//t.Log(a == e) // 长度不同的数组直接比较会报错
}

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
