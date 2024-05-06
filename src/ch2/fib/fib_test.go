package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 1
	a := 1
	b := 1
	fmt.Print(a, " ")
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tm := a
		a = b
		b = tm + a
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//temp := a
	//a = b
	//b = temp

	a, b = b, a // 交换两个变量的值.
	t.Log(a, b)
}
