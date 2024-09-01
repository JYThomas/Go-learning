# Golang基础

## 一、变量声明与赋值

~~~golang
// 变量声明
var name string
name = "Thomas"

var name = "Thomas"

var name, age, gender = "Thomas", 13, "male"

// 短变量声明
name := "Thomas"

// 指针
name := "Thomas"
x := &name  // 取地址符 x为指针变量
fmt.Println(*x)  // 访问指针变量: Thomas
~~~

## 二、类型声明

作用在于提升代码的可读性与增加类型安全性

~~~golang
// 提升代码可读性
type Money float64

func Pay(amount Money) {
    fmt.Println(amount)
}

// 将浮点类型声明为Money类型，而不是浮点数，让读的人知道是代表钱
var salary Money = 1000.50
Pay(salary)
~~~

~~~go
// 变量类型安全性 声明了特定类型 在传参传错的时候会引发变量类型错误
// 比如要传递两个int类型的变量分别到两个函数中
type Hours int
type Age int

func SetHours(h Hours) {
    return h
}

func SetAge(a Age){
    return a
}

func main(){
    // 正确的传递函数
    var hours Hours = 10
    var age Age = 25
    
    SetHours(hours)
    SetAge(age)
    
    // 编译错误 类型不匹配
    SetHours(age)
    SetAge(hours)
}
~~~

## 三、变量作用域

变量在程序中的可见范围

1.包级作用域

~~~go
// 19.scope_1.go
package main
// 这个变量可以在使用main包环境下的任何地方使用
var GlobalVar = "This is a global variable"

// 19.scope_2.go
package main
import (
    "fmt"
)
func main() {
    fmt.Println(GlobalVar)
}

// run: go run 19.scope_1.go 19.scope_2.go 
// 同时执行两个文件 才能输出 单独执行第二个文件 没法输出 第一个文件没有编译到内存中 
// 虽然是声明写的同一个main包 但还是会读取不到
~~~

2.函数级作用域：变量在函数中声明，只能在函数中访问

3.代码块级作用域：代码块级作用域指的是变量在 `{}` 括号内的作用域。

4.闭包：可以捕获并记住其**外层函数的变量和环境** (最外层的函数封装，嵌套闭包也可以访问最外层)，并在其内部函数中使用它们。不必将这些信息存储在全局变量中。

~~~go
func Counter() func() int {
    count := 0
    return func (){
        // 函数内部的匿名函数(闭包)，可读取其外层函数的变量和环境
        count++
        return count
    }
}

func main(){
    c := Counter()
    fmt.Println(c()) // 1
    fmt.Println(c())  // 2
    fmt.Println(c())   // 3
}
~~~

## 四、基本数据类型(整数、浮点数、布尔值、字符串、常量)

~~~go
// 查看变量类型
package main

import (
	"fmt"
)

func main(){
	var a int = 10
	fmt.Printf("%T\n", a)  // int
}

~~~



~~~go
// int、int8、int16、int32、int64：表示有符号整数类型。
// uint、uint8、uint16、uint32、uint64：表示无符号整数类型。

package main

import "fmt"

func main() {
    var a int = 42         // 默认 int 类型，具体大小取决于系统
    var b int8 = -128      // 8 位有符号整数，范围：-128 到 127
    var c uint16 = 65535   // 16 位无符号整数，范围：0 到 65535

    fmt.Println("int:", a)
    fmt.Println("int8:", b)
    fmt.Println("uint16:", c)
}
~~~

~~~go
// float32：32 位浮点数，单精度。
// float64：64 位浮点数，双精度（Go 中的默认浮点数类型）。

package main

import "fmt"

func main() {
    var x float32 = 3.14          // 单精度浮点数
    var y float64 = 2.7182818284  // 双精度浮点数

    fmt.Println("float32:", x)
    fmt.Println("float64:", y)
}
~~~

~~~go
// 布尔类型

package main

import "fmt"

func main() {
    var isTrue bool = true
    var isFalse bool = false

    fmt.Println("isTrue:", isTrue)
    fmt.Println("isFalse:", isFalse)
}
~~~

