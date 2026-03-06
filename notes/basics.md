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
```



##### 函数

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

// 匿名函数
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
```





1
