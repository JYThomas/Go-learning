package main

import (
	"fmt"
)

// 定义一个形状接口 其中包含计算不同形状的面积的方法
type Shape interface {
    Area() float64
}

// 定义一个矩形对象
type Rectangle struct {
    Type string
    Heigth float64
    Width float64
}

// 矩形对象实现Shape接口中的面积求解 （注意 Area方法要和接口中的一致，包括参数）
func (r Rectangle) Area ()(ResultArea float64) {
    ResultArea = r.Heigth * r.Width
    return ResultArea
}

// 定义一个圆形对象
type Circle struct {
    Type string
    Radius float64
}

// 圆形实现Shape接口中的面积求解
func (c Circle) Area() (ResultArea float64) {
    ResultArea = 3.14 * c.Radius * c.Radius
	return ResultArea
}

func main(){
    // 初始化矩形对象
    rect := Rectangle{Heigth:10, Width:20, Type: "Rectangle"}
    cir := Circle{Radius: 10, Type:"Circle"}

	fmt.Printf("The Area of %s is %f\n", rect.Type, rect.Area())
	fmt.Printf("The Area of %s is %f\n", cir.Type, cir.Area())
    
    // 定义实现了接口的对象的结构体切片
    Shapes := []Shape{rect, cir}
    for _, shape := range Shapes {
        
		fmt.Printf("Area: %f\n", shape.Area())
    }
}