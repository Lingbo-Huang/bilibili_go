package main

import "fmt"

/*
当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
*/

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "hello" {
		talk = "你好呀"
	} else {
		talk = "拜拜了您嘞"
	}
	return
}

func main() {
	// 指针接收者实现接口，那么只能给peo传入指针类型的值
	// 普通对象接收者实现接口，那么不管是传对象还是指针类型都可以
	// 因为Go语言中有对指针类型变量求值的语法糖，指针内部会自动求值解引用
	var peo People = &Student{}
	think := "hello"
	fmt.Println(peo.Speak(think))
}
