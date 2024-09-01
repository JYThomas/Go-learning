package main

import "fmt"

var x, y int = 10, 20

// 这种声明用于定义全局变量， 统一管理
var (
	a int 
	b bool
)

// 无声明方式定义变量不能在函数外部使用
// name := "Thomas"  // syntax error: non-declaration statement outside function body

func main(){
	name := "Thomas"
	// fmt.Println(x,y,a,b,name) // 10 20 0 false Thomas

	// 局部变量 声明了就要用 不管有没有赋值
	var intvar int

	// 全局变量允许声明而不使用
	fmt.Println(a,b,name)
}