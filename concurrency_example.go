package main

import (
	"fmt"
	"math/rand"
	"time"
)

func print1() {
	for i := 0; i < 10; i++ {
		fmt.Println("print 1:", 1)
		duration := time.Duration(rand.Intn(250)) * time.Millisecond
		time.Sleep(duration)
	}
}
func print2() {
	for i := 0; i < 10; i++ {
		fmt.Println("print 2:", 2)
		duration := time.Duration(rand.Intn(250)) * time.Millisecond
		time.Sleep(duration)
	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {

		c <- "----ping"

	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {

		c <- "pong----"

	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}
func pingerWithReferenceSharing(c *string) {
	for i := 0; ; i++ {

		*c = "----ping"

	}

}
func pongerWithReferenceSharing(c *string) {
	for i := 0; ; i++ {

		*c = "pong----"

	}

}
func printerWithReferenceSharing(c *string) {
	for {
		msg := *c
		fmt.Println(msg)
		time.Sleep(time.Second)
	}

}

func usingChannels() {
	c := make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	fmt.Scanln()
}
func usingReferenceSharing() {
	var c string
	go pingerWithReferenceSharing(&c)
	go pongerWithReferenceSharing(&c)
	go printerWithReferenceSharing(&c)
	fmt.Scanln()

}
func main() {
	// go print1()
	// go print2()
	usingChannels()
	// usingReferenceSharing()

}
