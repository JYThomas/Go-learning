package main 

import "fmt"

func main(){
	// map用于存放键值对 key value 类似于字典 第一个string是key类型，另一个string 是值的类型 
	// go中可定义除默认类型外自定义的聚合类型、重命名类型 所以这可自定义
	/*
		其中 KeyType 是键的类型，ValueType 是值的类型，initialCapacity 是可选的参数，用于指定 Map 的初始容量。
		Map 的容量是指 Map 中可以保存的键值对的数量，当 Map 中的键值对数量达到容量时，Map 会自动扩容。
		如果不指定 initialCapacity，Go 语言会根据实际情况选择一个合适的值。
	*/
	person := make(map[string]string)
	person["name"] = "Thomas"
	person["age"] = "24"
	fmt.Println(person)

	// 直接定义一个集合(字面量)
	prices := map[string]int{
		"Apple": 10,
		"Banana": 3,
		"dragonfruit": 6,
	}

	// 集合curd 
	// 查询
	apple_price := prices["Apple"]
	fmt.Println(apple_price)

	// 新增
	prices["pear"] = 4
	pear_price := prices["pear"]
	fmt.Println(pear_price)

	// 修改
	prices["pear"] = 11
	fmt.Println(prices["pear"])

	// 删除
	delete(prices, "pear")

	// 迭代遍历
	for k, v := range prices{
		fmt.Printf("key=%s, value=%d\n", k, v)
	}
}

// go run /home/withboom/Go-learning/gobasic/src/11.map.go