package stack

import (
	"reflect"
	"testing"
)

type Output struct {
	expectedStack []int
	expectedValue int
}

func TestStack(t *testing.T) {
	testcases := []struct {
		input  Input
		output Output
	}{
		{Input{[]int{1, 2, 3, 4, 5}, "push", 3}, Output{[]int{1, 2, 3, 4, 5, 3}, -1}},
		{Input{[]int{1, 2, 3, 4, 5}, "pop", -1}, Output{[]int{1, 2, 3, 4}, 5}},
		{Input{[]int{1, 2, 3, 4, 5}, "top", -1}, Output{[]int{1, 2, 3, 4, 5}, 5}},
		{Input{[]int{}, "pop", -1}, Output{[]int{}, -1}},
		{Input{[]int{}, "top", -1}, Output{[]int{}, -1}},
		{Input{[]int{}, "insert", -1}, Output{[]int{}, -1}},
	}

	for i, tc := range testcases {
		actualOutputStack, actualOutputValue := Stack(tc.input)
		if !reflect.DeepEqual(actualOutputStack, tc.output.expectedStack) || !reflect.DeepEqual(actualOutputValue, tc.output.expectedValue) {
			t.Errorf("Error in test case %v", i+1)
		}
	}
}

func TestPush(t *testing.T) {
	testcases := []struct {
		input       []int
		value       int
		outputSlice []int
		outputValue int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 3, 4, 5, 3}, -1},
		{[]int{}, 7, []int{7}, -1},
	}

	for i, tc := range testcases {
		actualOutputSlice, actualOutputValue := Push(tc.input, tc.value)
		if !reflect.DeepEqual(actualOutputSlice, tc.outputSlice) || !reflect.DeepEqual(actualOutputValue, tc.outputValue) {
			t.Errorf("Error in test case %v", i+1)
		}
	}
}

func TestPop(t *testing.T) {
	testcases := []struct {
		input       []int
		outputSlice []int
		outputValue int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4}, 5},
		{[]int{}, []int{}, -1},
	}

	for i, tc := range testcases {
		actualOutputSlice, actualOutputValue := Pop(tc.input)
		if !reflect.DeepEqual(actualOutputSlice, tc.outputSlice) || !reflect.DeepEqual(actualOutputValue, tc.outputValue) {
			t.Errorf("Error in test case %v", i+1)
		}
	}
}

func TestTop(t *testing.T) {
	testcases := []struct {
		input       []int
		outputSlice []int
		outputValue int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 5},
		{[]int{}, []int{}, -1},
	}

	for i, tc := range testcases {
		actualOutputSlice, actualOutputValue := Top(tc.input)
		if !reflect.DeepEqual(actualOutputSlice, tc.outputSlice) || !reflect.DeepEqual(actualOutputValue, tc.outputValue) {
			t.Errorf("Error in test case %v", i+1)
		}
	}
}
