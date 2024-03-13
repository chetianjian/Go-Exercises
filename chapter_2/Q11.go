package chapter_2

import "fmt"

/*
Q11. map 函数
map() 函数是一个接受一个函数和一个列表作为参数的函数。
函数应用于列表中的每个元素，而一个新的包含有计算结果的列表被返回。因此:
map(f(), (a1, a2, ..., an−1, an)) = (f(a1), f(a2), ..., f(an−1), f(an))

1. 编写 Go 中的简单的 map() 函数。它能工作于操作整数的函数就可以了。

2. 扩展代码使其工作于字符串列表。
*/

func Q11Answer() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 2. Answer 11.1.<<<<<<<<<<<<<<<<<<<")
	fmt.Println(intMap([]int{1, 2, 3}, func(a int) int { return a + 1 }))
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 2. Answer 11.2.<<<<<<<<<<<<<<<<<<<")
	fmt.Println(strMap([]string{"ab", "bc", "cd"}, func(s string) string { return string(s[1]) + string(s[0]) }))
}

func intMap(list []int, f func(int) int) []int {
	result := make([]int, 0)
	for _, val := range list {
		result = append(result, f(val))
	}
	return result
}

func strMap(list []string, f func(string) string) []string {
	result := make([]string, 0)
	for _, val := range list {
		result = append(result, f(val))
	}
	return result
}
