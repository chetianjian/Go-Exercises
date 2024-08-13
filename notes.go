package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello, world!")
	convertDataType()
	dataTypeConvert2String()
	getVarAddressType()
	getVarAddressModifyVal()
	erorr1Modify()
	error2Modify()
	error3Modify()
	selfIncrement()
	intDivision()
	modExample()
	holidaySolution()
	fahrenheit2Celsius()
	relationalOperators()
	ageComparison()
	exchangeTwoVars()
	gradeReward(65)
	number2Letter('d')
	ageMatch(29)
	switchFallthrough(10)
	forLoopPrint()
	secondTypeForLoop()
	thirdTypeForLoop()
	stringForLoopType1()
	stringForLoopType2()
	arrayCreation()
	arrayInitialization()
	arrayInitializationByIndex()
	multiArrayCreation()
	multiArrayEasierCreation()
	sliceCreation()
	sliceCap()
	sliceAppend()
	sliceCopy()
	mapMonthDayNum()
	deferTest()
	inputArgs(2, 3, 1, 5, 0)
	lambdaFunc()
	nilPrint()
	goroutineChan()
	goroutineChanSelect()
	afterFuncTest()
	test()
}

// 基本数据类型的转换
func convertDataType() {
	var (
		a  int32
		n1 float32
		n2 int8
		n3 int64
	)
	n1 = float32(a)
	n2 = int8(a)
	n3 = int64(a)
	fmt.Printf("a=%v n1=%v n2=%v n3=%v", a, n1, n2, n3)
}

// 基本数据类型转换为string使用
func dataTypeConvert2String() {
	var (
		num1   int
		num2   float64
		b      bool
		myChar byte
		str    string
	)
	num1 = 99
	num2 = 23.456
	b = true
	myChar = 'h'

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q\n", str, str)
}

// 获取变量的地址
// & 获取地址
// * 获取值
func getVarAddressType() {
	var i int = 10
	// i 的地址是什么，&i
	fmt.Println("var i address=", &i)

	// ptr 是一个指针变量
	// ptr 的指针类型：*int
	// ptr 的值： &i
	var ptr *int = &i // 使用 *ptr 获取 ptr 指向的值
	fmt.Println("ptr=", ptr)
	fmt.Println("ptr 的地址=", &ptr)
	fmt.Println("ptr 指向的值=", *ptr)
}

// 写一个程序，获取一个 int 变量 num 的地址，并显示到终端
// 将num的地址赋给指针 ptr，并通过 ptr 去修改 num 的值
// （通过指针修改变量的值）
func getVarAddressModifyVal() {
	var num int
	num = 99
	fmt.Println("num 的地址=", &num)
	var ptr *int
	ptr = &num
	*ptr = 10 // 修改指针 ptr 对应地址的值
	fmt.Println("修改指针 ptr 后的 num 值=", num)
}

func erorr1Modify() {
	var a int = 300
	var ptr *int = &a
	fmt.Println(ptr)
}

func error2Modify() {
	var a int = 300
	var ptr *int = &a
	fmt.Println(ptr)
}

func error3Modify() {
	var a int = 300
	var b int = 400
	var ptr *int = &a
	*ptr = 100
	ptr = &b
	*ptr = 200
	fmt.Printf("a=%d, b=%d, *ptr=%d", a, b, *ptr)
}

/*
值类型和引用类型的说明：
1. 值类型：基本数据类型 int、float、bool、string、数组、struct
	变量直接存储值，内存通常在栈中分配

2. 引用类型：指针、slice切片，map、管道chan、interface 等

	变量存储的是一个地址，这个地址对应的空间才是真正存储数据（值）。内存通常在堆上分配，当没有任何变量引用这个地址时，
		该地址对应的数据空间就成为一个垃圾，由 GC 来回收。
*/

func selfIncrement() {
	var a int
	a = 1
	a++
	fmt.Println("\n", a)
}

// 对于除号 /， 在整数之间做除法时，只保留整数部分而舍弃小数部分。例如：
func intDivision() {
	var a int
	a = 19
	a = a / 5
	fmt.Println(a)
}

func modExample() {
	var (
		a int
		b int
	)
	a = 13
	b = 5
	fmt.Println(a % b)
}

