// 4-14
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := make([]int, 3)

	copy(b, a)
	b[0] = 9

	fmt.Println(a)
	fmt.Println(b)
}
