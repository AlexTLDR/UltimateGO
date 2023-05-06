package main

import (
	"fmt"
	"unicode/utf8"
)

// UTF for copy paste -> ☺ <-

func main() {
	s := "G☺"
	banner(s, 6)
	for i, v := range s {
		fmt.Println(i, "=", v)
	}

	// byte (uint8)
	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)
	// rune (int32)
	u := s[0]
	fmt.Printf("%c of type %T\n", u, u)

	x, y := 1, "1"
	fmt.Printf("x=%v, y=%v\n", x, y)
	fmt.Printf("x=%#v, y=%#v\n", x, y) // Use %#v in debug/log

	fmt.Printf("%20s!\n", s)
	fmt.Println(isPalindrome(s))
}

func banner(text string, width int) {
	// padding := (width - len(text)) / 2 BUG: len is in bytes
	padding := (width - utf8.RuneCountInString(text)) / 2
	fmt.Println(len(text))
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func isPalindrome(s string) bool {
	// if len(s) == 1 {
	// 	return true
	// }
	for i := 0; i < len(s); i++ {
		if s[i] == s[len(s)-1-i] {
			return true
		}
	}
	return false
}
