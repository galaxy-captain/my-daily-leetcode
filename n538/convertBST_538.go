package n538

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	count := 0
	traverse(root, &count)
	return root
}

func traverse(root *TreeNode, count *int) {
	if root == nil {
		return
	}
	traverse(root.Right, count)
	*count += root.Val
	root.Val = *count
	traverse(root.Left, count)
}
