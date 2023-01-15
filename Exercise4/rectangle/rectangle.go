package rectangle

func CalculateArea(length int, width int) int {
	if length < 0 || width < 0 {
		return -1
	}
	return length * width
}
func CalculatePerimeter(length int, width int) int {
	if length < 0 || width < 0 {
		return -1
	}
	return 2 * (length + width)
}
