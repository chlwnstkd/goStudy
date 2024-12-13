// 7-7
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("고루틴 : ", i)
			time.Sleep(i * time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("메인함수 : ", i)
	}
}
