package main

import "fmt"

func test1() {
	fmt.Println("test1")
	a := &Person{"jacko", 20}
	b := Person{"john", 21} // 推荐这两种形式来创建实例
	a.SayHello()
	// Hello, my name is jacko.
	b.SayHello()
	// Hello, my name is john.
	a.ChangeName_("jack")
	a.SayHello()
	// Hello, my name is jacko.
	b = *a // 只有在赋值的时候才会有区别
	b.SayHello()
	// Hello, my name is jacko.
	a.ChangeName("jack")
	a.SayHello()
	// Hello, my name is jack.
	b.SayHello()
	// Hello, my name is jacko.
}

func test2() {
	fmt.Println("test2")
	a := Test{Person{"jacko", 20}, 1, "hello"}
	b := Test{Person{"john", 21}, 2, "world"}
	a.p.SayHello()
	// Hello, my name is jacko.
	b.p.SayHello()
	// Hello, my name is john.
}

func test3() {
	fmt.Println("test3")
	arr := []Person{{"jacko", 20}, {"john", 21}}
	for _, p := range arr {
		p.SayHello()
	}
	// Hello, my name is jacko.
	// Hello, my name is john.
	for _, p := range arr {
		p.ChangeName("jj")
	}
	for _, p := range arr {
		p.SayHello()
	}
	// Hello, my name is jacko.
	// Hello, my name is john.
}

func test4() {
	fmt.Println("test4")
	arr := []*Person{&Person{"jacko", 20}, &Person{"john", 21}}
	for _, p := range arr {
		p.SayHello()
	}
	// Hello, my name is jacko.
	// Hello, my name is john.
	for _, p := range arr {
		p.ChangeName("jj")
	}
	for _, p := range arr {
		p.SayHello()
	}
	// Hello, my name is jj.
	// Hello, my name is jj.
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
