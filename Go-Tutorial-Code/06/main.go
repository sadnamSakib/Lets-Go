package main

import (
	"fmt"
)

func return1() (int, int) {
	return 1, 2
}
func return2() (int, int) {
	return 3, 4
}

func main() {
	var (
		a, b, c, d int
	)
	// a, b = return1()
	// c, d = return2()
	for i := 0; i < 10; i++ {
		fmt.Println(a, b)
		fmt.Println("--------------")
		var a = a + 1
		fmt.Println(a, b)
		fmt.Println("--------------")
		a, b := return1()
		fmt.Println(a, b)
	}
	fmt.Println(a, b, c, d)
}
