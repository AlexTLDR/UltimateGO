package main

import (
	"fmt"
	"unicode/utf8"
)

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
