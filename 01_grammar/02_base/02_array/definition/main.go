package main

import "fmt"

var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

func main() {
	a := [3]int{1, 2}           // 未初始化元素为0
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度
	c := [5]int{2: 100, 4: 200} // 使用引号初始化元素
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}
	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)

	m := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	n := [...][2]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(m, n)
}
