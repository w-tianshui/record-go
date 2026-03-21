程序从main包开始运行

```go
import (
	"fmt"
    "math"
)
```





#### 数据类型

| C++                          | Go                                                           |
| ---------------------------- | ------------------------------------------------------------ |
| bool                         | bool                                                         |
| int, short, long, long long  | int, int8, int16, int32, int64                               |
| unsigned int, unsigned short | uint, uint8 (byte), uint16,<br />uint32 (rune), uint64, uintptr |
| float, double                | float32, float64                                             |
|                              | complex64, complex128                                        |
|                              | string                                                       |
|                              |                                                              |

```go
i := 42
f := float64(i)	// 显式类型转换

// 字符串
内建string类型，不可变
可通过 + 连接

// 数组
[5]int  // 数组大小属于类型 不可变
// 动态数组
slice  // 对底层数组的引用

// 映射（map）哈希表（Hash Table）
内建map类型，基于哈希实现

// 结构体
struct  // 只包含数据，方法附加在结构体外

内建函数类型，可以传递和返回

内建interface类型，实现多态的核心（隐式实现）

只支持安全指针，用于传递引用和性能优化，不可进行算数操作
```



##### 变量

```go
// 变量赋值
var x int = 10	// 显式声明
x := 10		// 声明并初始化变量的简写
var y = 10		// 自动推断有初始值的变量的类型
var z int		// 无初始值，初始化为 零值
var c, python, java = true, false, "no!"	// 一次声明多个变量
```



##### 常量

```go
// 字符、字符串、布尔、数值

// const用于声明常量，不能用:=
const s string = "constant"
const d = 3e20 / n	// 可执行任意精度运算，数值型常量没有确定的类型

```



#### 逻辑结构

##### 循环

```go
// 初始/条件/后置 初始和后置是可选的
for i := 0; i<5; i++ {
    ...
}
for i<=5 {
    ...
}
for {
    ...
    break
}
```



##### 分支

```go
if x==1 {
    ...
}else{
    ...
}

// switch的case可用","分隔多个值
switch i {
case 1, 01:
    fmt.Println("one")
case 2:
    fmt.Println("two")
}
// switch可不带表达式，case的表达式可以不是常量
switch {
case x == 1:
    fmt.Println(1)
default:
    fmt.Println(2)
}

// Type Switch
whatAmI := func(i interface{}) {	// interface{}可以接受任何类型的值（可写为any）
    switch t := i.(type) {	// type switch语法，提取i的类型
    case bool:
        fmt.Println("I'm a bool")
    case int:
        fmt.Println("I'm an int")
    default:
        fmt.Printf("Don't know type %T\n", t)
    }
}
whatAmI(true)	// bool
whatAmI(1)		// int
whatAmI("hey")	// string
```





#### 数组

```go
// 类型为 [5]int
var a [5]int
var a2 [2][3]int

// 直接声明并初始化
b := [5]int{1, 2, 3, 4, 5}
```



#### 切片

不同于数组，slice的类型于元素个数无关

```go
// 创建一个长度为3的string类型的slice（初始值为零值）
s := make([]string, 3)
s = append(s, "a")
// 截取切片
s2 := s[1:3]	// [1,3)
s3 := s[:3]		// [0,3)
s4 := s[2:]		// [2,3]
// 定义并初始化
t := []string{"a", "b", "c"}

// 二维切片，内部长度可不同
twoD := make([][]int, 3)
for i := 0; i<3; i++ {
    innerLen := i+1		// 数组个数为1, 2, 3
    twoD[i] = make([]int, innerLen)
    for j := 0; j<inner; j++{
        twoD [i][j] = i+j
    }
}
// [[0] [1 2] [2 3 4]]
```


#### Map

```go
`make(map[key-type]val-type)`
m := make(map[string]int)	// name[key] = val
m["k1"]=7		// 添加键值对
m["k1"]=13
fmt.Println(m)	// map[k1:7 k2:13]
delete(m, "k2")	// 通过key删除键值对
x := m["k2"]	// 访问不存在的key，返回零值
_, ok := m["k2"]	// 第二个返回值表示是否存在

// 声明并初始化
n := map[string]int{"k1":7, "k2":13}
```



#### 函数

