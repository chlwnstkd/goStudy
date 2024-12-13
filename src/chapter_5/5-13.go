// 5-13
package main

import (
	"fmt"
)

func main() {

	sum := func(a, b int) int {
		return a + b
	}

	r := sum(1, 2)
	fmt.Println(r)
}
