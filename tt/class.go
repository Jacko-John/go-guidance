package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.name)
}

func (p Person) ChangeName_(name string) {
	p.name = name
}

func (p *Person) ChangeName(name string) {
	p.name = name
}

type Test struct {
	p Person
	a int
	b string
}
