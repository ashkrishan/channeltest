package examples

import (
	"fmt"
)

func MultiChannel() {
	c1 := incrementor("Foo: ")
	c2 := incrementor("Bar: ")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final values: ", <-c3+<-c4)

}

func incrementor(s string) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for v := range c {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}
