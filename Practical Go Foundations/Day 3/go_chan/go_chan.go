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

		/*Fix 1: use a parameter
		go func(n int) {
			fmt.Println(n)
		}(i)
		*/

		//Fix 2: use a loop body variable
		i := i // "i" shadows "i" from the loop
		go func() {
			fmt.Println(i)
		}()

	}

	time.Sleep(10 * time.Millisecond)
	/*Channel semantics
	  - send & receive will block until opposite operation(*)
	  - receive from a closed channel will return zero value without blocking
	*/
	ch := make(chan string)
	go func() {
		ch <- "hi" // send
	}()

	msg := <-ch // receive
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("got:", msg)
	}

	msg = <-ch // ch is closed
	fmt.Printf("closed: %#v\n", msg)

	msg, ok := <-ch // ch is closed
	fmt.Printf("closed: %#v (ok=%v)\n", msg, ok)
}
