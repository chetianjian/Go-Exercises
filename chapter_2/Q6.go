package chapter_2

import "fmt"

/*
Q6. 整数排序
编写函数，返回其(两个)参数正确的(自然)数字顺序: f(7,2) -> 2,7
f(2,7) -> 2,7
*/

func Q6Answer() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 2. Answer 6.<<<<<<<<<<<<<<<<<<<")
	low, high := intSort(7, 2)
	fmt.Println("low: ", low, "high: ", high)
}

func intSort(int1 int, int2 int) (int, int) {
	if int1 < int2 {
		return int1, int2
	} else {
		return int2, int1
	}
}
