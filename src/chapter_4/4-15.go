// 4-15
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(len(a), cap(a))

	a = append(a, 6, 7)
	fmt.Println(len(a), cap(a))
}
