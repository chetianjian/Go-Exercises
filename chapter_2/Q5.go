package chapter_2

import "fmt"

/*
Q5. 平均值函数
编写一个函数用于计算一个float64类型的slice的平均值。
*/

func Q5Answer() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 2. Answer 5.<<<<<<<<<<<<<<<<<<<")
	calcSliceAverage([]float64{3.14, 5.8, 7.2, 10.5})
}

func calcSliceAverage(slice []float64) {
	var sum float64 = 0.0
	for _, val := range slice {
		sum += val
	}

	fmt.Println("Average of the slice: ", sum/float64(len(slice)))
}