// 假如还有 97 天假，问这是多少星期余多少天？
func holidaySolution() {
	var (
		holiday int
		weekNum int
		dayNum  int
	)
	holiday = 97
	weekNum = holiday / 7
	dayNum = holiday % 7
	fmt.Printf("week: %v, day: %v", weekNum, dayNum)
}

// 定义一个变量保存华氏温度，华氏温度转换摄氏温度等公式为：5 / 9 * (华氏温度 - 100)，请求出华氏温度对应的摄氏温度。
func fahrenheit2Celsius() {
	var degree float64
	degree = 132.4
	degree = 5.0 / 9 * (degree - 100) // 注意：是 5.0 / 9 ！！！
	fmt.Println("\n", degree)
}

/*
关系运算符
	1. 关系运算符的结果的都是 bool 型，要么是 true，要么是 false
	2. 关系表达式经常用在 if 结构的条件中或循环结构的条件中

	符号和 Python 一致
*/

func relationalOperators() {
	var n1 int = 9
	var n2 int = 8
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 <= n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 < n2)
}

/*
逻辑运算符
	用于连接多个条件（一般来说是关系表达式），最终的结果也是一个 bool 值

	与	&&
	或	||
	非	!
*/

func ageComparison() {
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok")
	}
}

/*
赋值运算符
	赋值运算符就是将某个运算后的值，赋给指定的变量
	=
	+=
	-=
	*=
	/=
	%=
	<<=		(左移后赋值)
	>>=		(右移后赋值)
	&=	(按位与后赋值)
	^=	(按位异或后赋值)
	|=	(按位或后赋值)
*/

func exchangeTwoVars() {
	var a int = 9
	var b int = 2
	fmt.Printf("交换前的情况是 a=%v, b=%v \n", a, b)

	t := a // 创建一个临时变量
	a = b
	b = t
	fmt.Printf("交换后的情况是 a=%v, b=%v \n", a, b)

	// 复合赋值的操作
	a += 17 // a = a + 17
	fmt.Println("a=", a)
}

/*
if statement
	GoLang 中的 if statement 基本格式为：

if a < b {
	...
} else if b > c {
	...
} else {
	...
}
注意到，与 Python 不同的地方在于 elif -> else if，且明确规定 else if / else 不能换行
*/

// 考试成绩为 100 分，奖励 BMW
// 为 (80, 99] 时，奖励 iphone
// 为 [60, 80] 时，奖励 ipad
// 其它时，什么也没有

func gradeReward(grade float32) {
	if grade == 100 {
		fmt.Println("BMW")
	} else if 80 < grade && grade <= 99 {
		fmt.Println("iphone")
	} else if 60 <= grade && grade <= 80 {
		fmt.Println("ipad")
	} else {
		fmt.Println("nothing")
	}
}

/*
switch 分支控制
	switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上到下逐试，直到匹配为止。
	匹配项后看也不需要加 break
*/

// 编写一个程序，该程序可以接收一个字符，比如：
// a, b, c, d, e, f, g
// a 表示星期一，b 表示星期二，...，根据用户的输入显示相应的信息
func number2Letter(key byte) {
	// var key byte
	// fmt.Println("请输入一个字符: a, b, c, d, e, f, g")
	// fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("Monday")
	case 'b':
		fmt.Println("Tuesday")
	case 'c':
		fmt.Println("Wednesday")
	case 'd':
		fmt.Println("Thursday")
	case 'e':
		fmt.Println("Friday")
	case 'f':
		fmt.Println("Saturday")
	case 'g':
		fmt.Println("Sunday")
	default:
		fmt.Println("Input error.")
	}
}

/*
switch 注意事项
	1. case / switch 后是一个表达式，即 常量、变量、一个有返回值的函数等
	2. case 后各个表达式的值的数据类型，必须和 switch 的表达式数据类型一致
	3. case 后面可以带多个表达式，使用逗号间隔。比如 case1, case2, ...
	4. case 后面的表达式如果是常量，则 case 之间不能重复
	5. case 后面不需要带 break，程序匹配到一个 case 就会执行对应的代码块，然后退出 switch，如果匹配不到则执行 default
	6. default 不是必须的
	7. switch 后也可以不带表达式，类似 if-else 分支使用
	8. switch 后也可以直接声明 / 定义一个变量，分号结束（类似于 Python 中的海象表达式，不推荐）
	9. switch 穿透 fallthrough，如果在 case 语句块后面增加 fallthrough，
		则会继续执行下一个 case，也叫 switch 穿透
	10. Type Switch: switch 语句还可以被用于 type-switch 来判断某个
		interface 变量中实际指向的变量类型
*/

