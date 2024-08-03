package main

import "fmt"

func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	v, ok := scoreMap["小明"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	// 遍历map时的元素顺序与添加键值对的顺序无关。
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	delete(scoreMap, "王五")

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}
