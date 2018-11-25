package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "|", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	a := gcd(13, 5)
	fmt.Println(a)
	fmt.Println(fib(10))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
		fmt.Println(x, y)
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
