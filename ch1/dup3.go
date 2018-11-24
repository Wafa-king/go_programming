package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		fmt.Println(11, filename)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
		}
		fmt.Println(string(data))
		fmt.Println("-----------")
		fmt.Println(strings.Split(string(data), "\n"))
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		fmt.Println(counts)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
