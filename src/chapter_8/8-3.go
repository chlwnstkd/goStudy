// 8-3
package main

import (
	"fmt"
	"reflect"
)

func h(args []reflect.Value) []reflect.Value {
	fmt.Println("Hello World!")
	return nil
}

func main() {
	var hello func()

	fn := reflect.ValueOf(&hello).Elem()

	fn.Set(v)

	hello()
}
