package rectangle

import "testing"

func Test_CalculateArea(t *testing.T) {
	str := []struct {
		length         int
		breadth        int
		expectedOutput int
	}{
		{1, 2, 2},
		{0, 9, 0},
		{-2, 8, -1},
	}
	for i := 0; i < len(str); i++ {
		actualOutput := CalculateArea(str[i].length, str[i].breadth)
		if actualOutput != str[i].expectedOutput {
			t.Errorf("Task Failed.Expected:%v,got:%v", str[i].expectedOutput, actualOutput)
		}
	}
}
func Test_CalculatePerimeter(t *testing.T) {
	str := []struct {
		length         int
		breadth        int
		expectedOutput int
	}{
		{1, 2, 6},
		{-1, 2, -1},
	}
	for i := 0; i < len(str); i++ {
		actualOutput := CalculatePerimeter(str[i].length, str[i].breadth)
		if actualOutput != str[i].expectedOutput {
			t.Errorf("Task Failed.Expected:%v,got:%v", str[i].expectedOutput, actualOutput)
		}
	}
}