```go
`func 函数名 (参数) 返回值类型`
func add(x int, y int) int {
	return x + y
}
// 省略参数相同类型
func add(x, y int) int {
	return x + y
}
// 多返回值
func swap(x, y string) (string, string) {
	return y, x
}
// 预先定义命名返回值 - 无参数的return会返回命名的返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 匿名函数 - 定义的时候不分配名称，只是临时用一下
add := func(x int, y int) int {
    return x + y
}
// 回调函数 - 用来被其他函数调用的函数，通常用匿名函数作为回调
```



#### 闭包

带有环境的函数，是回调的一种特殊形式

```go
// inSeq函数返回一个匿名函数，返回的函数隐藏变量i以形成闭包，使得i不会在intSeq返回后被销毁
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
nextInt := intSeq()		// 定义nextInt为一个匿名函数，nextInt在被调用的时候会让i++，并输出
fmt.Println(nextInt()) 	// i = 1	
fmt.Println(nextInt()) 	// i = 2
fmt.Println(nextInt()) 	// i = 3

nextInts := intSeq()	// 定义一个新的闭包，它含有自己的i
```





#### 0

##### 零值

```go
int, float	// 0
string		// ""
bool		// false
指针 切片 map channel 函数	// mil
结构体		// 各字段是各自的零值
```



##### Comma OK

用来安全地处理可能会失败/值不存在的情况，断言成功ok为true，断言失败ok为false

```go
`value, ok := ...`
// Type Assertion
`value, ok := interface变量.(目标类型)`
if s, ok := i.(string); ok {
    // i的类型是string
    // 断言成功，s存储字符串
}

// map的key是否存在
_, ok := m["k2"]	// m - map[string]int

// 通道读取
v, ok := <-ch
```



##### 泛型

~ 底层类型匹配

```go
func Index[S ~[]E, E comparable](s S, v E) int
[S ~[]E, E comparable]	// 类型参数列表，定义函数要使用的泛型类型（告诉编译器这个函数是针对一类满足特定规则的类型的通用模板）
S ~[]E	// S可以是原生切片，也可以是基于切片定义的自定义类型

```





#### 语法糖

对功能没有实质性影响，提高代码可读性

###### 1.短变量声明

同时完成声明和初始化，并自动推断类型

```go
name := "world"	// var name string = "world"
```



###### 2.for...range循环

无需关心长度、索引边界来获取索引和值

```go
// 遍历切片、数组、Map、字符串、通道
for index, value := range mySlice {
    fmt.Println(index, value)
}
// 等价于
for i:=0; i<len(mySlice); i++ {
    index := i
    value := mySlice[i]
    fmt.Println(index, value)
}

// 对map进行遍历
for key, value := range myMap {
    fmt.Println(key, value)
}
// range可以只遍历map的键
for key := range myMap {
    fmt.Println(key)
}

// 在字符串中迭代unicode码点
for i, c := range "go" {
    fmt.Println(i, c)
}
// 0 103 /n 1 111
```



###### 3.defer语句

使一个函数调用在包含它的函数执行结束前被执行，是资源清理利器，可以把资源申请和释放的代码写在一起。其他语言中，这通常需要``try...finally`来实现

```go
file, err := os.Open("test.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close() // 无论函数如何退出（正常返回、panic），这个都会被执行
```



###### 4.go关键字

启动Goroutine，创建一个goroutine并把它放入调度器中，隐藏了线程创建/复用、栈管理等。

```go
go myFunction() // 异步执行 myFunction，不阻塞当前线程
```



###### 5.类型切换

判断一个接口的多种可能类型时使用

```go
switch v := i.(type) {
case string:
    ...
case int:
    ...
}
// 等价于
if v, ok = i.(string); ok {
    ...
} else if v, ok = i.(int); ok {
    ...
}
```



###### 6.变长参数

variadic parameters
允许向一个函数传递任意数量的同类型参数。函数内部打包成一个切片，`...int` 相当于 `[]int`

- 变长参数一定是最后一个参数
- 传递切片需要进行解包

```go
`elems ...Type` // 类型前面加"..."
func sum(nums ...int) int {
    fmt.Printf("传入的参数 nums 的类型是: %T\n", nums) // 查看它的真实类型
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

numbers := []int{10, 20, 30}
// 解包切片，传入后还会再打包成切片
total := sum(numbers...) 
```

- `...any`/`...interface{}` 任意类型的任意数量的同类型参数
