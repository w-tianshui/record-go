##### `builtin`

Built-in Functions

```go
// return the length of v
func len(v Type) int

// return the capacity of v
func cap(v Type) int
// slice - the maximum length the slice can reach when resliced
// channel - the channel buffer capacity
```

| type              | length                                                   |
| ----------------- | -------------------------------------------------------- |
| array, slice, map | number of elements in v                                  |
| pointer to array  | number of elements in *v                                 |
| string            | number of bytes in v                                     |
| channel           | number of elements queued (unread) in the channel buffer |

```go
// allocates and initializes an object of type slice, map or chan (only)
func make(t Type, size ...IntegerType) Type

// appends elements to the end of a lice
func append(slice []Type, elems ...Type) []Type

// copy | source -> destination, returns the number of elements copied
func copy(dst, src []Type) int

// clear maps/slices （全设为零值）
func clear[T ~[]Type | ~map[Type]Type1](t T)
```



##### Executable

`package main`

- 这段代码不是给别人调用的，会编译成一个可以直接双击运行的二进制程序
- 这个包里必须存在一个名为main的函数 (`func main()`)，这是整个程序的启动入口点



##### `fmt`

```go
// fmt.Println
fmt.Println("随机数是 ", rand.Intn(10))
// fmt.Println/Sprintf 会自动寻找 String() 方法
today := time.Now().Weekday()
fmt.Println(today)	// 会调用today的String()方法，而不是打印出底层的数字
fmt.Println(a)	// 打印数组: [v1 v2 v3 ...]/[[0 1 2] [1 2 3]]


fmt.Printf("现在你有了 %g 个问题。\n", math.Sqrt(7))
fmt.Printf("类型：%T 值：%v\n", MaxInt, MaxInt)

fmt.Sprint(math.Sqrt(x))	// 将参数转化为字符串并返回这个字符串
```



###### 格式化动词

通用格式化动词

| 动词 | 描述                         |
| ---- | ---------------------------- |
| %v   | 值的默认格式，可打印任何类型 |
| %T   | 打印值的类型                 |

布尔与整数动词

| 动词 | 描述             |
| ---- | ---------------- |
| %d   | 十进制整数       |
| %b   | 二进制           |
| %o   | 八进制           |
| %x   | 十六进制         |
| %X   | 十六进制（大写） |
| %t   | 布尔值           |

浮点数与复数动词

| 动词 | 描述              |
| ---- | ----------------- |
| %g   | 根据精度选择%e/%f |
| %e   | 科学计数法（e）   |
| %E   | 科学计数法（E）   |
| %f   | 小数形式          |

字符串与字节动词

| 动词 | 描述                                 |
| ---- | ------------------------------------ |
| %s   | 字符串                               |
| %q   | 带引号的字符串（对特殊字符进行转义） |
| %x   |                                      |

指针动词

| 动词 | 描述 |
| ---- | ---- |
|      |      |
|      |      |



##### `math`

提供基础数学常量和函数

```go
math.Sin(x)		// 输入float64参数
math.Sqrt(7)
math.Pi
math.Pow(3,2)	// 9
```



##### `math/rand`

生成伪随机数

```go
rand.Intn(10)	// [0,10)
```



##### `math/cmplx`

```go
complx.Sqrt(-5+12i)
```



##### `time`

provides functionality for measuring and displaying time

```go
time.Now().Weekday()	// time.Weekday类型，周日为0
time.Saturday	// time.Weekday(6)
```



##### `slices`

defines various functions useful with slices of any type

```go
// 查找
`func Contains`
func Contains[S ~[]E, E comparable](s S, v E) bool
// reports whether v is present in s
slices.Contains(numbers, 2)

`func Index`
func Index[S ~[]E, E comparable](s S, v E) int
// return the index the first occurrence of v in s / -1
fmt.Println(slices.Index(numbers, 2))

`func Max`
func Max[S ~[]E, E cmp.Ordered](x S) E
// return the maximal value

// 修改
`func Insert`
func Insert[S ~[]E, E any](s S, i int, v ...E) S
`func Delete`
func Delete[S ~[]E, E any](s S, i, j int) S
`func Replace`
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S
`func Reverse`
func Reverse[S ~[]E, E any](s S)

// 排序
`func Sort`
func Sort[S ~[]E, E cmp.Ordered](x S)
`func SortFunc`
func SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int)
```



##### `bytes`

`[]byte`同样适用

```go
`func Split`
`func Replace`
`func Trim`
`func Contains`
```



##### 0

- 包中函数名大写——公开的；小写——私有的
