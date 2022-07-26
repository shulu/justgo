package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) testPerson() Person {
	p.Age = p.Age + 1
	fmt.Println(p)
	return p
}

func (p *Person) testPerson2() {
	p.Age = p.Age + 1
	fmt.Println(p)
}

func main() {

	var p1 = Person{"LISA", 20}
	fmt.Println(p1)
	p1.testPerson()
	fmt.Println(p1)
	var p2 = &Person{"tom", 21}
	fmt.Println(p2)
	p2.testPerson2()
	fmt.Println(p2)

}
