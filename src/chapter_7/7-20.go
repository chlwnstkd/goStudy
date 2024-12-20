// 7-20
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex)

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
	mutex.Lock()
	fmt.Println("broadcast")
	cond.Broadcast()
	mutex.Unlock()

}
