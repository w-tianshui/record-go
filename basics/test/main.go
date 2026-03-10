package main

import (
	"fmt"
)

func sum(nums ...int) int {
	fmt.Printf("传入的参数 nums 的类型是: %T\n", nums) // 查看它的真实类型
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	fmt.Println(sum(1, 2, 3))
	a := make([]int, 3)
	fmt.Printf("%T", a)
}
