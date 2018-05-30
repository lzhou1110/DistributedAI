# 基础
## 包，变量和函数
### 包
每个 Go 程序都是由包构成的。
程序从`main`包开始运行。

本程序通过导入路径`"fmt"`和`"math/rand"`来使用这两个包。

按照约定，包名与导入路径的最后一个元素一致。例如`"math/rand"`包中的源码均以`package rand`语句开始。

*注意：* 此程序的运行环境是固定的，因此 `rand.Intn`总是会返回相同的数字。 （要得到不同的数字，需为生成器提供不同的种子数，参见 `rand.Seed`。）
### 导入
此代码用圆括号组合了导入，这是“分组”形式的导入语句。[推荐]

```
import (
	"fmt"
	"math"
)
```
当然你也可以编写多个导入语句，例如：

```
import "fmt"
import "math"
```
### 导出名
在`Go`中，如果一个名字以大写字母开头，那么它就是已导出的。例如，`Pizza`就是个已导出名，`Pi`也同样，它导出自`math`包。

`pizza`和`pi`并未以大写字母开头，所以它们是未导出的。

在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。

执行代码，观察错误输出。

### 函数
函数可以没有参数或接受多个参数。
`add`接受两个`int`类型的参数。

注意**类型在变量名**之后。

```go
func add(x int, y int) int {
	return x + y
}
```
当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
`x int, y int`可以写作

```go
func add(x, y int) int{
    return x+y
}
```

