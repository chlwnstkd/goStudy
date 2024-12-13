// 5-11
package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func diff (a int, b int) int {
	return a - b
}

func main() {
	
	f := map[string]func(int, int) int{
		"sum" : sum,
		"diff" : diff,
	}
	
	fmt.Println(f["sum"](1, 2)
	fmt.Println(f["diff"](1, 2)
}
