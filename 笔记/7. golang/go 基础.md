## 一、诞生背景与设计理念



Go 程序的入口是 `main` 函数，这个函数位于 `main` 包中。Go 程序必须有一个 `main` 包和 `main` 函数来启动执行， `main` 包中的 `main` 函数会在程序运行时自动执行。

`package main`：每个 Go 程序都是由一个包（`package`）组成的，`main` 包表示这个程序是一个可执行程序。其他包主要是库，不会直接运行。

`func main() {}`：定义了程序的入口点，这个函数没有参数也没有返回值。Go 程序的执行从 `main` 函数开始。

### 1. 诞生背景

Go语言（又称Golang）于2007年由Google的Robert Griesemer、Rob Pike和Ken Thompson开始设计，2009年正式发布。它的诞生源于对现有编程语言的不满和现代软件开发需求的变化。

1. **现有语言的局限性**

- **C语言的问题：**

  - 适合单机时代，但在互联网时代显得过于底层

  - 缺乏现代语言特性（如垃圾回收、并发支持）

  - 内存管理复杂，容易出现内存泄漏和安全问题

- **C++的困境：**

  - 语言特性过于复杂，学习曲线陡峭

  - 开发效率低下，常常"花了100%的精力，却只有60%的开发效率"

  - 编译时间长，影响开发体验

  - 模板系统复杂，错误信息难以理解

- **Java和C#的问题：**

  - 虽然解决了内存管理问题，但继承了C++的复杂哲学

  - 运行时依赖重，部署复杂

  - 在多核时代，并发编程模型不够优雅

  - 无法充分利用现代硬件的多核特性

2. **现代软件开发的新需求**

软件开发面临新的挑战：

- **大规模分布式系统**：需要处理海量并发请求
- **多核处理器普及**：需要充分利用多核性能
- **快速迭代需求**：需要高效的开发和部署流程
- **云原生架构**：需要轻量级、易部署的应用



### 2. 设计理念

Go语言是**静态类型语言**，变量的类型在编译时确定，他的设计遵循以下核心理念：

1. **简洁性（Simplicity）**

Go语言的简洁性体现在语法设计上，Go语言只有25个关键字，相比之下，C++有超过80个关键字、Java有50+个关键字、Python有35个关键字。同时Go语言不支持类继承，而是通过组合和接口实现代码复用

2. **高效性（Efficiency）**

Go编译为机器码，性能接近C/C++

3. **并发支持（Concurrency）**

Go内置的并发原语



### 3. 垃圾回收机制

1. 自动内存管理

Go语言内置垃圾回收器，开发者无需手动管理内存：

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func createLargeSlice() []int {
    // 创建大型切片，内存会自动管理
    slice := make([]int, 1000000)
    for i := range slice {
        slice[i] = i
    }
    return slice
}

func memoryUsageDemo() {
    var m runtime.MemStats
    
    // 获取内存使用情况
    runtime.ReadMemStats(&m)
    fmt.Printf("初始内存使用: %d KB\n", m.Alloc/1024)
    
    // 创建大量对象
    for i := 0; i < 10; i++ {
        data := createLargeSlice()
        _ = data // 使用变量避免编译器优化
    }
    
    runtime.ReadMemStats(&m)
    fmt.Printf("分配后内存使用: %d KB\n", m.Alloc/1024)
    
    // 强制垃圾回收
    runtime.GC()
    runtime.ReadMemStats(&m)
    fmt.Printf("GC后内存使用: %d KB\n", m.Alloc/1024)
}

func main() {
    memoryUsageDemo()
    
    // 展示GC统计信息
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("GC次数: %d\n", m.NumGC)
    fmt.Printf("GC总暂停时间: %v\n", time.Duration(m.PauseTotalNs))
}
```

2. 低延迟GC

Go的垃圾回收器经过优化，具有以下特点：

- **并发收集**：GC与程序并行运行
- **低延迟**：暂停时间通常小于1ms
- **适应性**：根据程序行为自动调整



## 二、数据类型和声明

### 1. 基本数据类型

**整形**

- int：平台相关（32位或64位系统上为32位或64位）、int8、int16、int32、int64

- uint：平台相关（32位或64位系统上为32位或64位）、uint8、uint16、uint32、uint64

**字符串**：支持索引（返回byte）和切片操作，但不能直接修改。

- string：不可变的字节序列，通常存储UTF-8编码的文本。

**布尔型**：bool 表示真（true）或假（false）。

**浮点型**

float：平台相关（32位或64位系统上为32位或64位）

float32：32位浮点数，约6-7位有效数字。

float64：64位浮点数，约15-16位有效数字，Go中默认浮点类型。

**复数型**

complex64：由两个float32组成（实部和虚部）。

complex128：由两个float64组成，Go中默认复数类型。



byte：uint8的别名，常用于字节操作。

rune：int32的别名，表示Unicode码点，常用于字符处理。



### 2. 复合数据类型

**数组**：固定长度、相同类型的元素序列。长度是数组类型的一部分，定义后不可变。

**切片**：动态长度、可变数组，基于数组的引用类型。包含指针、长度（len）和容量（cap）。

**映射**：键值对集合，类似Python的字典。

**结构体**：自定义类型，聚合多个字段。

**指针**：存储变量的内存地址。

**接口**：定义方法集合，类型通过实现方法隐式满足接口。

**函数类型**：函数可以作为类型，用于回调或高阶函数。

**Channel**：用于goroutine间通信的类型。



### 3. 值类型和引用类型

**值类型的特点**：

- 变量直接存储数据本身
- 赋值时会**创建完全独立的副本**
- **修改副本不会影响原始数据**

**Go语言中的值类型包括**：

```go
// 基本数据类型
var a int = 10
var b float32 = 3.14
// 虽然string在Go中表现为值类型，但其底层实现类似引用类型（包含指向字符串数据的指针和长度）。不过由于字符串是不可变的，所以在使用上表现为值类型。
var c string = "hello"	
var d bool = true
var e byte = 'A'
var f rune = '中'

// 数组
var arr1 [3]int = [3]int{1, 2, 3}
var arr2 = arr1 // 完全复制，arr2 的修改不影响 arr1

// 结构体
type Person struct {
    Name string
    Age  int
}
var p1 = Person{Name: "Alice", Age: 30}
var p2 = p1 // 完全复制
```



**引用类型的特点**：

- 变量存储的是指向底层数据的引用
- 赋值时复制的是引用，不是数据本身
- 多个变量可能指向同一个底层数据

```go
// 切片（Slice）
var slice1 []int = []int{1, 2, 3}
var slice2 = slice1 // 共享底层数组
slice2[0] = 100     // 会影响 slice1

// 映射（Map）
var map1 map[string]int = map[string]int{"a": 1, "b": 2}
var map2 = map1     // 共享底层数据
map2["a"] = 100     // 会影响 map1

// 通道（Channel）
var ch1 chan int = make(chan int)
var ch2 = ch1       // 共享同一个通道

// 函数
var fn1 func() = func() { fmt.Println("hello") }
var fn2 = fn1       // 共享同一个函数

// 接口（Interface）
var i1 interface{} = "hello"
var i2 = i1         // 共享底层数据

// 指针
var x int = 10
var ptr1 *int = &x
var ptr2 = ptr1     // 共享指向同一个内存地址
```



接口类型在赋值时会复制底层的值和类型信息，但如果底层值是引用类型，那么引用会被共享。



### 3. 声明变量

1. **var**

```go
func main() {
	var x int
	var y int = 18
	var z = 18
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("z: %v\n", z)
}
```

2. **短变量声明**

短变量声明是Go中最常用的声明方式，仅限于函数内部，通过 `:=` 自动推导变量类型。

限制：

- 只能在函数内部使用（不能用于包级别变量）。
- 至少声明一个新变量（不能用于已声明的变量，除非与新变量一起声明）。

```go
package main

import "fmt"

name := "ld" // 不合法，短变量不能声明全局对象
func main() {
    x := 10	// 合法，自动推导为 int
    x, y := 20, 30 // 合法，y是新变量，x被重新赋值
    fmt.Println(x, y) // 输出: 20 30
    // x := 40 // 错误：x已声明，不能用 :=
}
```

3. 声明**多个变量**

```go
func main() {
	var x, y, z int = 0, 1, 2
	var (
		name string = "ld"
		age = 18
	)

	fmt.Println("Initial values:", x, y, z)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
}
```

4. **匿名变量 `_`**

匿名变量不占用命名空间，不会分配内存

```go
import "fmt"

func test() (string, int) {
	return "ld", 18
}

func main() {
	username, _ := test()
	fmt.Println("Username:", username)
	// fmt.Println("Age:", age)
}
```

5. 常量 **const**

Go中的常量使用 const 关键字声明，值在编译时确定且不可更改。

```go
func main() {
	const (
		a int = 1
		b
		c
	)
	fmt.Println(a, b, c) // 如果是 const 定义多个字段，后续字段不写的话，默认是第一个值
}
// 输出
1, 1, 1
```

`iota` 只能在 const 中使用，iota 会在每行 + 1，并赋值给该行没有初始值的常量

```go
package main

import "fmt"

func main() {
	const (
		a = iota
		b = 100
		c = iota
		d
	)
	fmt.Println(a, b, c, d)
}

// 输出
0 100 2 3
```

6. **默认初始值**

- int, float32, float64：0

- bool：false

- string：""（空字符串）

- 指针、切片、映射、通道、接口：nil

- 结构体：字段均为零值



### 4. 格式化输出

格式化动词以 % 开头，用于指定数据的输出格式。以下是一些常见的格式化动词：

| 动词  | 描述                     | 示例代码                  | 输出     |
| ----- | ------------------------ | ------------------------- | -------- |
| `%v`  | 默认格式（值的默认表示） | fmt.Printf("%v", 42)      | 42       |
| `%#v` | 确切结构和类型信息       | fmt.Printf("%s", "hello") | "hello"  |
| `%T`  | 值的类型                 | fmt.Printf("%T", 42)      | int      |
| `%s`  | 字符串                   | fmt.Printf("%s", "hello") | hello    |
| `%f`  | 浮点数（默认6位小数）    | fmt.Printf("%f", 3.14159) | 3.141590 |
| `%b`  | 二进制                   | fmt.Printf("%b", 42)      | 101010   |
| `%d`  | 十进制整数               | fmt.Printf("%d", 42)      | 42       |
| `%x`  | 十六进制（小写）         | fmt.Printf("%x", 42)      | 2a       |

格式化输出函数：

`fmt.Printf`：格式化并打印到标准输出（控制台）。

`fmt.Sprintf`：格式化并返回格式化后的字符串。

`fmt.Fprintf`：格式化并写入指定的 io.Writer（如文件或网络连接）。



### 5. golang 的 **`包`**

包是 Go 中一组相关源文件的集合，每个 Go 文件都必须声明所属的包。包名通常与文件所在目录名一致。

- 每个 Go 源文件必须以 package 关键字开头，声明所属包名。

- **一个文件夹**下面直接包含的文件**只能归属一个 package**，同样一个 package 的文件不能在多个文件夹下。

- 包名可以不和文件夹的名字一样，包名不能包含－符号。**包名为main的包为应用程序的入口包**，这种包编译后会得到一个可执行文件，而**编译不包含main包的源代码则不会得到可执行文件**

1. 导入包

```go
// 1. 为包指定别名，避免命名冲突或简化代码
import (
    f "fmt" // 别名为 f
)

// 将包的标识符直接引入当前命名空间（不推荐，易导致命名冲突）
import . "fmt" // 直接使用 Println 而非 fmt.Println

// 仅执行包的初始化逻辑，不使用其内容（常用于触发 init 函数）。匿名导入的包与其他方式导入的包一样都会被编译到可执行文件中。
import _ "database/sql" // 仅执行初始化
```

2. 使用第三方的包

```go
1. go mod init 项目名称

2. 引入第三方包

3. go mod tidy
// 超时请使用代理
$env:http_proxy = "socks5://127.0.0.1:7890"
$env:https_proxy = "socks5://127.0.0.1:7890"
```

3. 可见性

在 Go 语言中，**字段**、**函数**或**方法**的可见性由其名称的首字母大小写决定

**大写开头**（如 Name）：表示字段或方法是导出的（exported），即包外的代码可以访问（类似 public）。

**小写开头**（如 name）：表示字段或方法是未导出的（unexported），即只能在定义该字段的包内访问（类似 private）。







## 三、常见操作

### 1. 字符串操作

字符串是 string 类型的变量，底层是一个只读的字节切片（[]byte），但**不可直接修改**。

1. **获取长度**

`len(str)`：返回字符串的**字节长度**（不是字符数）。

```go
func main() {
	str1 := "hello"
	str2 := "hello我"
	fmt.Println(len(str1), len(str2)) // 输出 5 8
}
```

2. **访问字符**

**byte**（uint8 的别名）：表示字符串中的单个字节，通常用于 ASCII 字符。

**rune**（int32 的别名）：表示完整的 Unicode 字符。

字节访问：通过索引（如 str[0]）访问单个字节，返回 byte 类型。

字符（rune）访问：将字符串转为 []rune 切片，访问 Unicode 字符。

```go
str := "Hello, 世界"
fmt.Println(str[0])           // 输出：72（'H' 的 ASCII 值）
runes := []rune(str)
fmt.Println(runes[7])        // 输出：19990（'界' 的 Unicode 码点）
```

3. **字符串拼接**

使用 + 运算符拼接字符串，但性能较低（因创建新字符串）。

更高效的方式是使用 `strings.Builder` 或 `bytes.Buffer`。

```go
str1 := "Hello"
str2 := "World"
result := str1 + ", " + str2 // 低效拼接
fmt.Println(result)          // 输出：Hello, World

var builder strings.Builder
builder.WriteString("Hello")
builder.WriteString(", ")
builder.WriteString("World")
result := builder.String()
fmt.Println(result) // Hello, World
```

4. **`strings` 包常用函数**

`strings.Contains(s, substr string) bool`：检查字符串 s 是否包含子串 substr。

`strings.ContainsRune(s string, r rune) bool`：检查是否包含特定 Unicode 字符。

`strings.Index(s, substr string) int`：返回子串首次出现的索引，未找到返回 -1。

`strings.LastIndex(s, substr string) int`：返回子串最后出现的索引。

`strings.Split(s, sep string) []string`：按分隔符分割字符串。

`strings.Join(a []string, sep string) string`：将字符串切片用分隔符连接。

```go
str := "Hello, World"
fmt.Println(strings.Contains(str, "World")) // true
fmt.Println(strings.Index(str, "o"))        // 4
fmt.Println(strings.LastIndex(str, "o"))    // 8
```

5. **修改字符串**

- 将字符串转为 []byte 切片，修改后再转回字符串。
- 将字符串转为 []rune 切片，修改后再转回字符串。

```go
str := "Hello"
bytes := []byte(str)
bytes[0] = 'h'
fmt.Println(string(bytes)) // 输出：hello

str := "Hello, 你好"
tmp := []rune(str)
tmp[7] = '我'
fmt.Println(string(tmp)) // 输出：Hello, 我好
```

6. **字符串类型转换**

字符串到其他类型的转换：

1. `ParseInt(s string, base int, bitSize int) (int64, error)` ：将字符串解析为指定进制的有符号整数。

2. `ParseUint(s string, base int, bitSize int) (uint64, error)`
3. `ParseFloat(s string, bitSize int) (float64, error)`
4. `ParseBool(s string) (bool, error)`

 其他类型到字符串的转换

1. `FormatInt(i int64, base int) string`：将有符号整数格式化为指定进制的字符串。

2. `FormatUint(i uint64, base int) string`

3. `FormatFloat(f float64, fmt byte, prec int, bitSize int) string`

4. `FormatBool(b bool) string`



### 2. 数组

固定长度、相同类型的元素序列。长度是数组类型的一部分，定义后**数组的大小不可变**。

1. **定义数组**

```go
var arr [5]int  // 声明一个长度为5的整型数组
var arr = [3]int{1, 2, 3}  // 声明一个长度为3的数组，并初始化
var arr = [...]int{1, 2, 3, 4}  // 声明一个数组，长度由元素数量自动推导
arr := [3]int{1, 2, 3}  // 使用简短声明符声明并初始化数组

var arr = [3][2]string{
    {"北京", "上海"},
    {"广州", "深圳"},
    {"成都", "重庆"},
}
fmt.Printf("arr: %v\n", arr)
```



### 3. **`切片`**

切片是一个动态数组，可以增长和缩小。与数组不同，**切片的长度是动态的**，可以在运行时改变。切片是动态的，当添加元素超过当前容量时，Go 会自动扩容。（通过 append 添加时，容量不足会扩容，扩容策略最开始是上次容量的2倍）

1. **定义切片**

切片的定义类似数组，只不过不写大小了

```go
// 定义一个切片，没有初始化，是nil
var s []int

// 直接创建并初始化
s := []int{1, 2, 3, 4}

// 通过数组初始化
s := []int{1, 2, 3, 4}

// 通过make，创建长度为 5，容量为 10 的切片
s := make([]int, 5, 10)  
```

2. **添加一个或多个元素**

```go
s := []int{1, 2, 3}
s = append(s, 4, 5, 6)  // 同时添加多个元素
fmt.Println(s)  // 输出 [1 2 3 4 5 6]

s1 := []int{1, 2}
s2 := []int{3, 4, 5}
s1 = append(s1, s2...)  // 用 ... 来解包 s2 中的元素
fmt.Println(s1)  // 输出 [1 2 3 4 5]

s1 := []int{1, 2}
s2 := []int{3, 4, 5}
s3 := []int{6, 7, 8}
s1 = append(s1, append(s2, s3...)...)
fmt.Printf("s1: %v\n", s1)
```

3. **取切片的子集**

通过索引，你可以从切片中取出一个子切片。

```go
s := []int{1, 2, 3, 4, 5}
sub := s[1:4]  // 从下标 1 到下标 3（不包含 4）
fmt.Println(sub)  // 输出 [2 3 4]

sub := s[:3]  // 取前 3 个元素
fmt.Println(sub)  // 输出 [1 2 3]

sub := s[2:]  // 从下标 2 开始，到切片末尾
fmt.Println(sub)  // 输出 [3 4 5]

sub := s[:]	// 从下标0到末尾
fmt.Println(sub)
```

