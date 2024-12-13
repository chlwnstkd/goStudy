// 3-11
package main

import (
	"fmt"
)

func main() {
	var a int = i

	if a == 1 {
		goto ERROR1
	}
	if a == 2 {
		goto ERROR2
	}
	if a == 3 {
		goto ERROR1
	}
	fmt.Println(a)
	fmt.Println("Success")
	return
ERROR1:
	fmt.Println("ERROR 1")
ERROR2:
	fmt.Println("ERROR 2")
	return

}
