package main // 和java不同的是, package和目录不一定是完全相同的.

import (
	"fmt"
	"os"
)

func main() {
	// 获取命令行参数
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}
	fmt.Println("hello world")
	os.Exit(0) // 返回程序执行状态的方式
}