4. **获取长度和容量**

新的切片是通过对**数组或另一个切片**的切片操作（[start:end]）创建的。新的切片的容量取决于另一个切片的底层数组从起始索引到末尾的可用元素数量。

```go
a := []int{1, 2, 3, 4, 5}
sl := a[1:2]
fmt.Printf("sl: %v size：%v cap：%v\n", sl, len(sl), cap(sl))

// 输出
sl: [2] size：1 cap：4
```

5. **深拷贝**

```go
a := []int{1, 2, 3, 4, 5}
b := make([]int, len(a))
copy(b, a)	// dist, src return 拷贝的size
a[0] = 10
fmt.Println(b)
```

6. **排序**

```go
sort.Ints(s []int)	对整数切片按升序排序。
sort.Float64s(s []float64)	对浮点数切片按升序排序。
sort.Strings(s []string)	对字符串切片按字典序排序。
sort.Sort(data Interface)	对实现 sort.Interface 的数据进行排序。
sort.IntSlice(s []int)		将 []int 切片转换为实现 sort.Interface 接口的类型，从而可以利用 sort 包的排序功能
sort.Float64Slice(s []float64)： 
sort.StringSlice(s []string)：
sort.Reverse() 将排序规则反转（如从升序变为降序）。

sort.Sort(sort.Reverse(sort.IntSlice(s)))	// 降序
```

```go
sort.Search(n int, f func(int) bool) int：在已排序的数据中执行二分查找，返回满足 f(i) == true 的最小索引 i。

func main() {
    nums := []int{1, 2, 3, 4, 5}
    x := 3
    index := sort.Search(len(nums), func(i int) bool { return nums[i] >= x })
    if index < len(nums) && nums[index] == x {
        fmt.Printf("Found %d at index %d\n", x, index) // 输出: Found 3 at index 2
    } else {
        fmt.Printf("%d not found\n", x)
    }
}
```



### 4. 字典

map是一种内建的数据结构，用于存储键值对。它类似于其他语言中的哈希表、字典或散列表。`map` 可以通过键高效地查找、删除和更新值。

1. **定义map**

```go
var m map[string]int

m := make(map[string]int)
```

2. **添加或修改值**

```go
m["age"] = 30
m["height"] = 175
fmt.Println(m)  // 输出 map[age:30 height:175]
```

3. **查询键值对**

```go
value := m["age"]
fmt.Println(value)  // 输出 30

---
value, ok := m["weight"]
if ok {
    fmt.Println(value)
} else {
    fmt.Println("weight 不存在")
}
```

4. **删除键值对**

```go
delete(m, "age")
fmt.Println(m)  // 输出 map[height:175]
```

5. **遍历**

```go
m := map[string]int{"age": 30, "height": 175, "weight": 70}
for key, value := range m {
    fmt.Println(key, value)
}
```



## 四、控制结构

特殊之处

1. 不允许省略花括号
2. 花括号 `{` 必须跟 `if` 同一行

```go
if x > 0 
    fmt.Println("Positive")  // 这是错误的，缺少花括号
```

```go
if x > 0 {
    fmt.Println("Positive")	// 可以
}
```

```go
if x > 0 { fmt.Println("Positive") }	// 可以
```



### 1. if

1. 基本语法

`condition` 必须是一个布尔值表达式（`true` 或 `false`）

```go
x := 10
if x > 5 {
    fmt.Println("x 大于 5")
} else {
    fmt.Println("x 小于或等于 5")
}
```

2. 特殊之处

在 `if` 语句的条件表达式前，可以添加初始化语句来声明和初始化变量。这些变量的作用范围仅限于 `if` 语句内部。

```go
if x := 10; x > 5 {
    fmt.Println("x 大于 5")
} else {
    fmt.Println("x 小于或等于 5")
}
```



### 2. switch

1. 基本语法

```go
switch expression {
case value1:
    // expression == value1 时执行
case value2:
    // expression == value2 时执行
default:
    // expression 没有匹配的值时执行
}
```

- `expression` 是你要判断的表达式，可以是常量、变量、函数调用等。

- `case` 后面跟着要匹配的值或表达式。

- `default` 是可选的，当没有任何 `case` 匹配时执行。

2. 特殊之处

**不需要 `break`**： `switch` 语句中，每个 `case` 默认是结束的，不需要显式地写 `break`。

**多重匹配**： `case` 后面可以列出多个值，匹配任意一个值都会执行对应的代码块。

```go
x := 2
switch x {
case 1, 2, 3:
    fmt.Println("x 是 1、2 或 3")
default:
    fmt.Println("x 不是 1、2 或 3")
}
// 输出
x 是 1、2 或 3
```





### 3. 循环

3.1 **`for`**

1. 基本语法

```go
for init; condition; post {
    // 循环体
}
```

3. init，condition，可省略 post

```go
for i := 0; i < 10; {
    fmt.Println(i)
    if i == 5 {
        break // 手动退出
    }
    i++
}
```

4. condition，可省略 init 和 post

```go
x := 0
for x < 5 {
    fmt.Println(x)
    x++
}
```

5. 都可省略

```go
for {
    fmt.Println("这是一个无限循环")
    break // 可以使用 break 来退出循环
}
```



3.2 for **`range`**

用于遍历可迭代对象（如字符串、切片、数组、映射、通道）。

```go
for index, value := range iterable {
    // 循环体
}
```

使用 `for range` 或将字符串转为 `[]rune` 。可以用来处理包含 Unicode 字符的字符串（如中文、表情符号）。

```go
str := "Hello, 世界"
for i, r := range str {
    fmt.Printf("Index: %d, Char: %c\n", i, r)
}
// 输出：
// Index: 0, Char: H
// Index: 1, Char: e
// ...
// Index: 7, Char: 世
// Index: 10, Char: 界
```



### 4. 函数

使用 `func` 关键字定义函数。函数有一个名称、参数列表、返回类型和函数体。

1. 基本用法

```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // 函数体
    return result
}
```

```go
func swap(a int, b int) (int, int) {
    return b, a
}

// 前面的参数类型和最后相同的话，可以简写参数类型
func swap(a, b int) (int, int) {
    return b, a
}
```

2. 可以返回多个值

返回值有 **匿名返回值** 和 **命名返回值**

```go
// 两个int类型的匿名返回值
func calculate(a, b int) (int, int) {
	sum = a + b
	product = a * b
	return sum, product
}

func main() {
	s, p := calculate(3, 5)
	fmt.Println(s, p) // 输出：8 15
}

------
// 两个int类型的命名返回值
func calculate(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return
}

func main() {
	s, p := calculate(3, 5)
	fmt.Println(s, p) // 输出：8 15
}
```

3. 可变参数

```go
func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

---
result := sum(1, 2, 3, 4, 5)
fmt.Println(result)	// 15
```

4. 匿名函数

匿名函数是没有名称的函数，可以直接定义在代码中，通常赋值给变量或作为参数传递。

```go
func (参数列表) 返回值类型 {
    // 函数体
}
------
func main() {
    // 直接定义并调用
    func(x, y int) {
        fmt.Println(x + y)
    }(3, 5) // 输出: 8
}
------
func main() {
    // 赋值给变量
    add := func(x, y int) int {
        return x + y
    }
    result := add(3, 5)
    fmt.Println(result) // 输出: 8
}
```

闭包：闭包是一个函数及其捕获的外部作用域变量的组合。闭包能够“记住”定义时环境的变量，并在**函数执行时访问或修改这些变量**，即使外部作用域已经结束。

变量捕获：

- 匿名函数可以访问其定义时所在作用域的变量（包括外部函数的局部变量）。
- 捕获是按引用进行的，意味着闭包操作的是外部变量的内存地址，而不是值的副本。

捕获的变量不会随着外部函数的结束而销毁，而是与闭包的生命周期绑定，直到闭包不再被引用。

```go
func counter() func() int {
	count := 0          // 外部变量
	return func() int { // 匿名函数形成闭包
		count++ // 捕获并修改 count
		return count
	}
}

func main() {
	c1 := counter()   // 创建第一个闭包
	c2 := counter()   // 创建第二个闭包
	fmt.Println(c1()) // 输出: 1
	fmt.Println(c1()) // 输出: 2
	fmt.Println(c2()) // 输出: 1
}
```



### 5. **`label`**

1. 标签必须紧挨在 `for`、`switch/select` 之前，或者在不违背前面规则的前提下，在函数中任意地方

2. break 配合 **标签** 可以退出指定的外层循环，而不是仅退出最内层结构。

```go
LabelName:
for ... {
    for ... {
        break LabelName // 退出 LabelName 标记的循环
    }
}
```

```go
// 可以理解成最开层的循环叫做 OuterLoop, break OuterLoop; 会跳出最外层循环
OuterLoop:
	for i := range 3 {
		fmt.Printf("Outer Loop: %v\n", i)
		for j := range 3 {
			if i == 1 && j == 1 {
				break OuterLoop // 退出外层循环
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
// 输出
Outer Loop: 0
i=0, j=0
i=0, j=1
i=0, j=2
Outer Loop: 1
i=1, j=0
```

3. continue 配合 **标签** 可以进入指定的外层循环的下次迭代，而不是再进入最内层循环的下次迭代。

```go
OuterLoop:
	for i := range 3 {
		fmt.Printf("Outer Loop: %v\n", i)
		for j := range 3 {
			if i == 1 && j == 1 {
				continue OuterLoop // 退出外层循环
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
// 输出
Outer Loop: 0
i=0, j=0
i=0, j=1
i=0, j=2
Outer Loop: 1
i=1, j=0
Outer Loop: 2
i=2, j=0
i=2, j=1
i=2, j=2
```

4. goto 会直接跳转到 `OuterLoop` 标签所在的位置。

```go
OuterLoop:
	for i := range 3 {
		fmt.Printf("Outer Loop: %v\n", i)
		for j := range 3 {
			if i == 1 && j == 1 {
				goto OuterLoop // 退出外层循环
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
// 输出
会陷入死循环，因为 goto 到标签后，它并不会像continue那样让 i + 1
```





## 五、常用

### 1. **`type `**

type 关键字用于定义新的类型或为现有类型创建别名。

1. **类型别名**

为现有类型（int）创建一个别名，MyInt 和 int 是完全等价的，只是名称不同。MyInt 和 int 是同一个类型，底层完全相同，编译器不会将它们视为不同的类型。

```go
type MyInt = int

func main() {
    var a MyInt = 42
    fmt.Println(a) // 输出: 42
}
// myInt 等价于 int
```

2. **自定义类型**

基于现有类型（int）创建一个新类型，MyInt 是一个独立的类型，尽管底层数据结构与 int 相同。MyInt 和 int 是不同的类型，编译器会严格区分它们。

```go
type MyInt int // 定义新类型 MyInt，基于 int

func main() {
    var a MyInt = 42
    var b int = 10
    // fmt.Println(a + b) // 错误：类型不匹配
    fmt.Println(a + MyInt(b)) // 正确：需要显式转换
}
```

3. **函数**

```go
type Operation func(int, int) int

func Add(a, b int) int {
    return a + b
}

func main() {
    var op Operation = Add
    result := op(3, 4)
    fmt.Println(result) // 输出: 7
}
```



### 2. defer

defer 是一个关键字，用于延迟执行一个函数调用或方法调用，直到包含它的函数返回（无论是正常返回还是发生 panic）。defer 常用于资源清理（如关闭文件、释放锁）或确保某些操作在函数退出时执行。

- defer 后接一个函数调用（可以是命名函数、匿名函数或方法）。

- 延迟的函数调用会在包含它的函数返回之前执行。

1. 基本用法

```go
func main() {
    fmt.Println("开始")
    defer fmt.Println("延迟执行1")
    defer fmt.Println("延迟执行2")
    fmt.Println("结束")
}
// 输出：
// 开始
// 结束
// 延迟执行2
// 延迟执行1
```

2. defer的执行**顺序 - 栈结构（LIFO）**

```go
func deferOrder() {
    for i := 0; i < 3; i++ {
        defer fmt.Printf("defer %d\n", i)
    }
    fmt.Println("函数体执行完毕")
}
// 输出：
// 函数体执行完毕
// defer 2
// defer 1
// defer 0
```

3. 匿名返回值 和 命名返回值

defer无法修改匿名返回值，但可以修改命名返回值

```go
// 匿名返回值
func test(x int) int {
	n := x
	defer func() {
		n *= 2
	}()
	return n
}
func main() {
	res := test(10)
	fmt.Printf("res: %v\n", res)	// 10
}

// 命名返回值
func test(x int) (n int) {
	n = x
	defer func() {
		n *= 2
	}()
	return
}
func main() {
	res := test(10)
	fmt.Printf("res: %v\n", res)	// 20
}
```

对于匿名返回值执行的操作分为两步，1. 将x的值（5）**复制**到**返回值存储区域**、2. 执行return指令。在**复制完成后**，才执行defer函数，defer中修改的是局部变量x，而不是已经复制到返回值区域的值

对于命名返回值在函数开始时的 **n 就是返回值存储区域**，return 时，不需要复制值，n 本身就是，defer函数执行时，修改的就是这个返回值变量本身。



### 3. panic、recover、error

#### 3.1 panic

panic 是一个内置函数，用于触发程序的运行时异常，中断正常执行流程。panic 通常表示程序遇到无法继续运行的严重错误（如数组越界、空指针访问）。

1. 参数可以是任意类型，通常是字符串或错误类型，用于描述异常原因。
2. panic 会导致当前 goroutine 的调用栈展开（unwind），执行所有 defer 语句。如果没有 recover 捕获 panic，程序会崩溃并打印堆栈跟踪。

```go
func main() {
    fmt.Println("Start")
    panic("Something went wrong!") // 触发 panic
    fmt.Println("End") // 不会执行
}

// 输出
Start
panic: Something went wrong!
[堆栈跟踪信息]
```

```go
package main

import "fmt"

func main() {
    defer fmt.Println("Cleanup on panic")
    fmt.Println("Start")
    panic("Critical error")
}

// 输出
Start
Cleanup on panic
panic: Critical error
[堆栈跟踪信息]
```

#### 3.2 recover

recover 是一个内置函数，用于捕获和处理 panic，防止程序崩溃。recover 必须在 defer 函数中调用，因为只有延迟函数才能在 panic 传播时捕获异常。

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    fmt.Println("Start")
    panic("Critical error")
    fmt.Println("End") // 不会执行
}

// 输出
Start
Recovered from: Critical error
```

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if err := recover(); err != nil {
            log.Printf("Panic: %v", err)
            http.Error(w, "Internal Server Error", 500)
        }
    }()
    riskyOperation()
})
// 在服务器程序中捕获 panic，确保服务继续运行。记录 panic 信息以便调试。
```

#### 3.3 error

1. `errors.New` 定义错误返回信息

```go
import (
	"errors"
	"fmt"
)

// 返回两个，如果是error的话 说明分母为0了
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {

	if val, error := divide(10, 0); error == nil {
		fmt.Printf("val: %v\n", val)
	} else {
		fmt.Printf("error: %v\n", error)
	}
}
```

2. `fmt.Errorf()` 定义

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("invalid fenmu: %d", b)
	}
	return a / b, nil
}
```

3. **自定义 error 类型**

```go
type ValidationError struct {
    Field string
    Value interface{}
    Message string
}

// 实现 Error 接口
func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed for field '%s' with value '%v': %s", 
        e.Field, e.Value, e.Message)
}

func validateEmail(email string) error {
    if !strings.Contains(email, "@") {
        return &ValidationError{	// 创建一个ValidationError指针对象返回
            Field: "email",
            Value: email,
            Message: "must contain @ symbol",
        }
    }
    return nil
}

func main() {

	if err := validateEmail("hello"); err != nil {
        fmt.Printf("err: %v\n", err)	// 这里隐式的调用了 err.Error()
	}
}
```



### 4. time 包 和 定时器

1. time 包

```go
func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println("Current time:", now)
	fmt.Printf("now.Year(): %v\n", now.Year())
	fmt.Printf("now.Month(): %v\n", now.Month())
	fmt.Printf("now.Day(): %v\n", now.Day())
	fmt.Printf("now.Hour(): %v\n", now.Hour())
	fmt.Printf("now.Minute(): %v\n", now.Minute())
	fmt.Printf("now.Second(): %v\n", now.Second())

	fmt.Printf("time.Hour: %v\n", time.Hour)     //1 h
	fmt.Printf("time.Minute: %v\n", time.Minute) // 1 m
	fmt.Printf("time.Second: %v\n", time.Second) // 1 s

	future := now.Add(2 * time.Hour) // 在当前时间基础上增加2小时
	duration := future.Sub(now)      // 2 消失
	fmt.Println("Future time:", future)
	fmt.Println("Duration:", duration)                     // 输出: Duration: 2h0m0s
	fmt.Println("Is future after now?", future.After(now)) // 输出: true
}

