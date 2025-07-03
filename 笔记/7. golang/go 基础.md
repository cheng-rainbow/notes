根据Go语言开发者自述，近10 多年，从单机时代的C语言到现在互联网时代的Java，都没有令人满意的开发语言，而C++往往给人的感觉是，花了100%的经历，却只有60%的开发效率，产出比太低，Java和c#的哲学又来源于C++。并且，随着硬件的不断升级，这些语言不能充分的利用硬件及CPU。因此，一门高效、简洁、开源的语言诞生了。





 nsq：bitly开源的消息队列系统，性能非常高，目前他们每天处理数十亿条的消息

docker：基于Ixc 的一个虚拟打包工具，能够实现PAAS平台的组建

packer:用来生成不同平台的镜像文件，例如VM、vbox、AWs等，作者是vagrant的

skynet:分布式调度框架

Doozer:分布式同步工具，类似ZooKeeper

Heka：mazila开源的日志处理系统

cbfs：couchbase开源的分布式文件系统

tsuru：开源的PAAS平台，和SAE实现的功能一模一样

groupcache：memcahe作者写的用于Google下载系统的缓存系统



Go语言的设计目标是解决现代软件开发中的常见问题，尤其是在大型系统开发中。它的核心理念包括：

- 简洁性：语言规范简单，语法直观，易于学习和维护。Go避免了复杂的特性（如继承、运算符重载），代码可读性高。
- 高效性：Go是编译型语言，编译速度快，直接生成机器码，运行性能接近C/C++。
- 并发支持：内置goroutine和channel，提供轻量级并发模型，适合多核和分布式系统。
- 现代性：支持垃圾回收、内存安全，适合云计算和微服务架构。
- 易部署：编译为单一二进制文件，无需运行时依赖，跨平台部署简单。



2. Go语言的核心特性

2.1 静态类型与编译型

- Go是静态类型语言，变量类型在编译时确定，减少运行时错误。
- 编译速度极快，接近脚本语言的开发体验。
- 生成的二进制文件包含所有依赖，运行无需额外环境。

2.2 简洁的语法

- Go的语法设计极简，只有25个关键字（如if、for、go、chan等），远少于C++或Java。
- 没有类和继承，使用结构体和接口实现组合和多态。
- 自动推导类型（如:=操作符），减少样板代码。

2.3 垃圾回收

- Go内置垃圾回收器（GC），开发者无需手动管理内存。
- GC经过优化，低延迟，适合高并发场景。

2.4 内置并发支持

- Goroutine：轻量级线程，由Go运行时管理，非OS线程，创建和切换成本极低（几KB内存）。
- Channel：用于goroutine间通信的原生机制，支持安全的并发数据传递。
- 并发模型基于CSP（Communicating Sequential Processes），强调“通过通信共享内存”而非“通过共享内存通信”。

2.5 标准库强大

- Go的标准库覆盖网络、文件操作、加密、测试等，功能丰富，开箱即用。
- 例如：
  - net/http：实现高性能Web服务器。
  - encoding/json：JSON序列化和反序列化。
  - testing：内置单元测试和基准测试框架。

2.6 工具链完善

- Go提供强大的工具链：
  - go fmt：自动格式化代码，统一代码风格。
  - go vet：静态分析，检测潜在错误。
  - go test：运行测试用例。
  - go mod：模块化依赖管理。
  - golang.org/x/tools：提供额外工具，如静态检查和代码补全。

2.7 跨平台与部署

- 支持主流平台（Linux、Windows、macOS）和架构（x86、ARM等）。
- 单一二进制文件，部署简单，无需安装运行时或虚拟机。



## 二、数据类型和声明

golang的数据类型分为 **值类型** 和 **引用类型**

- 值类型：所有基本类型数据、**数组**、**结构体**

- 其他的为引用类型



### 1. 基本数据类型

**整形**

- int：平台相关（32位或64位系统上为32位或64位）
- int8：8位，范围 -128 到 127
- int16：16位，范围 -32,768 到 32,767
- int32：32位，范围 -2,147,483,648 到 2,147,483,647
- int64：64位，范围 -2^63 到 2^63-1



