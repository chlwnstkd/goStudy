// 5-17
package main

import (
	"fmt"
)

func HelloWorld() {
	defer func() {
		fmt.Println("world")
	}()
	func() {
		fmt.Println("Hello")
	}
}

func main() {
	HelloWorld()
}
