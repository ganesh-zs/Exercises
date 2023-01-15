package sum

import "testing"

func Test_Calculate(t *testing.T) {

	str := []struct {
		input  int
		output int
	}{
		{1, 1},
		{-10, -1},
		{0, 0},
	}
	for i := 0; i < len(str); i++ {
		actualOutput := Calculate(str[i].input)
		if actualOutput != str[i].output {
			t.Errorf("Test case %d failed, expected %v but got %v", i+1, str[i].output, actualOutput)
		}
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(i)
	}
}
