package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumArr(a [10]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func main() {
	rand.NewSource(time.Now().Unix())
	var b [10]int
	for i := 0; i < len(b); i++ {
		// 产生一个0到1000随机数
		b[i] = rand.Intn(1000)
	}
	sum := sumArr(b)
	fmt.Printf("sum=%d\n", sum)
}
