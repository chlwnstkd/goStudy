// 5-12
package main

import (
	"fmt"
)

func main() {
	
	func() {
		fmt.Println("Hello World!")
	}()
	
	func(s string) {
		fmt.Println(s)
	} {"Hellp World!"}
	
	r : func(a int, b int) int {
		return a + b
	} (1, 2)
	
	fmt.Println(r)
}
