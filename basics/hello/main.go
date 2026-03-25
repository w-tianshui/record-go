package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextInt := intSeq()    // 定义nextInt为一个函数，nextInt在被调用的时候会让自由变量/捕获变量i++，并输出
	fmt.Println(nextInt()) // i = 1
	fmt.Println(nextInt()) // i = 2
	fmt.Println(nextInt()) // i = 3
	fmt.Println("hello world")
}
