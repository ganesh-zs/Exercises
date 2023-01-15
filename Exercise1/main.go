package main

import (
	"Exercise1/Double"
	"Exercise1/concat"
	"Exercise1/perimeter"
	"fmt"
)

func main() {
	fmt.Println("Double value : ", Double.Calculate(10))
	fmt.Println("String Concatination : ", concat.Greet("Ganesh"))
	fmt.Println("Perimeter of Circle : ", perimeter.CirclePerimeter(10), " and Perimeter of Square : ", perimeter.SquarePerimeter(10))
	fmt.Println("Hai")
}
