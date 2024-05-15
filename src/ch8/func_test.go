package ch8

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
*
返回多个值的函数
*/
func returnMultiVal() (int, int) {
	return rand.Intn(10), 7
}

func TestFn(t *testing.T) {
	a, b := returnMultiVal()
	t.Log(a, b)

	// 构造带计时功能的函数
	tsFG := timeSpent(slowFunc)
	// 执行带计时功能的函数
	t.Log(tsFG(110))

	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5, 6))
}

/*
*
参数是一个函数, 返回值也是一个函数
*/
func timeSpent(inner1 func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now() // 获取开始时间
		ret := inner1(n)    //执行参数函数
		// 计算耗时并打印
		fmt.Println("time spent:", time.Since(start).Seconds())
		// 返回执行结果
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func Sum(ops ...int) int {
	ret := 0
	for _, i := range ops {
		ret += i
	}
	return ret
}