~~~go
// 字符串类型
package main

import "fmt"

func main() {
    var hello string = "Hello, World!"
    // 单箭头换行
    var multiLine string = `This is
a multi-line
string.`

    fmt.Println(hello)
    fmt.Println(multiLine)
}

~~~

~~~go
// 常量
package main

import "fmt"

const Pi = 3.14159
const Truth = true
const Greeting = "Hello, Go!"

func main() {
    fmt.Println("Pi:", Pi)
    fmt.Println("Truth:", Truth)
    fmt.Println(Greeting)
}
~~~

## 五、数据类型转换

### 1.整数类型之间的转换、浮点数与整数的转换

~~~go
package main

import "fmt"

func main(){
    var n int = 42
    var b int64 = int64(n)  // 将 int 转换为 int64
    var c int8 = int8(n)  // 将 int 转换为 int8
    
    var f float = 3.141
    vat i int = int(f)  // 将 float64 转换为 int //结果为3
}
~~~

### 2.字符串与其他类型的转换(**strconv包**)

~~~go
package main

import (
    "fmt"
    "strconv"
)

func main(){
    var i int = 32
    var f float64 = 3.14
    var b bool = true
    
    // 其他类型转字符串
    // int类型数字转换为字符串数字
    var s1 = strconv.Itoa(i) 
    // 参数说明：变量， 浮点数格式化方式，指定小数点要保留的位数(-1为自动选择全部)， 转换为字符串后浮点数的精度
    var s2 = strconv.FormatFloat(f, "g", -1, 64) 
    // 将 bool 转换为 string
    var s3 = strconv.FormatBool(b)
    
    // 字符串转其他类型
    i2, _ = strconv.Atoa("123")          // 将 string 转换为 int
    f2, _ = strconv.ParseFloat("3.14")   // 将 string 转换为 float64
    f3，_ = strconv.ParseBool("true")    // 将 string 转换为 bool
}
~~~

3.数组与切片的数据转换：要用循环处理，不能直接Format整个数组或切片



## 六、复合数据类型(数组、切片slice、映射map、结构体、JSON)

### 1.数组Array （固定大小的同类型元素集合。）

**var arr [3]int{1,2,3}**

~~~go
// 数组定义 var 变量名 [数组大小]数组元素类型

// 定义一个长度为 5 的整型数组
var arr [5]int  // [0,0,0,0,0]

// 定义并初始化数组
arr := [3]int{1, 2, 3}

// 自动推断长度
arr := [...]int{1,2,3,4,5}

// 多维数组
var matrix [2][2]int
~~~

~~~go
// 数组的增删查改操作
var arr [3]int{1,2,3}

// 读取数组元素
arr[1] // 2

// 修改数组元素
arr[1] = 2  // [2,2,3]

// 增加数组元素(在数组中数组的大小是固定的，要增加元素 需要将使用切片)
arr_slice := arr[:]
arr_slice = append(arr_slice, 60) // [2,2,3,60]

// 数组长度
len(arr)

// 删除元素(同样的。需要转换为切片slice进行操作) 在切片部分讲解

// 遍历数组
for index, value := range arr{
    fmt.Println(index, value)
}
~~~

### 2.切片 （动态大小的数组视图，支持动态扩展。）

**var sli = make([]int, 3, 5)**

~~~go
// 切片定义 	

// 字面量定义，但是不需要定义大小，类似于数组
var sli []int{1,2,4}

// make函数定义
// 创建一个长度为3，容量为5的切片
// 指定切片的容量主要用于优化性能，特别是在涉及到动态增长时。
/* 1.当你使用 append 函数向切片添加元素时，如果切片的容量不足以容纳新元素，Go 语言会自动创建一个更大的底层数组，并将旧数组的元素复制到新数组中。预定义容量可以减少这种动态扩展的次数，从而提高性能。*/
/* 2.如果切片的初始容量足够大，可以在添加大量元素时避免频繁的内存分配和复制操作。这样可以显著提升性能，尤其是在需要动态扩展切片的情况下。*/

