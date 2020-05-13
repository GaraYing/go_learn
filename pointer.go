package main

import "fmt"
// import "math"
// 指针

func main() {
	// 分配了一个整型数组，长度为10，容量为100，并返回前10个数组的切片
	v := make([]int, 10,100)
	fmt.Println(v)
	fmt.Println(v[3])
	
	m := make(map[int]string)
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"
	fmt.Println(m)
	fmt.Println(m[3])
	
	delete(m, 2)
	fmt.Println(m)
	
	 for key, val := range m{
        fmt.Printf("%d => %s \n", key, val)
        /*输出：(顺序在运行时可能不一样)
            three => 3
            one => 1
            two => 2*/
    }
}