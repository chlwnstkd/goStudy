// 7-1
package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello World!")
}

func main() {
	go hello()

	fmt.Scanln()
}
