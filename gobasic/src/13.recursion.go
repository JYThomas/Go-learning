package main 
import "fmt"

func factorial(n uint64)(result uint64){
	if(n > 0){
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int)(result int){
	if(n<2){
		return n 
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main(){
	target := 15
	// 计算15的阶乘 15*14*13*.....
	fmt.Println(factorial(uint64(target)))

	// 计算15以内的非伯纳切数列
	var k int
	for k=0; k<15; k++ {
		fmt.Println(fibonacci(k))
	}
}