// 输出
urrent time: 2025-07-02 23:00:02.4874845 +0800 CST m=+0.000000001
now.Year(): 2025
now.Month(): July
now.Day(): 2
now.Hour(): 23
now.Minute(): 0
now.Second(): 2
time.Hour: 1h0m0s
time.Minute: 1m0s
time.Second: 1s
Future time: 2025-07-03 01:00:02.4874845 +0800 CST m=+7200.000000001
Duration: 2h0m0s
Is future after now? true
```

2. 定时器

在 Go 语言中，**定时器（Timer）**和 **周期性触发器（Ticker）**是 time 包提供的两种核心机制，用于处理时间相关的任务。两者都基于 Go 的并发模型，通过通道（chan）与 goroutine 配合使用，提供了高效、简洁的定时功能。

`time.Timer`：用于在指定时间后触发一次事件，常用于超时控制或延迟执行。

`time.Ticker`：用于按固定时间间隔重复触发事件，常用于周期性任务。



### 5. 指针、make 、new

Go 语言中，指针（pointer）是一种重要的特性，用于直接操作内存地址，从而实现高效的数据传递和修改。相比 C/C++ 的指针，Go 限制了一些不安全的操作（如指针算术），同时保留了指针的核心功能。

1. 什么是指针

指针**是一个变量**，**存储**的是**另一个变量的内存地址**。

**定义**指针：`var p *T`（T 是任意类型，如 int、string、结构体等）。

**获取地址**：`&x` 返回变量 x 的内存地址。

**解引用**：`*p` 访问指针 `p` 指向的变量的值。

2. 定义并初始化

```go
func main() {

	var n *int
	fmt.Printf("n: %v\n", n)

	n = new(int)
	fmt.Printf("n: %v\n", n)

	x := 1
	n = &x
	*n = 10

	fmt.Printf("n: %v\n", x)
}
// 输出
n: <nil>
n: 0xc0001060e0
n: 10
```

3. 对于**结构体**，可以省略 **解引用** 的操作



`make` 和 `new`是两个是 golang 种用于内存分配的内置函数

3. `new(T)` 

   分配类型 T 的零值内存，并返回指向该内存的指针 `*T`。

   为任意类型（包括基本类型、结构体、指针等）分配内存。并返回指向该内存的指针。

4. `make(T, args)` 

   用于创建并初始化特定类型的对象，仅适用于切片（slice）、映射（map）和通道（chan）。

   分配内存并初始化内部数据结构（如切片的底层数组、映射的哈希表、通道的缓冲区）。

   返回初始化后的对象（不是指针），类型为 `T`。



## 六、结构体和接口

### 1. 结构体

结构体（struct）是Go语言中用来组合不同类型数据的复合数据类型。它类似于其他语言中的类或记录，但更加简洁和灵活。

#### 1.1 定义、实例化、访问修改

1. **定义结构体**

```go
// 结构体
type Person struct {
    Name string
    Age  int
}
```

2. **实例化**

实例化结构体，有两类方式，分别可以将结构体变量初始化为 **值类型** 和 **指针类型** 变量

```go
// 1. 直接声明
person := Person{
    Name:  "Alice",
    Age:   30,
    Email: "alice@example.com",
}

// 2. 零值初始化
var person Person
// person.Name == "", person.Age == 0, person.Email == ""

上面是值类型的
----------
下面是指针类型的

// 3. 使用 new 关键字
person := new(Person)
// 等价于
// person := &Person{}

// 4. 使用指针直接创建
person := &Person{
    Name:  "Bob",
    Age:   25,
    Email: "bob@example.com",
}
```

3. 访问修改字段

```go
通过点号（.）访问或修改结构体的字段，无论是指针类型还是值类型：
person := Person{Name: "Alice", Age: 30}
fmt.Println(person.Name) // 输出: Alice
person.Age = 31         // 修改 Age
fmt.Println(person.Age)  // 输出: 31

对于指针类型的结构体，Go 提供了语法糖，允许直接使用点号访问字段，无需显式解引用：
person := &Person{Name: "Bob", Age: 25}
fmt.Println(person.Name) // 输出: Bob
person.Age = 26         // 自动解引用，等价于 (*person).Age = 26
```



#### 1.2 匿名字段

```go
type Person struct {
    string // 匿名字段，类型即为字段名
    int    // 匿名字段
}

func main() {
    p := Person{"Alice", 30}
    fmt.Println(p.string) // 访问匿名字段
    fmt.Println(p.int)
    
    // 也可以这样访问
    fmt.Println(p)
}
```



#### 1.3 结构体的方法

1. **方法的定义**

```go
func (接收者变量 接收者类型) 方法名(参数列表) (返回值列表) {
    // 方法体
}
```

在Go语言中，**方法是属于某个类型的函数**。接收者就是用来指定这个方法属于哪个类型的。

普通函数的调用

```go
func Add(a, b int) int {
    return a + b
}

// 调用方式
result := Add(5, 3)
```

结构体方法（有接收者）

```go
type Calculator struct {
    name string
}

func (c Calculator) Add(a, b int) int {
    fmt.Printf("%s is calculating...\n", c.name)
    return a + b
}

// 调用方式
calc := Calculator{name: "MyCalc"}
result := calc.Add(5, 3)  // 注意：是通过实例调用的, 函数内部可以访问或者修改这个示例的字段
```

**示例**

```go
// 定义一个结构体
type Person struct {
    Name string
    Age  int
}

// 值接收者方法
func (p Person) SayHello() {
    fmt.Printf("Hello, I'm %s, %d years old\n", p.Name, p.Age)
}

// 指针接收者方法
func (p *Person) SetAge(age int) {
    p.Age = age
}

// 带返回值的方法
func (p Person) GetInfo() string {
    return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func main() {
    // 创建结构体实例
    person := Person{Name: "Alice", Age: 25}
    
    // 调用方法
    person.SayHello()           // Hello, I'm Alice, 25 years old
    
    person.SetAge(30)          // 修改年龄
    person.SayHello()          // Hello, I'm Alice, 30 years old
    
    info := person.GetInfo()
    fmt.Println(info)          // Name: Alice, Age: 30
}
```

2. **值接收者 vs 指针接收者**

**值接收者（Value Receiver）**：

- 接收结构体的副本
- 方法内部的修改不会影响原始结构体

**指针接收者（Pointer Receiver）**：

- 接收结构体的指针
- 可以修改原始结构体的字段
- 避免大结构体的拷贝开销

3. **方法链**

```go
type Builder struct {
    content string
}

func (b *Builder) AddText(text string) *Builder {
    b.content += text
    return b
}

func (b *Builder) AddNewLine() *Builder {
    b.content += "\n"
    return b
}

func (b *Builder) AddSpace() *Builder {
    b.content += " "
    return b
}

func (b *Builder) Build() string {
    return b.content
}

func main() {
    builder := &Builder{}
    
    result := builder.
        AddText("Hello").
        AddSpace().
        AddText("World").
        AddNewLine().
        AddText("This is a test").
        Build()
    
    fmt.Println(result)
}
```



#### 1.4 结构体嵌套（继承）

```go
type Address struct {
	email  string
	number string
	local  string
}

type Detail struct {
	email string
}

type Person struct {
	name string
	age  int64
	sex  string
	Address		// 匿名结构体，类似继承
	Detail		// 继承
}

func (p Person) eat() {	// Person 结构体的方法
	fmt.Printf("%v is eating\n", p.name)
}

type Man struct {
	Person			// 匿名结构体，类似继承
	isSmoking bool
}

func (p Man) smoke() {		// Man 结构体的方法
	fmt.Printf("%v is Smoking!\n", p.name)
}

func main() {
	man := new(Man)
	man.name = "ld"
	man.age = 18
	man.sex = "man"
	man.Address.email = "lidigolang@gmail.com"
	man.Detail.email = "lige8421@gmail.com"
	man.local = "CN"
	man.number = "16639562930"
	man.isSmoking = false

	fmt.Printf("person: %v\n", man)
	man.eat()
	man.smoke()
}
```



#### 1.5 标签（Tags）

1. 基本标签使用

```go
import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"-"`              // 忽略此字段
    Age      int    `json:"age,omitempty"`  // 如果为零值则忽略
}

func main() {
    user := User{
        ID:       1,
        Name:     "Alice",
        Email:    "alice@example.com",
        Password: "secret123",
        Age:      0,
    }
    
    // 序列化为JSON
    jsonData, err := json.Marshal(user)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Println("JSON:", string(jsonData))
    
    // 反序列化
    var user2 User
    err = json.Unmarshal(jsonData, &user2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Printf("Unmarshaled: %+v\n", user2)
}
```

2. 多种标签

```go
import (
    "encoding/json"
    "encoding/xml"
    "fmt"
)

type Product struct {
    ID          int     `json:"id" xml:"id" db:"product_id"`
    Name        string  `json:"name" xml:"name" db:"product_name"`
    Price       float64 `json:"price" xml:"price" db:"price"`
    Description string  `json:"description,omitempty" xml:"description,omitempty" db:"description"`
}

func main() {
    product := Product{
        ID:          1,
        Name:        "Laptop",
        Price:       999.99,
        Description: "High performance laptop",
    }
    
    // JSON序列化
    jsonData, _ := json.Marshal(product)
    fmt.Println("JSON:", string(jsonData))
    
    // XML序列化
    xmlData, _ := xml.Marshal(product)
    fmt.Println("XML:", string(xmlData))
}
```

3. 自定义标签处理

```go
import (
    "fmt"
    "reflect"
    "strings"
)

type ValidationTag struct {
    Name     string `validate:"required,min=2,max=50"`
    Email    string `validate:"required,email"`
    Age      int    `validate:"min=0,max=120"`
    Password string `validate:"required,min=8"`
}

func ValidateStruct(v interface{}) []string {
    var errors []string
    val := reflect.ValueOf(v)
    typ := reflect.TypeOf(v)
    
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        fieldType := typ.Field(i)
        
        tag := fieldType.Tag.Get("validate")
        if tag == "" {
            continue
        }
        
        rules := strings.Split(tag, ",")
        for _, rule := range rules {
            if rule == "required" && field.Kind() == reflect.String && field.String() == "" {
                errors = append(errors, fmt.Sprintf("%s is required", fieldType.Name))
            }
            if strings.HasPrefix(rule, "min=") && field.Kind() == reflect.String {
                // 简单的长度验证示例
                minLen := 2 // 简化处理
                if len(field.String()) < minLen {
                    errors = append(errors, fmt.Sprintf("%s is too short", fieldType.Name))
                }
            }
        }
    }
    
    return errors
}

func main() {
    user := ValidationTag{
        Name:     "A",
        Email:    "invalid-email",
        Age:      25,
        Password: "123",
    }
    
    errors := ValidateStruct(user)
    if len(errors) > 0 {
        fmt.Println("Validation errors:")
        for _, err := range errors {
            fmt.Println("-", err)
        }
    }
}
```



#### 1.6 结构体的比较和复制

在Go语言中，只有当结构体的所有字段都是**可比较类型**时，结构体才能进行比较。

1. **可比较的类型**

   - 基本数据类型（int、float、string、bool等）

   - 数组（元素类型可比较）

   - 指针
   - 接口（动态类型相同且可比较）

2. **不可比较的类型**

   - 切片（slice）

   - 映射（map）

   - 函数

   - 包含不可比较字段的结构体

```go
type Point struct {
    X, Y int
}

type Person struct {
    Name string
    Age  int
}

type PersonWithSlice struct {
    Name    string
    Age     int
    Hobbies []string // 包含不可比较的字段
}

func main() {
    // 可比较的结构体
    p1 := Point{X: 1, Y: 2}
    p2 := Point{X: 1, Y: 2}
    p3 := Point{X: 2, Y: 3}
    
    fmt.Println("p1 == p2:", p1 == p2) // true，所有字段都相等
    fmt.Println("p1 == p3:", p1 == p3) // false，字段值不同
    
    person1 := Person{Name: "Alice", Age: 30}
    person2 := Person{Name: "Alice", Age: 30}
    fmt.Println("person1 == person2:", person1 == person2) // true
    
    // 包含不可比较字段的结构体无法比较
    // personWithSlice1 := PersonWithSlice{Name: "Alice", Age: 30, Hobbies: []string{"reading"}}
    // personWithSlice2 := PersonWithSlice{Name: "Alice", Age: 30, Hobbies: []string{"reading"}}
    // fmt.Println(personWithSlice1 == personWithSlice2) // 编译错误：invalid operation
}
```

3. **比较规则**

- 两个结构体相等当且仅当它们的**所有对应字段**都相等
- 结构体比较是**逐字段比较**的
- 如果结构体包含不可比较的字段，则整个结构体不可比较

---



Go语言中结构体的赋值默认是**值复制**，但需要区分浅复制和深复制的概念。

1. 浅复制 vs 深复制

   - **浅复制**：复制结构体本身，但结构体中的引用类型字段（如切片、映射）仍然指向原始数据

   - **深复制**：完全复制结构体及其所有引用类型字段，创建完全独立的副本

```go

type Address struct {
	Street string
	City   string
}

type Person struct {
	Name    string
	Age     int
	Address Address  // 值类型，会完全复制
	Friends []string // 引用类型，浅复制时共享底层数组
}

// 浅复制方法
func (p Person) ShallowCopy() Person {
	return p // 直接返回，Go会自动进行值复制, 因为Person中有引用类型 切片，所以是浅拷贝，切片会共享
}

// 深复制方法
func (p Person) DeepCopy() Person {
	// 手动复制切片
	friends := make([]string, len(p.Friends))
	copy(friends, p.Friends)

	return Person{
		Name:    p.Name,
		Age:     p.Age,
		Address: p.Address, // Address是值类型，自动深复制
		Friends: friends,   // 使用新创建的切片
	}
}

func main() {
	original := Person{
		Name:    "Alice",
		Age:     30,
		Address: Address{Street: "123 Main St", City: "New York"},
		Friends: []string{"Bob", "Charlie"},
	}

	// 浅复制演示
	shallow := original.ShallowCopy()
	shallow.Name = "Alice Copy"           // 不影响原始数据
	shallow.Address.Street = "456 Oak St" // 不影响原始数据
	shallow.Friends[0] = "David"          // 影响原始数据！因为共享底层数组

	fmt.Println("=== 浅复制后 ===")
	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Shallow:  %+v\n", shallow)

	// 重新创建原始数据
	original = Person{
		Name:    "Alice",
		Age:     30,
		Address: Address{Street: "123 Main St", City: "New York"},
		Friends: []string{"Bob", "Charlie"},
	}

	// 深复制演示
	deep := original.DeepCopy()
	deep.Name = "Alice Deep Copy"
	deep.Address.Street = "789 Pine St"
	deep.Friends[0] = "Eva" // 不影响原始数据

	fmt.Println("\n=== 深复制后 ===")
	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Deep:     %+v\n", deep)
}
```

```
=== 浅复制后 ===
Original: {Name:Alice Age:30 Address:{Street:123 Main St City:New York} Friends:[David Charlie]}
Shallow:  {Name:Alice Copy Age:30 Address:{Street:456 Oak St City:New York} Friends:[David Charlie]}

=== 深复制后 ===
Original: {Name:Alice Age:30 Address:{Street:123 Main St City:New York} Friends:[Bob Charlie]}
Deep:     {Name:Alice Deep Copy Age:30 Address:{Street:789 Pine St City:New York} Friends:[Eva Charlie]}
```

2. 关键点总结

   - **值类型字段**（如int、string、结构体）在复制时会创建新的副本
   - **引用类型字段**（如切片、映射）在浅复制时共享底层数据
   - **深复制需要手动处理**引用类型字段，或使用JSON等序列化方法

   - **JSON方法简单但性能较低**，适合复杂嵌套结构但不频繁的场景

   - **手动深复制性能更好**，但需要处理每个引用类型字段



#### 1.7 结构体的内存布局

**相同字段的结构体，仅仅是字段定义的顺序不同，就会导致占用的内存大小不同**。Go编译器在**编译阶段**就会查看每个字段的类型和大小，根据目标平台的对齐要求计算布局，最终确定每个字段的偏移量和总大小。

CPU访问内存时，有一个重要的规则：**每种数据类型都有自己的对齐要求**。

1. 对齐要求

   - `bool`, `int8`, `uint8`: 可以放在任何位置

   - `int16`, `uint16`: 必须放在能被2整除的地址上

   - `int32`, `uint32`, `float32`: 必须放在能被4整除的地址上

   - `int64`, `uint64`, `float64`: 必须放在能被8整除的地址上

2. 内存布局对比

```go
package main

import (
    "fmt"
    "unsafe"
)

type BadLayout struct {
    a bool   // 1 byte
    b int64  // 8 bytes
    c bool   // 1 byte
    d int32  // 4 bytes
}

type GoodLayout struct {
    b int64  // 8 bytes
    d int32  // 4 bytes
    a bool   // 1 byte
    c bool   // 1 byte
}

func main() {
    fmt.Printf("BadLayout size: %d bytes\n", unsafe.Sizeof(BadLayout{}))   // 32 bytes
    fmt.Printf("GoodLayout size: %d bytes\n", unsafe.Sizeof(GoodLayout{})) // 16 bytes
    
    // 查看字段在内存中的偏移位置
    bad := BadLayout{}
    fmt.Printf("BadLayout field offsets:\n")
    fmt.Printf("  a: %d\n", unsafe.Offsetof(bad.a))  // 0
    fmt.Printf("  b: %d\n", unsafe.Offsetof(bad.b))  // 8 (需要8字节对齐)
    fmt.Printf("  c: %d\n", unsafe.Offsetof(bad.c))  // 16
    fmt.Printf("  d: %d\n", unsafe.Offsetof(bad.d))  // 20
    
    good := GoodLayout{}
    fmt.Printf("GoodLayout field offsets:\n")
    fmt.Printf("  b: %d\n", unsafe.Offsetof(good.b))  // 0
    fmt.Printf("  d: %d\n", unsafe.Offsetof(good.d))  // 8
    fmt.Printf("  a: %d\n", unsafe.Offsetof(good.a))  // 12
    fmt.Printf("  c: %d\n", unsafe.Offsetof(good.c))  // 13
}
```

3. 内存布局图解

**BadLayout的内存布局**（浪费空间）：

```
偏移  0   1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24...31
     [a] [padding--------------------] [b------------------------] [c] [padding--] [d---------] [padding--]
```

**GoodLayout的内存布局**（紧凑）：

```
偏移  0   1   2   3   4   5   6   7   8   9  10  11  12  13  14  15
     [b------------------------] [d---------] [a] [c] [padding--]
```

4. 为什么会有填充（Padding）？

   - **CPU访问效率**：CPU访问对齐的数据比访问未对齐的数据快得多

   - **硬件要求**：某些CPU架构要求数据必须对齐，否则会出错

   - **缓存行利用**：对齐的数据能更好地利用CPU缓存

5. 优化建议

   - **按大小排序**：将大字段放在前面，小字段放在后面

   - **相同大小分组**：相同大小的字段放在一起

   - **使用工具检查**：可以使用`go vet`或专门的工具检查结构体布局



#### 1.8 空结构体（Empty Struct）

空结构体是一个**不包含任何字段的结构体**，它有特殊的性质和用途。

1. 空结构体的特点

   - **零内存占用**：`unsafe.Sizeof(struct{}{})` 返回 0

   - **相同地址**：所有空结构体实例共享同一个内存地址

   - **不影响对齐**：在结构体中不会影响其他字段的对齐

```go
package main

import (
    "fmt"
    "unsafe"
)

type Empty struct{}

type WithEmpty struct {
    a int
    b Empty
    c int
}

func main() {
    fmt.Printf("Empty struct size: %d bytes\n", unsafe.Sizeof(Empty{}))     // 0 bytes
    fmt.Printf("WithEmpty struct size: %d bytes\n", unsafe.Sizeof(WithEmpty{})) // 16 bytes
    
    // 所有空结构体实例都指向同一个内存地址
    empty1 := Empty{}
    empty2 := Empty{}
    fmt.Printf("empty1 address: %p\n", &empty1)
    fmt.Printf("empty2 address: %p\n", &empty2)
    // 地址相同，因为它们不占用内存
}
```



### 2. 接口

**接口（interface）**是Go语言中定义一组方法签名的抽象类型。接口描述了对象应该做什么，而不是怎么做。

#### 2.1 定义、实现接口

接口的实现与方法的接收者类型有关：

- 值接收者：方法定义为 `func (t T) Method()`，**值和指针**都可以调用。
- 指针接收者：方法定义为 `func (t *T) Method()`，只有**指针**类型可以调用。

```go
package main

import "fmt"

// 1. 定义一个接口
type Speaker interface {
    Speak() string
}

// 2. 定义实现接口的结构体
type Dog struct {
    Name string
}

type Cat struct {
    Name string
}

// 3. Dog实现Speaker接口
func (d Dog) Speak() string {
    return fmt.Sprintf("%s says: Woof!", d.Name)
}

// Cat实现Speaker接口
func (c Cat) Speak() string {
    return fmt.Sprintf("%s says: Meow!", c.Name)
}

func main() {
    var speaker Speaker
    
    dog := Dog{Name: "Buddy"}
    cat := Cat{Name: "Whiskers"}
    
    speaker = dog
    fmt.Println(speaker.Speak())
    
    speaker = cat
    fmt.Println(speaker.Speak())
}
```

#### 2.1 接口的隐式实现

Go语言的接口实现是**隐式的**，不需要显式声明实现关系。只要类型实现了接口的所有方法，就自动满足该接口。

特点：

**1. 鸭子类型（Duck Typing）** Go遵循"如果它走起来像鸭子，叫起来像鸭子，那它就是鸭子"的原则。只要一个类型实现了接口的所有方法，它就自动满足该接口，无需显式声明。

**2. 解耦合程度更高** 在传统OOP语言中，类必须在定义时就知道要实现哪些接口。而Go中，你可以为已存在的类型"事后"定义接口：

1. 隐式实现的优势

```go
type Writer interface {
    Write([]byte) (int, error)
}

type FileWriter struct {
    filename string
}

// 只要实现了Write方法，就自动实现了Writer接口
func (fw FileWriter) Write(data []byte) (int, error) {
    // 写入文件的实现
    fmt.Printf("Writing to file %s: %s\n", fw.filename, string(data))
    return len(data), nil
}

type NetworkWriter struct {
    address string
}

func (nw NetworkWriter) Write(data []byte) (int, error) {
    fmt.Printf("Sending to %s: %s\n", nw.address, string(data))
    return len(data), nil
}

func WriteData(w Writer, data []byte) {
    w.Write(data)
}

func main() {
    fw := FileWriter{filename: "test.txt"}
    nw := NetworkWriter{address: "192.168.1.1"}
    
    WriteData(fw, []byte("Hello File"))
    WriteData(nw, []byte("Hello Network"))
}
```

#### 2.2 空接口（Empty Interface）

空接口`interface{}`不包含任何方法，因此所有类型都实现了空接口。

1. 空接口的使用

```go
func PrintAnything(value interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", value, value)
}

func main() {
    PrintAnything(42)
    PrintAnything("hello")
    PrintAnything([]int{1, 2, 3})
    PrintAnything(struct{Name string}{Name: "Test"})
}
---
Value: 42, Type: int
Value: hello, Type: string
Value: [1 2 3], Type: []int
Value: {Test}, Type: struct { Name string }
```

2. `any` 类型（Go 1.18+）

```go
// Go 1.18引入了any作为interface{}的别名
func ProcessData(data any) {
	switch v := data.(type) {	// 断言，判断 data 的类型，只能用在 switch 语句
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case []int:
		fmt.Printf("Int slice: %v\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
func main() {
	ProcessData(42)
	ProcessData("hello")
	ProcessData([]int{1, 2, 3})
	ProcessData(struct{ Name string }{Name: "Test"})
}
---
Integer: 42
String: hello
Int slice: [1 2 3]
Unknown type: struct { Name string }
```



#### 2.3 类型断言（Type Assertion）

类型断言用于从接口值中提取具体类型的值。

1. 基本语法

```go
// 语法：value, ok := interfaceValue.(Type)
var i interface{} = "hello"

// 安全的类型断言
if str, ok := i.(string); ok {
    fmt.Println("String value:", str)
} else {
    fmt.Println("Not a string")
}

// 不安全的类型断言（可能panic）
str := i.(string)
fmt.Println(str)
```

2. 类型断言的应用

Go 是强类型语言，编译器在编译时只知道 test 是一个**接口类型 Test**，并不知道它的动态类型是 **string**。（有些像 父类引用指向子类对象）

```go
type Test interface{}

func main() {
	var test Test = "hello"

	// 方法1：类型断言
	if str, ok := test.(string); ok {
		fmt.Printf("len: %d\n", len(str))
	}

	// 方法2：直接类型断言（如果确定是字符串）
	fmt.Printf("len: %d\n", len(test.(string)))

	fmt.Printf("test: %T %v\n", test, test)
}

```

#### 2.4 类型开关（Type Switch）

类型开关是类型断言的一种特殊形式，用于处理多种类型。

1. 基本语法

```go
func ProcessValue(value interface{}) {
    switch v := value.(type) {	// // 断言，判断 data 的类型，只能用在 switch 语句
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    case []int:
        fmt.Printf("Int slice with %d elements\n", len(v))
    case nil:
        fmt.Println("Nil value")
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```



#### 2.5 接口组合

Go支持接口组合，可以将多个接口组合成一个新接口。

1. 基本组合

```go
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type ReadWriter interface {
    Reader
    Writer
}

// 等价于
type ReadWriter2 interface {
    Read([]byte) (int, error)
    Write([]byte) (int, error)
}
```

2. 复杂组合示例

这种设计让你可以根据实际需求选择最小化的接口，而不是总是传递整个File对象。这是Go接口组合的优雅之处！

```go
// 基础接口 - 每个都定义一个特定能力
type Reader interface {
    Read([]byte) (int, error)    // 读取能力
}

type Writer interface {
    Write([]byte) (int, error)   // 写入能力
}

type Closer interface {
    Close() error                // 关闭能力
}

type Seeker interface {
    Seek(offset int64, whence int) (int64, error)  // 定位能力
}

// ReadWriteCloser = 读 + 写 + 关闭
type ReadWriteCloser interface {
    Reader   // 嵌入Reader接口
    Writer   // 嵌入Writer接口  
    Closer   // 嵌入Closer接口
}

// ReadWriteSeeker = 读 + 写 + 定位
type ReadWriteSeeker interface {
    Reader
    Writer
    Seeker
}

type File struct {
    name string
}

// File实现了所有4个方法，所以它自动满足所有相关接口
func (f *File) Read(p []byte) (n int, err error) { /*...*/ }
func (f *File) Write(p []byte) (n int, err error) { /*...*/ }
func (f *File) Close() error { /*...*/ }
func (f *File) Seek(offset int64, whence int) (int64, error) { /*...*/ }

func main() {
    file := &File{name: "example.txt"}
    
    // 根据需要，同一个file可以被当作不同的接口使用
    
    // 只需要读写功能时
    var rw ReadWriter = file
    rw.Read([]byte{})
    rw.Write([]byte("hello"))
    
    // 需要读写和关闭功能时
    var rwc ReadWriteCloser = file
    rwc.Read([]byte{})
    rwc.Write([]byte("world"))
    rwc.Close()
    
    // 需要读写和定位功能时
    var rws ReadWriteSeeker = file
    rws.Read([]byte{})
    rws.Write([]byte("test"))
    rws.Seek(0, 0)
}
```



## 七、Go 语言 Goroutines 和 Channel 详解

**进程 vs 线程 vs 协程**

**`进程 (Process)`**：操作系统分配资源的基本单位

- **特点**：拥有独立的内存空间、文件句柄等资源
- **通信**：通过 IPC（进程间通信）如管道、共享内存等
- **开销**：创建和切换成本高

**`线程 (Thread)`**：CPU 调度的基本单位

- **特点**：共享进程的内存空间，但有独立的栈
- **通信**：通过共享内存直接通信
- **开销**：比进程轻量，但仍有一定开销

**`协程 (Coroutine/Goroutine)`**：用户态的轻量级线程

- **特点**：由程序自己调度，不依赖操作系统
- **通信**：通过 Channel 等机制
- **开销**：非常轻量，可以创建大量协程



**并发 vs 并行**

**`并发 (Concurrency)`**：同时处理多个任务的能力

- **特点**：在单核上通过时间片轮转
- **关键**：任务的组织和调度

**`并行 (Parallelism)`**：同时执行多个任务的能力

- **特点**：需要多核处理器
- **关键**：任务的同时执行



### 1. Goroutine

Goroutine 是 Go 语言中的轻量级线程，由 Go 运行时管理。它比操作系统线程更轻量，创建和销毁的成本极低。

特点：

- **轻量级**：初始栈大小只有 2KB，可动态增长。   **高并发**：一个程序可以轻松创建成千上万个 goroutine
- **协作式调度**：由 Go 运行时调度器管理。   **无需手动管理**：自动垃圾回收

#### 1.1 创建 Goroutine

```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    for i := 0; i < 5; i++ {
        fmt.Println("Hello from goroutine:", i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // 启动 goroutine
    go sayHello()
    
    // 主 goroutine 继续执行
    for i := 0; i < 3; i++ {
        fmt.Println("Main goroutine:", i)
        time.Sleep(200 * time.Millisecond)
    }
    
    // 等待 goroutine 完成
    time.Sleep(1 * time.Second)
}
```

#### 1.2 匿名函数 Goroutine

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 使用匿名函数创建 goroutine
    go func() {
        fmt.Println("Anonymous goroutine")
    }()
    
    // 带参数的匿名函数
    go func(name string) {
        fmt.Println("Hello", name)
    }("World")
    
    time.Sleep(1 * time.Second)
}
```

#### 1.3 Goroutine 的调度

Go 运行时使用 M:N 调度模型：

- **M**：操作系统线程
- **N**：Goroutine
- **P**：处理器（逻辑处理器）

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    // 查看当前 GOMAXPROCS 设置
    fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
    
    // 查看当前 goroutine 数量
    fmt.Println("Goroutines:", runtime.NumGoroutine())
    
    // 启动多个 goroutine
    for i := 0; i < 10; i++ {
        go func(id int) {
            fmt.Printf("Goroutine %d\n", id)
        }(i)
    }
    
    fmt.Println("Goroutines after launch:", runtime.NumGoroutine())
    runtime.GC() // 手动触发垃圾回收
}
```





