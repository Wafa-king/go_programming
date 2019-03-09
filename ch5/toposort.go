package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"a": {"1"},
	"b": {"2"},
	"c": {
		"3",
		"4",
		"5",
	},
	"d": {"6"},
	"e": {"7"},
	"f": {"8"},
	"g": {"9"},
	"h": {"10"},
	"i": {"11", "12"},
	"j": {"11", "12"},
}

func main() {
	for i, cource := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, cource)
	}
}

func topoSort(m map[string][]string) string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
