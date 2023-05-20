package main

import (
	"context"
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
		time.Sleep(20 * time.Millisecond)
		ch2 <- 2
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	// select {} is like for but without consuming CPU

	select {
	case val := <-ch1:
		fmt.Println("ch1", val)

	case val := <-ch2:
		fmt.Println("ch2", val)
	//case <-time.After(5 * time.Millisecond):
	case <-ctx.Done():
		fmt.Println("timeout")
	}

}
