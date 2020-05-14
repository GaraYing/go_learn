package main

import "fmt"
// import "math"
// 指针


type Person struct{
	name string
	age int
}

type Skills []string

type Student struct{
	Person // 匿名字段，struct
	Skills // 匿名字段，自定义的类型string slice
	int    // 内置类型作为匿名字段
	speciality string
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

	mark := Student{Person:Person{"Mark",18},speciality:"Software Engineer"}
	fmt.Println("Mark's age is ", mark.age)
	mark.Person = Person{"Mark2",21}
	mark.Person.age -= 1
	fmt.Println("Mark's age is ", mark.Person.age)
	fmt.Println("Mark : ", mark)
	// 修改他的skill技能字段
	mark.Skills = []string{"pingpang"}
	fmt.Println("His skill are ", mark.Skills)
	mark.Skills = append(mark.Skills, "basketball","football")
	fmt.Println("Now, his skill are ", mark.Skills)
	// 修改匿名内置类型字段
	mark.int  = 120
	fmt.Println("His lucky number is ", mark.int)

	for i := 0; i < 5; i++ {
		// 被延期的函数以后进先出（LIFO）的顺行执行，因此将打印4 3 2 1 0
		defer fmt.Printf("%d ", i)
		fmt.Printf("%d ", i)
	}
}
