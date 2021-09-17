package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0)
	b := []int{1, 2}
	c := append(b, a...)
	fmt.Println(c)
}
