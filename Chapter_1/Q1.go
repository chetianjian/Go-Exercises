package main

import "fmt"

/*
Q1. For-loop
1. 创建一个基于 for 的简单循环。使其循环 10 次，并且使用 fmt 包打印出计数器的值。
2. 用 goto 改写 1. 的循。关键字 for 不可使用。
3. 再次改写这个循环，使其遍历一个 array，并将这个 array 打印在屏幕上。
*/

func main() {
	var iter_num int = 10
	var arr = [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Answer 1.<<<<<<<<<<<<<<<<<<<")
	forLoopQuestion1(iter_num)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Answer 2.<<<<<<<<<<<<<<<<<<<")
	forLoopQuestion2(iter_num)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Answer 3.<<<<<<<<<<<<<<<<<<<")
	forLoopQuestion3(arr)
}

func forLoopQuestion1(iter_num int) {
	for i := 0; i < iter_num; i++ {
		fmt.Println("Iteration count: ", i)
	}
}

func forLoopQuestion2(iter_num int) {
	var i int
	i = 0
ifLabel:
	if i < iter_num {
		goto loopLabel
	} else {
		return
	}

loopLabel:
	fmt.Println("Iteration count: ", i)
	i++
	goto ifLabel
}

func forLoopQuestion3(arr [6]int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}
