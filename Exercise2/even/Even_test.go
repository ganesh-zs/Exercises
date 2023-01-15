package even

import "testing"

func Test_Calculate(t *testing.T) {

	str := []struct {
		input  int
		output bool
	}{
		{1, false},
		{2, true},
		{5, false},
		{6, true},
	}
	for i := 0; i < len(str); i++ {
		actualOutput := IsEven(str[i].input)
		if actualOutput != str[i].output {
			t.Errorf("Test case %d failed, expected %v but got %v", i+1, str[i].output, actualOutput)
		}
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEven(i)
	}
}
