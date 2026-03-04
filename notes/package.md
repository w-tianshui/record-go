`main`

程序从main包开始运行。



##### `fmt`

```go
// fmt.Println
fmt.Println("随机数是 ", rand.Intn(10))
// fmt.Println/Sprintf 会自动寻找 String() 方法
today := time.Now().Weekday()
fmt.Println(today)	// 会调用today的String()方法，而不是打印出底层的数字


fmt.Printf("现在你有了 %g 个问题。\n", math.Sqrt(7))
fmt.Printf("类型：%T 值：%v\n", MaxInt, MaxInt)

fmt.Sprint(math.Sqrt(x))	// 将参数转化为字符串并返回这个字符串
```



##### 格式化动词

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
