// 6-9
package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello ~")
}

type Student struct {
	P      Person
	school string
	grade  int
}

func main() {

	var a Student

	s.p.greeting()
}
