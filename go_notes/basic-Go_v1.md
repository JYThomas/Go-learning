## Go基础



### 1.包声明

​		在Go语言中，**包声明用于指定当前文件所属的包**。包是组织和管理Go代码的基本单元，它可以包含多个Go源代码文件。一个包可以是一个库（library）或一个可执行程序（executable）。
​		在每个Go源代码文件的开头，都需要使用包声明来指定当前文件所属的包。包声明使用关键字package后面跟着包的名称。例如：

~~~shell
package main
~~~

​		在上面的例子中，`package main`表示当前文件属于名为"main"的包。"main"包是一个特殊的包，它用于定义可执行程序的入口函数`main()`。

**包声明通常与文件名保持一致**，例如，如果文件名是`example.go`，那么包声明应该是`package example`。

包声明的作用是：

1. 帮助组织和分类代码：通过将相关的代码文件放在同一个包中，可以更好地组织和管理代码，提高代码的可读性和可维护性。
2. 控制代码的可见性：在Go语言中，**标识符（变量、函数、类型等）的可见性是由标识符的首字母的大小写来决定的**。如果一个标识符的首字母是大写字母，那么它对于其他包是可导出的（可公开的），其他包可以访问它；如果一个标识符的首字母是小写字母，那么它对于其他包是不可导出的（不可公开的），其他包无法直接访问它。**包声明的名称也遵循这个规则**，如果包的名称是大写字母开头，那么其他包可以访问它；如果包的名称是小写字母开头，那么其他包无法直接访问它。

### 2.main函数

`main()`函数在Go语言中类似于Java中的主函数。

在Go语言中，`main()`函数是程序的入口函数。当程序开始执行时，会首先调用`main()`函数。在`main()`函数中，您可以编写程序的主要逻辑，包括调用其他函数、处理输入输出、进行计算等等。

~~~shell
`main()`函数的签名是固定的，没有参数和返回值。
程序只能有一个`main()`函数。在同一个包中，只能有一个`main()`函数作为程序的入口。
`main()`函数是程序执行的起点，它会负责调用其他函数、处理程序的逻辑和控制流程。它是程序的入口点。
~~~



### 3.变量声明与定义

#### (1)使用var声明变量并指定类型

~~~shell
var 变量名 类型
var age int
var name string
# 在这种方式中，先使用关键字`var`声明变量，然后指定变量名和类型。例如，`age`是一个整数类型的变量，`name`是一个字符串类型的变量。
~~~

#### (2)声明并初始化变量

~~~shell
var 变量名 类型 = 值
var age int = 25
var name string = "John"
# 在这种方式中，可以在声明变量的同时为其指定初始值。例如，`age`被初始化为25，`name`被初始化为"John"。
~~~

#### (3)简短变量声明(函数内部使用)

短变量声明方式只能用于函数内部局部变量，**不能在函数外使用**

~~~shell
变量名 := 值
age := 25
name := "John"
# 简短声明变量是一种更简洁的方式，通常在函数内部使用。在这种方式中，可以省略`var`关键字和变量的类型，Go语言会根据值的类型自动推断变量的类型。例如，`age`被推断为整数类型，`name`被推断为字符串类型。
~~~

#### (4)常量定义

​		常量的特点是一旦被定义后，其值就不能再被修改。常量在程序运行期间保持不变。因此，常量的值必须在声明时就赋予一个确定的值。

​		在实际使用中，常量的类型和值可以根据需要进行选择。常量可以是数值类型（整数、浮点数）、布尔类型（true、false）、字符串类型等。

​		需要注意的是，常量的命名规则和变量相同，按照Go语言的命名规范，常量的首字母大小写决定了其可见性。如果常量的首字母是大写字母，其他包可以访问该常量；如果首字母是小写字母，其他包无法直接访问该常量。

~~~shell
const 常量名 类型 = 值
const Pi int = 3
const Pi float64 = 3.14159

# 默认类型和指定类型的情况
const Pi = 3.14159  // 声明一个浮点数常量，默认类型为float64
const MaxSize int = 100  // 声明一个整数常量，指定类型为int
const Monday = "星期一"  // 声明一个字符串常量，默认类型为string
const IsEnabled = true  // 声明一个布尔常量，默认类型为bool
~~~

#### (5)变量作用域



### 4.Go数据类型与数据结构

#### (1)Go数据类型

~~~shell
基本数据类型：
	整数类型：包括int、int8、int16、int32、int64，以及无符号整数类型uint、uint8、uint16、uint32、uint64。
	浮点数类型：包括float32和float64。
	布尔类型：bool，取值为true或false。
	字符串类型：string，用于表示文本数据。
	字符类型：byte，用于表示ASCII字符，相当于uint8类型。
	复数类型：complex64和complex128，用于表示复数。

复合数据类型：
	数组（Array）：具有固定长度且元素类型相同的数据结构。
	切片（Slice）：动态长度的序列，是对数组的一层封装。
	映射（Map）：键值对的无序集合。
	结构体（Struct）：可以包含不同类型字段的自定义类型。
	接口（Interface）：定义了一组方法的集合，用于实现多态性。
	函数（Function）：用于封装一段可执行的代码块。
