// 6-6-1
package main

import (
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

func main() {
	rect := &Rectangle(20, 10)

	fmt.Println(rect)
}
