package employee

import "testing"

func TestFindAge(t *testing.T) {
	testcases := []struct {
		input          Employee
		expectedOutput int
	}{
		{Employee{"Ganesh", "Manchi", "01-09-2001"}, 21},
		{Employee{"Prasoon", "Sukhla", "02-03-2025"}, -1},
		{Employee{"Swaroop", "Swaroop", "02-03-2023"}, -1},
		{Employee{"Reshmi", "Gupta", "29-01-2023"}, -1},
		{Employee{"Mukundh", "Mukundh", "15-01-2023"}, 0},
		{Employee{"Naga", "Surya", "h5-01-2023"}, -1},
	}
	for i, tc := range testcases {
		actualOutput := FindAge(tc.input)
		if actualOutput != tc.expectedOutput {
			t.Errorf("Test case %d failed, expected %d but got %d", i+1, tc.expectedOutput, actualOutput)
		}
	}
}

func TestGreetEmployee(t *testing.T) {
	testcases := []struct {
		input          Employee
		expectedOutput string
	}{
		{Employee{"Ganesh", "Manchi", "01-09-2001"}, "Hello Ganesh Manchi"},
		{Employee{"Prasoon", "", "01-09-2001"}, "Hello Prasoon"},
		{Employee{"", "Reshmi", "01-09-2001"}, "Hello Reshmi"},
		{Employee{"", "", "02-03-2002"}, "Hello"},
	}
	for i, tc := range testcases {
		actualOutput := GreetEmployee(tc.input)
		if actualOutput != tc.expectedOutput {
			t.Errorf("Test case %d failed, expected %q but got %q", i+1, tc.expectedOutput, actualOutput)
		}
	}
}
