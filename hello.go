package main

import "fmt"
// import "math"

func main() {
	// var e string = "golang"
	var e = "golang"
        fmt.Println(e)
	var a = 12
	if a<100 {
		fmt.Println(a)
	}
	// var flag bool = false
	flag := false
	fmt.Printf("flag=",flag)
	b := true
	fmt.Println(!b)
	var c int
	fmt.Println(c)
	var f string
	fmt.Printf("%t\n",f=="")
	fmt.Printf("%t\n",1==2)
	fmt.Println("hello world")
	_,value := 34,24
	fmt.Printf("_ = %d\n",value)
	m := `hello
	world`
	fmt.Printf("%s\n",m)
	arr := [...]int{1,2,3,44}
	fmt.Println(arr)
}
