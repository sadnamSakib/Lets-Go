package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}
type Response struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
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
	var r = Response{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
	var r2 Response
	json.Unmarshal(j, &r2)

	fmt.Printf("%T %+v\n", r, r)
	fmt.Printf("%T %+v\n", r2, r2)
}
