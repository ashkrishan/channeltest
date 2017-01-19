package deadlockchallenge

import (
	"fmt"
)

func FactorialChallenge() {
	// f := factorial(4)
	// fmt.Println("Total:", f)
	c1 := factorial(4)
	c2 := factorial(6)
	c3 := factorial(40)
	fmt.Println(<-puller(c1))
	fmt.Println(<-puller(c2))
	fmt.Println(<-puller(c3))
}

func factorial(n int64) chan int64 {
	out := make(chan int64)
	go func() {
		for i := n; i > 0; i-- {
			out <- n
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int64) <-chan int64 {
	out := make(chan int64)
	var total int64 = 1
	go func() {
		for v := range c {
			total *= v
		}
		out <- total
		close(out)
	}()
	return out
}

// func factorial(n int) int {
// 	total := 1
// 	for i := n; i > 0; i-- {
// 		total *= i
// 	}
// 	return total
// }

/*
CHALLENGE #1:
-- Use goroutines and channels to calculate factorial

CHALLENGE #2:
-- Why might you want to use goroutines and channels to calculate factorial?
Ans: When concurrncy is required in the program for optimisation goroutines can be used as factorial
---- Formulate your own answer, then post that answer to this discussion area: https://goo.gl/flGsyX
---- Read a few of the other answers at the discussion area to see the reasons of others
*/
