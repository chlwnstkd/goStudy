// 4-9
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}

	a = append(a, 4, 5, 6)
	fmt.Println(a)
}
