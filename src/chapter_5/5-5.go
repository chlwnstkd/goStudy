// 5-5
package main

import (
	"fmt"
)

func SumAndDiff(a int, b int) (aum int, diff int) {
	sum = a + b
	diff = a - b
	return sum, diff
}

func main() {
	sum, diff := SumAndDiff(6, 2)
	fmt.Println(sum, diff)
}
