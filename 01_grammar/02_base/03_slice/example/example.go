package main

import "fmt"

func main() {
	// 10 : 2 表示下标为10的元素为2
	// []int 切片， [...]int 自动推断长度的数组
	data := [...]int{0, 1, 2, 3, 4, 10: 2}
	slice := data[:]
	for index, value := range slice {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}
	// [:2]表示下标0~1，后面:3表示新的切片的容量为3
	s := data[:2:3]
	// 往里加俩值，超出容量
	s = append(s, 100, 200)
	fmt.Println(s, data)
	fmt.Println(&s[0], &data[0])

	data2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(data2)
	s1 := data2[8:]
	s2 := data2[:5]
	fmt.Printf("slice s1: %v\n", s1)
	fmt.Printf("slice s2: %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1: %v\n", s1)
	fmt.Printf("copied slice s2: %v\n", s2)
	fmt.Println("last array data: ", data2)
}
