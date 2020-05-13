package main

import (
	"fmt"
	"math"
)

func main() {

    max := max(34,232)
	fmt.Println(max)
	mathFunc()
	printNum()
	ifFunc(100)
}

func max(a int, b int)int{
	if a>b {
		return a
	}
	return b
}

const s string = "constanat"

func mathFunc(){
	fmt.Println(s)
	const n = 500000000
	const d = 3e28 / n
	fmt.Println(d)
	// fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

func printNum(){
	for i := 0; i<10; i++{
		fmt.Printf("%d",i)
	}
}

func ifFunc(x int){
	if x>10{
	   fmt.Println("x is greater than 10")
	}else {
	   fmt.Println("x is less than 10")
	}
}
