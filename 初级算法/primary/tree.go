package primary

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeFromSlice(data []int) {
	var root *TreeNode

	for _, v := range data {
		InserTree(root, v)
	}
}
func InserTree(root *TreeNode, v int) {
	node := new(TreeNode)
	node.Val = v

	if root == nil {
		root = node
		return
	}

	if v > root.Val {
		if root.Left == nil {
			root.Left = node
		} else {
			InserTree(root.Left, v)
		}
	} else {
		if root.Right == nil {
			root.Right = node
		} else {
			InserTree(root.Right, v)
		}
	}

}

// 给定一个二叉树，找出其最大深度。
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
