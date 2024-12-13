// 6-11
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

func (p *Student) greeting() {
	fmt.Println("Hello Student~")
}

func main() {

	var a Student

	s.Person.greeting()

	s.greeting()
}
