package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case 3, 6, 0:
		fmt.Println("It's the weekend")
	default:
		fmt.Println((int)(time.Now().Weekday()))
		fmt.Println("It's a weekday")
	}
}
