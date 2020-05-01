package traverse

import (
	"testing"
)

func TestTreeNode_PreOrderTraverse(t *testing.T) {
	tests := []struct{
		node *TreeNode
		result []int
	} {
		{
			&TreeNode{
				5,
				&TreeNode{3, &TreeNode{9, nil, nil}, nil},
				&TreeNode{6, nil, nil},
			},
			[]int {5, 3, 9, 6},
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
				break
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
				break
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
				break
			}
		}
		PostOrderResult = PostOrderResult[0:0]
	}
}

func TestTreeNode_BreadthFirstTraverse(t *testing.T) {
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
				&TreeNode{10, nil, &TreeNode{10, &TreeNode{9, nil, nil}, &TreeNode{20, nil, nil}}},
			},
			[]int {4, 5, 10, 1, 10, 9, 20},
		},
		{
			&TreeNode{5, nil, &TreeNode{4, nil, nil}},
			[]int{5, 4},
		},
	}
	for _, test := range tests {
		result := test.node.BreadthFirstTraverse()
		for i, item := range result {
			if item != test.result[i] {
				t.Errorf("expect%v, but got %v", test.result, result)
				break
			}
		}
	}
}

func TestTreeNode_DeepFirstTraverse(t *testing.T) {
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
				&TreeNode{10, nil, &TreeNode{10, &TreeNode{9, nil, nil}, &TreeNode{20, nil, nil}}},
			},
			[]int {4, 5, 1, 10, 10, 9, 20},
		},
		{
			&TreeNode{5, nil, &TreeNode{4, nil, nil}},
			[]int{5, 4},
		},
	}
	for _, test := range tests {
		result := test.node.DeepFirstTraverse()
		for i, item := range result {
			if item != test.result[i] {
				t.Errorf("expect%v, but got %v", test.result, result)
				break
			}
		}
	}
}

func TestTreeNode_LevelOrder(t *testing.T) {
	tests := []struct{
		node *TreeNode
		result [][]int
	} {
		{
			nil,
			[][]int{},
		},
		{
			&TreeNode{
				4,
				&TreeNode{5, nil, &TreeNode{1, nil, nil}},
				&TreeNode{10, nil, &TreeNode{10, &TreeNode{9, nil, nil}, &TreeNode{20, nil, nil}}},
			},
			[][]int {
				{4},
				{5, 10},
				{1, 10},
				{9, 20},
			},
		},
		{
			&TreeNode{5, &TreeNode{10, nil, nil}, &TreeNode{4, nil, nil}},
			[][]int {
				{5},
				{10, 4},
			},
		},
	}
	for _, test := range tests {
		result := test.node.LevelOrder()
		if len(result) != len(test.result) {
			t.Errorf("expect%v, but got %v", test.result, result)
		}
		for i := 0; i < len(result); i++ {
			actualItems := result[i]
			expectItems := test.result[i]
			if len(actualItems) != len(expectItems) {
				t.Errorf("expect%v, but got %v", test.result, result)
				break
			}
			success := true
			for j := 0; j < len(actualItems); j++ {
				if actualItems[j] != expectItems[j] {
					t.Errorf("expect%v, but got %v", test.result, result)
					success = false
					break
				}
			}
			if !success {
				break
			}
		}
	}
}