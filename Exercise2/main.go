package main

import (
	"Exercise2/calculator"
	"Exercise2/even"
	"Exercise2/sum"
	"fmt"
)

func main() {
	fmt.Println(even.IsEven(11))
	fmt.Println(calculator.Calculate(10, 2, "/"))
	fmt.Println(sum.Calculate(5))
}
