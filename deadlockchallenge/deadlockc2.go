package deadlockchallenge

import (
	"fmt"
)

func DeadlockChallenge2() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

// Why does this only print zero?
// And what can you do to get it to print all 0 - 9 numbers?
