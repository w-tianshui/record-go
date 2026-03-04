#### 基础语法

```go
// 变量赋值
var x int = 10
var c, python, java = true, false, "no!"
x := 10
```

```go
// 逻辑结构
if 
if a>b {
    ...
}
// 初始化语句和后置语句是可选的
for i := 0; i<5; i++ {
    ...
}
// while
for i<5 {
    ...
}
```

```go
// 函数
func add(x int, y int) int {
	return x + y
}
// 省略相同类型
func add(x, y int) int {
	return x + y
}
// 多返回值
func swap(x, y string) (string, string) {
	return y, x
}
// 命名返回值 - 无参数的return会返回命名的返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

```go
import "fmt"
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
f := float64(i)	// 类型转换
// 常量 不能用:=声明
const Pi = 3.14

// 字符串
内建string类型，不可变

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
