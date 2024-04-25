package main

import "fmt"

var ( //全局变量申明
	v1 int
	v2 string
)

func main() {
	//var a string = "Runoob"
	//fmt.Println(a)
	//
	//var b, c int = 1, 2
	//fmt.Println(b, c)
	//
	//var d int //默认值0
	//fmt.Println(d)
	//
	//var e bool //默认值false
	//fmt.Println(e)
	//
	//var f string //默认值为 ""
	//fmt.Println(f)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	a := 300
	print(a)

	println(v2)
	print(v1)

	g, h := 100, "hello"
	println(g, h)

	a2 := a
	println(&a, &a2)
}
