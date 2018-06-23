# 基础
## 包，变量和函数
### 包
每个 Go 程序都是由**包**构成的。
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

```go
package main
import (
	"fmt"
	"math"
)
func main() {
	// fmt.Println(math.pi); // math.pi undefined 
    fmt.Printlin(math.Pi)
}
```

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
var p *int
value := *p
pointer := &value
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
    a, b := swap("hello", "world") // '' 与 ""不同
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
import (
	"fmt"
)
var c, python, java bool;
var i int;
// var x float32, y int16;   syntax error
func main() {
	fmt.Println(i, c, python, java);
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
总结： **声明变量**必须只用两种

- 使用`var`申明列表, 如`var x, y = 1, "dd";`;
- 而`:=`可以代替`var` 做简洁申明并且必须进行初始化， `python, java := "ss", 2;`

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
	k, w := 1, 2;
	fmt.Println(k, w);
	// fmt.Println 打印字符串并换行,fmt.Printf获取变量输出(推荐使用)
	// fmt.Println("Type: %T Value: %v\n", ToBe, ToBe);
	// fmt.Println("Type: %T value: %v\n", MaxInt, MaxInt);
	fmt.Printf("Type %T value: %v\n", ToBe, ToBe);
	fmt.Printf("Type: %T value: %v\n", MaxInt, MaxInt);
	fmt.Printf("Type: %T value: %v\n", z, z);
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
	fmt.Printf("%v %v %q %q\n", i, f, b, s)// %v 打印值；%T 打印类型，%q 打印全部信息
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

```go
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
### switch
`switch` 是编写一连串 `if - else` 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。
GO的switch语句类似于 `C、C++、Java、JavaScript 和 PHP` 中的，不过 **Go 只运行选定的 case**，而非之后所有的 case。 实际上，Go 自动提供了在这些语言中每个 case 后面所需的 `break` 语句。 除非以 `fallthrough` 语句结束，否则分支会自动终止。 
Go 的另一点**重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。**

```go
package main
import (
	"fmt"
	"runtime"
)
func main() {
	fmt.Print("Go runs on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "runtime":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
```
### switch 的求值顺序
switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。

```go
switch i {
case 0:
case f():
}
```
在 i==0 时 f 不会被调用。）
*注意：* Go 练习场中的时间总是从 2009-11-10 23:00:00 UTC 开始，该值的意义留给读者去发现。

```go
package main 
import (
	"fmt"
	"time"
)
func main() {
	fmt .Println("when is Saturday?");
	today := time.Now().Weekday();
	// fmt.Println(today);
	// fmt.Printf("%v %T\n", today, today);
	// fmt.Printf("%v %T\n", time.Now(), time.Now());
	switch time.Saturday {
		case today + 0 :
			fmt.Println("Today");
		case today + 1 :
			fmt.Println("Tomorrow");
		case today + 2 :
			fmt.Println("in two days");
		default:
			fmt.Println("too far away");
	}
}
```
### 没有条件的 switch
没有条件的 switch 同 `switch true` 一样。
这种形式能将一长串 `if-then-else` 写得更加清晰。

```go
package main
import (
	"time"
	"fmt"
)
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 18:
		fmt.Println("Good afternoon")
	case t.Hour() < 24:
		fmt.Println("Good night")
	default:
		fmt.Println("Good day")
	}
}
```
### defer
defer 语句会将函数推迟到外层函数返回之后执行。
推迟调用的函数其参数会立即求值，但**直到外层函数返回前该函数都不会被调用**。

```go
package main
import (
	"fmt"
)
func main() {
	defer fmt.Println("hello defer")//相对main函数
	fmt.Println("world")
}
```
### defer 栈
推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。

```go
package main
import (
	"fmt"
)
func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
```
## 更多类型 struct, slice and 映射
### 指针
Go 拥有指针。指针保存了值的内存地址(指针就是一个指向地址的数据)。**在go中获得该变量的内存地址 用&a, 别名而已，并没有占用内存空间。实际上他们是同一个东西，在内存中占用同样的一个存储单元。**
**go中所有的都是按值传递，对于复杂类型，传的是指针的拷贝**

1. 类型 `*T` 是指向 `T` 类型值的指针。其零值为 `nil`。
`var p *int`
2. `&` 操作符会生成一个指向其操作数的**指针**(`&`取地址值)。

```go
i :=42
p = &i
```
3. `*` 操作符表示指针指向的**底层值**(`*`取内容值)。

```go
fmt.Println(*p)// 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
// example
package main 
import (
	"fmt"
)
func main() {
	i := 43;
	// p := &i;
	var p *int = &i;
	j := *p;
	j = 1;
	*p = 3;
	fmt.Printf("value: %v type: %T\n", i, i);
	fmt.Printf("value: %v type: %T\n", p, p);
	fmt.Printf("value: %v type: %T\n", j, j);	
}
```
这也就是通常所说的“间接引用”或“重定向”。
与 C 不同，Go 没有指针运算。

```go
package main
import (
	"fmt"
)
func main() {
	// i, j := 43, 8782;
	// // var p *int = &i;
	// p := &i; // pointer p to i
	// fmt.Printf("address: %v value: %v\n", p, *p);
	// fmt.Printf("%p\n", &j)
	
	// Learnin pointer
	var m map[int]string = map[int]string {
		0: "00",
		1: "11",
	};
	mm := m; //deep copy m value to mm
	var values = 4; // 如果是简单的类型，深拷贝
	values2 := values;
	fmt.Printf("values value: %v address:  %p\n", values, &values);
	fmt.Printf("values2 value: %v address:  %p\n", values2, &values2); 

	var mapLiu map[string]string = map[string]string {
		"0": "www",
		"1": "xxxx",
	};
	mapLiu2 := mapLiu; // 如果是复杂类型， 浅拷贝
	fmt.Printf("mapLiu value: %v address:  %p\n", mapLiu, mapLiu);
	fmt.Printf("mapLiu2 value: %v address:  %p\n", mapLiu2, mapLiu2); 

	fmt.Printf("m value: %v address: %p\n",m, m);
	fmt.Printf("mm value: %v address: %p\n", mm, mm);
	// fmt.Printf("m value: %v address: %v\n",m, &m);
	// fmt.Printf("mm value: %v address: %v\n", mm, &mm);
	
	changeMap(m); //(1) go中所有的都是按值传递，对于复杂类型，传的是指针的拷贝``
	// changeMap(&m); // (2) 直接传指针 也是传指针的拷贝
	fmt.Printf("m value: %v address: %p\n",m, m);
	fmt.Printf("mm value: %v address: %p\n", mm, mm);

	// （3）
	// 形参 和 实参
	// param := 3;
	// fmt.Printf("param value: %v address: %p\n",param, &param);
	// changeParam(param);
	// fmt.Printf("param value: %v address: %p\n",param, &param);
}
// (1)
func changeMap(mmm map[int]string) {
	mmm[1] = "eeee";
	fmt.Printf("changeMap func value: %v address: %p\n", mmm, &mmm);
}

// (2)
// func changeMap(mmmm *map[int]string) {
// 	// temp := *mmmm;
// 	// temp[0] = "啛啛喳喳";
// 	// fmt.Printf("func changeMap value: %v address: %p\n", temp, &temp);
// 	mmmm = nil;
// 	// *mmmm = nil;
// 	fmt.Printf("func changeMap value: %v address: %p\n", mmmm, mmmm);
// }

// (3)
// func changeParam(x int) {
// 	x = 2222;
// 	fmt.Printf("func changeParam value: %v address: %p\n", x, &x);
// }
```
### 结构体
一个结构体（struct）就是一个 字段的集合。

```go
package main
import (
	"fmt"
)
type Vertex struct {
	X int
	Y int
}
func main() {
	fmt.Println(Vertex{1, 2})
}
```
### 结构体字段
结构体字段使用点号来访问。 

```go
package main
import (
	"fmt"
)
type Vertex struct {
	X int
	Y int
}
func main() {
	v := Vertex{1, 2}
	v.X = 3
	fmt.Println(v)
}
```
### 结构体指针
1. 结构体字段可以通过**结构体指针**来访问。
2. 如果我们有一个指向结构体的指针 p，那么可以通过 `(*p).X` 来访问其字段 X。不过这么写太啰嗦了，所以语言也允许我们使用**隐式间接引用**，直接写 p.X 就可以。

```go
package main
import (
	"fmt"
)
type Vertex struct {
	X int
	Y int
}
func main() {
	v := Vertex{1, 2}
	p := &v
	// (*p).X = 4
	// fmt.Println(*p)
	p.X = 3e3
	fmt.Println(*p)
}
```
### 结构体文法
1. 结构体文法通过直接列出字段的值来新分配一个结构体。
2. 使用 `Name:` 语法可以仅列出部分字段。（字段名的顺序无关。）
3. 特殊的前缀 `&` 返回一个指向结构体的指针。

```go
package main 
import (
	"fmt"
)
type Vertex  struct {
	X, Y int
}
var (
	v1 = Vertex{1, 2}; // has type Vertex
	v2 = Vertex{X: 1}; // Y: 0 is implicit
	v3 = Vertex{};		// X:0 Y:0
	// p = &Vertex{1, 2} // has type *Vertex
	p = &v1;
)
func main() {
	var pointer *int;
	fmt.Println(pointer);
	fmt.Println(v1, v2, v3, p);
	fmt.Printf("v1 address: %p value: %v\n", &v1, v1);
	fmt.Printf("p address: %p value: %v\n", p, *p); // 指针为同一个地址
}
```

### 数组
类型 `[n]T` 表示拥有 `n 个 T` 类型的值 的数组。
表达式 会将变量 `a` 声明为拥有有 `10` 个整数的数组。
将数组看作一个特殊的struct，结构的字段名对应数组的索引，同时成员的数目固定

```go
var a [10]int
```
数组的长度是其类型的一部分，因此数组不能改变大小。这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。

```go
// array
package main 
import (
	"fmt"
)
func main() {
	var a [2]string; 
	a[0] = "hellp";
	a[1] = "ekedd";
	fmt.Println(a);
	fmt.Println(a[0], a[1]);
	primes := [6]int{2, 3, 5, 7, 9, 11};
	fmt.Println(primes);
	var m map[int]string = map[int]string {
		1: "www0",
		2: "sss",
	};
	fmt.Println(m);
}
```

### 切片(slices)
每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。
类型`[]T`表示一个元素类型为`T`的切片
切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔:
`a[low : high] 即 var s []int = primes[1 : 4]`
他会选择一个半开区间，左闭右开(美国人的习惯)
以下表达式创建了一个切片，包含了a中下标从1到3的元素:
`a[1:3]`

```go
//slices
package main
import (
	"fmt"
)
func main() {
	// var primes [6]int = [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(primes)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1 : 4]
	fmt.Println(s)
}
```

### 切片就像数组的引用(slices are like references to arrays)
切片**并不存储任何数据**，它只是描述了底层数组中的一段(引用)。更改切片的元素会修改其底层数组中对应的元素。与它共享底层数组的切片都会观测到这些修改。

```go
// slices are like references to arrays
package main
import (
	"fmt"
)
func main() {
	// var names [4]string = [4]string{}
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a := names[0 : 2]
	b := names[1 : 3]
	fmt.Println(a, b)
	b[0] = "XXXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```
### 切片文法(slices literals)
切片文法类似于没有长度的数组文法。
这是一个数组文法：
`[3]bool{true, true, false}`
下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：
`[]bool{true, true, false}`

```go
// slices literals
package main
import (
	"fmt"
)
func main() {
	// create array
	var array [4]int = [4]int{1, 2, 3, 4}
	fmt.Printf("value: %v, Type: %T\n", array, array)
	array2 := [4]bool{false, true, false, true}
	fmt.Printf("value: %v Type: %T\n", array2, array2)
	// create slices
	var q []int = []int{3, 4, 6, 7}
	fmt.Printf("value: %v Type:%T\n", q, q)
	q2 := []bool{true, false, true, true, false}
	fmt.Printf("value: %v Type: %T\n", q2, q2)
	// create slices with struct
	var s1 []struct{
		i int
		b bool
	} = []struct {
		i int
		b bool
	}{
		{2, false},
		{3, true},
		{1, true},
		{6, false},
	}
	fmt.Printf("value:%v Type: %T\n", s1, s1)
	s := []struct{
		i int
		b bool
	}{
		{1, true},
		{3, false},
		{5, false},
	}
	fmt.Printf("value:%v Type: %T\n", s, s)
}
```
### 切片的默认行为
在进行切片时，你可以利用它的默认行为来忽略上下界。切片下界的默认值为 0，上界则是该切片的长度。
对于数组
`var a [10]int`
来说，以下切片是等价的：

```go
a[0:10]
a[:10]
a[0:]
a[:]
```

```go
// slices defaults
package main
import (
	"fmt"
)
func main() {
	s := []int{2, 3, 4, 6, 8}
	s1 := s[1:4] 
	fmt.Println(s, s1)
	s2 := s[:2]
	fmt.Println(s, s2)
	s2[0] = 100
	// slices are like references to array
	fmt.Println(s, s1)
	fmt.Println(s, s2)
}
```

### 切片的长度与容量
切片拥有 **长度** 和 **容量**。
切片的长度就是它所包含的元素个数。
**切片的容量**是从**切片后的第一个元素开始数，到其底层数组元素末尾的个数**。
切片 s 的长度和容量可通过表达式 `len(s)` 和 `cap(s)` 来获取。 len()获取的切片的得到的长度，cap是指员原来数组的长度。
你可以通过重新切片来扩展一个切片，给它提供足够的容量。试着修改示例程序中的切片操作，向外扩展它的容量，看看会发生什么。

```go
// slices go
package main 
import (
	"fmt"
)
func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s);
}
func main() {
	s := []int{3, 4, 5,7, 8, 10};
	printSlices(s);
	// Slices the slice to give it zero length
	s = s[:0];
	printSlices(s);
	// extend th length
	s = s[:4];
	printSlices(s);
	// drop its first two values
	s = s[:2];
	printSlices(s); //len=2 cap=6 [3 4]
	s = s[:6];
	printSlices(s); //len=6 cap=6 [3 4 5 7 8 10]
	s = s[3:4];
	printSlices(s); //len=1 cap=3 [7]
}
```
### nil切片
切片的的零值是nil
nil切片的长度和容量为0切没有底层数组。

```go
// nil slices
package main
import (
	"fmt"
)
func main() {
	// var s = []int //error []int is not expression
	// var s []int = []int{1, 2, 3} <==> s := []int{1, 2, 3}
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
}
```
### 用 make 创建切片
切片可以用内建函数 `make` 来创建，这也是你创建动态数组的方式。
`make` 函数会分配一个**元素为零值**的数组并**返回一个引用了它的切片**：
`a := make([]int, 5)  // len(a)=5`
要指定它的容量，需向 make 传入第三个参数：

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

```go
// use make to create slices
package main
import (
	"fmt"
)
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d value=%v type=%T\n", s, len(x), cap(x), x, x)
}
func main() {
	// x := [3]string{"Лайка", "Белка", "Стрелка"}
	// s := x[:] // a slice referencing the storage of x
	// 创建 一个元素值为0的数组并返回一个引用了它的切片
	a := make([]int, 5) // len(a) = 5 cap(a) = 5
	printSlice("a", a)
	// 创建 一个指定容量， 需要向make传入第三个参数
	// b := make([]int, 3, 5)	// len(b)=3, cap(b)=5
	b := make([]int, 0, 5);	// len(b)=0, cap(b)=5
	printSlice("b", b); 	// b len=0 cap=5 value=[] type=[]int
	c := b[:3];
	printSlice("c", c);	// c len=3 cap=5 value=[0 0 0] type=[]int
	d := c[2:5];
	printSlice("d", d); // d len=3 cap=3 value=[0 0 0] type=[]int
}
```
[Note]一个切片是一个数组片段的描述。它包含了指向数组的指针，片段的长度， 和容量（片段的最大长度）。前面使用 `make([]byte, 5)` 创建的切片变量 s ,其中`[]int`指向数组指针；长度是切片引用的元素数目。容量是底层数组的元素数目（从切片指针开始）。 [reference](https://blog.go-zh.org/go-slices-usage-and-internals)
### 切片的切片
切片可包含任何类型，甚至包括其它的切片。

```go
// slices of slices
package main 
import (
	"fmt"
	"strings"
)
func main() {
	// create a tic-tac-toe board
	board := [][]string {
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	};
	// 
	board[0][0] = "X";
	board[2][2] = "O";
	board[1][2] = "X";
	board[1][0] = "O";
	board[0][2] = "X";
	for i:=0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "));
	}
}
```
### 向切片追加元素
为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 `append` 函数。内建函数的文档对此函数有详细的介绍。
`func append(s []T, vs ...T) []T`
append 的第一个参数 `s` 是一个元素类型为 `T` 的切片，其余类型为 T 的值将会追加到该切片的末尾。
append 的结果是一个包含原切片所有元素加上新添加元素的切片。
当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。

```go
// append
package main 
import (
	"fmt"
)
func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d value=%v type=%T\n", len(s), cap(s), s, s);
}
func main() {
	var s []int;
	printSlices(s);
	// append works on nil slices
	s = append(s, 0);
	printSlices(s);
	// the slices grows as needed
	s = append(s, 1);
	printSlices(s);
	// we can add more than one element at time
	s = append(s, 2, 4, 5, 6);
	printSlices(s);
	// s1 := []int{7, 4, 2}; // 不是同一种类型
	// s = append(s, s1);
	// printSlices(s);
}
```
### Range
`for` 循环的 `range` 形式可遍历切片或映射。
当使用 `for` 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。

```go
// Range
package main
import (
	"fmt"
)
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```
可以将下标或值赋予 `_` 来忽略它。
若你只需要索引，去掉 `, value` 的部分即可。

```go
package main
import(
	"fmt"
)
func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // ==2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```
### 练习：切片
实现 `Pic`。它应当返回一个长度为 `dy` 的切片，其中每个元素是一个长度为 dx，元素类型为 `uint8` 的切片。当你运行此程序时，它会将每个整数解释为灰度值（好吧，其实是蓝度值）并显示它所对应的图像。

图像的选择由你来定。几个有趣的函数包括 `(x+y)/2, x*y, x^y, x*log(y)` 和 `x%(y+1)`。

（提示：需要使用循环来分配 `[][]uint8` 中的每个 `[]uint8`；请使用 `uint8(intValue)` 在类型之间转换；你可能会用到 `math` 包中的函数。）

```go
package main

import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	var rgb [][]uint8
	fmt.Println(dx, dy)
	for i := 0; i<dy; i++ {
		for j:=0; j <dx; j++ {
			rgb[i][j] = uint8((i+j)/2)
		}
	}
	return rgb
}

func main() {
	pic.Show(Pic)
}
```
### 映射(map)
映射将key映射到value。
映射的零值为 `nil` 。`nil` 映射既没有键，也不能添加键。
`make`函数会返回给定类型的映射(map)，并将其初始化备用。
`map type is type: map[string]main.Vertex`

```go
// map
package main
import (
	"fmt"
)
type Vertex struct {
	Lat, Long float64
}
// initialize map
var m map[string]Vertex
func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		// 40.22234, -74.33113 // error
		40.22234, -74.33113,
	}
	fmt.Printf("value: %v type: %T\n",m, m)
	fmt.Println(m["Bell Labs"])
}
```
### map的文法(写法)
映射的文法与结构体相似，不过必须有键名。

```go
// map literals
package main
import (
	"fmt"
)
type Vertex struct {
	Lat, Long float64
}
// var s []int
var s = []int{1, 2, 3}
// map 类型 map[string]Vertex
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.441245, -68.421167889,
	},
	"Google": Vertex{
		37.3425, -124.341,
	},
}
func main() {
	fmt.Println(m)
}
```
若顶级类型只是一个类型名，你可以在文法的元素中省略它。

```go
package main
import (
	"fmt"
)
type Vertex struct {
	Lat, Long float64
}
var m = map[string]Vertex{
	// "Bell Labs": Vertex{
	// 	30.4422, -44.44212,
	// },
	// "Google": Vertex {
	// 	39.442, -55.224,
	// },
	"Bell": {-22.3444, 33.312},
	"Google":{-33.442, 33.51},
}
func main() {
	fmt.Println(m)
}
```

### 修改映射(map)

在` map m`中插入或者修改元素：

`m[key] = elem`
获取元素:

`elem = m[key]`
删除元素:

`delete(m, key)`
通过双赋值检测某个键是否存在:

`elem, ok = m[key]`
若key在m中，ok为true；否则, ok为false.
若key不在map中，那么elem是该map的零值。
即当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。
**注**: 若elem或者ok还未声明，你可以使用短变量声明:

` elem, ok := m[key]`

```go
package main 
import (
	"fmt"
)
func main() {
	m := make(map[string]int);
	m["Answer"] = 42;
	fmt.Println("The value:", m["Answer"]);
	m["Answer"] = 48;
	fmt.Println("Thw value:", m["Answer"]);
	delete(m, "Answer");
	fmt.Println("The value:", m["Answer"]);
	v, ok := m["Answer"];
	fmt.Println("The value:", v, "Present?", ok);
}
```
### 函数值
函数也是值。它们可以像其它值一样传递。
函数值可以用作函数的参数或返回值。

```go
package main 
import (
	"math"
	"fmt"
)
func compute(fn func(float64, float64) float64) float64 {
	fmt.Println(fn);
	return fn(3, 4);
}
func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y);
	}
	fmt.Println(hypot(5, 12));
	fmt.Println(compute(hypot));
	fmt.Println(compute(math.Pow));
}
```
### 函数的闭包
Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被“绑定”在了这些变量上。
例如，函数`adder`返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。

```go
package main
import (
	"fmt"
)
func adder() func(int) int {
	sum := 0;
	return func(x int) int {
		sum += x;
		return sum;
	}
}
func main() {
	pos, neg := adder(), adder();
	for i:=0; i<10; i++ {
		fmt.Println(
			pos(i),
			neg(-i),
		);
	}
}
```

### 练习：斐波纳契闭包
实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。

```go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last = 0;
	return func(x int) int {
		
	};
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```
## 方法与接口
### 方法
**Go 没有类**。不过你可以为**结构体类型定义方法**。
方法就是一类带特殊的 **接收者** 参数的函数。
方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
在此例中，Abs 方法拥有一个名为 `v`，类型为 `Vertex` 的接收者。

```go
package main
import (
	"fmt"
	"math"
)
type Vertex struct {
	X, Y float64
}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y);
}
func main() {
	v := Vertex{3, 4};
	fmt.Println(v.Abs());
}
```
### 方法即是函数(Methods(class) are functions)
记住：方法只是个带接收者参数的函数。
现在这个 Abs 的写法就是个正常的函数，功能并没有什么变化。

```go
package main
import (
	"fmt"
	"math"
)
type Vertex struct {
	X, Y float64
}
// define a method
// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y);
// }
// function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y);
}
func main() {
	v := Vertex {3, 4};
	// fmt.Println(v.Abs());
	fmt.Println(Abs(v));
}
```
你也可以为非结构体类型声明方法。如例子中的一个带Abs方法的数值类型MyFloat。
你只能为在同一个包内定义的类型的接受者声明方法，而不能为其他包内定义的类型(包括int之类的内建类型)的接受者声明方法。
(注：就是**接收者的类型定义和方法声明必须在同一个包内**；不能为内建类型声明方法。)

```go
package main
import (
	"math"
	"fmt"
)
type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		 return float64(-f);
	}
	return float64(f);
}
func main() {
	f := MyFloat(-math.Sqrt2);
	fmt.Println(f.Abs());
}
// example 2
package main
import (
	"fmt"
)
type MyInt int
func (Myint MyInt) Abs() int {
	if Myint < 0 {
		return int(-Myint);
	}
	return int(Myint);
}
func main() {
	myint := MyInt(-2);
	fmt.Println(myint.Abs());
}
```
通用：声明的方法为在类中添加方法，必须是初始化类，生成对象调用方法。
### 指针接收者
你可以为指针接收者声明方法。
这意味着对于某类型 T，接收者的类型可以用 *T 的文法。（此外，T 不能是像 *int 这样的指针。）
例如，这里为 *Vertex 定义了 Scale 方法。
指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）。由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
试着移除第 16 行 Scale 函数声明中的 *，观察此程序的行为如何变化。
若使用值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作。（对于函数的其它参数也是如此。）Scale 方法必须用指针接受者来更改 main 函数中声明的 Vertex 的值。

```go
package main 
import (
	"math"
	"fmt"
)
type Vertex struct {
	X, Y float64
}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y);
}
func (v *Vertex) Scale(f float64) {
	// (*v).X = (*v).X * f;
	// (*v).Y = (*v).Y * f;
	v.X *= f;
	v.Y *= f;
	// 
	fmt.Printf("2 value: %v type: %T address: %p\n", v, v, v);
}
// func (v Vertex) Scale(f float64) {
// 	v.X = v.X * f;
// 	v.Y = v.Y * f;
// 	fmt.Printf("2 value: %v type: %T address: %p\n", v, v, &v);
// }
func main() {
	v := Vertex{3, 4};
	fmt.Printf("1 value: %v type: %T address: %p\n", v, v, &v);
	v.Scale(10);
	fmt.Printf("3 value: %v type: %T address %p\n", v, v, &v);
	fmt.Println(v.Abs());
}
```
### 指针与函数
把 Abs 和 Scale 方法重写为函数。接着移除Scale中的*。

```go
package main
import (
	"math"
	"fmt"
)
type Vertex struct {
	X, Y float64
}
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y);
}
func Scale(v *Vertex, f float64) {
	v.X =  v.X * f;
	v.Y = v.Y * f;
} 
func main() {
	v := Vertex{3, 4};
	Scale(&v, 10);
	fmt.Println(Abs(v));
}
```
### 方法与指针重定向
比较前两个程序，你大概会注意到带指针参数的函数必须接受一个指针：

```go
var v Vertex
ScaleFunc(v, 5)  // 编译错误！
ScaleFunc(&v, 5) // OK
```
而以指针为接受者的方法被调用时，接受者既能为值又能为指针：

```go
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```
对于语句 `v.Scale(5)`，即便 v 是个值而非指针，带指针接收者的方法也能被直接调用。 也就是说，由于 Scale 方法有一个指针接收者，为方便起见，Go 会将语句 `v.Scale(5)` 解释为 `(&v).Scale(5)`。

```go
package main 
import (
	"fmt"
)
type Vertex struct {
	X, Y float64
}
// define method
func (v *Vertex) Scale(f float64) {
	// (*v).X *= f;
	// (*v).Y *= f;
	v.X *= f;
	v.Y *= f;
}
// define function
func ScaleFunction(v *Vertex, f float64) {
	v.X *= f;
	v.Y *= f;
}
func main() {
	// v := Vertex{3, 4}
	// v.Scale(2)
	// ScaleFunction(&v, 10)
	// p := &Vertex{4, 3}
	// p.Scale(3)
	// ScaleFunction(p, 8)
	// fmt.Println(v, p)
	v := Vertex{3, 4};
	v.Scale(10);
	fmt.Println("1:",v);
	ScaleFunction(&v, 2);
	fmt.Println("2:", v);
	// p := &v;
	p := &Vertex{4, 3};
	p.Scale(2);
	fmt.Println("3:", p);
	ScaleFunction(p, 10);
	fmt.Println("4:", p);
}
```
### 方法与指针重定向（续）
同样的事情也发生在相反的方向。
接受一个值作为参数的函数必须接受一个指定类型的值：

```go
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // 编译错误！
```
而以值为接收者的方法被调用时，接收者既能为值又能为指针：

```go
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```
这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs()。

```go
package main
import (
	"fmt"
	"math"
)
type Vertex struct {
	X, Y float64
}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y);
}
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y);
}
func main() {
	v := Vertex{3, 4};
	fmt.Println(v.Abs());
	fmt.Println(AbsFunc(v));
	p := &Vertex{4, 3};
	fmt.Println(p.Abs());
	fmt.Println(AbsFunc(*p));
}
```
### 选择值或指针作为接收者
使用指针接收者的原因有二：
首先，方法能够**修改其接收者指向的值**。
其次，这样可以**避免在每次调用方法时复制该值**。若值的类型为大型结构体时，这样做会更加高效。
在本例中，`Scale` 和 `Abs` 接收者的类型为 `*Vertex`，即便 Abs 并不需要修改其接收者。
通常来说，所有给定类型的方法都应该有值或**指针接收者**，但并不应该二者混用。（我们会在接下来几页中明白为什么。）

```go
package main
import (
	"fmt"
	"math"
)
type Vertex struct {
	X, Y float64
}
// define method
func (v *Vertex)Scale(f float64) {
	v.X *= f;
	v.Y *= f;
}
func (v *Vertex)Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y);
}
// define function
func ScaleFunc(v *Vertex, f float64){
	v.X *= f;
	v.Y *= f;
}
func AbsFunc(v *Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y);
}
func main() {
	v :=Vertex{3, 4};
	v.Scale(2); 
	// (&v).Scale(2); equal
	fmt.Println(v);
	// ScaleFunc(v, 20); cannot use v(type Vertex) as type *Vertex
	ScaleFunc(&v, 20);
	fmt.Println(v);

	v2 := &Vertex{4, 5};
	// fmt.Printf("Before scaling: %v, Abs: %v\n", v2, v2.Abs());
	// fmt.Printf("Before scaling: %p, Abs: %v\n", v2, v2.Abs());
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v2, v2.Abs());
	v2.Scale(5);
	fmt.Printf("After Scaling: %+v, Abs: %v\n", v2, v2.Abs());
}
```
### 接口
**接口类型** 是由一组方法签名定义的集合
接口类型的值可以保存任何实现了这些方法的值。

```go
package main
import (
	"fmt"
	"math"
)
type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f);
	}
	return float64(f);
}
type Vertex struct {
	X, Y float64
}
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y);
} 
// define interface
type Abser interface {
	Abs() float64
}
func main() {
	// initlize
	var a Abser;
	f := MyFloat(-math.Sqrt2);
	v := Vertex{3, 4};
	a = f; // a MyFloat implements Abser
	a = &v; // a *Vertex implements Abser
	// 下面一行，v是一个Vertex(而不是 *Vertex)，所以没有实现Abser
	a = v;
	fmt.Println(a.Abs());
}
```
### 接口与隐式实现
类型通过**实现**一个**接口的所有方法** 来实现该接口。既然无需专门显示声明，也就没有“implement”关键字。
隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。
因此，也就无需再每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。

```go
package main
import (
	"fmt"
)
type I interface {
	M();
}
type T struct {
	S string;
}
// This method means type T implements the inteface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S);
}
func main() {
	var i I = T{"hellosdsfs"};
	i.M();
}
```
### 接口值
在内部，接口值可以看做包含**值和具体类型**的元组： `(value, type)`
接口值保存了一个具体底层类型的具体值。
接口值调用方法时会执行其底层类型的同名方法。

```go
package main
import (
	"fmt"
	"math"
)
// 1.声明接口
type I interface {
	M();
}
// 2.声明一个结构体T
type T struct {
	S string
}
// 3.结构体T声明结构体的方法，实现接口的方法，隐式实现
func (t *T) M() {
	fmt.Println(t.S);
}
type F float64
func (f F) M() {
	fmt.Println(f);
}
func main() {
	var i I;
	i = &T{"Hello"};
	describe(i);
	i.M();
   // 接口 有具体类型的类的去实现，方法也相应的类的方法去实现。
	i = F(math.Pi);
	describe(i);
	i.M();
}
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i);
}
```
### 底层值为 nil 的接口值
即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。
在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。
*注意：* 保存了 nil 具体值的接口其自身并不为 nil。

```go
package main
import (
	"fmt"
)
// 1.声明一个接口I
type I interface {
	M();
}
// 2.声明具体类型T 结构体
type T struct {
	S string
}
// 3.具体类型结构体T(reciever) 实现接口的方法(隐式实现)
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>");
		return;
	}
	fmt.Println(t.S);
}
func main() {
	var i I;
	var t *T;
	i = t;
	describe(i);
	i.M();
	i = &T{"Hellowqww"};
	describe(i);
	i.M();
}
func describe(i I) {
	fmt.Printf("(%v %T)\n", i, i);
}
```
### nil 接口值
nil 接口值既不保存值也不保存具体类型。
为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。

```go
package main
import (
	"fmt"
)
type I interface {
	M();
}
func main() {
	var i I;
	describe(i);
	i.M();
}
func describe(i I) {
	fmt.Printf("(%v %T)\n",i, i);
}
```
### 空接口
指定了零个方法的接口值被称为 *空接口：*  `interface{}`
空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
空接口被用来处理未知类型的值。例如，`fmt.Print` 可接受类型为 `interface{}` 的任意数量的参数。

```go
package main
import (
	"fmt"
)
func main() {
	var i interface{};
	describe(i);
	i = 42;
	describe(i);
	i = "hello";
	describe(i);
}
func describe(i interface{}) {
	fmt.Printf("(%v %T)\n", i, i);
}
```
### 类型断言
**类型断言** 提供了访问接口值底层具体值的方式。
`t := i.(T)`
该语句断言接口值 i 保存了具体类型 T，并将其底层类型为 T 的值赋予变量 t。
若 i 并未保存 T 类型的值，该语句就会触发一个panic。
为了 判断 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。
`t, ok := i.(T)`
若 i 保存了一个 T，那么 t 将会是其底层值，而 ok 为 true。
否则，ok 将为 false 而 t 将为 T 类型的零值，程序并不会产生恐慌。
请注意这种语法和读取一个映射时的相同之处。

```go
package main
import (
	"fmt"
)
func main() {
	// 声明一个接口， 
	var i interface{} = "hello";
	fmt.Printf("(%v %T)\n",i, i);
	s := i.(string);
	fmt.Println(s);
	s, ok := i.(string);
	fmt.Println(s, ok);
	f, ok := i.(float64);
	fmt.Println(f, ok);
	f = i.(float64);
	fmt.Println(f);
}
```
### 类型选择
类型选择 是一种按顺序从几个类型断言中选择分支的结构。
类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）， 它们针对给定接口值所存储的值的类型进行比较。

```go
switch v := i.(type) {
case T:
    // v 的类型为 T
case S:
    // v 的类型为 S
default:
    // 没有匹配，v 与 i 的类型相同
}
```
类型选择中的声明与类型断言 `i.(T)` 的语法相同，只是具体类型 T 被替换成了关键字 type。
此选择语句判断接口值 i 保存的值类型是 T 还是 S。在 T 或 S 的情况下，变量 v 会分别按 T 或 S 类型保存 i 拥有的值。在默认（即没有匹配）的情况下，变量 v 与 i 的接口类型和值相同。

```go
package main
import (
	"fmt"
)
func do (i interface{}) {
	switch v := i.(type) {
		case int: 
			fmt.Printf("Twice %v is %v\n", v, v*2);
		case string:
			fmt.Printf("%q is %v bytes long\n", v, len(v));
		default:
			fmt.Printf("I don't know about type %T!\n", v);
	}
}
func main() {
	do(11);
	do("hell9");
	do(true);
}
```
### Stringer
fmt 包中定义的 **Stringer** 是最普遍的接口之一。

```go
type Stringer interface {
    String() string
}
```
Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。

```go
package main
import (
	"fmt"
)
// type Stringer interface {
// 	String() string;
// }
type Person struct {
	Name string
	Age int
}
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age);
}
func main() {
	a := Person{"Arthur Dent", 43};
	z := Person{"Zaphod Beelebrox", 9001};
	fmt.Println(a, z);
}
```
### 练习：Stringer
通过让 `IPAddr` 类型实现 `fmt.Stringer` 来打印点号分隔的地址。
例如，`IPAddr{1, 2, 3, 4}` 应当打印为 `"1.2.3.4"`。

```go
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipAddr IPAddr) String() string {
	
	return fmt.Sprintf("%v.%v.%v.%v\n", ipAddr[0],ipAddr[1],ipAddr[2],ipAddr[3]);
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```
### 错误
Go 程序使用 `error` 值来表示错误状态。
与 `fmt.Stringer` 类似，`error` 类型是一个内建接口：

```go
type error interface {
    Error() string
}
```
（与 fmt.Stringer 类似，fmt 包在打印值时也会满足 error。）
通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil 来进行错误处理。

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```
error 为 nil 时表示成功；非 nil 的 error 表示失败。

```go
package main
import (
	"time"
	"fmt"
)
type MyError struct {
	When time.Time
	What string
}
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What);
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	};
}
func main() {
	if err := run(); err != nil {
		fmt.Println(err);
	}
}
```