// 注意 该切片的样子为[0,0,0]
var sli = make([]int, 3, 5)
// 不指定切片容量大小，新增时会以长度大小的两倍扩大切片容量
var sli = make([]int, 3)

// 利用数组创建切片(范围取值)
var arr [5]int{1,2,3,4,5}
var sli = arr[1:3]  // [2,3,4]

// 切片长度
len(sli)
~~~

~~~go
// 切片的增删查改

var sli = []int{1,2,3}

// 访问切片元素
sli[2] // 3

// 修改切片元素
sli[2] = 10 // [1,2,10]

// 追加元素(新增)
sli = append(sli, 20)  // [1,2,10,20]

// 删除元素（通常通过创建新的切片来完成，不会直接修改原切片的大小。）
// 删除索引为2的元素
IndexToDelete := 2
result := append(sli[:IndexToDelete], sli[IndexToDelete+1:]...)  // [1,2,20] 跳过下标为2的截取
// ... 运算符将 slice[index+1:] 展开为单独的参数 相当于每个sli[IndexToDelete+1:] 都执行一次append到sli[:IndexToDelete]中

// 遍历切片
for index, value := range sli{
    fmt.Println(index, value)
}
~~~

### 3.映射map （无序的键值对集合，键是唯一的。）

**var m = map["string"]int** 

~~~go
// 映射定义

func main(){
    // 字面量定义映射 "string"代表映射的key类型 int代表映射的value类型
    var m = map["string"]int{
        "Apple": 10,
        "banana": 6,
        "watermalen": 4
    }
    
    // 使用make函数定义映射
    var m2 = make(map["string"]int)
    m2["Apple"] = 12
}
~~~

~~~go
// 映射增删查改
var m = map["string"]int{
    "Apple": 10,
    "banana": 3
}

// 访问映射元素
m["Apple"]  // 10

// 判断元素key是否存在
price, exists := m["orange"]
if exists{
    fmt.Println(price)
}else{
    fmt.Println("Not Found Key")
}

// 修改键值对
m["Apple"] = 20  // {"Apple": 20, "banana": 3}

// 新增键值对
m["grape"] = 14  // {"Apple": 20, "banana": 3, "grape": 14}

// 删除键值对
delete(m, "Apple")  // {"banana": 3, "grape": 14}

// 遍历映射
for key, value := range m{
    fmt.Println(key, value)
}
~~~

### 4.结构体struct (由多个字段组成的复合数据类型，字段可以有不同的类型。)

在 Go 语言中，结构体（struct）是一种聚合数据类型，可以将不同类型的字段组合在一起，用于表示复杂的数据结构。结构体类似于其他编程语言中的“类”或“记录”，但 Go 语言中的结构体不包含方法，方法需要通过关联函数来实现。

**type Person struct {}** 

~~~go
// go字面量初始化
type Person struct {
    Name string // 姓名
    City string // 城市
    Age int     // 年龄
}

// 实例化结构体
p := Person{
    Name: "Thomas",
    City: "Shenzhen",
    Age: 25
}

// 零值初始化
var p Person
// p.Name == ""
// p.City == ""
// p.Age == 0

// 点运算符访问
p.City  // Shenzhen
~~~

~~~go
// 修改结构体值
p.City = "Los Angles"

// 删除和新增操作不支持 结构体就是定义好结构的一个数据结构
~~~

~~~go
// 结构体嵌套

type Contact struct {
    Phone string
    Email string
}

type Employee struct {
    Name string
    Age int
    Contact Contact
}

employee := Employee {
    Name: "Thomas",
    Age: 25
    Contact: Contact{
        Phone: "13978514254",
        Email: "13978514254@163.com"
    }
}

// 访问属性
fmt.Println(employee.Name) // Thomas
fmt.Println(employee.Contact.Phone) // 13978514254
~~~



### 5.Json （轻量级的数据交换格式，用于数据的编码和解码。）

**json.Marshal(p) ** 

**json.Unmarshal([]byte(JsonStr), &p2)** 

~~~go
// JSON字符串与结构体的相互转换
// encoding/json: JSON 数据的编码（序列化）和解码（反序列化）
import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string
    Age int
}

