package Double

import "testing"

func Test_Calculate(t *testing.T) {

	testcases := []struct {
		input  int
		output int
	}{
		{1, 2},
		{0, 0},
		{-20, -40},
	}
	for i, tc := range testcases {
		actualOutput := Calculate(tc.input)
		if actualOutput != tc.output {
			t.Errorf("Test case %d failed, expected %d but got %d", i+1, tc.output, actualOutput)
		}
	}
}
