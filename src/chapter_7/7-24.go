// 7-24
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	var data int32 = 0
	wg := new(sync,WaitGroup)
	
	for i :=; i < 2000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, 1)
		}()
	}
	
	for i:=0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, -1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(data)
}