p := Person{
    Name: "Thomas",
    Age: 14
}

// 结构体转Json
JsonData, _ := json.Marshal(p)
fmt.Println(string(JsonData))  // {"Name":"Thomas","Age":14} 要转为字符串输出
// JsonData [123 34 110 97 109 101 34 58 34 66 111 98 34 44 34 97 103 101 34 58 50 53 125] JSON 格式的字节切片

// Json字符串转结构体
var p2 Person
JsonStr := `{"Name":"Thomas","Age":14}`
json.unmarshal([]byte(JsonStr), &p2)
fmt.Printf("Decoded Person: %+v\n", p2)  // {"Name":"Thomas","Age":14}
~~~



## 七、控制流

~~~go
// if else
package main
func main(){
    i := 10
    if i > 0 {
        fmt.Println("大于0")
    else{
        fmt.Println("小于0")
    }
}
~~~

~~~go
// if, else if, else
package main
func main(){
    i := 10
    if i > 0 {
        fmt.Println("大于0")
    }else if i < 0{
        fmt.Println("小于0")
    } else{
        fmt.Println("等于0")
    }
}
~~~

~~~go
// switch case
package main
func main(){
    day := 6
    switch day {
    case 1:
        fmt.Println(1)
    case 2:
        fmt.Println(2)
    default:
        fmt.Println("Other")
    }
}
~~~

~~~go
// for
// 计数循环
for init; condition; post {
    
}

for i:=0; i<10; i++ {
    
}

// 无限循环
for {
    
}
count := 0
for {
    fmt.Println("无限循环中")
    count++
    if count >= 3 {
        break
    }
}

// while循环
for condition {
    
}

var condition = true
for condition {
    
}

// 遍历数组、切片、映射
for index, value := range collection{
    // index 是当前索引，value 是当前值
}
~~~



## 八、异常处理

Go语言采用了一种简单而独特的错误处理机制。它鼓励明确地检查错误，而不是依赖异常机制。

1.基础的错误处理

~~~go
// 函数通常会返回一个值和一个错误对象。
// 如果错误对象为nil，表示没有错误；否则，它包含了关于该错误的信息。
result, err := someFunction()
if err != nil {
   log.Fatal(err)
}
~~~

2.错误处理的场景

~~~go
// 1.文件读取操作
// 在进行文件操作时，可能会遇到诸如文件不存在、权限不足等各种错误。
package main

import (
	"fmt"
    "io/ioutil"
)

func main(){
    FileName := "example.text"
    content, err := ioutil.ReadFile(FileName)
    if err != nil {
        fmt.Printf("Error reading file %s: %s", FileName, err)
    	return 
    }
    fmt.Printf("File Content: %s", content)
}
~~~

~~~go
// 网络请求
// 在进行网络请求时，可能会遇到连接超时、网络不可达等错误。

package main
import (
	"fmt"
    "net/http"
    "time"
)

func main(){
    client := http.Client{
        Timeout: 5 * time.Second, // 设置请求超时时间
    }
    resp, err := client.Get("http://example.com")
    if err != nil {
        fmt.Printf("Error making http request: %s\n", err)
        return // 出现错误及时停止执行函数
    }
    defer resp.Body.Close()
    fmt.Println("Response status:", resp.Status)
}
~~~

~~~go
// 数据库操作
// 在执行数据库查询、更新时，可能会遇到数据库连接失败、SQL 语法错误等问题。

package main

import (
    "database/sql"
    "fmt"
    "github.com/go-sql-driver/mysql"
)

func main(){
    db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database")
    if err != nil {
        fmt.Printf("Error connecting to database: %s\n", err)
        return  // 执行遇到错误，提前结束
    }
    defer db.Close()
    
    rows, err := db.Query("select * from users")
    if err != nil {
        fmt.Printf("Error executing SQL Query: %s\n", err)
    	return 
    }
    defer rows.Close
    
    // 处理查询结果
}
~~~

~~~go
// API调用
~~~

[掌握Go语言：Go语言精细错误，清晰、高效的错误处理实践](https://blog.csdn.net/wenbingy/article/details/137481620?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-5-137481620-blog-135625264.235^v43^control&spm=1001.2101.3001.4242.4&utm_relevant_index=8) 



## 九、函数

在 Go 语言中，函数是基本的代码结构单元，用于将可复用的逻辑封装起来。Go 语言的函数支持多种特性，包括参数传递、返回值、多返回值、匿名函数、闭包等。

~~~go
func FunctionName() returnType {
    // 函数体
    return result
} 
~~~

基本用法

~~~go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(3, 4)
    fmt.Println("Result:", result)
}

~~~

多返回值

~~~go
package main

import (
	"fmt"
    "error"
)

func division(a int, b int)(result int, err error){
    if b == 0 {
        return 0, error.New("divide by zero")
    }
    return a/b, nil
}

func main(){
    a := 10
    b := 2
    result, err := division(a,b)
    if err != nil {
        fmt.Printf("division operation fail: %s\n", err)
        return
    }
    fmt.Println(result)
}
~~~

匿名函数

~~~go
package main

import (
	"fmt"
)

func main(){
    // 定义匿名函数并立即调用
    func(message string){
        fmt.Println(message)
    }("Hello World")
    
    // 将匿名函数结果复制给变量 变量函数
    add := func(a int, b int) (result int) {
        return a + b
    }
    result := add(1, 2)
    fmt.Println(result)
}
~~~

闭包

~~~go
package main

import "fmt"

func main(){
    counter := 0
    increasement := func()(autoplus, int){
        counter = counter + 1
        return counter
    }
    
    fmt.Println(increasement()) // 1
    fmt.Println(increasement()) // 2
    fmt.Println(increasement()) // 3
}
~~~

defer关键字：延迟函数的执行，函数A包含defer修饰的函数 那么等到函数A执行完之后 才会执行defer修饰的函数

~~~go
// 延迟执行代码块，常用于资源释放或异常处理。
package main

import "fmt"

func main(){
    defer fmt.Println("Hello defer")  // 后输出
    fmt.Println("run first")  // 先输出
}
~~~

~~~go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close() // 确保文件在函数结束时关闭

    // 处理文件内容
    fmt.Println("File opened successfully")
}
~~~

