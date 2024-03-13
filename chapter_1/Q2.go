package chapter_1

import "fmt"

/*
Q2. Fizz-Buzz
编写一个程序，打印从 1 到 100 的数字。
当是 3 的倍数就打印 “Fizz” 代替数字，
当是 5 的倍数就打印 “Buzz”。
当数字同时是 3 和 5 的倍数 时，打印 “FizzBuzz”。
*/

func Q2Answer() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 1. Answer 2.<<<<<<<<<<<<<<<<<<<")
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		}
	}
}
