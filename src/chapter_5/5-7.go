// 5-7
package main

import (
	"fmt"
)

func sum(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}

func main() {
	n := []int{1, 2, 3, 4, 5}
	r := sum(n...)
	fmt.Println(r)
}
