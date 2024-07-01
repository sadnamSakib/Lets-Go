package main

import "fmt"

// "encoding/json"

// "time"

// type Employee struct {
// 	Name   string
// 	Number int
// 	Boss   *Employee
// 	Hired  time.Time
// }
// type Response struct {
// 	Page   int      `json:"page"`
// 	Fruits []string `json:"fruits"`
// }
func fib() func() int {
	a, b := 0, 1
	return func() int {
		c := 0
		a, b = b, a+b
		fmt.Println("C: ", c)
		return a
	}
}

func fib2() int {
	a, b := 0, 1
	a, b = b, a+b
	return a
}

func main() {
	// f := fib()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }
	fmt.Println(846 / 3)
	// var e = Employee{
	// 	Name:   "Apurbo",
	// 	Number: 1,
	// 	Hired:  time.Now(),
	// }
	// b := Employee{
	// 	Name:   "Boss",
	// 	Number: 2,
	// 	Hired:  time.Now(),
	// }
	// e.Boss = &b
	// fmt.Printf("%T %+v\n", e, e)
	// var r = Response{
	// 	Page:   1,
	// 	Fruits: []string{"apple", "peach", "pear"},
	// }
	// j, _ := json.Marshal(r)
	// fmt.Println(string(j))
	// var r2 Response
	// json.Unmarshal(j, &r2)

	// fmt.Printf("%T %+v\n", r, r)
	// fmt.Printf("%T %+v\n", r2, r2)
	// var mp map[int]int
	// mp[1] = 10
	// fmt.Println(mp)
}