// 7. Example
func ageMatch(age int) {
	switch {
	case age == 10:
		fmt.Println("age=10")
	case age == 20:
		fmt.Println("age=20")
	default:
		fmt.Println("没有匹配到")
	}
}

func switchFallthrough(num int) {
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough // 默认只能穿透一层
	case 20:
		fmt.Println("ok2")
		fallthrough
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
	}
}

/*
什么时候使用 switch，什么时候使用 if？
	1. 如果判断的具体数值不多，而且符合整数、浮点数、字符、字符串这几种类型，
		建议使用 switch 语句，简洁高效
	2. 其它情况：对区间判断和结果为 bool 类型的判断，使用 if，if 的使用范围更广
*/

/*
for 循环

	语法格式：

	for 循环变量初始化；循环条件；循环变量迭代 {
		循环操作
	}
*/

func forLoopPrint() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello, world!")
	}
}

/*
for 循环的循环条件是返回一个布尔值的表达式，例如上例中的: i<=10
for 循环还有第二种写法，类似于 Python 中的 while loop （GoLang 中没有 while loop 的关键字）
*/

func secondTypeForLoop() {
	var j int
	j = 1 // 循环变量初始化
	for j <= 10 {
		fmt.Println("hello, GoLang!")
		j++
	}
}

// for 循环的第三种写法
func thirdTypeForLoop() {
	var k int
	k = 1
	for { // 等价于 for ;; {}，是一个死循环
		if k <= 10 {
			fmt.Println("hello, Aristeia!", k)
		} else {
			break // 在 for 循环内部写退出循环的条件，满足则 break 循环
			// 如果不加这个条件，则 for loop 成为死循环
		}
		k++
	}
}

// 传统遍历字符串的形式
func stringForLoopType1() {
	var str string = "hello world!"
	// 这说明 GoLang 中的索引也是从 0 开始的
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}
}

func stringForLoopType2() {
	var str string
	str = "Aristeia"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}
}

/*
array, slices & map
可以利用 array 在列表中进行多个值的排序，或者使用更加灵活的 slice，
字典或哈希类型同样可以使用，在 Go 中叫做 map
*/

/*
array 由 [n]<type> 定义，n 表示 array 的长度，<type> 表示希望存储
的内容的类型。对 array 的元素赋值或索引是由方括号完成的，如下例：
*/

func arrayCreation() {
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Println(arr) // [42 13 0 0 0 0 0 0 0 0]
}

// 不想用 0 初始化数组
func arrayInitialization() {
	a := [...]int{1, 2, 3}
	fmt.Println(a)
}

// 根据下标初始化数组
func arrayInitializationByIndex() {
	// index 1 为 2.0，index 3 为 7.0
	balance := [5]float32{1: 2.0, 3: 7.0}
	fmt.Println(balance)
}

// 注意，如果使用多维数组，所有项目必须都指定
func multiArrayCreation() {
	arr := [3][2]int{[...]int{1, 2}, [...]int{3, 4}, [...]int{5, 6}}
	fmt.Println(arr)
}

/*
更简便的写法来自于更新日志：
	array、slice 和 map 的复合声明变得更加简单。
	使用复合声明的 array、slice 和 map，元素复合声明的类型与外部一致，
	则可以省略。
*/

func multiArrayEasierCreation() {
	arr := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(arr)
}

/*
slice 与 array 接近，但是在新的元素加入的时候可以增加长度。
slice 总是指向底层的一个 array。
slice 是一个指向 array 的指针，这是与 array 不同的地方。
slice 是引用类型，这意味着当赋值某个 slice 到另一个变量，两个引用会指向同一个 array
*/

// 创建一个保存了 10 个元素的 slice。
func sliceCreation() {
	sl := make([]int, 10)
	fmt.Println(sl)
}

func sliceCap() {
	arr := []int{2, 3, 5, 7, 11, 13}
	sli := arr[1:4] // 虽然索引在1:4的位置，但GoLang会截取索引 1 后的所有元素
	fmt.Println(sli)
	fmt.Println(len(sli))
	fmt.Println(cap(sli)) // 因此 cap(sli) = 5
}

