package chapter_2

import "fmt"

/*
Q13. 冒泡排序
编写一个针对int类型的slice冒泡排序的函数。
*/

func Q13Answer() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 2. Answer 13.<<<<<<<<<<<<<<<<<<<")
	unsortedList := []int{45, 32, 8, 33, 12, 22, 19, 97}
	fmt.Println("Input unsorted list: ", unsortedList)
	fmt.Println(bubbleSort(unsortedList))
}

func bubbleSort(list []int) []int {
	length := len(list)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if list[j] > list[j+1] {
				tmp := list[j]
				list[j] = list[j+1]
				list[j+1] = tmp
			}
		}
	}
	return list
}
