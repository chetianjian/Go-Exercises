package chapter_4

import "fmt"

/*
Q18. 使用 interface 的 map 函数
使用练习 Q11. 的答案，利用 interface 使其更加通用。让它至少能同时工作于
int 和 string。
*/

func Q18Answer() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 4. Answer 18.1.<<<<<<<<<<<<<<<<<<<")
	m := []data{1, 2, 3, 4}
	s := []data{"a", "b", "c", "d"}
	mf := Map(m, mult2)
	sf := Map(s, mult2)
	fmt.Printf("%v\n", mf)
	fmt.Printf("%v\n", sf)
}

type data interface{}

func mult2(d data) data {
	switch d := d.(type) {
	case int:
		return d * 2
	case string:
		return d + d
	}
	return d
}

func Map(d []data, f func(data) data) []data {
	m := make([]data, 0)
	for _, val := range d {
		m = append(m, f(val))
	}
	return m
}
