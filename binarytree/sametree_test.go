package binarytree

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		root1  *Node
		root2  *Node
		result bool
	}{
		{
			&Node{
				0,
				nil,
				nil,
			},
			&Node{
				0,
				nil,
				nil,
			},
			true,
		},
		{
			&Node{
				0,
				&Node{
					0,
					nil,
					nil,
				},
				nil,
			},
			&Node{
				0,
				nil,
				nil,
			},
			false,
		},
		{
			nil,
			nil,
			true,
		},
		{
			&Node{
				0,
				&Node{
					0,
					nil,
					nil,
				},
				nil,
			},
			&Node{
				0,
				&Node{
					0,
					nil,
					nil,
				},
				nil,
			},
			true,
		},
	}
	for i, test := range tests {
		if ok := isSameTree(test.root1, test.root2); ok != test.result {
			t.Errorf("第 %d 个案例, expect %v, but found %v", i+1, test.result, ok)
		}
	}

}
