package calculator

import "testing"

func Test_Calculate(t *testing.T) {
	s := []struct {
		input1         float32
		input2         float32
		operator       string
		expextedOutput float32
	}{
		{1, 2, "+", 3},
		{2, 3, "*", 6},
		{10, 2, "/", 5},
		{3, 4, "+", 7},
		{5, 3, "*", 15},
		{16, 2, "/", 8},
		{13, 2, "+", 15},
		{20, 3, "*", 60},
	}
	for i := 0; i < len(s); i++ {
		actualOutput := Calculate(s[i].input1, s[i].input2, s[i].operator)
		if actualOutput != s[i].expextedOutput {
			t.Errorf("Test case %d failed, expected %v but got %v", i+1, s[i].expextedOutput, actualOutput)
		}
	}
}

func BenchmarkCalculate(b *testing.B) {
	s := []struct {
		input1   float32
		input2   float32
		operator string
		output   float32
	}{
		{1, 2, "+", 3},
		{2, 3, "*", 6},
		{10, 2, "/", 5},
	}
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s); j++ {
			Calculate(s[j].input1, s[j].input2, s[j].operator)
		}
	}

}
