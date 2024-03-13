package chapter_2

import "fmt"

/*
Q8. 栈
1. 创建一个固定大小保存整数的栈。它无须超出限制的增长。
定义 push 函数：将数据放入栈；以及 pop 函数：从栈中取得内容。
栈应当是后进先出(LIFO) 的。
*/

func Q8Answer() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 2. Answer 8.1.<<<<<<<<<<<<<<<<<<<")
	stack := Stack{
		data: make([]int, 0),
	}
	stack.push(1)
	fmt.Println(stack)
	stack.push(3)
	fmt.Println(stack)
	val := stack.pop()
	fmt.Println(stack)
	fmt.Println("value poped: ", val)
}

type Stack struct {
	data []int
}

func (stack *Stack) pop() int {
	num := stack.data[0]
	stack.data = stack.data[1:]
	return num
}

func (stack *Stack) push(num int) {
	stack.data = append(stack.data, num)
}
