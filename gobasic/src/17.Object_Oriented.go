package main

import "fmt"

// 面向对象思想在golang中的实现

// 定义一个接口
type Animal interface {
	// 方法名称(接受参数类型) 方法返回值 
	Speak(string) string 
}

// 定义一个基础动物结构体
type AnimalBase struct {
	Name string
}

// 定义这个基础结构体的方法
// 方法关键字func (方法接受参数类型) 方法名() 方法返回值
func (a AnimalBase) Eat() string {
	return a.Name + " is Eating"
}

// 定义一个实例结构体 使用嵌入的方式实现继承思想
type Dog struct{
	AnimalBase  // 嵌入抽象结构体
	Attr_1 string // 自己特定的结构体属性
}

// 结构体Dog的方法
func (d Dog) Speak(sounds string) string{
	return d.Name + " says " + sounds
}

// 定义Cat结构体 同样嵌入基类(基结构体 继承基类的属性与方法)
type Cat struct {
	AnimalBase
	Attr_2 string 
}

// 结构体Cat的方法 实现接口
func (c Cat) Speak(sounds string) string {
	return c.Name + " says " + sounds
}

// 接口多态性处理"函数"
func MakeSpeak(a Animal, sounds string){
	fmt.Println(a.Speak(sounds))
}


func main(){
	// 初始化结构体
	dog := Dog{
		// 初始化实例属性 其中嵌入属性需要用类型声明的方式确定value, 在这里可以将AnimalBase{Name: "dd"} 看成是一个自定义的类型(复杂类型)
		AnimalBase: AnimalBase{Name: "dd"},
		Attr_1: "yellow",
	}

	cat := Cat{
		AnimalBase: AnimalBase{Name: "cc"},
		Attr_2: "white",
	}

	// 实例继承基类
	fmt.Println(dog.Eat())
	fmt.Println(cat.Eat())

	// 实例实现接口
	MakeSpeak(dog, "woof")
	MakeSpeak(cat, "meow")
}