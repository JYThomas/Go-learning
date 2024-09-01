package main
import (
	"fmt"
	"errors"
)

func divide(a, b float64)(float64, error){
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

func main(){
	result, err := divide(10, 2)
	
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}

	result, err = divide(10, 0)
	// golang的异常处理用返回值处理 用errors.New()创建异常对象 手动控制异常对象的生成 用于处理异常
	// err如果是异常对象的话 这时候异常对象不是nil 处理异常
	// nil用于表示指针、切片、映射、通道、函数和接口等类型的零值或未初始化的值 errors.New()不是零值 这时候触发err != nil为true 处理异常
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
}


