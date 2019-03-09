package main

import (
	"fmt"
)

func main() {
	fmt.Println(HasPrefix("abcdefg", "abc"))
	fmt.Println(Contains("abcdefg", "bc"))
}

// HasPrefix 是否以prefix开头
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// Contains 是否包含substr
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