多个defer 最后一个defer先执行 以栈的方式 先入后出

~~~go
package main

import "fmt"

func main() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")

    fmt.Println("Main function")
}

Main function
Third defer
Second defer
First defer
~~~

**defer + panic + recover版的异常处理**  

在 Go 语言中，`panic` 和 `recover` 是用于异常处理的机制。`panic` 用于引发一个异常，通常表示程序遇到了不可恢复的错误。而 `recover` 则用于捕获并处理这些异常，从而防止程序崩溃。

**使用场景** 

- **恢复不可恢复的错误**：使用 `recover` 可以从可能导致程序崩溃的 `panic` 中恢复，**继续执行程序**。这对于构建健壮的、不会因为单个错误而崩溃的服务特别有用。
- **局部处理**：使用 `panic` 和 `recover` 可以在特定模块或函数内捕获并处理异常，而不影响整个程序的执行。

~~~go
package main

func main(){
    defer func(){
        if r := recover(); r != nil {
            fmt.Println("Handle error：", r)
        }
    }
    fmt.Println("About to panic")
    panic("something worong")
    fmt.Println("output finished?  This line will never be executed")
}

/*
	main函数执行
	defer修饰的匿名函数将最后执行
	
	打印About to panic
	执行panic函数 程序终止。
	
	执行defer函数(因为panic main函数已经结束了 该defer的匿名函数执行了)
	捕获panic的输出，执行匿名函数 输出Handle error：something worong

	main函数执行结束(panic后的输出将不会执行)
*/

/*
慎用 panic：panic 应该用于真正的不可恢复错误，例如程序的内部逻辑错误或无法处理的外部条件。对于大多数普通错误，应该使用返回值（如 error）来处理，而不是 panic。
recover 只能在 defer 中有效：recover 必须在 defer 函数中使用，否则无法捕获 panic。
*/
~~~

