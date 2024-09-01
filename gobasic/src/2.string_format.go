package main

import (
	"fmt"
)

func main(){
	var name = "Thomas"
	var age = 24
	/*格式化字符串输出 换行输出*/
	fmt.Println(fmt.Sprintf("My name is %s and I'm %d years old", name, age))
	// %d 表示整型数字，%s 表示字符串 简化版本 这里的占位符类似于C
	fmt.Printf("My name is %s and I'm %d years old", name, age)
}