#### 与C语言对比
[参考](https://blog.go-zh.org/gos-declaration-syntax)

1. C系列语言定义`int x`， 是定义一个表达式， 包含变量， 声明表达式的类型，从右向左读取。
2. C系列之外的语言则是:

```
x: int
p: pointer to int
a: array[3] of int
```
从左向右读取，写出一个变量， 声明的类型(Note：指针类型除外)

3. GO中的指针
    在GO语法中，将括号放到类型的左边，但表达式则是将括号放在右边

```GO
var a []int
x = a[1]
```
### 多值返回
函数可以返回任意数量的返回值。

```go
package main
import "fmt"
func swap(a, b string) (string, string){
    return b, a
}
func main(){
    a, b := swap("hello", "world")
    fmt.Println("a, b:", a, b)
}
```
### 命名返回值
1. Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
2. 返回值的名称应当具有一定的意义，它可以作为文档使用。
3. 没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
4. 直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。

```go
package main
import "fmt"
func split(sum int) (x, y int) {
	x = sum * 4 /9
	y = sum - x
	return
}
func main() {
	fmt.Println(split(17))
}
```
### 变量
`var`语句用于**声明一个变量列表**，跟函数的参数列表一样，类型在最后。

就像在这个例子中看到的一样，var 语句可以出现在包或函数级别。

```go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```
### 变量的初始化
变量声明可以包含初始值，每个变量对应一个。

如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。

```go
package main
import "fmt"
var i, j int = 1, 2
func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```
### 短变量声明
在函数中，简洁**赋值语句**`:=`可在类型明确的地方代替`var`声明。

函数外的每个语句都必须以关键字开始`（var, func 等等）`，因此 `:= `结构**不能**在函数外使用。

```go
package main
import "fmt"
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```
### 基本类型
GO的基本类型

```go
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // uint8 的别名
rune // int32 的别名 表示一个 Unicode 码点
float32 float64
complex64 complex128
```

1. 本例展示了几种类型的变量。 **同导入语句一样，变量声明也可以“分组”成一个语法块**。
2. `int, uint` 和 `uintptr` 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。 

```
package main
import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64-1
	z complex128 = cmplx.Sqrt(-5+12i)
)

var i, j float32 = 1, 2
var a, b, x, y = 1, 2, true, false // var 申明一个变量列表(类型可省略)
func main() {
	k, w := 1, 2
	fmt.Println(k, w)
	// fmt.Println 打印字符串并换行,fmt.Printf获取变量输出(推荐使用)
	// fmt.Println("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Println("Type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value: %v\n", z, z)
}
```
### 零值
没有明确初始值的变量声明会被赋予它们的**零值**
零值：数值类型为 0；布尔类型为false；字符串为" "空字符串

```go
package main
import "fmt"
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```
### 类型转换
表达式`T(v)`将值`v`转换为类型`T`
一些关于数值的转换：

```go 
var i int = 42
var f float64 = float64(i)
var u int = uint(f)
// 或者，更加简单的形式：
i := 42
f := float64(i)
u := uint(f)
```
与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。试着移除例子中 float64 或 uint 的转换看看会发生什么。

```go
package main
import (
	"fmt"
	"math"
)
func main() {
	var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var u uint = uint(f)
	// var f = math.Sqrt(float64(x*x + y*y))
	// var u = uint(f)
	f := math.Sqrt(float64(x*x + y*y))
	u := uint(f)
	// fmt.Printf(x, f, u)  fmt.Printf针对string
	fmt.Println(x, f, u)
}
```
### 类型推导
在声明一个变量而不指定其类型时(即不使用带`:=`或者`var=`表达式)，变量的类型由右值推导得出。当右值声明了类型时，新变量的类型与其相同：

```go
var i int
j := i // j 也是一个 int
```
不过当右边包含未指明类型的数值常量时，新变量的类型就可能是`int, float64, complex128`, 取决于常量的精度类型:

```go
i := 2  // int
f := 3.142 // float64
g := 0.867 + 0.5i //complex128
package main
import "fmt"
func main() {
	v := 42 // 修改这里！
	fmt.Printf("v is of type %T\n", v)
}
```
### 常量
常量的声明与变量类似，只不过是使用`const`关键字。
常量可以是字符、字符串、布尔值或数值。
常量**不能用**`:=`语法声明。

```go
package main
import "fmt"
const Pi = 3.14
func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```
### 数值常量
数值常量是高精度的 值。
一个未指定类型的常量由上下文来决定其类型。
`int`类型最大可以存储一个 64 位的整数，有时会更小。

```
package main
import (
	"fmt"
)
const (
	// create a huge number by shifting 1 bit left 100 spaces
	// In other words, the binary number that is 1 followed by 100 zeros
	Big = 1 << 100 
	// shift it right again in 99 places, so we end up with 1<<1 or 2. the binary number
	Small = Big >> 99
)
func needInt(x int) int { return 10*x + 1}
func needFloat(x float64) float64 { return x * 0.1}
func main() {
	// fmt.Println(Big)
	fmt.Println(Small)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```

## 流程控制语句
### for
Go 只有一种循环结构：`for`循环。
基本的 for 循环由三部分组成，它们用分号隔开：

1. 初始化语句：在第一次迭代前执行
2. 条件表达式：在每次迭代前求值
3. 后置语句：在每次迭代的结尾执行

[注意] 初始化语句通常为一句**短变量声明**，该变量声明仅在 for 语句的作用域中可见。
一旦条件表达式的布尔值为 false，循环迭代就会终止。
和`C、Java、JavaScript`之类的语言不同，`Go`的`for`语句后面没有小括号，大括号 { } 则是必须的。
初始化语句和后置语句是可选的。

```go
package main
import (
	"fmt"
)
func main() {
	// var sum int = 1
	// var sum = 1
	// var (
	// 	sum int = 1
	// )
	sum := 0;
	for i:=1; i <= 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
// exchange
package main
import (
	"fmt"
)
func main() {
	// var sum = 0
	// var sum int = 0
	// var (
	// 	sum int = 1
	// )
	sum := 1
	for ;sum< 10; {
		sum += sum
		fmt.Println(sum)
	}
}
```
### for 是 Go 中的 “while”
此时你可以去掉分号，因为 C 的`while`在 Go 中叫做`for`。

```go
package main
import (
	"fmt"
)
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```
### 无限循环
如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。

```go
package main
import (
    "fmt"
)
func main() {
    for {
    }
}
```
### if
Go 的 if 语句与 for 循环类似，表达式外无需小括号 `( ) `，而大括号 `{ }` 则是必须的。

```go
package main 
import (
	"fmt"
	"math"
)
func sqrt(x float64) string {
	if x < 0 {
		// return string(math.Sqrt(-x)) + "i" error
		// return math.Sqrt(x) cannot convert float64 into string
		return sqrt(-x)+"i"
	} 
	return fmt.Sprint(math.Sqrt(x)) // use Sprint to convert float64 into string
}
func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```
### if 的简短语句
同 for 一样， if 语句**可以在条件表达式前执行一个简单的语句**。
该语句声明的变量作用域仅在 if 之内。

```go
package main
import (
	"fmt"
	"math"
)
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```
### if 和 else
在 if 的**简短语句中声明的变量同样可以在任何对应的 else 块中使用**。

（在 main 的 `fmt.Println` 调用开始前，两次对 pow 的调用均已执行并返回。）

```go
package main
import (
	"math"
	"fmt"
)
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
} 
```
### 练习：循环与函数
我们来实现一个平方根函数：用牛顿法实现平方根函数。
计算机通常使用循环来计算 x 的平方根。从某个猜测的值 z 开始，我们可以根据 z² 与 x 的近似度来调整 z，产生一个更好的猜测：`z -= (z*z - x) / (2*z)`
重复调整的过程，猜测的结果会越来越精确，得到的答案也会尽可能接近实际的平方根。

```go
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x)/(2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
```

