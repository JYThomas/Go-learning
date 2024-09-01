package main
import (
	"fmt"
)

func main(){
	// 初始化数组 []数组大小 数组类型 {数组值}
	arr_int := [5]int{1, 2, 3, 4, 5}

	// 当不知道数组大小时 用[...]替代
	arr_float := [...]int{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}

	fmt.Println(arr_int)
	fmt.Println(arr_float)

	// 访问数组元素
	for i := 0; i<len(arr_int); i++{
		// 数组名+下标
		fmt.Println(arr_int[i])
	}
	// 注意 Go的数组是不可动态变化的 要想实现动态变化的列表 使用切片slice
}