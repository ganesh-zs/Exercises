package temperature

import "testing"

func TestCelToFer(t *testing.T) {
	tests := []struct {
		input          CelToFer
		expectedOutput float64
	}{
		{CelToFer{100}, 212},
		{CelToFer{0}, 32},
		{CelToFer{-100}, -148},
		{CelToFer{3}, 37.4},
	}
	for i, tc := range tests {
		actualOutput := tc.input.Convertion()
		if actualOutput != tc.expectedOutput {
			t.Errorf("Test case %d failed, expected %v but got %v", i+1, tc.expectedOutput, actualOutput)
		}
	}
}

func TestFerToCel(t *testing.T) {
	tests := []struct {
		input          FerToCel
		expectedOutput float64
	}{
		{FerToCel{100}, 37.77777777777778},
		{FerToCel{-100}, -73.33333333333334},
		{FerToCel{0}, -17.77777777777778},
	}
	for i, tc := range tests {
		actualOutput := tc.input.Convertion()
		if actualOutput != tc.expectedOutput {
			t.Errorf("Test case %d failed, expected %v but got %v", i+1, tc.expectedOutput, actualOutput)
		}
	}
}
