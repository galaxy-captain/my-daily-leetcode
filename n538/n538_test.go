package n538

import (
	"my-daily-leetcode/helper"
	"testing"
)

func TestN538(t *testing.T) {
	tree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: -4,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	expect := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	helper.AssertObject(t, expect, convertBST(tree))
}
