package n617

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 != nil {
		return t2
	}
	merge(false, nil, t1, t2)
	return t1
}

func merge(isLeft bool, parent *TreeNode, t1 *TreeNode, t2 *TreeNode) {

	if t1 == nil && t2 == nil {
		return
	} else if t1 != nil && t2 == nil {
		merge(true, t1, t1.Left, nil)
		merge(false, t1, t1.Right, nil)
	} else if t1 == nil && t2 != nil {
		if parent != nil {
			t1 = &TreeNode{Val: t2.Val}
			if isLeft {
				parent.Left = t1
			} else {
				parent.Right = t1
			}
		}
		merge(true, t1, nil, t2.Left)
		merge(false, t1, nil, t2.Right)
	} else if t1 != nil && t2 != nil {
		t1.Val += t2.Val
		merge(true, t1, t1.Left, t2.Left)
		merge(false, t1, t1.Right, t2.Right)
	}

}
