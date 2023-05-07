package main

import "fmt"

type Item struct {
	X int
	Y int
}

func main() {
	var i1 Item
	fmt.Println(i1)
	fmt.Printf("i1: %#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{
		Y: 10,
		//X:20,
	}
	fmt.Printf("i3: %#v\n", i3)
}
