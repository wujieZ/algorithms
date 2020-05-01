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