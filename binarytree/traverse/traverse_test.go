package traverse

import "testing"

func TestTreeNode_PreOrderTraverse(t *testing.T) {
	tests := []struct{
		node *TreeNode
		result []int
	} {
		{
			&TreeNode{
				5,
				&TreeNode{3, nil, nil},
				&TreeNode{6, nil, nil},
			},
			[]int {5, 3, 6},
		},
		{
			&TreeNode{
				5, nil, nil,
			},
			[]int{5},
		},
		{
			nil,
			[]int{},
		},
	}

	for _, test := range tests {
		test.node.PreOrderTraverse()
		for i, item := range PreOrderResult {
			if item != test.result[i] {
				t.Errorf("expect%v, but got %v", test.result, PreOrderResult)
			}
		}
		PreOrderResult = PreOrderResult[0:0]
	}
}

func TestTreeNode_InOrderTraverse(t *testing.T) {
	tests := []struct{
		node *TreeNode
		result []int
	} {
		{
			nil,
			[]int{},
		},
		{
			&TreeNode{
				4,
				&TreeNode{5, nil, &TreeNode{1, nil, nil}},
				&TreeNode{10, nil, nil},
			},
			[]int {5, 1, 4, 10},
		},
	}

	for _, test := range tests {
		test.node.InOrderTraverse()
		for i, item := range InOrderResult {
			if item != test.result[i] {
				t.Errorf("expect%v, but got %v", test.result, InOrderResult)
			}
		}
		InOrderResult = InOrderResult[0:0]
	}
}

func TestTreeNode_PostOrderTraverse(t *testing.T) {
	tests := []struct{
		node *TreeNode
		result []int
	} {
		{
			nil,
			[]int{},
		},
		{
			&TreeNode{
				4,
				&TreeNode{5, nil, &TreeNode{1, nil, nil}},
				&TreeNode{10, nil, nil},
			},
			[]int {1, 5, 10, 4},
		},
		{
			&TreeNode{5, nil, &TreeNode{4, nil, nil}},
			[]int{4, 5},
		},
	}

	for _, test := range tests {
		test.node.PostOrderTraverse()
		for i, item := range InOrderResult {
			if item != test.result[i] {
				t.Errorf("expect%v, but got %v", test.result, PostOrderResult)
			}
		}
		PostOrderResult = PostOrderResult[0:0]
	}
}