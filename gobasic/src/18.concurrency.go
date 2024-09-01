package main

import (
	"fmt"
	"time"
)

// golang 并发处理 goroutine

// 定义一个函数，用于计算一个数字的平方，并将结果发送到通道
// 要平方计算的参数， 计算结果要放到名字为resultChan的通道chan中 这个chan存储的结果类型为int类型
func square(number int, resultChan chan int){
	/*
		在 Go 语言中，time.Second 的类型是 time.Duration，而 number 是一个 int 类型。要将 number 乘以 time.Second，你需要先将 number 转换为 time.Duration 类型。
	*/
	time.Sleep(time.Duration(number) * time.Second)
	result := number * number
	resultChan <- result // 将结果发送到通道中
}

func main(){
	// 创建一个通道 用于接受结果
	reschan := make(chan int)

	// 启动多个goroutine 计算结果的平方 计算函数中 计算越大的数假设需要的时间越久 但是因为是同时并发开始的 所以不会是单线程的耗时
	target := []int{1,2,3,4,5}
	
	for _, number := range target{
		// go关键字启动goroutine
		go square(number, reschan)
	}

	for range target{
		// 从通道中按照target的顺序 获取计算结果
		result := <-reschan
		fmt.Println(result)
	}
}

/*
	按照平方函数的计算 如果不设置go并发 那么程序的总耗时将是单线程的耗时相加 1+2+3+4+5=15秒 但是程序总共执行时间为5秒
	所以程序是并发执行的 五次计算是相同时间开始 那么程序的总耗时将是最长时间的那个5秒的耗时 所以程序的总耗时为5秒
*/