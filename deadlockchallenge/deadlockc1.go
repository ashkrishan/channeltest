package deadlockchallenge

import (
	"fmt"
)

func DeadlockChallenge1() {
	c := make(chan int)
	go func() {
		c <- 1
		close(c)
	}()
	fmt.Println(<-c)
}
