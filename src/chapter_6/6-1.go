// 6-1
package main

import (
	"fmt"
)

func main() {

	var numPrt *int = new(int)

	*numPrt = 1

	fmt.Println(*numPrt)
}
