package main

import (
	"fmt"
	"unicode"
)

func main() {
	for _, input := range []string{
		"the quick brown fox jumps over the lazy dog",
} {
		if isPangram(input) {
			fmt.Println("pangram")
		} else {
			fmt.Println("not pangram")
		}
	}
}

func isPangram(s string) bool {
	mymap := mymap()
	for _, r := range []rune(s) {
		if unicode.IsLetter(r) {
			t := unicode.ToLower(r)
			if _, ok := mymap[t]; ok {
				mymap[t] = true
			} else {
				return false
			}
		}
	}
	for _, v := range mymap {
		if !v {
			return false
		}
	}
	return true
}

func mymap() map[rune]bool {
	m := make(map[rune]bool, 28)
	for i := 'a'; i <= 'z'; i++ {
		m[i] = false
	}
	return m
}
