package main

import "fmt"

// slice 的操作

func printSlice(s []int) {
	fmt.Println(s, "len", len(s), " cap ", cap(s))

}
func main() {

	//     ============================ 创建 ==========================

	var s []int

	for i := 0; i < 100; i++ {
		//printSlice(s)
		s = append(s, i*2+1)
	}

	//fmt.Println(s)
	// 给切片赋值
	s1 := []int{1, 2, 3, 4, 5, 6}
	printSlice(s1)
	// 指定切片 len , cap 通过 make 进行创建
	// 创建一个 16长度的切片
	s2 := make([]int, 16)
	printSlice(s2)
	// 创建一个 len 16, cap 32 的切片
	s3 := make([]int, 16, 32)
	printSlice(s3)
	//     ============================ 拷贝 ==========================
	copy(s2, s1)
	printSlice(s2)
	// 删除元素
	fmt.Println("Deleting elments from slice")
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println("s2", s2)
	// 弹出头部和尾部的值
	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("front", front)
	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("tail", tail)
}
