package concat

import "testing"

func Test_Greet(t *testing.T) {
	testcases := []struct {
		input          string
		expectedOutput string
	}{
		{"Ganesh", "Hello,Ganesh!"},
		{"", "Hello,!"},
		{"Prasoon Sukhla", "Hello,Prasoon Sukhla!"},
	}
	for i, tc := range testcases {
		actualOutput := Greet(tc.input)
		if actualOutput != tc.expectedOutput {
			t.Errorf("Test case %d failed, expected %q but got %q", i+1, tc.expectedOutput, actualOutput)
		}
	}
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Ganesh")
	}
}
