package chapter_2

import "fmt"

/*
Q14. 函数返回一个函数
1. 编写一个函数返回另一个函数，返回的函数的作用是对一个整数+2。
	函数的名称叫做 plusTwo。然后可以像下面这样使用:
		p := plusTwo()
    	fmt.Printf("%v\n", p(2))
		应该打印 4。

2. 使 1 中的函数更加通用化，创建一个 plusX(x) 函数，返回一个函数用于对整数加上 x。
*/

func Q14Answer() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 2. Answer 14.1.<<<<<<<<<<<<<<<<<<<")
	p1 := plusTwo()
	var val int = 5
	fmt.Println("Input val: ", val)
	fmt.Println(p1(val))

	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 2. Answer 14.2.<<<<<<<<<<<<<<<<<<<")
	var x int = 17
	fmt.Println("Input x: ", x)
	fmt.Println("Input val: ", val)
	p2 := plusX(x)
	fmt.Println(p2(val))
}

func plusTwo() func(p int) int {
	return func(p int) int { return p + 2 }
}

func plusX(x int) func(p int) int {
	return func(p int) int { return p + x }
}
