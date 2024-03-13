package chapter_2

import "fmt"

/*
Q9. 变参
编写函数接受整数类型变参，并且每行打印一个数字。
*/

func Q9Answer() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 2. Answer 9.<<<<<<<<<<<<<<<<<<<")
	uncertainIntArgs(5, 1, 4, 9, 3)
}

func uncertainIntArgs(arg ...int) {
	for _, v := range arg {
		fmt.Println(v)
	}
}
