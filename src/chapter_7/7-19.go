// 7-19
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var rwMutex = new(sync.RWMutex)

	c := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(n int) {
			rwMutex.Lock()
			c <- true
			fmt.Println("wait begin : ", n)
			cond.Wait()
			fmt.Println("wait end : ", n)
			rwMutex.Unlock()
		}(i)
	}
	for i := 0; i < 3; i++ {
		<-c
	}
	for i := 0; i < 3; i++ {
		mutex.Lock()
		fmt.Println("signal : ", i)
		cond.Signal()
		mutex.Unlock()
	}
	fmt.Scanln()
}
