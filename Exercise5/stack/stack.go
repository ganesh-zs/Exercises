package stack

type Input struct {
	Stack     []int
	Operation string
	Value     int
}

// Push return slice by appending value
func Push(stack []int, value int) ([]int, int) {
	stack = append(stack, value)
	return stack, -1
}

// Pop returns last inserted value by removing it
func Pop(stack []int) ([]int, int) {
	popValue := -1
	n := len(stack)
	if len(stack) > 0 {
		popValue = stack[n-1]
		stack = stack[:n-1]
	}
	return stack, popValue
}

// Top returns last inserted value
func Top(stack []int) ([]int, int) {
	topValue := -1
	n := len(stack)
	if len(stack) > 0 {
		topValue = stack[n-1]
	}
	return stack, topValue
}

// Stack returns slice by performing operation
func Stack(input Input) ([]int, int) {
	switch input.Operation {
	case "push":
		return Push(input.Stack, input.Value)
	case "pop":
		return Pop(input.Stack)
	case "top":
		return Top(input.Stack)
	default:
		return input.Stack, -1
	}
}
