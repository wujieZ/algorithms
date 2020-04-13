package stack

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	tests := []struct{
		original *Stack
		items []int
		arrLen int
	} {
		{
			&Stack{},
			[]int{1, 2},
			2,
		},
		{
			&Stack{1, 2, 3},
			[]int{1},
			4,
		},
	}
	for _, test := range tests {
		result := test.original.Push(test.items...)
		if result != test.arrLen {
			t.Errorf("expectd len is %d, got %d", test.arrLen, result)
		}
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct{
		original *Stack
		success bool
		item int
		arrLen int
	} {
		{
			&Stack{},
			false,
			0,
			0,
		},
		{
			&Stack{1, 2, 3},
			true,
			3,
			2,
		},
	}
	for _, test := range tests {
		result, ok := test.original.Pop()
		if ok != test.success ||
			result != test.item ||
			len(*test.original) != test.arrLen {
			t.Errorf("pop fail, expected return %d, but got %d", test.item, result)
		}
	}
}

func TestStackIsEmpty(t *testing.T)  {
	tests := []struct{
		arr *Stack
		isEmpty bool
	} {
		{
			&Stack{},
			true,
		},
		{
			&Stack{1, 2, 3, 4, 5, 6},
			false,
		},
	}
	for _, test := range tests {
		result := test.arr.isEmpty()
		if result != test.isEmpty {
			t.Errorf("%v, isEmpty expected %v, but got %v", *test.arr, test.isEmpty, result)
		}
	}
}