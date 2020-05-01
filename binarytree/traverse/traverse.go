package traverse

type TreeNode struct {
	Value int
	LeftNode *TreeNode
	RightNode * TreeNode
}

// 前序遍历
var PreOrderResult []int
func (node *TreeNode) PreOrderTraverse(){
	if node == nil {
		return
	}
	PreOrderResult = append(PreOrderResult, node.Value)
	node.LeftNode.PreOrderTraverse()
	node.RightNode.PreOrderTraverse()
}

// 中序遍历
var InOrderResult []int
func (node *TreeNode) InOrderTraverse(){
	if node == nil {
		return
	}
	node.LeftNode.InOrderTraverse()
	InOrderResult = append(InOrderResult, node.Value)
	node.RightNode.InOrderTraverse()
}

// 后序遍历
var PostOrderResult []int
func (node *TreeNode) PostOrderTraverse(){
	if node == nil {
		return
	}
	node.LeftNode.PostOrderTraverse()
	node.RightNode.PostOrderTraverse()
	PostOrderResult = append(PostOrderResult, node.Value)
}

// 广度优先遍历
func (node *TreeNode) BreadthFirstTraverse() []int{
	var result []int
	if node == nil {
		return result
	}
	var q []*TreeNode
	q = append(q, node)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		result = append(result, node.Value)
		if node.LeftNode != nil {
			q = append(q, node.LeftNode)
		}
		if node.RightNode != nil {
			q = append(q, node.RightNode)
		}
	}
	return result
}

// 深度优先遍历(非递归写法)
func (node *TreeNode) DeepFirstTraverse() []int{
	var result []int
	if node == nil {
		return result
	}
	var stack []*TreeNode
	stack = append(stack, node)
	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[0:len(stack) - 1]
		result = append(result, node.Value)
		if node.RightNode != nil {
			stack = append(stack, node.RightNode)
		}
		if node.LeftNode != nil {
			stack = append(stack, node.LeftNode)
		}
	}
	return result
}

// 层次遍历(跟广度优先差不多，分层次， 二维数组)
func (node *TreeNode) LevelOrder() [][]int{
	result := make([][]int, 0)
	if node == nil {
		return result
	}
	var q []*TreeNode
	q = append(q, node)
	for len(q) > 0 {
		l := len(q)
		var arr []int
		for i := 0; i < l; i++ {
			node := q[i]
			arr = append(arr, node.Value)
			if node.LeftNode != nil {
				q = append(q, node.LeftNode)
			}
			if node.RightNode != nil {
				q = append(q, node.RightNode)
			}
		}
		result = append(result, arr)
		q = q[l:]
	}
	return result
}