**总结** 

- **`panic`**：用于引发异常，导致程序崩溃。
- **`recover`**：用于捕获和处理异常，防止程序崩溃，必须在 `defer` 中使用。
- **组合使用**：`panic` 和 `recover` 组合使用，可以提供一种强大的错误处理机制，确保在异常情况下程序能够安全地处理并继续运行。



## 十、方法

在 Go 语言中，方法是与特定类型相关联的函数。方法与函数类似，但它们属于某个类型（即接收者类型）。方法可以访问和修改其接收者的状态，并且与结构体、接口等类型紧密相关。

定义方法：方法定义的语法与函数类似，不同之处在于方法的定义包括一个接收者（receiver），它表示该方法是哪个类型的方法。

~~~go
func (receiver ReceiverType) MethodName(parameters) returnType {
    // 方法体
}
~~~

~~~go
package main

import "fmt"

// 首先定义一个结构体 矩形
type Rectangle struct {
    Width int
    Length int
}

// 为矩形结构体定义一个计算面积的函数 （值接收者：接收者是类型的一个值副本，方法中对接收者的修改不会影响原始值。）
func (r Rectangle) Area()(ResultArea int) {
    return r.Width * r.Length
}

// 修改矩形结构体的数值(指针传递)  
// （指针接收者：接收者是类型的一个指针，方法中对接收者的修改会影响原始值。通常使用指针接收者来避免复制大量数据或修改原始数据。）
func (r *Rectangle) Resize(Width int, Length int) {
    r.Width = Width
    r.Length = Length
}

func main(){
    rect := Rectangle{Width:5, Length:10}
    ResultArea := rect.Area()
    fmt.Println(ResultArea) // 50
    
    // 修改长宽
    ResizeWidth := 10
    ResizeLength := 10
    rect.Resize(ResizeWidth, ResizeLength)
    ResultArea := rect.Area()
    fmt.Println(ResultArea)  // 100
}
~~~



## 十一、接口

在 Go 语言中，接口（interface）是一种类型，它定义了一组方法的集合。接口允许**不同类型的对象共享行为**而不需要继承，这种设计使得 Go 语言的类型系统非常灵活。**接口是 Go 语言中实现多态的核心机制，通过接口可以实现灵活的代码设计和模块化。** 

~~~go
// 定义方式
type InterfaceName interface {
    MethodName1(parameter1 type1, parameter2 type2) returnType1
    MethodName2(parameter1 type1) returnType2
}
~~~

示例：多态性、解耦合、可拓展性。

~~~shell
多态性：不同类型的对象共享行为
解耦合：接口只管统一行为，具体的实现由不同对象的方法实现去完成
可拓展性：新增功能是只需实现接口定义，不需要修改已有代码
~~~

~~~go
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

The Area of Rectangle is 200.000000
The Area of Circle is 314.000000
Area: 200.000000
Area: 314.000000
~~~

~~~shell
总结：
多态性：不同的形状结构体去实现共有的方法，计算面积
解耦性：接口只定义共有方法，具体的面积算法在各形状Area实现中具体完成
可拓展性：当有新的形状对象进来(三角形)，只需定义好三角形结构体，然后让三角形实现具体的面积计算方法，即可实现共有行为的共享，期间不需要修改其他形状的面积计算方法。
~~~

参考：[浅谈Golang接口：作用、应用场景及实际应用](https://www.cnblogs.com/cheyunhua/p/17875734.html) 



## 十二、并发编程

1.Goroutine的基本使用

~~~go
package main

import (
	"fmt"
	"time"
)

func count(n int, animal string) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, animal)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	/*Version 1 : 顺序执行*/
	count(3, "sheep")
	count(2, "cow")
}

1 sheep
2 sheep
3 sheep
1 cow
2 cow
~~~

~~~go
// 并发执行
package main

import (
	"fmt"
	"time"
)

func count(n int, animal string) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, animal)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	/*Version 2 : 并发执行*/
	go count(3, "sheep")
	count(2, "cow")
}

