package main

import "fmt"

type Employee struct {
	Name   string    `json:"name"`
	Number int       `json:"number"`
	Boss   *Employee `json:"boss"`
}

func main() {
	var e = Employee{
		Name:   "Apurbo",
		Number: 1,
	}
	b := Employee{
		Name:   "Boss",
		Number: 2,
	}
	e.Boss = &b
	fmt.Printf("%T %+v\n", e, e)
}
