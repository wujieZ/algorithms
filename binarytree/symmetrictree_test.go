package binarytree

import "testing"

func TestIsSymmetricTree(t *testing.T) {
	tests := []struct {
		root   *Node
		result bool
	}{
		{
			&Node{
				0,
				nil,
				nil,
			},
			true,
		},
		{
			nil,
			true,
		},
		{
			&Node{
				0,
				&Node{
					1,
					nil,
					nil,
				},
				&Node{
					1,
					nil,
					nil,
				},
			},
			true,
		},
		{
			&Node{
				0,
				&Node{
					1,
					nil,
					nil,
				},
				&Node{
					2,
					nil,
					nil,
				},
			},
			false,
		},
		{
			&Node{
				1,
				&Node{
					2,
					&Node{
						3,
						nil,
						nil,
					},
					&Node{
						4,
						nil,
						nil,
					},
				},
				&Node{
					2,
					&Node{
						4,
						nil,
						nil,
					},
					&Node{
						3,
						nil,
						nil,
					},
				},
			},
			true,
		},
		{
			&Node{
				1,
				&Node{
					2,
					&Node{
						3,
						nil,
						nil,
					},
					&Node{
						4,
						nil,
						nil,
					},
				},
				&Node{
					2,
					&Node{
						3,
						nil,
						nil,
					},
					&Node{
						4,
						nil,
						nil,
					},
				},
			},
			false,
		},
	}

	for i, test := range tests {
		if ok := isSymmetricTree(test.root); ok != test.result {
			t.Errorf("第 %d 个案例, expect %v, but found %v", i+1, test.result, ok)
		}
	}
}
