#### 基础语法

```go
// 变量赋值
var x int = 10
x := 10
```

```go
// 逻辑结构
if a>b {
    ...
}

for i := 0; i<5; i++ {
    ...
}
```

```go
// 函数
func main(){
    ...
}
```

#### 数据类型

| C++                          | Go                                        |
| ---------------------------- | ----------------------------------------- |
| bool                         | bool                                      |
| int, short, long, long long  | int, int8, int16, int32, int64            |
| unsigned int, unsigned short | uint, uint8, uint16,<br />uint 32, uint64 |
| float, double                | float32, float64                          |
| char  (通常是1字节)          | byte (/uint8, 主要用于数据流)             |
|                              | string                                    |
|                              |                                           |

```go
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