// 注意 输出结果是乱序的
1 cow
1 sheep
2 sheep
2 cow
~~~

~~~go
// 尝试同时并发执行
package main

import (
	"fmt"
	"time"
)

func count(n int, animal string) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, animal)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	/*Version 3 : 尝试同时并发执行*/
	go count(3, "sheep")
	go count(2, "cow")
}

// 无输出， 程序直接结束
// 原因在于 协程在后台执行 但是main函数(主线程)先执行完毕了，那么协程就不会执行 改进方案:WaitGroup 
~~~

~~~go
// WaitGroup并发执行
// 并发执行
package main

import (
	"fmt"
	"time"
	"sync"  // 引入WaitGroup等待协程运行完毕再终止主线程(main)
)

func count(n int, animal string) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, animal)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// 创建一个WaitGroup对象(计数器)
	var wg sync.WaitGroup
	// 声明有多少个任务要被执行
	wg.Add(2)

	// 匿名函数封装count函数运行 定义之后即刻运行
	go func() {
		count(3, "sheep")
		// 计数器wg-1
		wg.Done()
	}()

	go func() {
		count(2, "cow")
		// 计数器wg-1
		wg.Done()
	}()

	// 让主线程main函数在此等待goroutine执行完毕之后 再往下执行
	wg.Wait()

	fmt.Println("goroutine done, main continue...")
}

// 输出
1 cow
1 sheep
2 sheep
2 cow
3 sheep
goroutine done, main continue...
~~~

2.Goroutine之间的通信

任务之间、主任务之间互相交流、传输数据，注意：**在其他的编程语言中，线程之间的相互交流通常是共享内存来完成，Go的实现是通过交流来实现共享内存** 

~~~go
// 假设要计算两个协程一共计算了多少个动物
package main

import (
	"fmt"
	"time"
	"sync"  // 引入WaitGroup等待协程运行完毕再终止主线程(main)
)

// 定义全局变量counter
var counter = 0

func count(n int, animal string) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, animal)
		counter = counter + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// 创建一个WaitGroup对象(计数器)
	var wg sync.WaitGroup
	// 声明有多少个任务要被执行
	wg.Add(2)

	// 匿名函数封装count函数运行 定义之后即刻运行
	go func() {
		count(3, "sheep")
		// 计数器wg-1
		wg.Done()
	}()

	go func() {
		count(2, "cow")
		// 计数器wg-1
		wg.Done()
	}()

	// 让主线程main函数在此等待goroutine执行完毕之后 再往下执行
	wg.Wait()

	fmt.Println(counter)

	fmt.Println("goroutine done, main continue...")
}

// 输出
1 cow
1 sheep
2 sheep
2 cow
3 sheep
5
goroutine done, main continue...
~~~

通过在具体的实现函数中修改全局变量，当使用协程去并发执行这个函数的时候，因为Goroutine有可能同时分配到多个CPU核心，这时候当counter为1时，数羊的协程执行counter+1=2, 数牛的协程同时也执行counter+1=2，最后counter变量的内存存放的结果为2，但是应该结果为3，因为两个协程各自数了一次。所以设计的想法就是**通过交流共享内存** 利用Channel实现。

channel：协程通信的通道，一个channel只能用来发送和接收一种数据类型，channel是双向的。But过，注意channel的收发消息，都会造成代码运行阻塞，也就是说 一个协程向channel发送一个信息，他会一直等待这个消息被接收了，才会继续往下执行。反过来，如果向从channel中接收消息，channel的另一边没有发送消息，那么接收消息这端也会一直等待，造成代码运行阻塞。

~~~go
// 主线程与协程之间的通信
package main

import (
	"fmt"
	"time"
)

