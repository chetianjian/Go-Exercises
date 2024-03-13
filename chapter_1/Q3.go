package chapter_1

import (
	"fmt"
	"unicode/utf8"
)

/*
Q3. 字符串
1. 建立一个 Go 程序打印下面的内容（到 100 个字符）
A
AA
AAA
AAAA
AAAAA
AAAAAA
...

2. 建立一个程序统计字符串里的字符数量：
	asSASA ddd dsjkdsjs dk
同时输出这个字符串的字节数。提示：看看 unicode/utf8 包

3. 扩展 / 修改上一个问题的程序，替换位置 4 开始的三个字符为 “abc”。

4. 编写一个 Go 程序可以逆转字符串，例如 “foobar” 被打印成 “raboof”。
   提示: 不幸的是你需要知道一些关于转换的内容。
*/

func Q3Answer() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 1. Answer 3.1.<<<<<<<<<<<<<<<<<<<")
	printAPyramid()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 1. Answer 3.2.<<<<<<<<<<<<<<<<<<<")
	countCharacters()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 1. Answer 3.3.<<<<<<<<<<<<<<<<<<<")
	stringReplace()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>Chapter 1. Answer 3.4.<<<<<<<<<<<<<<<<<<<")
	stringInvert("foobar")
}

func printAPyramid() {
	var strA string = ""
	var counter int = 0
	for {
		strA += "A"
		fmt.Println(strA)
		counter += len(strA)
		if counter >= 100 {
			break
		}
	}
}

func countCharacters() {
	var str string = "asSASA ddd dsjkdsjs dk"
	var dict = map[string]int{}
	for _, char := range str {
		char := string(char)
		_, present := dict[char]
		if char == " " {
			continue
		}
		if !present {
			dict[char] = 1
		} else {
			dict[char] += 1
		}
	}
	for char, val := range dict {
		fmt.Println(char, val)
	}
	fmt.Println("Total character number: ", utf8.RuneCountInString(str))
	fmt.Println("Total byte number: ", len(str))
}

func stringReplace() {
	var str string = "asSASA ddd dsjkdsjs dk"
	var r = []rune(str)
	copy(r[4:4+3], []rune("abc"))
	fmt.Println("After replacing: ", string(r))
}

func stringInvert(str string) {
	var emptyStr string = ""
	var lengthOfStr int = utf8.RuneCountInString(str)
	for i := 1; i <= lengthOfStr; i++ {
		emptyStr += string(str[lengthOfStr-i])
	}
	fmt.Println("Inverted string: ", emptyStr)
}
