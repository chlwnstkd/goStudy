// 7-14
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)
	go func() {
		for {
			i := <-c1
			c1 <- 10
			fmt.Println("c1 : ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		for {
			c2 <- "Hello, world"
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		for{
			select {
			case c1 <- 10
			case s := <-c2:
				fmt.Println("c2 : ", s)
		}
	}()

	time.Sleep(10 * time.Second)
}
