package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "我爱小琳琳"
	// 字符串长度
	fmt.Println("length", utf8.RuneCountInString(s))
	// 遍历每个字符
	for i, ch := range []rune(s) {

		fmt.Printf("(%d %c)", i, ch)

	}
}
