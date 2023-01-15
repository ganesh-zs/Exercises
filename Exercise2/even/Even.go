package even

// IsEven takes an integer input and returns whether it is even or not
func IsEven(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
