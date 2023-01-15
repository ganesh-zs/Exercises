package main

import (
	"Exercise3/Reverse"
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	Reverse.ReverseName(x)
	fmt.Println(x)
}