~~~

除了以上列出的常用数据类型，Go语言还提供了一些其他的数据类型，如指针（Pointer）、通道（Channel）等，用于更高级的编程需求。

~~~shell
# 整数类型
package main

import "fmt"

func main() {
    var num int = 10
    fmt.Println(num)
}
~~~

~~~shell
# 浮点类型
package main

import "fmt"

func main() {
    var num float64 = 3.14
    fmt.Println(num)
}
~~~

~~~shell
# 布尔类型
package main

import "fmt"

func main() {
    var isTrue bool = true
    fmt.Println(isTrue)
}
~~~

~~~shell
# 字符串类型
package main

import "fmt"

func main() {
    var message string = "Hello, World!"
    fmt.Println(message)
}
~~~

~~~shell
# 字符类型 
package main

import "fmt"

func main() {
    var ch byte = 'A'
    fmt.Println(ch)
}
~~~

~~~shell
# 复数类型
package main

import "fmt"

func main() {
    var complexNum complex128 = 3 + 2i
    fmt.Println(complexNum)
}
~~~

#### (2)Go数据类型转换

​		在Go语言中，可以使用类型转换（Type Conversion）将一个数据类型的值转换为另一个数据类型。类型转换的基本语法如下：

~~~shell
目标类型的值 = 目标类型(待转换的值)
~~~

~~~shell
# 整数类型转换为浮点数类型：
package main

import "fmt"

func main() {
    var num int = 10
    var floatNum float64 = float64(num)
    fmt.Println(floatNum)
}
~~~

~~~shell
# 浮点数类型转换为整数类型
package main

import "fmt"

func main() {
    var floatNum float64 = 3.14
    var num int = int(floatNum)
    fmt.Println(num)
}
~~~

~~~shell
# 整数类型转换为字符串类型
package main

import "fmt"
import "strconv"

func main() {
    var num int = 42
    var str string = strconv.Itoa(num)
    fmt.Println(str)
}
~~~

~~~shell
# 字符串类型的整数转换为int类型
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    num, err := strconv.Atoi(str)
    if err == nil {
        fmt.Println(num)
    } else {
        fmt.Println("转换失败:", err)
    }
}

# 字符串类型的浮点数转换为float64类型
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "3.14"
    num, err := strconv.ParseFloat(str, 64)
    if err == nil {
        fmt.Println(num)
    } else {
        fmt.Println("转换失败:", err)
    }
}

```使用strconv.ParseFloat函数将字符串变量str转换为float64类型的数值。函数的第二个参数64表示将结果转换为64位浮点数。如果转换成功，将打印输出转换后的浮点数值；如果转换失败，将打印输出转换失败的错误信息```
~~~

#### (3)Go数据结构定义及其操作

##### （3.1）数组(Array) 

**固定长度的数据结构,动态操作的情况在切片上** 

~~~shell
# 数组声明
var arrayName [length]dataType
声明一个包含 5 个整数的数组： var numbers [5]int
~~~

~~~shell
# 数组初始化
使用索引分配值： numbers := [5]int{1, 2, 3, 4, 5}。
省略长度的方式： numbers := [...]int{1, 2, 3, 4, 5}，Go 会根据提供的元素数量自动推断长度。
使用索引和值的方式： numbers := [5]int{0: 1, 2: 3}，这样可以指定特定索引位置的值。 # [0,1,3,0,0]
~~~

~~~shell
# 多维数组
var matrix [rows][columns]dataType
rows 是行数，columns 是列数，dataType 是元素的数据类型。
~~~

##### （3.2）数组操作

~~~shell
# 访问数组元素
value := numbers[2]
~~~

~~~shell
# 数组长度
length := len(numbers)
~~~

~~~shell
# 数组遍历(索引循环)
numbers := [5]int{1, 2, 3, 4, 5}

for index := 0; index < len(numbers); index++ {
    fmt.Println(numbers[index])
}

# 数组遍历(元素循环)
numbers := [5]int{1, 2, 3, 4, 5}

for index, value := range numbers {
    fmt.Println(index, value)
}
~~~

##### （3.3）切片(Slice)

​		在 Go 语言中，可以使用切片（slice）来定义不知道长度大小的动态数组。切片是对数组的一个引用，它提供了方便的动态操作和自动扩容功能。

​		切片由三个部分组成：指向底层数组的指针、切片的长度和切片的容量。容量表示切片可以容纳的元素个数，而长度表示切片**当前拥有的元素个数**。

​		切片的容量设置相当于为切片提供了一定的预留空间，以便在需要时进行动态扩容，而不需要每次添加元素时都重新分配内存。这是切片相对于数组的一项重要特性。

~~~shell
# 定义切片
var sliceName []dataType
```sliceName 是切片的名称，dataType 是切片中元素的数据类型。

var numbers []int  // 定义一个整数切片
var names []string  // 定义一个字符串切片
~~~

~~~shell
# 切片初始化(make函数方式)
sliceName := make([]dataType, length, capacity)
numbers = make([]int, 5)  // 初始化长度为 5 的整数切片(初始值为0)  # 赋值方式靠循环
```其中，`length` 是切片的初始长度，`capacity` 是切片的初始容量。初始容量是可选的，如果不指定初始容量，切片将会根据需要进行自动扩容。

