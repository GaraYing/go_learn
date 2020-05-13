package main

import "fmt"
// import "math"
// 指针


type Person struct{
	name string
	age int
}

func main() {
    // 初始化 
	person := Person{"Json", 23}
	// person := Person{name:"Json", age:23}
	fmt.Println(person)
	person.name = "Jack"
	fmt.Println(person)
	pPerson := &person
	pPerson.age = 25
	fmt.Println(person)
	for i := 0; i < 5; i++ {
		// 被延期的函数以后进先出（LIFO）的顺行执行，因此将打印4 3 2 1 0
		defer fmt.Printf("%d ", i)
		fmt.Printf("%d ", i)
	}
}
