package chapter_2

import "fmt"

/*
Q10. 斐波那契
斐波那契数列以: 1, 1, 2, 3, 5, 8, 13, ...开始。
或者用数学形式表达: x1 = 1; x2 = 1; xn = xn−1 + xn−2; ∀n > 2。
编写一个接受 int 值的函数，并给出这个值得到的斐波那契数列。
*/

func Q10Answer() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 2. Answer 10.<<<<<<<<<<<<<<<<<<<")
	result := fibonacci(100)
	fmt.Println(result)
}

// Use dynamic programming
func fibonacci(n int) int {
	dp := make([]int, 0)
	dp = append(dp, 1)
	dp = append(dp, 1)
	for i := 2; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}
