package main
import "fmt"

func main(){
	// 初始化切片 []不加大小定义
	slice_int := []int{1,2,3,4,5}
	fmt.Println(slice_int[:])  // [1 2 3 4 5]
	fmt.Println(slice_int[:2])  // [1 2]

	// 默认初始化一个大小为3的切片 数组元素全为0
	myslice := make([]int, 3)
	fmt.Println(myslice)  // [0 0 0]


	// 默认初始化一个大小为3,容量为5的切片 数组元素全为0
	myslice2 := make([]int, 3, 5)
	fmt.Println(myslice2, len(myslice2), cap(myslice2))


	// 定义一个空切片
	var nilslice []int
	// 判断是否为空
	if nilslice == nil{
		fmt.Println(nilslice)
	}


	// ======================= 切片操作 ========================
	// (1)新增元素
	myslice3 := make([]int, 3)
	myslice3[0] = 1
	myslice3[1] = 2
	myslice3[2] = 3
	// 向切片中新增数组元素
	myslice3 = append(myslice3, 4)
	fmt.Println(myslice3, len(myslice3), cap(myslice3))  // [1 2 3 4] 4 5
	myslice3 = append(myslice3, 5)
	fmt.Println(myslice3, len(myslice3), cap(myslice3))  // [1 2 3 4 5] 5 5
	/*
		尝试向切片中添加第 6 个元素时，切片的长度超过了其初始容量，Go 自动分配了一个更大的底层数组（新的容量是原容量的两倍）
	*/
	myslice3 = append(myslice3, 6)
	fmt.Println(myslice3, len(myslice3), cap(myslice3))  // [1 2 3 4 5 6] 6 10

	// 同时追加多个元素
	myslice3 = append(myslice3, 7,8,9)
	fmt.Println(myslice3, len(myslice3), cap(myslice3))  // [1 2 3 4 5 6 7 8 9] 9 12

	// (2)修改元素
	myslice3[0] = 100
	fmt.Println(myslice3, len(myslice3), cap(myslice3))  // [100 2 3 4 5 6 7 8 9] 9 12

	// (3)删除元素()原理是使用slice切片原理[]
	// 删除第一个元素
	myslice3 = myslice3[1:]
	fmt.Println(myslice3, len(myslice3), cap(myslice3))  // [2 3 4 5 6 7 8 9] 9 12

	// 删除下标为4的元素
	index := 4
	myslice3 = append(myslice3[:index], myslice3[index+1:]...) // 注意后面要加上 ...
	fmt.Println(myslice3, len(myslice3), cap(myslice3))  // [2 3 4 5 7 8 9] 7 11

	// 删除多个元素
	start := 2
    end := 5
	myslice3 = append(myslice3[:start], myslice3[end:]...) // 注意后面要加上 ...
	fmt.Println(myslice3, len(myslice3), cap(myslice3))  // [2 3 8 9] 4 11

	// (5)切片复制 直接赋值 操作的是同一块内存地址
	myslice4 := myslice3
	fmt.Println(myslice3)
	fmt.Println(myslice4)
	fmt.Println(myslice3[0])  // 2
	fmt.Println(myslice4[0])  // 2
	fmt.Println(&myslice3[0])  // 0xc000094068
	fmt.Println(&myslice4[0])  // 0xc000094068

	myslice3[0] = 82
	fmt.Println(myslice3[0])  // 82
	fmt.Println(myslice4[0])  // 82
	fmt.Println(&myslice3[0])  // 0xc000094068
	fmt.Println(&myslice4[0])  // 0xc000094068

	// 创建一个空切片
	myslice5 := make([]int, 5)
	copy(myslice5, myslice4)
	fmt.Println(myslice4[0])  // 82
	fmt.Println(myslice5[0])  // 82
	myslice4[0] = 821
	fmt.Println(myslice4[0])  // 821
	fmt.Println(myslice5[0])  // 82
	fmt.Println(&myslice4[0])  // 0xc000094068
	fmt.Println(&myslice5[0])  // 0xc0000ae090

	/*
		总结 在切片复制的时候 直接使用=复制 操作的是同一块内存地址，
		而使用copy函数 则是新开辟一段内存地址 从首个元素起始指针位置可以看出
	*/

}