func sliceAppend() {
	var s0 = []int{0, 0}      // 注意，与 array 不同，slice 的创建不需要 ...
	s1 := append(s0, 2)       // s1 = []int{0, 0, 2}
	s2 := append(s1, 3, 5, 7) // s2 = []int{0, 0, 2, 3, 5, 7}
	s3 := append(s2, s0...)   // s3 = []int{0, 0, 2, 3, 5, 7, 0, 0}
	fmt.Println("s0=", s0)
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("s3=", s3)
}

func sliceCopy() {
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([]int, 6)
	// copy 函数从 a[0:] 复制到 s 中，复制的元素数量 a[0:] 和 s 的长度最小值
	// a[0:] = []int{0, 1, 2, 3, 4, 5, 6, 7}，8个元素，s 只有 6 个元素，故复制前 6 个
	n1 := copy(s, a[0:]) // s = []int{0, 1, 2, 3, 4, 5}, n1 = 6
	n2 := copy(s, s[2:]) // s = []int{2, 3, 4, 5, 4, 5}, n2 = 4
	fmt.Println("n1=", n1)
	fmt.Println("n2=", n2)
}

/*
map: 字符串作为索引的数组（类似于 Python 中的 dictionary）
*/

func mapMonthDayNum() {
	var monthdays = map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31, "Apr": 30,
		"May": 31, "Jun": 30, "Jul": 31, "Aug": 31,
		"Sep": 30, "Oct": 31, "Nov": 30, "Dec": 31, // 最后的逗号是必须的
	}
	fmt.Println(monthdays)

	// 闰年重写 Feb
	monthdays["Feb"] = 29
	fmt.Println(monthdays)

	// Loop map
	for month, day := range monthdays {
		fmt.Println(month, day)
	}

	// 检查 map 中某元素是否存在
	_, present := monthdays["hello"]
	fmt.Println(present) // 如果存在返回 true

	// 从 map 中移除元素
	delete(monthdays, "Mar")

	fmt.Println(monthdays)
}

// In GoLang, the statement after the keyword `defer`, would be
// executed before the function completes, following the
// `First In Last Out` rule.
func deferTest() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	} // would print 4 3 2 1 0
}

// args 函数接受不定数量的参数，这些参数的类型全部是 int
// 在这个例子中，变量 arg 是一个 int 类型的 slice
func inputArgs(arg ...int) {
	for k, v := range arg {
		fmt.Println(k, v)
	}
}

// 匿名函数（类似于 Python 中的 lambda 函数）
func lambdaFunc() {
	add := func(num1 int, num2 int) int {
		return num1 + num2
	}
	res := add(2, 3)
	fmt.Println(res)
}

func nilPrint() {
	var ptr *int
	fmt.Println(ptr)
}

/*
使用 goroutine 时，如果不等待 goroutine 的执行，程序立即终止，并且
任何正在执行的 goroutine 都会停止。为了修复这个，需要一些能够同 goroutine
通讯的机制，这一机制通过 channels 的形式使用。channel 类似于 Python
中的 Queue，但是是阻塞式的循环队列，相当于queue.put() 而不是
queue.put_nowait()

chan <- 1：向 chan 中压入 1
<- chan：从 chan 中压出一个元素
a := <- chan：从 chan 中压出一个元素，并赋值给变量 a
*/

// 声明一个 channel
var ci chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	ci <- 1
}

func goroutineChan() {
	ci = make(chan int) // var ci chan int //
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting, but not too long.")
	<-ci
	<-ci
}

/*
丑陋的地方在于，不得不对 channel 进行两次压出操作（因为有两个 goroutine），
如果不知道启动了多少个 goroutine 怎么办呢？

select 关键字，通过 select 可以监听 channel 上输入的数据。
*/

func goroutineChanSelect() {
	ci = make(chan int)
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting, but not too long.")
	for i := 0; ; {
		select {
		case <-ci:
			i++
			if i > 1 {
				return // 不能用 break，因为 break 的是 select
			}
		}
	}
}

func afterFuncTest() {
	duration := time.Duration(3) * time.Second
	f := func() {
		fmt.Println("Function called by AfterFunc() after 3 seconds.")
	}
	Timer1 := time.AfterFunc(duration, f)
	defer Timer1.Stop()
	time.Sleep(10 * time.Second) // to block the main rountine
}

type e interface {
	Add()
	Put()
	Pop()
}

func test() {
	t := float64(time.Second)
	fmt.Println(t)
}