- uint：平台相关（32位或64位系统上为32位或64位）
- uint8：8位，范围 0 到 255（等同于byte）
- uint16：16位，范围 0 到 65,535
- uint32：32位，范围 0 到 4,294,967,295
- uint64：64位，范围 0 到 2^64-1



- byte：uint8的别名，常用于字节操作。
- rune：int32的别名，表示Unicode码点，常用于字符处理。

**字符串**

string：不可变的字节序列，通常存储UTF-8编码的文本。

支持索引（返回byte）和切片操作，但不能直接修改。

**布尔型**：bool 表示真（true）或假（false）。

**浮点型**

float：平台相关（32位或64位系统上为32位或64位）

float32：32位浮点数，约6-7位有效数字。

float64：64位浮点数，约15-16位有效数字，Go中默认浮点类型。

**复数型**

complex64：由两个float32组成（实部和虚部）。

complex128：由两个float64组成，Go中默认复数类型。



### 2. 复合数据类型

**数组**：固定长度、相同类型的元素序列。长度是数组类型的一部分，定义后不可变。

**切片**：动态长度、可变数组，基于数组的引用类型。包含指针、长度（len）和容量（cap）。

**映射**：键值对集合，类似Python的字典。

**结构体**：自定义类型，聚合多个字段。

**指针**：存储变量的内存地址。

**接口**：定义方法集合，类型通过实现方法隐式满足接口。

**函数类型**：函数可以作为类型，用于回调或高阶函数。

**Channel**：用于goroutine间通信的类型。



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

| 动词 | 描述                     | 示例代码                  | 输出     |
| ---- | ------------------------ | ------------------------- | -------- |
| %v   | 默认格式（值的默认表示） | fmt.Printf("%v", 42)      | 42       |
| %T   | 值的类型                 | fmt.Printf("%T", 42)      | int      |
| %s   | 字符串                   | fmt.Printf("%s", "hello") | hello    |
| %f   | 浮点数（默认6位小数）    | fmt.Printf("%f", 3.14159) | 3.141590 |
| %b   | 二进制                   | fmt.Printf("%b", 42)      | 101010   |
| %d   | 十进制整数               | fmt.Printf("%d", 42)      | 42       |
| %x   | 十六进制（小写）         | fmt.Printf("%x", 42)      | 2a       |

格式化输出函数：

`fmt.Printf`：格式化并打印到标准输出（控制台）。

`fmt.Sprintf`：格式化并返回格式化后的字符串。

`fmt.Fprintf`：格式化并写入指定的 io.Writer（如文件或网络连接）。



### 5. golang 的包





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

```go
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
// 简写，隐式返回
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

4. **类型嵌套**

Go 通过类型嵌入实现类似继承的功能。嵌入允许将一个类型的字段或方法直接“嵌入”到另一个类型中。嵌入的类型的方法和字段可以直接访问

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return "I am " + a.Name
}

type Dog struct {
    Animal // 嵌入 Animal 结构体
    Breed  string
}

func main() {
    d := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Golden Retriever"}
    fmt.Println(d.Speak()) // 输出: I am Buddy
}
```



### 2. type **`结构体和接口`**

在 Go 语言中，**字段**、**函数**或**方法**的可见性由其名称的首字母大小写决定

**大写开头**（如 Name）：表示字段或方法是导出的（exported），即包外的代码可以访问（类似 public）。

**小写开头**（如 name）：表示字段或方法是未导出的（unexported），即只能在定义该字段的包内访问（类似 private）。

1. **定义结构体**

```go
type Person struct {
    Name string
    Age  int
    Email string
}
```

2. **实例化结构体**

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

3. **访问修改字段**

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

4. **定义方法**

除了可以对结构体定义方法，可以给 **自定义类型** 定义方法

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) SayHello() string {
    return "Hello, my name is " + p.Name
}

