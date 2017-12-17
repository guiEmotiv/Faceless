package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("dasdsa")
	fmt.Println("soy un texto", facevoid(12))
}
func facevoid(a int) (s string) {
	s = strconv.Itoa(a)
	return
}
