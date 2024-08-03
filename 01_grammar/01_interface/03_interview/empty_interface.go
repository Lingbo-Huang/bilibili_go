package main

import "fmt"

/*
空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
空接口类型的变量可以存储任意类型的变量。
用空接口作为函数的参数可以接收任意类型的参数
用空接口可以实现能保存任意值的字典
*/

func justfyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string, value is %v\n", v)
	case int:
		fmt.Printf("x is a int, value is %v\n", v)
	case bool:
		fmt.Printf("x is a bool, value is %v\n", v)
	default:
		fmt.Printf("unsupport type!")
	}
}

func main() {
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "菜小波"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

	var x interface{}
	x = "pprof.cn"
	v, ok := x.(string) // 类型断言
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("断言失败")
	}

	justfyType(x)

}
