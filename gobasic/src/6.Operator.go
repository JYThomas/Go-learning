package main

import "fmt"

/*
	// 运算符：算术运算符(+-*% ++ --)、关系运算符(==、!=、>、<、>=、<=)、逻辑运算符(&& || !)、赋值运算符()、其他运算符(取地址&a、指针变量*a)
	// 运算符优先级 和常用编成语言类似
*/

func main(){
	var myname string = "Thomas"
	// %T占位符是用于输出变量的类型
	fmt.Printf("type of myname is %T\n", myname)
}