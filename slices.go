// 切片
package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 =", s2)

	fmt.Println("After  updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")

	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Println("s1", s1)
	fmt.Println("s2", s2) // 扩展到 s2[2] = s1[4]  = 7  , 底层数组的7

	// 打印切片长度 ，扩展数组长度
	arr[0], arr[2] = 0, 2
	fmt.Println("arr", arr)
	fmt.Printf("s1=%v , len(s1)=%d , cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v , len(s2)=%d , cap(s2)=%d\n", s2, len(s2), cap(s2))
	// append 的操作 , 必须接收 append 的返回值 ,  超越 cap 会重新分配更大的底层数组
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3", s3)
	fmt.Println("s4", s4)
	fmt.Println("s5", s5)

}
