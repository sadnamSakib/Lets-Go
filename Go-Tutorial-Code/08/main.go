package main

import (
	"fmt"
)

func passByValue(b [3]int) int {
	b[0] = 0

	return b[1]

}
func passByReference(b []int) int {
	b[0] = 0
	fmt.Printf("b@ %p\n", b)
	return b[1]
}

func passMap(m map[int]int) {
	m[3] = 0
	m = make(map[int]int)
	m[4] = 4
	fmt.Println("m inside function:", m)
}

func passMapByPointer(m1 *map[int]int, m2 *map[int]int) {
	(*m1)[3] = 0
	(*m2)[4] = 4
	m1, m2 = m2, m1
	fmt.Println("m1 inside function:", *m1)
	fmt.Println("m2 inside function:", *m2)
}
func testVariableScope(a *int) {
	*a = 10
	b := 5
	a = &b
}

func lineBreak() {
	fmt.Println("----------------------------------------------------")
}

func main() {
	a := [3]int{1, 2, 3}
	v := passByValue(a)

	fmt.Println(a, v)
	lineBreak()
	b := []int{1, 2, 3}
	v = passByReference(b)
	fmt.Printf("b@ %p\n", b)
	fmt.Println(b, v)
	lineBreak()

	m := map[int]int{
		1: 5,
		2: 4,
		3: 3,
		4: 2,
		5: 1,
	}
	passMap(m)
	fmt.Println("m", m)

	lineBreak()
	m1 := map[int]int{
		1: 50,
		2: 40,
		3: 30,
		4: 20,
		5: 10,
	}
	m2 := map[int]int{
		1: 5,
		2: 4,
		3: 3,
		4: 2,
		5: 1,
	}
	passMapByPointer(&m1, &m2)
	fmt.Println("m1 outside function:", m1)
	fmt.Println("m2 outside function:", m2)

	val := 1
	testVariableScope(&val)
	fmt.Println(val)

}
