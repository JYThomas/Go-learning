package main

import "fmt"

func main(){
	// if else
	flag := true
	if flag {
		fmt.Println("TRUE")
	// else不能换行 。。。
	}else{
		fmt.Println("FALSE")
	}

	// if elseif else
	A := 10
	if A > 100{
		fmt.Println(1)
	}else if A > 50{
		fmt.Println(2)
	}else{
		fmt.Println(3)
	}

	// switch case
	score := 78
	var grade string 
	switch{
	case score > 90: 
		grade = "A"
	case score > 80: 
		grade = "B"
	case score > 70: 
		grade = "C"
	default: 
		grade = "Nothing"
	}
	// 注意 这里定义了grade变量 尽管他被赋值了 但是还不是被使用 grade declared and not used 需要输出
	fmt.Println(grade)

	// for
	sum := 0
	for i:=1;i<10;i++{
		sum += i
	}
	fmt.Println(sum)

	//for 枚举
	arr := []string{"google", "firefox"}  // 定义字符数组
	for index,value := range arr{
		fmt.Println(index,value)
		// 0 google
		// 1 firefox
	}
}