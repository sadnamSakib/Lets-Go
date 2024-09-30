package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var data = []int{}

func fillArray() {
	for i := 0; i < 160000000; i++ {
		data = append(data, i)
	}

}

func sumArrayConcurrently(s int, e int, wg *sync.WaitGroup, result chan int) {
	sum := 0
	defer wg.Done()
	for i := s; i < e; i++ {
		sum += data[i]
	}
	result <- sum

}

func sumArray(s int, e int) int {
	sum := 0
	for i := s; i < e; i++ {
		sum += data[i]
	}
	return sum
}

func ReadArray() {
	fillArray()
	result := make(chan int)
	wg := sync.WaitGroup{}
	start := time.Now()
	fmt.Println("Summation without concurrency:", sumArray(0, len(data)))
	elapsed := time.Since(start)
	fmt.Printf("Time taken without concurrency: %v \n", elapsed)

	start = time.Now()
	numGoroutines := 4
	chunkSize := len(data) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go sumArrayConcurrently(i*chunkSize, (i+1)*chunkSize, &wg, result)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	sum := 0

	for r := range result {
		sum += r
		fmt.Println("Sum from goroutine:", r)
	}

	elapsed = time.Since(start)

	fmt.Println("Sum using concurrency:", sum)
	fmt.Printf("Time taken with concurrency: %v", elapsed)
}