### 2. Channel

Channel 是 Go 语言中用于 goroutine 之间通信的管道。它遵循 "Don't communicate by sharing memory; share memory by communicating" 的设计哲学。

特点：

- **类型安全**：channel 有特定的类型
- **同步通信**：可以用于 goroutine 同步
- **内存安全**：避免竞态条件

#### 2.1 Channel 的创建和基本使用

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 创建一个 int 类型的 channel
    ch := make(chan int)
    
    // 启动一个 goroutine 发送数据
    go func() {
        ch <- 42 // 发送数据到 channel
        fmt.Println("Data sent")
    }()
    
    // 从 channel 接收数据
    value := <-ch
    fmt.Println("Received:", value)
    
    time.Sleep(100 * time.Millisecond)
}
```

#### 2.2 Channel 的方向性

```go
package main

import "fmt"

// 只能发送的 channel
func sender(ch chan<- int) {
    ch <- 1
    ch <- 2
    close(ch)
}

// 只能接收的 channel
func receiver(ch <-chan int) {
    for value := range ch {
        fmt.Println("Received:", value)
    }
}

func main() {
    ch := make(chan int)
    
    go sender(ch)
    receiver(ch)
}
```

#### 2.3 缓冲 Channel

```go
package main

import "fmt"

func main() {
    // 创建缓冲大小为 3 的 channel
    ch := make(chan int, 3)
    
    // 可以发送 3 个值而不阻塞
    ch <- 1
    ch <- 2
    ch <- 3
    
    fmt.Println("Buffer length:", len(ch))
    fmt.Println("Buffer capacity:", cap(ch))
    
    // 接收数据
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

#### 2.4 Channel 的关闭

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 3)
    
    // 发送数据
    ch <- 1
    ch <- 2
    ch <- 3
    
    // 关闭 channel
    close(ch)
    
    // 使用 range 循环读取所有数据
    for value := range ch {
        fmt.Println("Received:", value)
    }
    
    // 检查 channel 是否关闭
    value, ok := <-ch
    if !ok {
        fmt.Println("Channel is closed")
    }
    fmt.Println("Value:", value) // 零值
}
```



### 3. Select

#### 3.1 基本 Select 使用

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "from ch1"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "from ch2"
    }()
    
    // select 会选择第一个准备好的 channel
    select {
    case msg1 := <-ch1:
        fmt.Println("Received:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Received:", msg2)
    }
}
```

#### 3.2 带默认分支的 Select

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "delayed message"
    }()
    
    select {
    case msg := <-ch:
        fmt.Println("Received:", msg)
    default:
        fmt.Println("No message received")
    }
    
    // 等待消息
    time.Sleep(3 * time.Second)
    select {
    case msg := <-ch:
        fmt.Println("Received:", msg)
    default:
        fmt.Println("No message received")
    }
}
```

#### 3.3 超时控制

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "message"
    }()
    
    select {
    case msg := <-ch:
        fmt.Println("Received:", msg)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout!")
    }
}
```



### 4. 常见并发模式

#### 4.1 Worker Pool 模式

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Second) // 模拟工作
        results <- job * 2
    }
}

func main() {
    const numWorkers = 3
    const numJobs = 5
    
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    
    var wg sync.WaitGroup
    
    // 启动 workers
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, jobs, results, &wg)
    }
    
    // 发送任务
    for i := 1; i <= numJobs; i++ {
        jobs <- i
    }
    close(jobs)
    
    // 等待所有 worker 完成
    wg.Wait()
    close(results)
    
    // 收集结果
    for result := range results {
        fmt.Println("Result:", result)
    }
}
```

#### 4.2 生产者-消费者模式

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 5; i++ {
        fmt.Printf("Producing: %d\n", i)
        ch <- i
        time.Sleep(100 * time.Millisecond)
    }
    close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for value := range ch {
        fmt.Printf("Consuming: %d\n", value)
        time.Sleep(200 * time.Millisecond)
    }
}

func main() {
    ch := make(chan int, 2) // 带缓冲的 channel
    var wg sync.WaitGroup
    
    wg.Add(2)
    go producer(ch, &wg)
    go consumer(ch, &wg)
    
    wg.Wait()
}
```

#### 4.3 扇入扇出模式

```go
package main

import (
    "fmt"
    "sync"
)

// 扇出：一个输入，多个输出
func fanOut(input <-chan int, workers int) []<-chan int {
    outputs := make([]<-chan int, workers)
    
    for i := 0; i < workers; i++ {
        output := make(chan int)
        outputs[i] = output
        
        go func(out chan<- int) {
            defer close(out)
            for value := range input {
                out <- value * value // 计算平方
            }
        }(output)
    }
    
    return outputs
}

// 扇入：多个输入，一个输出
func fanIn(inputs []<-chan int) <-chan int {
    output := make(chan int)
    var wg sync.WaitGroup
    
    for _, input := range inputs {
        wg.Add(1)
        go func(in <-chan int) {
            defer wg.Done()
            for value := range in {
                output <- value
            }
        }(input)
    }
    
    go func() {
        wg.Wait()
        close(output)
    }()
    
    return output
}

func main() {
    // 创建输入 channel
    input := make(chan int)
    
    // 启动数据生成器
    go func() {
        defer close(input)
        for i := 1; i <= 5; i++ {
            input <- i
        }
    }()
    
    // 扇出到 3 个 worker
    outputs := fanOut(input, 3)
    
    // 扇入收集结果
    result := fanIn(outputs)
    
    // 打印结果
    for value := range result {
        fmt.Println("Result:", value)
    }
}
```

### 5. 同步原语

Go语言提供了丰富的同步原语来处理并发编程中的数据竞争和协程间通信问题。这些同步原语主要分布在`sync`包中，是构建安全并发程序的基础工具。

#### 5.1 互斥锁（Mutex）

互斥锁是最基本的同步原语，用于保护共享资源，确保同一时间只有一个goroutine可以访问临界区。

1. 基本用法

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Get() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

func main() {
    counter := &Counter{}
    
    // 启动多个goroutine并发访问
    for i := 0; i < 1000; i++ {
        go counter.Increment()
    }
    
    time.Sleep(time.Second)
    fmt.Println("Counter value:", counter.Get())
}
```

2. 注意事项

- 必须成对使用Lock()和Unlock()
- 建议使用defer来确保锁的释放
- 不可重入，同一goroutine重复加锁会死锁
- 零值可用，无需初始化

#### 5.2 读写锁（RWMutex）

读写锁允许多个goroutine同时读取，但写入时需要独占访问。适用于读多写少的场景。

1. 基本用法

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type SafeMap struct {
    mu   sync.RWMutex
    data map[string]int
}

func NewSafeMap() *SafeMap {
    return &SafeMap{
        data: make(map[string]int),
    }
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.mu.RLock()
    defer sm.mu.RUnlock()
    val, ok := sm.data[key]
    return val, ok
}

func (sm *SafeMap) Set(key string, value int) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    sm.data[key] = value
}

func (sm *SafeMap) Delete(key string) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    delete(sm.data, key)
}
```

