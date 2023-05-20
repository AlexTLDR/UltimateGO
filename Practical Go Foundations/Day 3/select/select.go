package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	// myfunc := func(duration int) {
	// 	time.Sleep(time.Duration(duration) * time.Millisecond)
	// 	ch1 <- 1
	// }

	// func(duration int) {
	// 	time.Sleep(time.Duration(duration) * time.Millisecond)
	// 	ch1 <- 1
	// }(10)

	// go myfunc(10000)

	// go myfunc(20000)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Millisecond)
		ch2 <- 2
	}()

	select {
	case val := <-ch1:
		fmt.Println("ch1", val)

	case val := <-ch2:
		fmt.Println("ch2", val)
	}

}
