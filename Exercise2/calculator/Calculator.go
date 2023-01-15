package calculator

// Calculate takes 2 integers and 1 operator then it returns an integer by performing operation on both numbers
func Calculate(number1 float32, number2 float32, operator string) float32 {
	switch operator {
	// case for addition
	case "+":
		return number1 + number2
	// case for subtraction
	case "-":
		return number1 - number2
	// case for multiplication
	case "*":
		return number1 * number2
	// case for division
	case "/":
		return number1 / number2
	// none of the valid operation
	default:
		return -1.0
	}
}
