package main

import "fmt"

var aa = 11

//var aa := 11  /包变量下是不可以使用 := 定义变量

var (
	cc = 22
	dd = 33
)

func variableZeroValue() {

	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDDeduction() {
	var a, b, c, s = 3, 4, true, "def" // 定义变量要 一一 对应初始值, 可以自动推导值
	fmt.Println(a, b, c, s)            // 定义了的变量要使用，否则报错
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def" // := 初始定义变量的时候使用

	fmt.Println(a, b, c, s) // 定义了的变量要使用，否则报错

}

func main() {

	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDDeduction()
	variableShorter()
	fmt.Println(aa, cc, dd)
}
