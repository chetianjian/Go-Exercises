package chapter_2

import "fmt"

/*
Q12. 最小值和最大值
1. 编写一个函数，找到 int slice ([]int) 中的最大值。
2. 编写一个函数，找到 int slice ([]int) 中的最小值。
*/

func Q12Answer() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 2. Answer 12.1.<<<<<<<<<<<<<<<<<<<")
	var unsortedList = []int{-1, 2, 0, -3, 5, -7, 10, 1}
	fmt.Println("Input list is: ", unsortedList)
	fmt.Println("Max value is: ", findMaxValue(unsortedList))
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 2. Answer 12.2.<<<<<<<<<<<<<<<<<<<")
	fmt.Println("Input list is: ", unsortedList)
	fmt.Println("Min value is: ", findMinValue(unsortedList))
}

func findMaxValue(list []int) int {
	maxVal := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] > maxVal {
			maxVal = list[i]
		}
	}
	return maxVal
}

func findMinValue(list []int) int {
	minVal := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] < minVal {
			minVal = list[i]
		}
	}
	return minVal
}
