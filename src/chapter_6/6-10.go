// 6-10
package main

import (
	"fmt"
)

type Student struct {
	Person
	school string
	grade  int
}

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello ~")
}

func main() {

	var a Student

	s.Person.greeting()

	s.greeting()
}
