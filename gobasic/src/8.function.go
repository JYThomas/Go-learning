package main
import (
	"fmt"
)

func main(){
	a := 10
	b := 20
	result := max(a,b)
	fmt.Println(result)

	intvar := 10
	strvar := "Hello"
	Res_intvar, Res_strvar := swap(intvar, strvar)
	fmt.Println(Res_intvar, Res_strvar)
	
	res_a, res_b := add_one(&a,&b)
	fmt.Println(res_a, res_b)
}

// 返回单个值
// 函数声明func; 函数名max; 参数列表a,b int为参数变量类型; 外部int为返回值类型
func max(a,b int) int {
	if a > b{
		return a 
	}else{
		return b
	}
}

// 函数返回多个值
func swap(intvar int, strvar string) (string, int)  {
	/*
		intvar = "this"
		strvar = 100
		注意 这种swap方法不能使用 原先的数据类型是什么就得是什么数据类型进行赋值覆盖 否则会报错
		cannot use strvar (variable of type string) as int value in assignment
		cannot use intvar (variable of type int) as string value in assignment
	*/
	return strvar, intvar
}

// 值传递与引用传递(值传递不改变原先变量值 引用传递会改变原先变量值(传递的是地址))
// 形参类型前加* 返回变量也带* 传递实参时加上取地址符号&
func add_one(a,b *int) (int, int) {
	*a += 1
	*b += 1
	return *a,*b
}
