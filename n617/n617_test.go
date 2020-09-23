package n617

import (
	"my-daily-leetcode/helper"
	"testing"
)

func TestN617(t *testing.T) {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	t2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	expect := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	helper.AssertObject(t, expect, mergeTrees(t1, t2))
}
