package binarytree

// leetcode, 判断一颗二叉树是否对称, 难度: 简单

func isSymmetricTree(root *Node) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(left *Node, right *Node) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Value != right.Value {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}
