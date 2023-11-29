package main

import "fmt"

func main() {
	// 创建 map
	m1 := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println("m1", m1)
	m2 := make(map[string]int)
	fmt.Println("m2", m2) // m2 == empty map
	var m3 map[string]int
	fmt.Println("m3", m3) // m3 == nil
	// 遍历的方法
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 取值
	courseName, ok := m1["course"]
	fmt.Println(courseName, ok)

	if causeName, ok := m1["cause"]; ok {

		fmt.Println(causeName)
	} else {
		fmt.Println("key dose not exist")
	}
}
