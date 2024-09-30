package concurrency

import (
	"fmt"
	"time"
)

func printFromChannel(c chan string) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case msg := <-c:
			fmt.Println(msg)
		default:
			fmt.Println("No message received")
		}
	}
}
func SelectTest() {
	c := make(chan string)
	go printFromChannel(c)
	for {
		var i int
		fmt.Scanln(&i)
		if i == 0 {
			c <- "Hello"
		} else {
			c <- "World"
		}

	}
}