2. 性能特点

- 读操作可以并发执行
- 写操作需要独占访问
- 适合读多写少的场景
- 读写锁的开销比普通互斥锁稍大

#### 5.3 条件变量（Cond）

条件变量用于goroutine间的等待和通知，通常与互斥锁配合使用。

1. 基本用法

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Queue struct {
    mu    sync.Mutex
    cond  *sync.Cond
    items []int
}

func NewQueue() *Queue {
    q := &Queue{
        items: make([]int, 0),
    }
    q.cond = sync.NewCond(&q.mu)
    return q
}

func (q *Queue) Put(item int) {
    q.mu.Lock()
    defer q.mu.Unlock()
    
    q.items = append(q.items, item)
    q.cond.Signal() // 唤醒一个等待的goroutine
}

func (q *Queue) Get() int {
    q.mu.Lock()
    defer q.mu.Unlock()
    
    for len(q.items) == 0 {
        q.cond.Wait() // 等待条件满足
    }
    
    item := q.items[0]
    q.items = q.items[1:]
    return item
}

func main() {
    queue := NewQueue()
    
    // 消费者goroutine
    go func() {
        for i := 0; i < 5; i++ {
            item := queue.Get()
            fmt.Printf("Got item: %d\n", item)
        }
    }()
    
    // 生产者
    for i := 0; i < 5; i++ {
        queue.Put(i)
        time.Sleep(time.Millisecond * 100)
    }
}
```

2. 关键方法

- `Wait()`: 释放锁并等待条件满足
- `Signal()`: 唤醒一个等待的goroutine
- `Broadcast()`: 唤醒所有等待的goroutine

#### 5.4 等待组（WaitGroup）

WaitGroup用于等待一组goroutine完成执行，是协调多个goroutine的常用工具。

1. 基本用法

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // 标记任务完成
    
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    // 启动5个worker
    for i := 1; i <= 5; i++ {
        wg.Add(1) // 增加计数
        go worker(i, &wg)
    }
    
    wg.Wait() // 等待所有goroutine完成
    fmt.Println("All workers completed")
}
```

2. 注意事项

- Add()必须在goroutine启动前调用
- Done()通常使用defer调用
- Wait()会阻塞直到计数器为0
- 不可重用，完成后需要创建新的WaitGroup

#### 5.5 一次性执行（Once）

Once确保某个操作只执行一次，常用于单例模式或初始化操作。

1. 基本用法

```go
package main

import (
    "fmt"
    "sync"
)

var (
    instance *Singleton
    once     sync.Once
)

type Singleton struct {
    value string
}

func GetInstance() *Singleton {
    once.Do(func() {
        fmt.Println("Creating singleton instance")
        instance = &Singleton{value: "I'm a singleton"}
    })
    return instance
}

func main() {
    var wg sync.WaitGroup
    
    // 并发调用GetInstance
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            s := GetInstance()
            fmt.Printf("Goroutine %d got: %s\n", id, s.value)
        }(i)
    }
    
    wg.Wait()
}
```

2. 使用场景

- 单例模式实现
- 全局配置初始化
- 资源的一次性分配
- 确保某个昂贵操作只执行一次

#### 5.6 内存池（Pool）

Pool用于重用对象，减少内存分配和垃圾回收的开销。

1. 基本用法

```go
package main

import (
    "fmt"
    "sync"
)

type Buffer struct {
    data []byte
}

func (b *Buffer) Reset() {
    b.data = b.data[:0]
}

var bufferPool = sync.Pool{
    New: func() interface{} {
        return &Buffer{
            data: make([]byte, 0, 1024),
        }
    },
}

func processData(input []byte) {
    // 从池中获取buffer
    buffer := bufferPool.Get().(*Buffer)
    defer bufferPool.Put(buffer) // 使用完毕后放回池中
    
    buffer.Reset()
    buffer.data = append(buffer.data, input...)
    
    // 处理数据
    fmt.Printf("Processed %d bytes\n", len(buffer.data))
}

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            data := []byte(fmt.Sprintf("data from goroutine %d", id))
            processData(data)
        }(i)
    }
    
    wg.Wait()
}
```

2. 使用原则

- 对象必须是无状态的或可重置的
- 适用于频繁创建和销毁的对象
- Put()前应该重置对象状态
- 不要依赖Pool中对象的持久性

#### 5.7 原子操作（Atomic）

原子操作提供了对基本数据类型的无锁操作，性能高且线程安全。

1. 基本用法

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

type AtomicCounter struct {
    value int64
}

func (c *AtomicCounter) Increment() {
    atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Get() int64 {
    return atomic.LoadInt64(&c.value)
}

func (c *AtomicCounter) Set(val int64) {
    atomic.StoreInt64(&c.value, val)
}

func main() {
    counter := &AtomicCounter{}
    var wg sync.WaitGroup
    
    // 并发递增
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    fmt.Printf("Final counter value: %d\n", counter.Get())
}
```

2. 支持的操作

- `Add`: 原子加法
- `Load`: 原子读取
- `Store`: 原子写入
- `Swap`: 原子交换
- `CompareAndSwap`: 原子比较并交换

#### 5.8 Map（并发安全的Map）

sync.Map是Go提供的并发安全的映射类型，适用于特定场景。

1. 基本用法

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var m sync.Map
    var wg sync.WaitGroup
    
    // 并发写入
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            key := fmt.Sprintf("key%d", id)
            m.Store(key, id*10)
        }(i)
    }
    
    wg.Wait()
    
    // 遍历所有键值对
    m.Range(func(key, value interface{}) bool {
        fmt.Printf("%s: %v\n", key, value)
        return true
    })
    
    // 读取特定键
    if value, ok := m.Load("key5"); ok {
        fmt.Printf("key5 = %v\n", value)
    }
    
    // 删除键
    m.Delete("key5")
    
    // 加载或存储
    actual, loaded := m.LoadOrStore("key11", 110)
    fmt.Printf("key11 = %v, was loaded = %v\n", actual, loaded)
}
```

2. 适用场景

- 键值对写入一次，读取多次
- 多个goroutine读取、写入和覆盖不同键的条目
- 不适合频繁写入的场景

#### 5.9 同步原语选择指南

1. 性能对比

| 同步原语  | 适用场景          | 性能特点     |
| --------- | ----------------- | ------------ |
| Mutex     | 通用互斥访问      | 中等开销     |
| RWMutex   | 读多写少          | 读操作低开销 |
| Atomic    | 简单数值操作      | 最低开销     |
| Channel   | 通信和同步        | 通信开销     |
| WaitGroup | 等待多个goroutine | 协调开销     |



## 八、反射和读写文件

### 1. 反射（Reflection）

反射是Go语言的一个强大特性，它允许程序在运行时**检查和操作对象的类型和值**。Go的反射机制主要通过`reflect`包来实现，提供了检查类型信息、动态调用方法、操作结构体字段等功能。

#### 1.1 反射基础概念

1. 反射的核心类型

Go反射的核心是两个类型：`reflect.Type`和`reflect.Value`。

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.14
    
    // 获取类型信息
    t := reflect.TypeOf(x)
    fmt.Printf("Type: %v\n", t)        // Type: float64
    fmt.Printf("Kind: %v\n", t.Kind()) // Kind: float64
    
    // 获取值信息
    v := reflect.ValueOf(x)
    fmt.Printf("Value: %v\n", v)           // Value: 3.14
    fmt.Printf("Type: %v\n", v.Type())     // Type: float64
    fmt.Printf("Kind: %v\n", v.Kind())     // Kind: float64
    fmt.Printf("Float: %v\n", v.Float())   // Float: 3.14
}
```

2. Type vs Kind

- **Type**: 具体的类型名称，如`int`、`string`、`MyStruct`
- **Kind**: 类型的底层种类，如`Int`、`String`、`Struct`

```go
package main

import (
    "fmt"
    "reflect"
)

type MyInt int
type Person struct {
    Name string
    Age  int
}

func main() {
    var mi MyInt = 42
    var p Person = Person{Name: "Alice", Age: 30}
    
    // MyInt的Type和Kind
    fmt.Printf("MyInt Type: %v, Kind: %v\n", 
        reflect.TypeOf(mi), reflect.TypeOf(mi).Kind())
    // 输出: MyInt Type: main.MyInt, Kind: int
    
    // Person的Type和Kind
    fmt.Printf("Person Type: %v, Kind: %v\n", 
        reflect.TypeOf(p), reflect.TypeOf(p).Kind())
    // 输出: Person Type: main.Person, Kind: struct
}
```

#### 1.2 基本类型的反射

1. 基本数据类型

```go
package main

import (
    "fmt"
    "reflect"
)

func analyzeValue(x interface{}) {
    v := reflect.ValueOf(x)
    t := reflect.TypeOf(x)
    
    fmt.Printf("Value: %v, Type: %v, Kind: %v\n", v, t, v.Kind())
    
    switch v.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        fmt.Printf("Integer value: %d\n", v.Int())
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        fmt.Printf("Unsigned integer value: %d\n", v.Uint())
    case reflect.Float32, reflect.Float64:
        fmt.Printf("Float value: %f\n", v.Float())
    case reflect.String:
        fmt.Printf("String value: %s\n", v.String())
    case reflect.Bool:
        fmt.Printf("Boolean value: %t\n", v.Bool())
    }
    fmt.Println("---")
}

func main() {
    analyzeValue(42)
    analyzeValue(3.14)
    analyzeValue("hello")
    analyzeValue(true)
    analyzeValue(uint64(100))
}
```

2. 指针类型的反射

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    x := 42
    ptr := &x
    
    v := reflect.ValueOf(ptr)
    t := reflect.TypeOf(ptr)
    
    fmt.Printf("Pointer Type: %v, Kind: %v\n", t, v.Kind())
    
    // 检查是否是指针
    if v.Kind() == reflect.Ptr {
        fmt.Printf("Points to type: %v\n", v.Type().Elem())
        
        // 获取指针指向的值
        elem := v.Elem()
        fmt.Printf("Value pointed to: %v\n", elem.Int())
        
        // 检查是否可以设置值
        if elem.CanSet() {
            elem.SetInt(100)
            fmt.Printf("After setting: %v\n", x)
        }
    }
}
```

#### 1.3结构体反射

1. 结构体字段操作

```go
package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name    string `json:"name" validate:"required"`
    Age     int    `json:"age" validate:"min=0,max=150"`
    Email   string `json:"email" validate:"email"`
    private string // 私有字段
}

func (p Person) GetInfo() string {
    return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p *Person) SetAge(age int) {
    p.Age = age
}

func analyzeStruct(s interface{}) {
    v := reflect.ValueOf(s)
    t := reflect.TypeOf(s)
    
    // 如果是指针，获取其指向的元素
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
        t = t.Elem()
    }
    
    if v.Kind() != reflect.Struct {
        fmt.Println("Not a struct")
        return
    }
    
    fmt.Printf("Struct name: %s\n", t.Name())
    fmt.Printf("Number of fields: %d\n", v.NumField())
    
    // 遍历结构体字段
    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        
        fmt.Printf("Field %d:\n", i)
        fmt.Printf("  Name: %s\n", field.Name)
        fmt.Printf("  Type: %s\n", field.Type)
        fmt.Printf("  Tag: %s\n", field.Tag)
        fmt.Printf("  Value: %v\n", value.Interface())
        fmt.Printf("  CanSet: %v\n", value.CanSet())
        
        // 获取特定标签
        if jsonTag := field.Tag.Get("json"); jsonTag != "" {
            fmt.Printf("  JSON tag: %s\n", jsonTag)
        }
        if validateTag := field.Tag.Get("validate"); validateTag != "" {
            fmt.Printf("  Validate tag: %s\n", validateTag)
        }
        
        fmt.Println()
    }
}

func main() {
    p := Person{
        Name:    "Alice",
        Age:     30,
        Email:   "alice@example.com",
        private: "secret",
    }
    
    analyzeStruct(p)
    
    // 使用指针来修改字段
    fmt.Println("--- Modifying struct fields ---")
    pPtr := &p
    v := reflect.ValueOf(pPtr).Elem()
    
    // 修改Name字段
    nameField := v.FieldByName("Name")
    if nameField.IsValid() && nameField.CanSet() {
        nameField.SetString("Bob")
        fmt.Printf("After setting Name: %s\n", p.Name)
    }
    
    // 修改Age字段
    ageField := v.FieldByName("Age")
    if ageField.IsValid() && ageField.CanSet() {
        ageField.SetInt(35)
        fmt.Printf("After setting Age: %d\n", p.Age)
    }
}
```

2. 结构体方法调用

```go
package main

import (
    "fmt"
    "reflect"
)

type Calculator struct {
    Name string
}

func (c Calculator) Add(a, b int) int {
    return a + b
}

func (c Calculator) Multiply(a, b int) int {
    return a * b
}

func (c *Calculator) SetName(name string) {
    c.Name = name
}

func (c Calculator) GetName() string {
    return c.Name
}

func callMethods(obj interface{}) {
    v := reflect.ValueOf(obj)
    t := reflect.TypeOf(obj)
    
    fmt.Printf("Object type: %s\n", t)
    fmt.Printf("Number of methods: %d\n", v.NumMethod())
    
    // 遍历所有方法
    for i := 0; i < v.NumMethod(); i++ {
        method := t.Method(i)
        fmt.Printf("Method %d: %s\n", i, method.Name)
        fmt.Printf("  Type: %s\n", method.Type)
        fmt.Printf("  Num inputs: %d\n", method.Type.NumIn())
        fmt.Printf("  Num outputs: %d\n", method.Type.NumOut())
        fmt.Println()
    }
    
    // 调用特定方法
    if method := v.MethodByName("Add"); method.IsValid() {
        args := []reflect.Value{
            reflect.ValueOf(10),
            reflect.ValueOf(20),
        }
        results := method.Call(args)
        fmt.Printf("Add(10, 20) = %v\n", results[0].Int())
    }
    
    if method := v.MethodByName("GetName"); method.IsValid() {
        results := method.Call(nil)
        fmt.Printf("GetName() = %v\n", results[0].String())
    }
}

func main() {
    calc := Calculator{Name: "MyCalculator"}
    
    // 调用值方法
    fmt.Println("--- Value methods ---")
    callMethods(calc)
    
    // 调用指针方法
    fmt.Println("--- Pointer methods ---")
    callMethods(&calc)
    
    // 动态调用SetName方法
    v := reflect.ValueOf(&calc)
    if method := v.MethodByName("SetName"); method.IsValid() {
        args := []reflect.Value{reflect.ValueOf("NewCalculator")}
        method.Call(args)
        fmt.Printf("After SetName: %s\n", calc.Name)
    }
}
```

#### 1.4 切片和数组反射

1. 切片操作

```go
package main

import (
    "fmt"
    "reflect"
)

func analyzeSlice(s interface{}) {
    v := reflect.ValueOf(s)
    t := reflect.TypeOf(s)
    
    if v.Kind() != reflect.Slice {
        fmt.Println("Not a slice")
        return
    }
    
    fmt.Printf("Slice type: %s\n", t)
    fmt.Printf("Element type: %s\n", t.Elem())
    fmt.Printf("Length: %d\n", v.Len())
    fmt.Printf("Capacity: %d\n", v.Cap())
    
    // 遍历切片元素
    for i := 0; i < v.Len(); i++ {
        elem := v.Index(i)
        fmt.Printf("Element %d: %v (type: %s)\n", i, elem.Interface(), elem.Type())
    }
    
    // 修改切片元素
    if v.Len() > 0 && v.Index(0).CanSet() {
        switch v.Index(0).Kind() {
        case reflect.Int:
            v.Index(0).SetInt(999)
        case reflect.String:
            v.Index(0).SetString("modified")
        }
    }
}

func createSlice(elementType reflect.Type, length int) interface{} {
    // 动态创建切片
    sliceType := reflect.SliceOf(elementType)
    slice := reflect.MakeSlice(sliceType, length, length)
    
    // 填充初始值
    for i := 0; i < length; i++ {
        elem := slice.Index(i)
        switch elementType.Kind() {
        case reflect.Int:
            elem.SetInt(int64(i * 10))
        case reflect.String:
            elem.SetString(fmt.Sprintf("item_%d", i))
        }
    }
    
    return slice.Interface()
}

func main() {
    // 分析整数切片
    intSlice := []int{1, 2, 3, 4, 5}
    fmt.Println("--- Integer slice ---")
    analyzeSlice(intSlice)
    fmt.Printf("After modification: %v\n", intSlice)
    
    // 分析字符串切片
    stringSlice := []string{"hello", "world", "golang"}
    fmt.Println("\n--- String slice ---")
    analyzeSlice(stringSlice)
    fmt.Printf("After modification: %v\n", stringSlice)
    
    // 动态创建切片
    fmt.Println("\n--- Dynamically created slices ---")
    dynamicIntSlice := createSlice(reflect.TypeOf(int(0)), 5)
    fmt.Printf("Dynamic int slice: %v\n", dynamicIntSlice)
    
    dynamicStringSlice := createSlice(reflect.TypeOf(string("")), 3)
    fmt.Printf("Dynamic string slice: %v\n", dynamicStringSlice)
}
```

2. 数组操作

