// 5-14
package main

import (
	"fmt"
)

func main() {

	a, b := 3, 5

	f := func(x int) int {
		return a*x + b
	}

	y := f(5)

	fmt.Println(y)
}
