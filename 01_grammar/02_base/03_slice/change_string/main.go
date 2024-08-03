package main

import "fmt"

func main() {
	str := "hello world"
	s := []byte(str) // 有中文用rune
	s[6] = 'G'
	s = s[:8]
	s = append(s, '!')
	str = string(s)
	fmt.Println(str)
}
