package main

import "fmt"
// import "math"
// 闭包

func main() {
	nextNumFunc := nextNum()
    for i:= 0; i<100; i++ {
      fmt.Println(nextNumFunc())
    }
}

func nextNum() func() int{
	i,j := 1,1
	return func() int{
		var temp = i+j
		i,j = j,temp
		return temp
	}
}