package main

import "fmt"

func main() {
	s := "G☺"
	banner(s, 6)
	for i, v := range s {
		fmt.Println(i, "=", v)
	}

	// byte (uint8)
	// rune (int32)

	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)
	u := s[0]
	fmt.Printf("%c of type %T\n", u, u)
}

func banner(text string, width int) {
	padding := (width - len(text)) / 2
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
