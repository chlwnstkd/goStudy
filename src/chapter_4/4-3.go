// 4-3
package main

import (
	"fmt"
)

func main() {
	a := [5]int{32, 29, 78, 16, 8}

	for i, value := range a {
		fmt.Println(i, value)
	}
}
