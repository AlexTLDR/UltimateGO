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
	  	- Buffered channel has cap(ch) non-blocking send operations
	  - receive from a closed channel will return zero value without blocking
	  - send to a closed channel will panic
	  - closing a closed channel will panic
	  - send/receive to a nil channel will block forever

	  https://www.353solutions.com/channel-semantics
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

	/* for/range does this
	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("got:", msg)
	}
	*/

	msg = <-ch // ch is closed
	fmt.Printf("closed: %#v\n", msg)

	msg, ok := <-ch // ch is closed
	fmt.Printf("closed: %#v (ok=%v)\n", msg, ok)

	// ch <- "hi" // ch is closed => panic

	values := []int{15, 8, 42, 16, 4, 23}
	fmt.Println(sleepSort(values))
}

/*
For every value "n" in values, spin a goroutine that will
- sleep "n" milliseconds
- send "n" over a channel

In the function body, collect values from the channel to a slice and return it.
*/
func sleepSort(values []int) []int {
	ch := make(chan int)
	for _, n := range values {
		n := n
		go func() {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}()
	}
	var out []int
	// for i := 0; i < len(values); i++ {

	// }
	for range values {
		n := <-ch
		out = append(out, n)
	}
	return out
}
