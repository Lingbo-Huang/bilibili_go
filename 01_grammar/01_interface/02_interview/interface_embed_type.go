package embedding_type

import "fmt"

type WashingMachine interface {
	wash()
	dry()
}

type dryer struct{}

func (d dryer) dry() {
	fmt.Println("甩一甩")
}

type haier struct {
	dryer
}

// haier 的方法 wash 和嵌入的类型dryer 的方法 dry 一起使得实现接口

func (h haier) wash() {
	fmt.Println("洗一洗")
}
