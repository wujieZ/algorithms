package queue

import "testing"


func TestQueuePush(t *testing.T) {
	tests := []struct{
		original *Queue
		items []int
		arrLen int
	} {
		{
			&Queue{},
			[]int{1, 2},
			2,
		},
		{
			&Queue{1, 2, 3},
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

func TestQueueShift(t *testing.T) {
	tests := []struct{
		original *Queue
		success bool
		item int
		arrLen int
	} {
		{
			&Queue{},
			false,
			0,
			0,
		},
		{
			&Queue{1, 2, 3},
			true,
			1,
			2,
		},
	}
	for _, test := range tests {
		result, ok := test.original.Shift()
		if ok != test.success ||
			result != test.item ||
			len(*test.original) != test.arrLen {
			t.Errorf("pop fail, expected return %d, but got %d", test.item, result)
		}
	}
}

func TestQueueGetLen(t *testing.T)  {
	tests := []struct{
		arr *Queue
		len int
	} {
		{
			&Queue{},
			0,
		},
		{
			&Queue{1, 2, 3, 4, 5, 6},
			6,
		},
	}
	for _, test := range tests {
		result := test.arr.getLen()
		if result != test.len {
			t.Errorf("expected return %d, but got %d", test.len, result)
		}
	}

}