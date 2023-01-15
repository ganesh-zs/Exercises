package Reverse

import (
	"reflect"
	"testing"
)

func Test_ReverseName(t *testing.T) {

	str := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{}, []int{}},
	}

	for i := 0; i < len(str); i++ {
		if !reflect.DeepEqual(ReverseName(str[i].input), str[i].output) {
			t.Errorf("Error in test case %v", i+1)
		}
	}
}

func BenchmarkReverseName(b *testing.B) {

	str := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{6, 7, 8}, []int{8, 7, 6}},
		{[]int{0, 9, 8, 4, 5}, []int{5, 4, 8, 9, 0}},
	}

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(str); j++ {
			ReverseName(str[j].input)
		}
	}
}
