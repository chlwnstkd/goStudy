// 7-18
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var data = []int()
	var rwMutex = new(sync.RWMutex)

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock()
			data += 1
			fmt.Println("write : ", data)
			time.Sleep(10 * time.Millisecond)
			rwMutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock()
			fmt.Println("read 1 : ", data)
			time.Sleep(10 * time.Second)
			rwMutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock()
			data += 1
			fmt.Println("read 2 : ", data)
			time.Sleep(10 * time.Second)
			rwMutex.Unlock()
		}
	}()

	time.Sleep(10 * time.Second)
}
