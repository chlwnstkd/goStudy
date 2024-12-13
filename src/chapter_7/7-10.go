// 7-10
package main

import (
	"fmt"
)

func producer(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	c <- 100
}

func consumer(c chan<- int) {
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(i)
}

func main() {
	c := make(chan int)

	go producer(i)
	go consumer(i)

	fmt.Scanln()
}
