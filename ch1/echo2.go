package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// 	sep += arg + s
	// 	s = " "
	// }
	fmt.Println(strings.Join(os.Args[1:], " "))
	// fmt.Println(sep)

}