```go
package main

import (
    "fmt"
    "reflect"
)

func analyzeArray(a interface{}) {
    v := reflect.ValueOf(a)
    t := reflect.TypeOf(a)
    
    if v.Kind() != reflect.Array {
        fmt.Println("Not an array")
        return
    }
    
    fmt.Printf("Array type: %s\n", t)
    fmt.Printf("Element type: %s\n", t.Elem())
    fmt.Printf("Length: %d\n", v.Len())
    
    // 遍历数组元素
    for i := 0; i < v.Len(); i++ {
        elem := v.Index(i)
        fmt.Printf("Element %d: %v\n", i, elem.Interface())
    }
}

func main() {
    // 固定大小数组
    intArray := [5]int{1, 2, 3, 4, 5}
    fmt.Println("--- Integer array ---")
    analyzeArray(intArray)
    
    // 字符串数组
    stringArray := [3]string{"go", "java", "python"}
    fmt.Println("\n--- String array ---")
    analyzeArray(stringArray)
    
    // 修改数组元素（需要传递指针）
    v := reflect.ValueOf(&intArray).Elem()
    if v.Index(0).CanSet() {
        v.Index(0).SetInt(100)
        fmt.Printf("After modification: %v\n", intArray)
    }
}
```

#### 1.5 Map反射

1. Map操作

```go
package main

import (
    "fmt"
    "reflect"
)

func analyzeMap(m interface{}) {
    v := reflect.ValueOf(m)
    t := reflect.TypeOf(m)
    
    if v.Kind() != reflect.Map {
        fmt.Println("Not a map")
        return
    }
    
    fmt.Printf("Map type: %s\n", t)
    fmt.Printf("Key type: %s\n", t.Key())
    fmt.Printf("Value type: %s\n", t.Elem())
    fmt.Printf("Length: %d\n", v.Len())
    
    // 遍历Map
    keys := v.MapKeys()
    for _, key := range keys {
        value := v.MapIndex(key)
        fmt.Printf("Key: %v, Value: %v\n", key.Interface(), value.Interface())
    }
    
    // 设置新的键值对
    if v.Type().Key().Kind() == reflect.String && v.Type().Elem().Kind() == reflect.Int {
        newKey := reflect.ValueOf("new_key")
        newValue := reflect.ValueOf(999)
        v.SetMapIndex(newKey, newValue)
        fmt.Println("Added new key-value pair")
    }
    
    // 删除键值对
    if len(keys) > 0 {
        v.SetMapIndex(keys[0], reflect.Value{}) // 传递零值删除
        fmt.Println("Deleted first key-value pair")
    }
}

func createMap(keyType, valueType reflect.Type) interface{} {
    mapType := reflect.MapOf(keyType, valueType)
    mapValue := reflect.MakeMap(mapType)
    
    // 添加一些示例数据
    if keyType.Kind() == reflect.String && valueType.Kind() == reflect.Int {
        mapValue.SetMapIndex(reflect.ValueOf("one"), reflect.ValueOf(1))
        mapValue.SetMapIndex(reflect.ValueOf("two"), reflect.ValueOf(2))
        mapValue.SetMapIndex(reflect.ValueOf("three"), reflect.ValueOf(3))
    }
    
    return mapValue.Interface()
}

func main() {
    // 分析现有Map
    m := map[string]int{
        "apple":  10,
        "banana": 20,
        "orange": 30,
    }
    
    fmt.Println("--- Original map ---")
    analyzeMap(m)
    fmt.Printf("After operations: %v\n", m)
    
    // 动态创建Map
    fmt.Println("\n--- Dynamically created map ---")
    dynamicMap := createMap(reflect.TypeOf(string("")), reflect.TypeOf(int(0)))
    fmt.Printf("Dynamic map: %v\n", dynamicMap)
    
    // 复杂Map
    complexMap := map[int][]string{
        1: {"a", "b"},
        2: {"c", "d"},
    }
    fmt.Println("\n--- Complex map ---")
    analyzeMap(complexMap)
}
```

#### 1.6 接口反射

1. 接口类型检查

```go
package main

import (
    "fmt"
    "reflect"
)

type Speaker interface {
    Speak() string
}

type Writer interface {
    Write(string) error
}

type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello, I'm " + p.Name
}

func (p Person) Write(text string) error {
    fmt.Printf("%s writes: %s\n", p.Name, text)
    return nil
}

func analyzeInterface(i interface{}) {
    v := reflect.ValueOf(i)
    t := reflect.TypeOf(i)
    
    fmt.Printf("Value: %v\n", v)
    fmt.Printf("Type: %s\n", t)
    fmt.Printf("Kind: %s\n", v.Kind())
    
    // 检查是否实现了特定接口
    speakerType := reflect.TypeOf((*Speaker)(nil)).Elem()
    writerType := reflect.TypeOf((*Writer)(nil)).Elem()
    
    fmt.Printf("Implements Speaker: %v\n", t.Implements(speakerType))
    fmt.Printf("Implements Writer: %v\n", t.Implements(writerType))
    
    // 如果是接口类型，获取底层具体类型
    if v.Kind() == reflect.Interface {
        concrete := v.Elem()
        fmt.Printf("Concrete type: %s\n", concrete.Type())
        fmt.Printf("Concrete value: %v\n", concrete.Interface())
    }
}

func main() {
    person := Person{Name: "Alice"}
    
    fmt.Println("--- Analyzing concrete type ---")
    analyzeInterface(person)
    
    fmt.Println("\n--- Analyzing interface type ---")
    var speaker Speaker = person
    analyzeInterface(speaker)
    
    fmt.Println("\n--- Analyzing empty interface ---")
    var empty interface{} = person
    analyzeInterface(empty)
    
    // 类型断言使用反射
    fmt.Println("\n--- Type assertion with reflection ---")
    v := reflect.ValueOf(empty)
    if v.Type() == reflect.TypeOf(Person{}) {
        p := v.Interface().(Person)
        fmt.Printf("Successfully asserted to Person: %v\n", p)
    }
}
```

#### 1.7 函数反射

1. 函数类型分析和调用

```go
package main

import (
    "fmt"
    "reflect"
)

// 各种函数类型
func add(a, b int) int {
    return a + b
}

func greet(name string) string {
    return "Hello, " + name
}

func processNumbers(numbers []int, processor func(int) int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = processor(num)
    }
    return result
}

func multiply(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func analyzeFunction(f interface{}) {
    v := reflect.ValueOf(f)
    t := reflect.TypeOf(f)
    
    if v.Kind() != reflect.Func {
        fmt.Println("Not a function")
        return
    }
    
    fmt.Printf("Function type: %s\n", t)
    fmt.Printf("Number of inputs: %d\n", t.NumIn())
    fmt.Printf("Number of outputs: %d\n", t.NumOut())
    
    // 分析输入参数
    fmt.Println("Input parameters:")
    for i := 0; i < t.NumIn(); i++ {
        fmt.Printf("  Param %d: %s\n", i, t.In(i))
    }
    
    // 分析输出参数
    fmt.Println("Output parameters:")
    for i := 0; i < t.NumOut(); i++ {
        fmt.Printf("  Return %d: %s\n", i, t.Out(i))
    }
    
    // 检查是否是可变参数函数
    if t.IsVariadic() {
        fmt.Println("This is a variadic function")
    }
}

func callFunction(f interface{}, args ...interface{}) {
    v := reflect.ValueOf(f)
    t := reflect.TypeOf(f)
    
    if v.Kind() != reflect.Func {
        fmt.Println("Not a function")
        return
    }
    
    // 检查参数数量
    if len(args) != t.NumIn() {
        fmt.Printf("Expected %d args, got %d\n", t.NumIn(), len(args))
        return
    }
    
    // 准备参数
    in := make([]reflect.Value, len(args))
    for i, arg := range args {
        in[i] = reflect.ValueOf(arg)
    }
    
    // 调用函数
    results := v.Call(in)
    
    // 处理返回值
    fmt.Printf("Function call results: ")
    for i, result := range results {
        if i > 0 {
            fmt.Print(", ")
        }
        fmt.Print(result.Interface())
    }
    fmt.Println()
}

func createFunction() interface{} {
    // 动态创建函数
    fnType := reflect.FuncOf(
        []reflect.Type{reflect.TypeOf(int(0)), reflect.TypeOf(int(0))}, // 输入参数类型
        []reflect.Type{reflect.TypeOf(int(0))},                        // 输出参数类型
        false, // 不是可变参数
    )
    
    fn := reflect.MakeFunc(fnType, func(args []reflect.Value) []reflect.Value {
        a := args[0].Int()
        b := args[1].Int()
        return []reflect.Value{reflect.ValueOf(a * b)}
    })
    
    return fn.Interface()
}

func main() {
    fmt.Println("--- Analyzing add function ---")
    analyzeFunction(add)
    
    fmt.Println("\n--- Analyzing greet function ---")
    analyzeFunction(greet)
    
    fmt.Println("\n--- Calling functions ---")
    callFunction(add, 10, 20)
    callFunction(greet, "Alice")
    
    fmt.Println("\n--- Higher-order function ---")
    doubler := multiply(2)
    analyzeFunction(doubler)
    callFunction(doubler, 5)
    
    fmt.Println("\n--- Dynamically created function ---")
    dynamicFn := createFunction()
    analyzeFunction(dynamicFn)
    callFunction(dynamicFn, 6, 7)
    
    // 使用reflect.ValueOf直接调用
    fmt.Println("\n--- Direct function call ---")
    fnValue := reflect.ValueOf(add)
    result := fnValue.Call([]reflect.Value{
        reflect.ValueOf(100),
        reflect.ValueOf(200),
    })
    fmt.Printf("Direct call result: %v\n", result[0].Int())
}
```

#### 1.8 反射的实际应用

1. 结构体验证器

```go
package main

import (
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("Field '%s': %s", e.Field, e.Message)
}

type User struct {
    Name     string `validate:"required,min=2,max=50"`
    Email    string `validate:"required,email"`
    Age      int    `validate:"min=18,max=100"`
    Password string `validate:"required,min=8"`
    Website  string `validate:"url,optional"`
}

func validateStruct(s interface{}) []ValidationError {
    var errors []ValidationError
    
    v := reflect.ValueOf(s)
    t := reflect.TypeOf(s)
    
    // 如果是指针，获取其指向的元素
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
        t = t.Elem()
    }
    
    if v.Kind() != reflect.Struct {
        return []ValidationError{{Field: "root", Message: "not a struct"}}
    }
    
    // 遍历结构体字段
    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        
        // 跳过未导出字段
        if !value.CanInterface() {
            continue
        }
        
        // 获取验证标签
        validateTag := field.Tag.Get("validate")
        if validateTag == "" {
            continue
        }
        
        // 解析验证规则
        rules := strings.Split(validateTag, ",")
        fieldErrors := validateField(field.Name, value, rules)
        errors = append(errors, fieldErrors...)
    }
    
    return errors
}

func validateField(fieldName string, value reflect.Value, rules []string) []ValidationError {
    var errors []ValidationError
    
    for _, rule := range rules {
        rule = strings.TrimSpace(rule)
        if rule == "" {
            continue
        }
        
        // 解析规则和参数
        parts := strings.Split(rule, "=")
        ruleName := parts[0]
        var ruleParam string
        if len(parts) > 1 {
            ruleParam = parts[1]
        }
        
        // 应用验证规则
        if err := applyValidationRule(fieldName, value, ruleName, ruleParam); err != nil {
            errors = append(errors, *err)
        }
    }
    
    return errors
}

func applyValidationRule(fieldName string, value reflect.Value, ruleName, param string) *ValidationError {
    switch ruleName {
    case "required":
        if isEmptyValue(value) {
            return &ValidationError{Field: fieldName, Message: "is required"}
        }
    case "optional":
        if isEmptyValue(value) {
            return nil // 可选字段为空时不验证其他规则
        }
    case "min":
        minVal, err := strconv.Atoi(param)
        if err != nil {
            return &ValidationError{Field: fieldName, Message: "invalid min parameter"}
        }
        
        switch value.Kind() {
        case reflect.String:
            if len(value.String()) < minVal {
                return &ValidationError{Field: fieldName, Message: fmt.Sprintf("minimum length is %d", minVal)}
            }
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            if value.Int() < int64(minVal) {
                return &ValidationError{Field: fieldName, Message: fmt.Sprintf("minimum value is %d", minVal)}
            }
        }
    case "max":
        maxVal, err := strconv.Atoi(param)
        if err != nil {
            return &ValidationError{Field: fieldName, Message: "invalid max parameter"}
        }
        
        switch value.Kind() {
        case reflect.String:
            if len(value.String()) > maxVal {
                return &ValidationError{Field: fieldName, Message: fmt.Sprintf("maximum length is %d", maxVal)}
            }
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            if value.Int() > int64(maxVal) {
                return &ValidationError{Field: fieldName, Message: fmt.Sprintf("maximum value is %d", maxVal)}
            }
        }
    case "email":
        if value.Kind() == reflect.String {
            email := value.String()
            if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
                return &ValidationError{Field: fieldName, Message: "must be a valid email"}
            }
        }
    case "url":
        if value.Kind() == reflect.String {
            url := value.String()
            if url != "" && !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
                return &ValidationError{Field: fieldName, Message: "must be a valid URL"}
            }
        }
    }
    
    return nil
}

func main() {
    // 有效的用户
    validUser := User{
        Name:     "Alice",
        Email:    "alice@example.com",
        Age:      25,
        Password: "securepassword123",
        Website:  "https://example.com",
    }
    
    fmt.Println("--- Validating valid user ---")
    if errors := validateStruct(validUser); len(errors) > 0 {
        fmt.Println("Validation errors:")
        for _, err := range errors {
            fmt.Printf("  %s\n", err.Error())
        }
    } else {
        fmt.Println("User is valid")
    }
    
    // 无效的用户
    invalidUser := User{
        Name:     "A",                    // 太短
        Email:    "invalid-email",        // 无效邮箱
        Age:      15,                     // 太小
        Password: "123",                  // 太短
        Website:  "not-a-url",           // 无效URL
    }
    
    fmt.Println("\n--- Validating invalid user ---")
    if errors := validateStruct(invalidUser); len(errors) > 0 {
        fmt.Println("Validation errors:")
        for _, err := range errors {
            fmt.Printf("  %s\n", err.Error())
        }
    } else {
        fmt.Println("User is valid")
    }
}
```

2. ORM查询构建器

```go
package main

import (
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

type QueryBuilder struct {
    tableName string
    fields    []string
    where     []string
    orderBy   []string
    limit     int
}

func NewQueryBuilder(model interface{}) *QueryBuilder {
    t := reflect.TypeOf(model)
    if t.Kind() == reflect.Ptr {
        t = t.Elem()
    }
    
    // 从结构体类型获取表名
    tableName := strings.ToLower(t.Name()) + "s"
    
    return &QueryBuilder{
        tableName: tableName,
        fields:    []string{"*"},
    }
}

func (qb *QueryBuilder) Select(fields ...string) *QueryBuilder {
    qb.fields = fields
    return qb
}

func (qb *QueryBuilder) Where(condition string, args ...interface{}) *QueryBuilder {
    // 简单的参数替换
    for i, arg := range args {
        placeholder := fmt.Sprintf("$%d", i+1)
        var value string
        
        v := reflect.ValueOf(arg)
        switch v.Kind() {
        case reflect.String:
            value = fmt.Sprintf("'%s'", v.String())
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            value = strconv.FormatInt(v.Int(), 10)
        case reflect.Float32, reflect.Float64:
            value = strconv.FormatFloat(v.Float(), 'f', -1, 64)
        case reflect.Bool:
            value = strconv.FormatBool(v.Bool())
        default:
            value = fmt.Sprintf("'%v'", arg)
        }
        
        condition = strings.Replace(condition, placeholder, value, -1)
    }
    
    qb.where = append(qb.where, condition)
    return qb
}

func (qb *QueryBuilder) OrderBy(field string, direction string) *QueryBuilder {
    qb.orderBy = append(qb.orderBy, fmt.Sprintf("%s %s", field, direction))
    return qb
}

func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
    qb.limit = limit
    return qb
}

func (qb *QueryBuilder) Build() string {
    var query strings.Builder
    
    // SELECT
    query.WriteString("SELECT ")
    query.WriteString(strings.Join(qb.fields, ", "))
    
    // FROM
    query.WriteString(" FROM ")
    query.WriteString(qb.tableName)
    
    // WHERE
    if len(qb.where) > 0 {
        query.WriteString(" WHERE ")
        query.WriteString(strings.Join(qb.where, " AND "))
    }
    
    // ORDER BY
    if len(qb.orderBy) > 0 {
        query.WriteString(" ORDER BY ")
        query.WriteString(strings.Join(qb.orderBy, ", "))
    }
    
    // LIMIT
    if qb.limit > 0 {
        query.WriteString(" LIMIT ")
        query.WriteString(strconv.Itoa(qb.limit))
    }
    
    return query.String()
}

// 结构体到Map的转换
func structToMap(s interface{}) map[string]interface{} {
    result := make(map[string]interface{})
    
    v := reflect.ValueOf(s)
    t := reflect.TypeOf(s)
    
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
        t = t.Elem()
    }
    
    if v.Kind() != reflect.Struct {
        return result
    }
    
    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        
        if !value.CanInterface() {
            continue
        }
        
        // 获取字段名（可以从标签获取）
        fieldName := field.Name
        if dbTag := field.Tag.Get("db"); dbTag != "" {
            fieldName = dbTag
        } else {
            fieldName = strings.ToLower(fieldName)
        }
        
        result[fieldName] = value.Interface()
    }
    
    return result
}

// 从Map填充结构体
func mapToStruct(m map[string]interface{}, s interface{}) error {
    v := reflect.ValueOf(s)
    if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
        return fmt.Errorf("destination must be a pointer to struct")
    }
    
    v = v.Elem()
    t := v.Type()
    
    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        fieldValue := v.Field(i)
        
        if !fieldValue.CanSet() {
            continue
        }
        
        // 获取字段名
        fieldName := field.Name
        if dbTag := field.Tag.Get("db"); dbTag != "" {
            fieldName = dbTag
        } else {
            fieldName = strings.ToLower(fieldName)
        }
        
        if mapValue, exists := m[fieldName]; exists {
            // 类型转换
            if err := setFieldValue(fieldValue, mapValue); err != nil {
                return fmt.Errorf("error setting field %s: %v", field.Name, err)
            }
        }
    }
    
    return nil
}

func setFieldValue(fieldValue reflect.Value, value interface{}) error {
    valueReflect := reflect.ValueOf(value)
    
    // 如果类型完全匹配
    if valueReflect.Type() == fieldValue.Type() {
        fieldValue.Set(valueReflect)
        return nil
    }
    
    // 类型转换
    switch fieldValue.Kind() {
    case reflect.String:
        fieldValue.SetString(fmt.Sprintf("%v", value))
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        if valueReflect.CanInt() {
            fieldValue.SetInt(valueReflect.Int())
        } else {
            return fmt.Errorf("cannot convert %v to int", value)
        }
    case reflect.Float32, reflect.Float64:
        if valueReflect.CanFloat() {
            fieldValue.SetFloat(valueReflect.Float())
        } else {
            return fmt.Errorf("cannot convert %v to float", value)
        }
    case reflect.Bool:
        if valueReflect.Kind() == reflect.Bool {
            fieldValue.SetBool(valueReflect.Bool())
        } else {
            return fmt.Errorf("cannot convert %v to bool", value)
        }
    default:
        return fmt.Errorf("unsupported field type: %v", fieldValue.Kind())
    }
    
    return nil
}

// 示例模型
type User struct {
    ID    int    `db:"id"`
    Name  string `db:"name"`
    Email string `db:"email"`
    Age   int    `db:"age"`
}

func main() {
    // 查询构建器示例
    fmt.Println("--- Query Builder ---")
    
    user := User{}
    query := NewQueryBuilder(user).
        Select("id", "name", "email").
        Where("age > $1", 18).
        Where("name LIKE $1", "%Alice%").
        OrderBy("name", "ASC").
        Limit(10).
        Build()
    
    fmt.Println("Generated SQL:", query)
    
    // 结构体到Map转换
    fmt.Println("\n--- Struct to Map ---")
    sampleUser := User{
        ID:    1,
        Name:  "Alice",
        Email: "alice@example.com",
        Age:   25,
    }
    
    userMap := structToMap(sampleUser)
    fmt.Printf("User map: %+v\n", userMap)
    
    // Map到结构体转换
    fmt.Println("\n--- Map to Struct ---")
    data := map[string]interface{}{
        "id":    2,
        "name":  "Bob",
        "email": "bob@example.com",
        "age":   30,
    }
    
    var newUser User
    if err := mapToStruct(data, &newUser); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("New user: %+v\n", newUser)
    }
}
```

