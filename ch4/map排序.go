package main

import (
	"fmt"
	"sort"
)

func main() {
	var names []string
	ages := map[string]int{
		"alice":   32,
		"charlie": 18,
	}
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
