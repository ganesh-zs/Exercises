package perimeter

const pi = 3.14

// CirclePerimeter returns Perimeter of Circle
func CirclePerimeter(radius float64) float64 {
	if radius < 0 {
		return -1
	}
	return 2 * pi * radius
}

// SquarePerimeter returns Perimeter of Square
func SquarePerimeter(side float64) float64 {
	if side < 0 {
		return -1
	}
	return 4 * side
}
