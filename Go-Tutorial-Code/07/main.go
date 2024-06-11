package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Formatting
	// a, b := 12, 345
	// c, d := 1.2, 3.45
	// fmt.Printf("a: %d, b: %d\n", a, b)
	// fmt.Printf("a: %x, b: %X\n", a, b)
	// fmt.Printf("a: %#x, b: %#X\n", a, b)
	// fmt.Printf("c: %f, d: %.2f\n", c, d)
	// fmt.Println()

	// fmt.Printf("|%6d|%6d|\n", a, b)
	// fmt.Printf("|%06d|%06d|\n", a, b)
	// fmt.Printf("|%-6d|%-6d|\n", a, b)
	// fmt.Printf("|%6.2f|%6f|\n", c, d)

	// fmt.Println()

	// s := []int{1, 2, 3}
	// arr := [3]rune{'a', 'b', 'c'}
	// fmt.Printf("%T\n", s)
	// fmt.Printf("%v\n", s)
	// fmt.Printf("%#v\n", s)
	// fmt.Println()

	// fmt.Printf("%T\n", arr)
	// fmt.Printf("%q\n", arr)
	// fmt.Printf("%v\n", arr)
	// fmt.Printf("%#v\n", arr)
	// fmt.Println()

	// m := map[string]int{
	// 	"one": 1,
	// 	"two": 2,
	// }
	// fmt.Printf("%T\n", m)
	// fmt.Printf("%v\n", m)
	// fmt.Printf("%#v\n", m)

	// st := "a string"
	// bt := []byte(st)
	// fmt.Printf("%T\n", st)
	// fmt.Printf("%q\n", st)
	// fmt.Printf("%v\n", st)
	// fmt.Printf("%#v\n", st)
	// fmt.Println()

	// fmt.Printf("%T\n", bt)
	// fmt.Printf("%q\n", bt)
	// fmt.Printf("%v\n", bt)
	// fmt.Printf("%#v\n", bt)
	// File I/O

	for _, fname := range os.Args[1:] {
		var lc, wc, cc int
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file:", err)
			continue
		}
		// if _, err := io.Copy(os.Stdout, file); err != nil {
		// 	fmt.Fprintln(os.Stderr, err)
		// 	continue
		// }
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			line := scan.Text()
			lc++
			wc += len(strings.Fields(line))
			cc += len(line)
		}
		fmt.Printf(" %7d %7d %7d %s\n", lc, wc, cc, fname)
		file.Close()
	}

}
