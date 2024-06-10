package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func t1() {
	s := "ওহে বিশ্ব"
	s += " hello"
	s[3] = 'x'
	fmt.Printf("%8T %[1]v %d\n", s, len(s))
	fmt.Printf("%8T %[1]v\n", []rune(s))
	fmt.Printf("%8T %[1]v %d \n", []byte(s), len([]byte(s)))
}

func main() {
	// t1()
	if len(os.Args) < 3 {
		fmt.Println("Usage: main <input-file> <output-file>")
		os.Exit(-1)
	}
	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)
		fmt.Println(t)
	}
}
