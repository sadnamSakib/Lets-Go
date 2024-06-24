package main

import (
	"fmt"
	"unicode/utf8"
)

func t1() {
	s := "ও "
	s += " hello"
	fmt.Printf("%T %[1]v If we use len function to find length of the string : %d\n", s, len(s))
	fmt.Printf("If we use utf8 package to count the number of cahracters : %d\n", utf8.RuneCountInString(s))
	fmt.Printf("%T %[1]v Length of the array of rune : %d\n", []rune(s), len([]rune(s)))
	fmt.Printf("%T %[1]v  Length of array of bytes : %d\n", []byte(s), len([]byte(s)))
}

func main() {
	t1()
	// if len(os.Args) < 3 {
	// 	fmt.Println("Usage: main <input-file> <output-file>")
	// 	os.Exit(-1)
	// }
	// old, new := os.Args[1], os.Args[2]
	// scan := bufio.NewScanner(os.Stdin)

	// for scan.Scan() {
	// 	s := strings.Split(scan.Text(), old)
	// 	t := strings.Join(s, new)
	// 	fmt.Println(t)
	// }
}