#### 1.9 反射的性能考虑

1. 性能测试

```go
package main

import (
    "fmt"
    "reflect"
    "time"
)

type TestStruct struct {
    ID   int
    Name string
    Age  int
}

func directAccess(s *TestStruct) {
    s.ID = 1
    s.Name = "Test"
    s.Age = 25
}

func reflectionAccess(s interface{}) {
    v := reflect.ValueOf(s).Elem()
    v.FieldByName("ID").SetInt(1)
    v.FieldByName("Name").SetString("Test")
    v.FieldByName("Age").SetInt(25)
}

func benchmarkDirect(iterations int) time.Duration {
    start := time.Now()
    
    for i := 0; i < iterations; i++ {
        s := &TestStruct{}
        directAccess(s)
    }
    
    return time.Since(start)
}

func benchmarkReflection(iterations int) time.Duration {
    start := time.Now()
    
    for i := 0; i < iterations; i++ {
        s := &TestStruct{}
        reflectionAccess(s)
    }
    
    return time.Since(start)
}

func benchmarkReflectionCached(iterations int) time.Duration {
    // 缓存反射信息
    t := reflect.TypeOf(TestStruct{})
    idField, _ := t.FieldByName("ID")
    nameField, _ := t.FieldByName("Name")
    ageField, _ := t.FieldByName("Age")
    
    start := time.Now()
    
    for i := 0; i < iterations; i++ {
        s := &TestStruct{}
        v := reflect.ValueOf(s).Elem()
        
        v.FieldByIndex(idField.Index).SetInt(1)
        v.FieldByIndex(nameField.Index).SetString("Test")
        v.FieldByIndex(ageField.Index).SetInt(25)
    }
    
    return time.Since(start)
}

func main() {
    iterations := 1000000
    
    fmt.Printf("Running %d iterations...\n", iterations)
    
    directTime := benchmarkDirect(iterations)
    reflectionTime := benchmarkReflection(iterations)
    cachedTime := benchmarkReflectionCached(iterations)
    
    fmt.Printf("Direct access: %v\n", directTime)
    fmt.Printf("Reflection access: %v\n", reflectionTime)
    fmt.Printf("Cached reflection: %v\n", cachedTime)
    
    fmt.Printf("Reflection is %.2fx slower than direct access\n", 
        float64(reflectionTime)/float64(directTime))
    fmt.Printf("Cached reflection is %.2fx slower than direct access\n", 
        float64(cachedTime)/float64(directTime))
}
```

2. 优化建议

```go
package main

import (
    "fmt"
    "reflect"
    "sync"
)

// 类型信息缓存
var typeCache = make(map[reflect.Type]*TypeInfo)
var typeCacheMutex sync.RWMutex

type TypeInfo struct {
    Type   reflect.Type
    Fields map[string]reflect.StructField
}

func getTypeInfo(t reflect.Type) *TypeInfo {
    typeCacheMutex.RLock()
    info, exists := typeCache[t]
    typeCacheMutex.RUnlock()
    
    if exists {
        return info
    }
    
    // 创建新的类型信息
    info = &TypeInfo{
        Type:   t,
        Fields: make(map[string]reflect.StructField),
    }
    
    // 缓存字段信息
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        info.Fields[field.Name] = field
    }
    
    // 存储到缓存
    typeCacheMutex.Lock()
    typeCache[t] = info
    typeCacheMutex.Unlock()
    
    return info
}

// 优化的字段设置函数
func setFieldOptimized(v reflect.Value, fieldName string, value interface{}) error {
    typeInfo := getTypeInfo(v.Type())
    
    field, exists := typeInfo.Fields[fieldName]
    if !exists {
        return fmt.Errorf("field %s not found", fieldName)
    }
    
    fieldValue := v.FieldByIndex(field.Index)
    if !fieldValue.CanSet() {
        return fmt.Errorf("field %s cannot be set", fieldName)
    }
    
    return setFieldValue(fieldValue, value)
}

func main() {
    fmt.Println("--- Performance Optimization Tips ---")
    
    fmt.Println("1. 缓存反射信息")
    fmt.Println("2. 使用FieldByIndex而不是FieldByName")
    fmt.Println("3. 预先检查CanSet()")
    fmt.Println("4. 避免重复的TypeOf调用")
    fmt.Println("5. 考虑使用代码生成替代反射")
    
    type TestStruct struct {
        Name string
        Age  int
    }
    
    s := TestStruct{}
    v := reflect.ValueOf(&s).Elem()
    
    // 使用优化的字段设置
    if err := setFieldOptimized(v, "Name", "Alice"); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    if err := setFieldOptimized(v, "Age", 25); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    fmt.Printf("Result: %+v\n", s)
}
```

#### 1.10 反射的最佳实践

1. 何时使用反射

```go
package main

import (
    "fmt"
    "reflect"
)

// 好的反射使用场景
func goodReflectionUseCases() {
    fmt.Println("--- Good Use Cases for Reflection ---")
    
    // 1. 通用序列化/反序列化
    fmt.Println("1. JSON/XML serialization")
    
    // 2. 数据验证框架
    fmt.Println("2. Data validation frameworks")
    
    // 3. ORM映射
    fmt.Println("3. ORM mapping")
    
    // 4. 配置绑定
    fmt.Println("4. Configuration binding")
    
    // 5. 测试框架
    fmt.Println("5. Testing frameworks")
    
    // 6. 依赖注入
    fmt.Println("6. Dependency injection")
}

// 应该避免的反射使用
func avoidReflectionUseCases() {
    fmt.Println("\n--- Avoid Reflection For ---")
    
    // 1. 简单的类型断言
    fmt.Println("1. Simple type assertions")
    
    // 2. 已知类型的操作
    fmt.Println("2. Operations on known types")
    
    // 3. 性能关键路径
    fmt.Println("3. Performance-critical paths")
    
    // 4. 简单的条件逻辑
    fmt.Println("4. Simple conditional logic")
}

// 反射的替代方案
func alternativesToReflection() {
    fmt.Println("\n--- Alternatives to Reflection ---")
    
    // 1. 接口
    fmt.Println("1. Interfaces")
    
    // 2. 类型断言
    fmt.Println("2. Type assertions")
    
    // 3. 代码生成
    fmt.Println("3. Code generation")
    
    // 4. 泛型（Go 1.18+）
    fmt.Println("4. Generics")
}

func main() {
    goodReflectionUseCases()
    avoidReflectionUseCases()
    alternativesToReflection()
}
```

2. 错误处理和安全性

```go
package main

import (
    "fmt"
    "reflect"
)

// 安全的反射操作
func safeReflectionOperations() {
    fmt.Println("--- Safe Reflection Operations ---")
    
    var data interface{} = "hello"
    
    // 1. 检查nil
    if data == nil {
        fmt.Println("Data is nil")
        return
    }
    
    v := reflect.ValueOf(data)
    
    // 2. 检查有效性
    if !v.IsValid() {
        fmt.Println("Value is not valid")
        return
    }
    
    // 3. 类型检查
    if v.Kind() != reflect.String {
        fmt.Printf("Expected string, got %v\n", v.Kind())
        return
    }
    
    // 4. 安全的类型转换
    if str, ok := data.(string); ok {
        fmt.Printf("String value: %s\n", str)
    }
    
    // 5. 检查是否可以设置
    if v.CanSet() {
        fmt.Println("Value can be set")
    } else {
        fmt.Println("Value cannot be set")
    }
}

// 处理反射错误
func handleReflectionErrors() {
    fmt.Println("\n--- Handling Reflection Errors ---")
    
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    var data interface{} = 42
    v := reflect.ValueOf(data)
    
    // 这会引发panic，因为int不能调用String()方法
    // 在实际代码中应该先检查类型
    if v.Kind() == reflect.String {
        fmt.Println(v.String())
    } else {
        fmt.Printf("Value is not a string, it's %v\n", v.Kind())
    }
}

func main() {
    safeReflectionOperations()
    handleReflectionErrors()
}
```

#### 1.11 总结

反射是Go语言的强大特性，但需要谨慎使用：

**优点：**

- 提供运行时类型信息
- 支持动态操作
- 适用于通用库和框架
- 可以实现灵活的序列化/反序列化

**缺点：**

- 性能开销较大
- 编译时类型检查缺失
- 代码可读性降低
- 运行时错误风险

**最佳实践：**

- 优先使用接口和类型断言
- 缓存反射信息以提高性能
- 进行充分的错误检查
- 在性能关键路径避免使用反射
- 考虑使用代码生成替代反射

通过合理使用反射，可以构建更加灵活和通用的Go程序，但要始终权衡灵活性和性能的关系。1 JSON序列化器

```go
package main

import (
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}

type Address struct {
    Street string `json:"street"`
    City   string `json:"city"`
}

func toJSON(v interface{}) string {
    return toJSONValue(reflect.ValueOf(v), 0)
}

func toJSONValue(v reflect.Value, indent int) string {
    t := v.Type()
    
    switch v.Kind() {
    case reflect.String:
        return fmt.Sprintf(`"%s"`, v.String())
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return strconv.FormatInt(v.Int(), 10)
    case reflect.Float32, reflect.Float64:
        return strconv.FormatFloat(v.Float(), 'f', -1, 64)
    case reflect.Bool:
        return strconv.FormatBool(v.Bool())
    case reflect.Slice:
        if v.IsNil() {
            return "null"
        }
        var items []string
        for i := 0; i < v.Len(); i++ {
            items = append(items, toJSONValue(v.Index(i), indent+1))
        }
        return "[" + strings.Join(items, ",") + "]"
    case reflect.Map:
        if v.IsNil() {
            return "null"
        }
        var pairs []string
        for _, key := range v.MapKeys() {
            keyStr := toJSONValue(key, indent+1)
            valueStr := toJSONValue(v.MapIndex(key), indent+1)
            pairs = append(pairs, keyStr+":"+valueStr)
        }
        return "{" + strings.Join(pairs, ",") + "}"
    case reflect.Struct:
        var fields []string
        for i := 0; i < v.NumField(); i++ {
            field := t.Field(i)
            fieldValue := v.Field(i)
            
            // 跳过未导出字段
            if !fieldValue.CanInterface() {
                continue
            }
            
            // 获取JSON标签
            jsonTag := field.Tag.Get("json")
            if jsonTag == "-" {
                continue
            }
            
            fieldName := field.Name
            omitEmpty := false
            
            if jsonTag != "" {
                parts := strings.Split(jsonTag, ",")
                if parts[0] != "" {
                    fieldName = parts[0]
                }
                for _, part := range parts[1:] {
                    if part == "omitempty" {
                        omitEmpty = true
                    }
                }
            }
            
            // 处理omitempty
            if omitEmpty && isEmptyValue(fieldValue) {
                continue
            }
            
            fieldJSON := fmt.Sprintf(`"%s":%s`, fieldName, toJSONValue(fieldValue, indent+1))
            fields = append(fields, fieldJSON)
        }
        return "{" + strings.Join(fields, ",") + "}"
    case reflect.Ptr:
        if v.IsNil() {
            return "null"
        }
        return toJSONValue(v.Elem(), indent)
    case reflect.Interface:
        if v.IsNil() {
            return "null"
        }
        return toJSONValue(v.Elem(), indent)
    default:
        return `"<unsupported type>"`
    }
}

func isEmptyValue(v reflect.Value) bool {
    switch v.Kind() {
    case reflect.String:
        return v.String() == ""
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return v.Int() == 0
    case reflect.Float32, reflect.Float64:
        return v.Float() == 0
    case reflect.Bool:
        return !v.Bool()
    case reflect.Slice, reflect.Map:
        return v.Len() == 0
    case reflect.Ptr, reflect.Interface:
        return v.IsNil()
    }
    return false
}

func main() {
    person := Person{
        Name:  "Alice",
        Age:   30,
        Email: "alice@example.com",
    }
    
    fmt.Println("--- JSON Serialization ---")
    fmt.Println(toJSON(person))
    
    // 测试omitempty
    personWithoutEmail := Person{
        Name: "Bob",
        Age:  25,
    }
    fmt.Println(toJSON(personWithoutEmail))
    
    // 测试复杂结构
    data := map[string]interface{}{
        "person": person,
        "numbers": []int{1, 2, 3},
        "active": true,
    }
    fmt.Println(toJSON(data))
}
```



### 2. 文件操作（File Operations）

Go语言提供了强大而灵活的文件操作功能，主要通过`os`、`io`、`io/ioutil`（Go 1.16后推荐使用`io`和`os`）、`bufio`等包来实现。Go的文件操作设计简洁，支持同步和异步操作，提供了从基本的文件读写到高级的文件系统操作的完整功能。

#### 2.1 文件操作基础

1. 文件的打开和关闭

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // 打开文件进行读取
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close() // 确保文件被关闭
    
    // 获取文件信息
    fileInfo, err := file.Stat()
    if err != nil {
        fmt.Printf("Error getting file info: %v\n", err)
        return
    }
    
    fmt.Printf("File name: %s\n", fileInfo.Name())
    fmt.Printf("File size: %d bytes\n", fileInfo.Size())
    fmt.Printf("File mode: %v\n", fileInfo.Mode())
    fmt.Printf("Is directory: %v\n", fileInfo.IsDir())
    fmt.Printf("Modification time: %v\n", fileInfo.ModTime())
    
    // 使用OpenFile进行更灵活的文件操作
    file2, err := os.OpenFile("example.txt", os.O_RDONLY, 0644)
    if err != nil {
        fmt.Printf("Error opening file with OpenFile: %v\n", err)
        return
    }
    defer file2.Close()
    
    fmt.Println("File opened successfully with OpenFile")
}
```

2. 文件打开模式

```go
package main

import (
    "fmt"
    "os"
)

func demonstrateOpenModes() {
    // 只读模式
    file1, err := os.OpenFile("readonly.txt", os.O_RDONLY, 0644)
    if err != nil {
        fmt.Printf("Error opening readonly file: %v\n", err)
    } else {
        fmt.Println("Readonly file opened")
        file1.Close()
    }
    
    // 只写模式，如果文件不存在则创建
    file2, err := os.OpenFile("writeonly.txt", os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Printf("Error opening writeonly file: %v\n", err)
    } else {
        fmt.Println("Writeonly file opened")
        file2.Close()
    }
    
    // 读写模式
    file3, err := os.OpenFile("readwrite.txt", os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        fmt.Printf("Error opening readwrite file: %v\n", err)
    } else {
        fmt.Println("Readwrite file opened")
        file3.Close()
    }
    
    // 追加模式
    file4, err := os.OpenFile("append.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        fmt.Printf("Error opening append file: %v\n", err)
    } else {
        fmt.Println("Append file opened")
        file4.Close()
    }
    
    // 截断模式（清空文件内容）
    file5, err := os.OpenFile("truncate.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        fmt.Printf("Error opening truncate file: %v\n", err)
    } else {
        fmt.Println("Truncate file opened")
        file5.Close()
    }
}

func main() {
    demonstrateOpenModes()
}
```

#### 2.2 文件读取操作

1. 基本读取方法

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func basicFileReading() {
    // 方法1: 使用Read方法
    file, err := os.Open("sample.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // 读取指定字节数
    buffer := make([]byte, 1024)
    n, err := file.Read(buffer)
    if err != nil && err != io.EOF {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
    
    // 方法2: 读取所有内容
    file.Seek(0, 0) // 重置文件指针到开始位置
    content, err := io.ReadAll(file)
    if err != nil {
        fmt.Printf("Error reading all content: %v\n", err)
        return
    }
    
    fmt.Printf("Full content: %s\n", string(content))
}

func readFileInChunks() {
    file, err := os.Open("largefile.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    buffer := make([]byte, 256) // 256字节的缓冲区
    totalBytes := 0
    
    for {
        n, err := file.Read(buffer)
        if err != nil {
            if err == io.EOF {
                fmt.Println("Reached end of file")
                break
            }
            fmt.Printf("Error reading file: %v\n", err)
            return
        }
        
        totalBytes += n
        fmt.Printf("Read chunk of %d bytes\n", n)
        // 处理读取的数据
        // processChunk(buffer[:n])
    }
    
    fmt.Printf("Total bytes read: %d\n", totalBytes)
}

func main() {
    // 创建示例文件
    sampleContent := "Hello, World!\nThis is a sample file for reading operations.\nIt contains multiple lines of text."
    err := os.WriteFile("sample.txt", []byte(sampleContent), 0644)
    if err != nil {
        fmt.Printf("Error creating sample file: %v\n", err)
        return
    }
    
    basicFileReading()
    readFileInChunks()
}
```

