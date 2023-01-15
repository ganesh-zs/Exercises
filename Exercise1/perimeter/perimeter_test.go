package perimeter

import "testing"

func TestCirclePerimeter(t *testing.T) {
	testcases := []struct {
		input          float64
		expectedOutput float64
	}{
		{1, 6.28},
		{0, 0},
		{-10, -1},
	}
	for i, tc := range testcases {
		actualOutput := CirclePerimeter(tc.input)
		if actualOutput != tc.expectedOutput {
			t.Errorf("Test case %d failed, expected %v but got %v", i+1, tc.expectedOutput, actualOutput)
		}
	}
}

func TestSquarePerimeter(t *testing.T) {
	testcases := []struct {
		input          float64
		expectedOutput float64
	}{
		{1, 4},
		{0, 0},
		{-10, -1},
	}
	for i, tc := range testcases {
		actualOutput := SquarePerimeter(tc.input)
		if actualOutput != tc.expectedOutput {
			t.Errorf("Test case %d failed, expected %v but got %v", i+1, tc.expectedOutput, actualOutput)
		}
	}
}
