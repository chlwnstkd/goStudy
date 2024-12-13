// 8-2
package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	a, b int
}

func main() {
	var f float32 = 1.3
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)

	fmt.Println(t.Name())
	fmt.Println(t.Size())
	fmt.Println(t.Kind() == reflect.Float64)
	fmt.Println(t.Kind() == reflect.Int64)

	fmt.Println(v.Type())
	fmt.Println(v.Kind() == reflect.Float64)
	fmt.Println(v.Kind() == reflect.Int64)
	fmt.Println(v.Float())
}
