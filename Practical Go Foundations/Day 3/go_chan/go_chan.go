package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		/* BUG: All go routines use the "i" for the loop
		go func() {
			fmt.Println(i)
		}()
		*/

		//Fix 1: use a parameter
		go func(n int) {
			fmt.Println(n)
		}(i)
	}

	time.Sleep(10 * time.Millisecond)
}