func count(n int, animal string, c chan string) {
	for i := 0; i < n; i++ {
		// 向channel中传递信息
		c <- animal
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// 定义一个channel 用于协程通信 这里是主线程与协程之间的通信
	// 创建一个channel，通信的数据类型为字符串类型
	c := make(chan string)
	go count(3, "sheep", c)
	// 从channel中收听信息
	message := <- c
	fmt.Println(message)
}
// 输出 只输出了一个，但是channel中有三条消息，进一步改进 for 不停读取channel消息
sheep
~~~

~~~go
// 并发执行
package main

import (
	"fmt"
	"time"
)

func count(n int, animal string, c chan string) {
	for i := 0; i < n; i++ {
		// 向channel中传递信息
		c <- animal
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// 定义一个channel 用于协程通信 这里是主线程与协程之间的通信
	// 创建一个channel，通信的数据类型为字符串类型
	c := make(chan string)
	go count(3, "sheep", c)
	// 从channel中收听信息
	for {
		message := <- c
		fmt.Println(message)
	}
}

// 输出
sheep
sheep
sheep
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/home/withboom/Go-learning/gobasic/Concurrency/test_goroutine.go:24 +0x7b
exit status 2

/*
	在输出完成后，结果报错，因为for循环一直读取channel中的消息，但是消息在前面已经输出完成了，channel中没有消息了 这时候就报	 错，提示所有协程都休眠了 没有协程往channel中传递消息
*/
~~~

~~~go
// 并发执行
package main

import (
	"fmt"
	"time"
)

func count(n int, animal string, c chan string) {
	for i := 0; i < n; i++ {
		// 向channel中传递信息
		c <- animal
		time.Sleep(time.Millisecond * 500)
	}
	// 当消息传递完毕后，手动关闭channel （注意点）
	close(c)
}

func main() {
	// 定义一个channel 用于协程通信 这里是主线程与协程之间的通信
	// 创建一个channel，通信的数据类型为字符串类型
	c := make(chan string)
	go count(3, "sheep", c)
	// 从channel中收听信息range实现 （注意点）
	for message := range c {  
		fmt.Println(message)
	}
}

//输出

sheep
sheep
sheep
~~~

当有多个协程与其各自的channel执行任务时，如何处理这些协程与channel之间传来的数据？

~~~go
// 并发执行
package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个channel 用于协程通信 这里是主线程与协程之间的通信
	// 创建一个channel，通信的数据类型为字符串类型
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "sheep"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "cow"
			time.Sleep(time.Millisecond * 2000)
		}
	}()
	 
	// 从channel中收听信息
    /*
    	注意点： 这里用for循环不停接收channel中的输出的信息，但是因为channel的消息放入和消息接收会阻塞代码执行
    	这里的话因为每个函数中都设置了休眠，所以导致输出的结果不是四个羊(0.5秒)一头牛(2秒)  而是按顺序在for循环中
    	执行c1的接收 然后c2接收，c2接收后等2秒 再回到c1接收(卡的这2秒没有人从c1接收) 实际上干等c2接收之后 c1才继续执行。 
    */
	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}

// 输出
sheep
cow
sheep
cow
sheep
cow
sheep
cow

~~~

有什么办法处理上面的白等情况？Golang中设置了Select语法来处理，监听，如果有channel可以接收 则去接收那个channel, 而不必等待一个channel完之后再去执行另一个channel的接收

~~~go
// 并发执行
package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个channel 用于协程通信 这里是主线程与协程之间的通信
	// 创建一个channel，通信的数据类型为字符串类型
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "sheep"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "cow"
			time.Sleep(time.Millisecond * 2000)
		}
	}()
	 
	// 从channel中收听信息 使用select解决阻塞问题
	// 当有多个channel发送消息的时候，使用select保证第一时间接收channel中的数据，而不需要等待其他channel的顺序执行阻塞时间
    // 不受不同channel消息频率不同的影响
	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}
}

// 输出
cow   // 第一个cow和下面的sheep同时执行 因为cow要等2秒 所以下面的sheep还会输出三次 再到cow 所以是一头牛四头羊
sheep
sheep
sheep
sheep
cow
sheep
sheep
sheep
sheep
...
~~~


goroutine参考:[用多线程给代码提速800% —— Golang高并发教程+实战](https://www.bilibili.com/video/BV1qT4y1c77u/?spm_id_from=333.337.search-card.all.click&vd_source=bc3fd2b42eb94a414c6746120ff5ddc6)




