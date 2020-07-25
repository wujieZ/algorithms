package binarytree

// leetcode, 判断两棵二叉树是否相同, 难度: 简单

type Node struct {
	Value       int
	Left, Right *Node
}

func isSameTree(root1 *Node, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Value != root2.Value {
		return false
	}
	return isSameTree(root1.Left, root2.Left) && isSameTree(root1.Right, root2.Right)
}