func (p *Person) SetAge(age int) {
    p.Age = age
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    fmt.Println(person.SayHello()) // 输出: Hello, my name is Alice
    person.SetAge(31)
    fmt.Println(person.Age)        // 输出: 31
}
```

5. **匿名字段**

Go 支持结构体嵌套，通过匿名字段（embedded struct）实现类似继承的功能。匿名字段是指在结构体中声明另一个结构体类型，但不指定字段名。

```go
type Address struct {
    City  string
    State string
}

type Person struct {
    Name    string
    Age     int
    Address // 匿名字段
}

------
访问匿名字段的方式有两种：
person := Person{
    Name:    "Alice",
    Age:     30,
    Address: Address{City: "New York", State: "NY"},
}
fmt.Println(person.City)  // 输出: New York

fmt.Println(person.Address.City) // 输出: New York
```

6.  **结构体标签**

结构体字段可以附加标签（tag），通常用于序列化（如 JSON、XML）或数据库映射。标签是一个字符串，定义在字段类型的后面，使用反引号（`）。

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`	// 定义转换成json后的字段名
    Age  int    `json:"age"`
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    data, _ := json.Marshal(person)
    fmt.Println(string(data)) // 输出: {"name":"Alice","age":30}
}
```



**接口**

```go
// 接口
type Speaker interface {
    Speak() string
}

type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}

func main() {
    var s Speaker = Person{Name: "Alice"}
    fmt.Println(s.Speak()) // 输出: Hello, my name is Alice
}
```



### 3. defer

defer 是一个关键字，用于延迟执行一个函数调用或方法调用，直到包含它的函数返回（无论是正常返回还是发生 panic）。defer 常用于资源清理（如关闭文件、释放锁）或确保某些操作在函数退出时执行。

1. defer 后接一个函数调用（可以是命名函数、匿名函数或方法）。
2. 延迟的函数调用会在包含它的函数返回之前执行。

```go
func main() {
    fmt.Println("Start")
    defer fmt.Println("Deferred call") // 延迟执行
    fmt.Println("End")
}

// 输出
Start
End
Deferred call
```



### 4. panic

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



### 5. recover

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



### 6. time 包

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







## 六、进阶

### 1. 定时器

在 Go 语言中，**定时器（Timer）**和 **周期性触发器（Ticker）**是 time 包提供的两种核心机制，用于处理时间相关的任务。两者都基于 Go 的并发模型，通过通道（chan）与 goroutine 配合使用，提供了高效、简洁的定时功能。

`time.Timer`：用于在指定时间后触发一次事件，常用于超时控制或延迟执行。

`time.Ticker`：用于按固定时间间隔重复触发事件，常用于周期性任务。



### 2. 指针

Go 语言中，指针（pointer）是一种重要的特性，用于直接操作内存地址，从而实现高效的数据传递和修改。相比 C/C++ 的指针，Go 限制了一些不安全的操作（如指针算术），同时保留了指针的核心功能。

1. 什么是指针

指针**是一个变量**，**存储**的是**另一个变量的内存地址**。

**定义**指针：`var p *T`（T 是任意类型，如 int、string、结构体等）。

**获取地址**：`&x` 返回变量 x 的内存地址。

**解引用**：`*p` 访问指针 `p` 指向的变量的值。

2. 定义并初始化

```go
var p *int // 声明一个 int 指针，初始值为 nil

x := 42
p = &x // p 指向 x 的地址

p = new(int) // 分配 int 类型的内存，初始值为 0
*p = 42      // 设置指针指向的值
```



### 3. make 和 new

`make` 和 `new`是两个是 golang 种用于内存分配的内置函数

`new(T)` 

1. 分配类型 T 的零值内存，并返回指向该内存的指针 `*T`。
2. 为任意类型（包括基本类型、结构体、指针等）分配内存。并返回指向该内存的指针。

`make(T, args)` 

1. 用于创建并初始化特定类型的对象，仅适用于切片（slice）、映射（map）和通道（chan）。
2. 分配内存并初始化内部数据结构（如切片的底层数组、映射的哈希表、通道的缓冲区）。
3. 返回初始化后的对象（不是指针），类型为 `T`。



