package main

import (
	"fmt"
	"errors"
)

func main() {
	max := max(34,232)
	fmt.Println(max)
	printNum()
	fmt.Println()
	err := errors.New("this is a error!")
	if err != nil{
		fmt.Print(err)
	}
}

func max(a int, b int)int{
	if a>b {
		return a
	}
	return b
}

const s string = "constanat"
func printNum(){
	for i := 0; i<10; i++{
		fmt.Printf("%d",i)
	}
}