# 切片初始化(字面量方式)
sliceName := []dataType{value1, value2, ...}
names = []string{"Alice", "Bob", "Charlie"}  // 初始化字符串切片并指定初始值 # 长度为字符串元素的数量
```在大括号中列出切片的初始元素值，每个值之间用逗号分隔。
~~~

​		需要注意的是，切片是引用类型，它们在内存中的存储是一个指向底层数组的指针。切片的长度可以使用 `len()` 函数获取，容量可以使用 `cap()` 函数获取。

##### （3.4）切片操作

（3.4.1）**切片长度与容量获取** 

~~~shell
使用 len() 函数获取切片的长度：length := len(numbers)
使用 cap() 函数获取切片的容量：capacity := cap(numbers)
~~~

（3.4.2）**切片元素追加** 

~~~shell
numbers := []int{1, 2, 3}
numbers = append(numbers, 4, 5)  // 向切片追加元素
~~~

（3.4.3）**切片元素截取** 

~~~shell
numbers[start:end]：截取切片，包括索引 start 对应的元素，但不包括索引 end 对应的元素。 # 左闭右开
~~~

（3.4.4）**切片元素复制**

~~~shell
source := []int{1, 2, 3}
destination := make([]int, len(source))
copy(destination, source)
~~~

（3.4.5）**切片元素删除** 

~~~shell
numbers := []int{1, 2, 3, 4, 5}
index := 2
numbers = append(numbers[:index], numbers[index+1:]...)
~~~

（3.4.6）**切片遍历** 

~~~shell
numbers := []int{1, 2, 3}
for index, value := range numbers {
    // 使用 index 和 value 进行操作
}
~~~

（3.4.7）**切片比较** 

~~~shell
slice1 := []int{1, 2, 3}
slice2 := []int{1, 2, 3}
equal := reflect.DeepEqual(slice1, slice2)

# 当两个切片进行相等性比较时，它们需要满足以下条件才会被认为是相同的：
长度相等：两个切片的长度必须相同。
内容相同：两个切片的每个元素在相同位置上的值必须相等。

# 切片相同 vs 切片不同
如果两个切片满足上述条件，那么它们被认为是相同的，比较的结果为 true。反之，如果两个切片的长度不相等或者它们的某个元素在相同位置上的值不相等，那么它们被认为是不相同的，比较的结果为 false。
~~~

~~~shell
slice1 := []int{1, 2, 3}
slice2 := []int{1, 2, 3}
slice3 := []int{4, 5, 6}

fmt.Println(slice1 == slice2)  // true
fmt.Println(slice1 == slice3)  // false
~~~

##### （3.5）指针



##### （3.6）函数封装及其调用



##### （3.7）结构体



##### 

### 5.Go流程语句



### 6.Go异常处理



### 7.包结构及第三方包导入



### 8.命名规范及注意事项

以下是一个 Golang 项目的典型目录和文件结构实例，展示了包命名和文件命名的规范和实践。

~~~shell
/myapp
├── cmd
│   ├── server
│   │   ├── main.go
│   │   ├── server_test.go
│   └── client
│       ├── main.go
│       └── client_test.go
├── pkg
│   ├── http
│   │   ├── server.go
│   │   ├── server_test.go
│   │   ├── client.go
│   │   └── client_test.go
│   ├── database
│   │   ├── connection.go
│   │   └── connection_test.go
│   └── utils
│       ├── converter.go
│       └── converter_test.go
└── internal
    ├── config
    │   ├── config.go
    │   └── config_test.go
    └── model
        ├── user.go
        └── user_test.go

~~~

在这个项目中，我们遵循以下几点：

1. `cmd` 目录用于存放应用程序的主要入口（main）代码。我们可以看到两个子目录：`server` 和 `client`，各自包含自己的 `main.go` 文件，这是 Go 项目的常见结构，允许我们构建多个独立的可执行文件。
2. `pkg` 目录用于存放可以被其他应用程序重用的代码。在这个示例中，我们有 `http`、`database` 和 `utils` 三个包，各自包含相关的 Go 代码。
3. `internal` 目录用于存放只能被当前应用程序使用的代码。在这里，我们有 `config` 和 `model` 两个包。
4. 文件名全部小写，使用下划线（_）分隔单词，如 `server.go` 和 `connection_test.go`。
5. 测试文件以 `_test.go` 结尾，对应于它们要测试的文件，如 `server_test.go` 和 `config_test.go`。

~~~shell
# 在上述目录结构中，根据默认约定，每个文件的 "包名" 应该与 "其所在的文件夹名" 对应。

针对给出的目录结构，可以推断以下文件的包名：
main.go 文件所在的文件夹是 server，因此包名应为 server。
server_test.go 文件所在的文件夹是 server，因此包名应为 server。
main.go 文件所在的文件夹是 client，因此包名应为 client。
client_test.go 文件所在的文件夹是 client，因此包名应为 client。
~~~