## 五、Goroutines 和 Channel

Go 的并发模型是其最强大且独特的特点之一。Go 提供了简洁而高效的并发编程方式，能够让你轻松地编写并发程序。Go 的并发模型主要是通过 **goroutines** 和 **channels** 来实现的。下面我会详细讲解这两者，以及它们如何协同工作来实现并发。

### 1. **Goroutines：轻量级线程**

在 Go 中，**goroutine** 是一种轻量级的线程。它由 Go 运行时（runtime）管理，能够并发执行任务。与操作系统线程相比，goroutine 的开销非常小，可以在同一程序中并发执行成千上万的 goroutine。

创建和管理 goroutine 的方式非常简单，只需要使用 `go` 关键字来启动一个新的 goroutine。任何函数都可以作为 goroutine 被调用。在一个 Go 程序中，所有的函数都是可以并发执行的。你只需要使用 `go` 关键字来启动它。

```go
func printMessage(message string) {
    fmt.Println(message)
}

func main() {
    go printMessage("Hello from goroutine!")  // 启动一个 goroutine
    printMessage("Hello from main thread!")   // 主线程输出
}
```

Goroutine 的特点

- **轻量级**：Goroutine 的创建和销毁比操作系统线程轻得多。Go 会自动管理 goroutine 的调度和执行。
- **内存共享**：Goroutine 可以共享内存，但需要同步机制来防止竞争条件。
- **自动调度**：Go 的运行时调度器会自动调度多个 goroutine 的执行。程序员不需要管理线程池或调度。



### 2. **Channels：沟通的桥梁**

`channel` 是 Go 中的一个核心概念，它提供了 goroutine 之间的通信机制。goroutines 通过 channels 交换数据，从而实现同步和数据传递。Go 的 channel 是类型安全的（即你只能通过特定类型的 channel 传递数据），并且支持阻塞、同步等特性。

1. 创建一个 Channel

```
go


复制编辑
ch := make(chan int)  // 创建一个传递整数的 channel
```

这里 `ch` 是一个 channel，它用来传递 `int` 类型的数据。你可以通过 `make(chan Type)` 来创建一个指定类型的 channel。

2. 向 Channel 发送数据

使用 `<-` 运算符可以将数据发送到 channel 中：

```go
ch <- 42  // 将数据 42 发送到 channel ch
```

3. 从 Channel 接收数据

同样，使用 `<-` 运算符来接收来自 channel 的数据：

```go
value := <-ch  // 从 channel ch 中接收数据
fmt.Println(value)  // 输出 42
```

4. Channels 的阻塞特性

- 如果一个 goroutine 尝试从 channel 中读取数据时，channel 中没有数据，它会阻塞，直到数据可用。
- 如果一个 goroutine 向 channel 发送数据时，另一个 goroutine 没有准备好接收，它会阻塞，直到接收方准备好。

这种特性使得 goroutines 之间的通信既可以同步也可以实现高效的数据交换。

5. Buffered Channels（缓冲区 Channel）

Go 还支持 **缓冲区 channel**，这使得可以向 channel 发送数据而不需要立即有接收方。通过为 channel 指定缓冲区大小，你可以在缓冲区满之前将数据发送到 channel 中，接收方再从缓冲区取数据。







Go 程序的入口是 `main` 函数，这个函数位于 `main` 包中。Go 程序必须有一个 `main` 包和 `main` 函数来启动执行。

`main` 函数是 Go 程序的入口函数，每个 Go 程序都需要一个 `main` 包，并且 `main` 包中的 `main` 函数会在程序运行时自动执行。

`package main`：每个 Go 程序都是由一个包（`package`）组成的，`main` 包表示这个程序是一个可执行程序。其他包主要是库，不会直接运行。

`func main() {}`：定义了程序的入口点，这个函数没有参数也没有返回值。Go 程序的执行从 `main` 函数开始。
