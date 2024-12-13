// 9-8
package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{10, 5, 3, 7, 6}
	b := []float64(4.2, 7.6, 5.5, 1.3, 9.9)
	c := []string("Maria", "Andrew", "John")

	sort.Ints(a)
	fmt.Println(a)

	sort.Float64s(b)
	fmt.Println(b)

	sort.Strings(c)
	fmt.Println(c)
}
