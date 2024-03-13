package chapter_1

import "fmt"

/*
Q4. 平均值
编写计算一个类型是 float64 的 slice 的平均值的代码。
*/

func Q4Answer() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 1. Answer 4.<<<<<<<<<<<<<<<<<<<")
	calcSliceAverage([]float64{3.14, 5.8, 7.2, 10.5})
}

func calcSliceAverage(slice []float64) {
	var sum float64 = 0.0
	for _, val := range slice {
		sum += val
	}

	fmt.Println("Average of the slice: ", sum/float64(len(slice)))
}