2. 使用bufio进行缓冲读取

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func bufferedFileReading() {
    file, err := os.Open("sample.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // 创建缓冲读取器
    reader := bufio.NewReader(file)
    
    // 方法1: 逐行读取
    fmt.Println("--- Reading line by line ---")
    lineNum := 1
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err.Error() == "EOF" {
                if len(line) > 0 {
                    fmt.Printf("Line %d: %s", lineNum, line)
                }
                break
            }
            fmt.Printf("Error reading line: %v\n", err)
            return
        }
        fmt.Printf("Line %d: %s", lineNum, line)
        lineNum++
    }
    
    // 重置文件指针
    file.Seek(0, 0)
    reader.Reset(file)
    
    // 方法2: 读取指定字节数
    fmt.Println("\n--- Reading specific bytes ---")
    buffer := make([]byte, 10)
    n, err := reader.Read(buffer)
    if err != nil {
        fmt.Printf("Error reading bytes: %v\n", err)
        return
    }
    fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
    
    // 方法3: 读取直到特定字符
    fmt.Println("\n--- Reading until delimiter ---")
    file.Seek(0, 0)
    reader.Reset(file)
    
    text, err := reader.ReadString(',')
    if err != nil {
        fmt.Printf("Error reading until comma: %v\n", err)
    } else {
        fmt.Printf("Read until comma: %s\n", text)
    }
}

func scannerReading() {
    file, err := os.Open("sample.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // 使用Scanner进行逐行读取
    scanner := bufio.NewScanner(file)
    
    fmt.Println("--- Using Scanner ---")
    lineNum := 1
    for scanner.Scan() {
        fmt.Printf("Line %d: %s\n", lineNum, scanner.Text())
        lineNum++
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error scanning file: %v\n", err)
    }
    
    // 使用自定义分割函数
    file.Seek(0, 0)
    scanner = bufio.NewScanner(file)
    
    // 按单词分割
    scanner.Split(bufio.ScanWords)
    
    fmt.Println("\n--- Reading word by word ---")
    wordNum := 1
    for scanner.Scan() {
        fmt.Printf("Word %d: %s\n", wordNum, scanner.Text())
        wordNum++
    }
}

func main() {
    // 创建示例文件
    sampleContent := "Hello, World!\nThis is a sample file for reading operations.\nIt contains multiple lines of text."
    err := os.WriteFile("sample.txt", []byte(sampleContent), 0644)
    if err != nil {
        fmt.Printf("Error creating sample file: %v\n", err)
        return
    }
    
    bufferedFileReading()
    scannerReading()
}
```

#### 2.3 文件写入操作

1. 基本写入方法

```go
package main

import (
    "fmt"
    "os"
)

func basicFileWriting() {
    // 方法1: 使用os.Create创建文件
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()
    
    // 写入字符串
    content := "Hello, Go file operations!\n"
    n, err := file.WriteString(content)
    if err != nil {
        fmt.Printf("Error writing string: %v\n", err)
        return
    }
    fmt.Printf("Wrote %d bytes\n", n)
    
    // 写入字节数组
    data := []byte("This is byte data.\n")
    n, err = file.Write(data)
    if err != nil {
        fmt.Printf("Error writing bytes: %v\n", err)
        return
    }
    fmt.Printf("Wrote %d bytes\n", n)
    
    // 方法2: 使用os.WriteFile一次性写入
    fullContent := "This is a complete file content.\nWritten using WriteFile function.\n"
    err = os.WriteFile("complete.txt", []byte(fullContent), 0644)
    if err != nil {
        fmt.Printf("Error writing complete file: %v\n", err)
        return
    }
    fmt.Println("Complete file written successfully")
}

func appendToFile() {
    // 追加模式打开文件
    file, err := os.OpenFile("append.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        fmt.Printf("Error opening file for append: %v\n", err)
        return
    }
    defer file.Close()
    
    // 追加内容
    appendContent := "This line is appended.\n"
    n, err := file.WriteString(appendContent)
    if err != nil {
        fmt.Printf("Error appending to file: %v\n", err)
        return
    }
    fmt.Printf("Appended %d bytes\n", n)
    
    // 多次追加
    for i := 1; i <= 3; i++ {
        line := fmt.Sprintf("Appended line %d\n", i)
        file.WriteString(line)
    }
    
    fmt.Println("Multiple lines appended successfully")
}

func main() {
    basicFileWriting()
    appendToFile()
}
```

2. 使用bufio进行缓冲写入

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func bufferedFileWriting() {
    file, err := os.Create("buffered_output.txt")
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()
    
    // 创建缓冲写入器
    writer := bufio.NewWriter(file)
    defer writer.Flush() // 确保缓冲区内容被写入
    
    // 写入多行内容
    lines := []string{
        "First line with buffered writer\n",
        "Second line with buffered writer\n",
        "Third line with buffered writer\n",
    }
    
    for i, line := range lines {
        n, err := writer.WriteString(line)
        if err != nil {
            fmt.Printf("Error writing line %d: %v\n", i+1, err)
            return
        }
        fmt.Printf("Buffered write: %d bytes\n", n)
    }
    
    // 写入字节数据
    data := []byte("Byte data through buffered writer\n")
    n, err := writer.Write(data)
    if err != nil {
        fmt.Printf("Error writing byte data: %v\n", err)
        return
    }
    fmt.Printf("Buffered write: %d bytes\n", n)
    
    // 手动刷新缓冲区
    err = writer.Flush()
    if err != nil {
        fmt.Printf("Error flushing buffer: %v\n", err)
        return
    }
    
    fmt.Println("Buffered writing completed successfully")
}

func writeWithFormat() {
    file, err := os.Create("formatted_output.txt")
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()
    
    writer := bufio.NewWriter(file)
    defer writer.Flush()
    
    // 使用fmt.Fprintf进行格式化写入
    name := "Alice"
    age := 30
    salary := 50000.50
    
    fmt.Fprintf(writer, "Employee Information:\n")
    fmt.Fprintf(writer, "Name: %s\n", name)
    fmt.Fprintf(writer, "Age: %d\n", age)
    fmt.Fprintf(writer, "Salary: $%.2f\n", salary)
    
    // 写入表格数据
    fmt.Fprintf(writer, "\n%-10s %-5s %-10s\n", "Name", "Age", "Salary")
    fmt.Fprintf(writer, "%-10s %-5s %-10s\n", "----", "---", "------")
    
    employees := []struct {
        Name   string
        Age    int
        Salary float64
    }{
        {"Alice", 30, 50000.50},
        {"Bob", 25, 45000.75},
        {"Charlie", 35, 60000.00},
    }
    
    for _, emp := range employees {
        fmt.Fprintf(writer, "%-10s %-5d $%-9.2f\n", emp.Name, emp.Age, emp.Salary)
    }
    
    fmt.Println("Formatted writing completed successfully")
}

func main() {
    bufferedFileWriting()
    writeWithFormat()
}
```

#### 2.4 文件指针操作

1. 文件指针移动

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func fileSeekOperations() {
    // 创建测试文件
    content := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    err := os.WriteFile("seek_test.txt", []byte(content), 0644)
    if err != nil {
        fmt.Printf("Error creating test file: %v\n", err)
        return
    }
    
    file, err := os.Open("seek_test.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // 获取文件大小
    fileInfo, _ := file.Stat()
    fileSize := fileInfo.Size()
    fmt.Printf("File size: %d bytes\n", fileSize)
    
    // 从当前位置读取
    buffer := make([]byte, 5)
    n, err := file.Read(buffer)
    if err != nil {
        fmt.Printf("Error reading: %v\n", err)
        return
    }
    fmt.Printf("Read from start: %s\n", string(buffer[:n]))
    
    // 获取当前位置
    currentPos, err := file.Seek(0, io.SeekCurrent)
    if err != nil {
        fmt.Printf("Error getting current position: %v\n", err)
        return
    }
    fmt.Printf("Current position: %d\n", currentPos)
    
    // 移动到文件开始
    newPos, err := file.Seek(0, io.SeekStart)
    if err != nil {
        fmt.Printf("Error seeking to start: %v\n", err)
        return
    }
    fmt.Printf("Moved to position: %d\n", newPos)
    
    // 移动到文件末尾
    newPos, err = file.Seek(0, io.SeekEnd)
    if err != nil {
        fmt.Printf("Error seeking to end: %v\n", err)
        return
    }
    fmt.Printf("Moved to end, position: %d\n", newPos)
    
    // 从末尾向前移动10个字节
    newPos, err = file.Seek(-10, io.SeekEnd)
    if err != nil {
        fmt.Printf("Error seeking from end: %v\n", err)
        return
    }
    fmt.Printf("Moved to position: %d\n", newPos)
    
    // 读取最后10个字符
    n, err = file.Read(buffer)
    if err != nil {
        fmt.Printf("Error reading from end: %v\n", err)
        return
    }
    fmt.Printf("Read from end: %s\n", string(buffer[:n]))
    
    // 移动到中间位置
    middlePos := fileSize / 2
    newPos, err = file.Seek(middlePos, io.SeekStart)
    if err != nil {
        fmt.Printf("Error seeking to middle: %v\n", err)
        return
    }
    fmt.Printf("Moved to middle position: %d\n", newPos)
    
    n, err = file.Read(buffer)
    if err != nil {
        fmt.Printf("Error reading from middle: %v\n", err)
        return
    }
    fmt.Printf("Read from middle: %s\n", string(buffer[:n]))
}

func randomFileAccess() {
    file, err := os.OpenFile("random_access.txt", os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // 写入初始数据
    initialData := make([]byte, 100)
    for i := range initialData {
        initialData[i] = byte('A' + i%26)
    }
    
    _, err = file.Write(initialData)
    if err != nil {
        fmt.Printf("Error writing initial data: %v\n", err)
        return
    }
    
    // 在随机位置写入数据
    positions := []int64{10, 25, 50, 75, 90}
    replacement := "XYZ"
    
    for _, pos := range positions {
        _, err = file.Seek(pos, io.SeekStart)
        if err != nil {
            fmt.Printf("Error seeking to position %d: %v\n", pos, err)
            continue
        }
        
        _, err = file.WriteString(replacement)
        if err != nil {
            fmt.Printf("Error writing at position %d: %v\n", pos, err)
            continue
        }
        
        fmt.Printf("Wrote '%s' at position %d\n", replacement, pos)
    }
    
    // 读取整个文件内容
    _, err = file.Seek(0, io.SeekStart)
    if err != nil {
        fmt.Printf("Error seeking to start: %v\n", err)
        return
    }
    
    content, err := io.ReadAll(file)
    if err != nil {
        fmt.Printf("Error reading all content: %v\n", err)
        return
    }
    
    fmt.Printf("Final content: %s\n", string(content))
}

func main() {
    fileSeekOperations()
    fmt.Println("\n--- Random File Access ---")
    randomFileAccess()
}
```

#### 2.5 文件和目录信息

1. 获取文件信息

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
)

func getFileInfo(filename string) {
    // 方法1: 使用os.Stat
    info, err := os.Stat(filename)
    if err != nil {
        fmt.Printf("Error getting file info: %v\n", err)
        return
    }
    
    fmt.Printf("File: %s\n", filename)
    fmt.Printf("  Name: %s\n", info.Name())
    fmt.Printf("  Size: %d bytes\n", info.Size())
    fmt.Printf("  Mode: %v\n", info.Mode())
    fmt.Printf("  ModTime: %v\n", info.ModTime())
    fmt.Printf("  IsDir: %v\n", info.IsDir())
    
    // 检查文件权限
    mode := info.Mode()
    fmt.Printf("  Permissions: %s\n", mode.Perm())
    fmt.Printf("  Is regular file: %v\n", mode.IsRegular())
    fmt.Printf("  Is directory: %v\n", mode.IsDir())
    fmt.Printf("  Is symlink: %v\n", mode&os.ModeSymlink != 0)
    
    // 方法2: 使用os.Lstat (不跟随符号链接)
    linfo, err := os.Lstat(filename)
    if err == nil {
        fmt.Printf("  Lstat IsDir: %v\n", linfo.IsDir())
        fmt.Printf("  Lstat Size: %d\n", linfo.Size())
    }
}

func checkFileExistence(filename string) {
    // 检查文件是否存在
    if _, err := os.Stat(filename); err == nil {
        fmt.Printf("File %s exists\n", filename)
    } else if os.IsNotExist(err) {
        fmt.Printf("File %s does not exist\n", filename)
    } else {
        fmt.Printf("Error checking file %s: %v\n", filename, err)
    }
    
    // 检查是否为目录
    if info, err := os.Stat(filename); err == nil {
        if info.IsDir() {
            fmt.Printf("%s is a directory\n", filename)
        } else {
            fmt.Printf("%s is a file\n", filename)
        }
    }
}

func compareFiles(file1, file2 string) {
    info1, err1 := os.Stat(file1)
    info2, err2 := os.Stat(file2)
    
    if err1 != nil || err2 != nil {
        fmt.Printf("Error comparing files: %v, %v\n", err1, err2)
        return
    }
    
    fmt.Printf("Comparing %s and %s:\n", file1, file2)
    fmt.Printf("  Size: %d vs %d\n", info1.Size(), info2.Size())
    fmt.Printf("  ModTime: %v vs %v\n", info1.ModTime(), info2.ModTime())
    
    if info1.Size() == info2.Size() {
        fmt.Println("  Files have the same size")
    } else {
        fmt.Println("  Files have different sizes")
    }
    
    if info1.ModTime().Equal(info2.ModTime()) {
        fmt.Println("  Files have the same modification time")
    } else {
        fmt.Println("  Files have different modification times")
    }
}

func main() {
    // 创建测试文件
    testFile := "test_info.txt"
    content := "This is a test file for information gathering."
    err := os.WriteFile(testFile, []byte(content), 0644)
    if err != nil {
        fmt.Printf("Error creating test file: %v\n", err)
        return
    }
    
    // 等待一秒钟，然后创建第二个文件
    time.Sleep(1 * time.Second)
    testFile2 := "test_info2.txt"
    err = os.WriteFile(testFile2, []byte(content), 0644)
    if err != nil {
        fmt.Printf("Error creating second test file: %v\n", err)
        return
    }
    
    // 获取文件信息
    getFileInfo(testFile)
    fmt.Println()
    
    // 检查文件存在性
    checkFileExistence(testFile)
    checkFileExistence("nonexistent.txt")
    fmt.Println()
    
    // 比较文件
    compareFiles(testFile, testFile2)
    
    // 清理
    os.Remove(testFile)
    os.Remove(testFile2)
}
```

#### 2.6 目录操作

1. 目录创建和删除

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func directoryOperations() {
    // 创建单个目录
    err := os.Mkdir("testdir", 0755)
    if err != nil {
        fmt.Printf("Error creating directory: %v\n", err)
    } else {
        fmt.Println("Directory 'testdir' created successfully")
    }
    
    // 创建多级目录
    err = os.MkdirAll("parent/child/grandchild", 0755)
    if err != nil {
        fmt.Printf("Error creating nested directories: %v\n", err)
    } else {
        fmt.Println("Nested directories created successfully")
    }
    
    // 获取当前工作目录
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Printf("Error getting current directory: %v\n", err)
    } else {
        fmt.Printf("Current working directory: %s\n", cwd)
    }
    
    // 改变工作目录
    err = os.Chdir("testdir")
    if err != nil {
        fmt.Printf("Error changing directory: %v\n", err)
    } else {
        fmt.Println("Changed to testdir")
        
        // 获取新的工作目录
        newCwd, _ := os.Getwd()
        fmt.Printf("New working directory: %s\n", newCwd)
        
        // 返回上级目录
        os.Chdir("..")
    }
    
    // 删除空目录
    err = os.Remove("testdir")
    if err != nil {
        fmt.Printf("Error removing directory: %v\n", err)
    } else {
        fmt.Println("Directory 'testdir' removed successfully")
    }
    
    // 删除目录及其所有内容
    err = os.RemoveAll("parent")
    if err != nil {
        fmt.Printf("Error removing directory tree: %v\n", err)
    } else {
        fmt.Println("Directory tree 'parent' removed successfully")
    }
}

func listDirectory(dirPath string) {
    // 方法1: 使用os.ReadDir (Go 1.16+)
    entries, err := os.ReadDir(dirPath)
    if err != nil {
        fmt.Printf("Error reading directory: %v\n", err)
        return
    }
    
    fmt.Printf("Contents of %s:\n", dirPath)
    for _, entry := range entries {
        info, err := entry.Info()
        if err != nil {
            fmt.Printf("Error getting info for %s: %v\n", entry.Name(), err)
            continue
        }
        
        fileType := "FILE"
        if entry.IsDir() {
            fileType = "DIR "
        }
        
        fmt.Printf("  %s %8d %s %s\n", 
            fileType, 
            info.Size(), 
            info.ModTime().Format("2006-01-02 15:04:05"), 
            entry.Name())
    }
    
    // 方法2: 使用filepath.Walk遍历目录树
    fmt.Printf("\nWalking directory tree from %s:\n", dirPath)
    err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        // 计算相对路径的层级
        rel, _ := filepath.Rel(dirPath, path)
        level := len(filepath.SplitList(rel)) - 1
        if rel == "." {
            level = 0
        }
        
        indent := ""
        for i := 0; i < level; i++
```
