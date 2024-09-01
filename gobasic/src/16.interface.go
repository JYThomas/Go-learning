package main

import "fmt"

// golang接口 结构体

/*
	// 首先在自定义数据结构中需要明确的概念 
	// 在接口、结构体的概念中， 需要将函数的概念转换为方法的概念 
	// 方法是是附加到某个类型上的函数，语法与函数类似，但需要一个特殊的接收者参数。
	// 接收者可以是结构体类型或任何其他自定义类型。通过这种方式，方法可以访问接收者的字段和其他方法。
*/

// 定义接口, 接口中声明了一个函数
// type 接口名称 interface {}
type Animal interface {
	Speak() string
}

// 定义一个结构体 
type Cat struct {}
// type 结构体每次 struct{}
// 定义一个与上述结构体绑定的函数(方法)
func (c Cat) Speak() string{
	return "Meow"
}

// 定义一个结构体 
type Dog struct {}
// 定义一个与上述结构体绑定的函数(方法)
func (d Dog) Speak() string{
	return "Woof"
}

// type Pig struct {}
// func (p Pig) Speak2() string{
// 	return "ppp"
// }

// 定义一个函数 接受实现了Animal接口类型的参数（这个说法不太准确）
// 接收一个可抽象为Animal接口的结构体 这个结构体实现了这个Animal接口的方法(准确一点的理解)
func MakeAnimalSpeak(a Animal){
	fmt.Println(a.Speak())
}


/*
	接口类型的用法：当多个结构体都涉及同样的功能，但具体的实现又不一样，这时候为了统一规范大家相同的部分，可以将这个功能抽象成接口，
	通过这个抽象的接口，来统一管理这个相同的功能，通过传递的结构体参数不同，实现了类似功能不同结构体(对象)的多种形态(多态性)
*/

func main(){
	CAT := Cat{}
	DOG := Dog{}
	// PIG := Pig{}

	MakeAnimalSpeak(CAT)
	MakeAnimalSpeak(DOG)
	// MakeAnimalSpeak(PIG)
}