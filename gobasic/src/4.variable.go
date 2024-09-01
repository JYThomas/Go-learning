package main

import "fmt"

func main(){
	// 变量定义
	var name string = "Thomas"
	var age int = 24
	// 注意： 上面定义的变量 定义了一定要用 Go语言里面定义了变量而不使用是一个报错 而不是一个警告，强规则校验
	fmt.Println(name, age)

	// 同时定义多个相同类型的变量
	var myname, hername = "Thomas", "Boom"
	fmt.Println(myname, hername)

	// 根据赋值情况定义变量类型
	var thisname = "ThomasJY"
	fmt.Println(thisname)

	// =========================================
	// 声明变量但不赋值 默认值的情况
	var intvar int  // 0
	var float64var float64 // 0.0
	var strvar string // ""
	var boolvar bool // false
	
	fmt.Println(intvar,float64var,strvar,boolvar)

	// 快速声明变量(系统推断类型) 
	//这种不带声明格式的只能在函数体中出现 
	intvar2 := 1
	/*
		相当于先声明 然后再赋值 一步完成
		var intVal int 
		intVal = 1 
	*/
	fmt.Println(intvar2)

	/*
		var intVal int 
		intVal :=1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明(这里声明+赋值 两次声明就报错)
	*/
	

}