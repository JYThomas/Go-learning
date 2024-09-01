package main

import "fmt"

const PI float64 = 3.14  // 声明全局常量

const (
	myINT int = 10
	myFLOAT = 6.5  // 变量类型可以省略
	// SAY ：= "Hello"  不能在外部使用简短声明
)

func main(){
	const pi float64 = 3.1415926  // 声明局部常量

	fmt.Println(pi)
	fmt.Println(PI)

	fmt.Println(myINT)
	fmt.Println(myFLOAT)


}