package sum

// Calculate returns sum of 1 to that number
func Calculate(number int) int {
	sum := 0
	if number < 0 {
		return -1
	}
	for i := 1; i <= number; i++ {
		sum += i
	}
	return sum
}
