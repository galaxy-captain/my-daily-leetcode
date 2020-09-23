package n257

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var record [][]string

func Do() {
	result := binaryTreePaths(&TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
					},
				},
			},
		},
	})
	fmt.Println(strings.Join(result, ","))
}

func binaryTreePaths(root *TreeNode) []string {
	//record = make([][]string, 0)
	if root == nil {
		return nil
	}
	traverse(root, make([]string, 0))
	result := make([]string, 0)
	for _, path := range record {
		result = append(result, strings.Join(path, "->"))
	}
	return result
}

func traverse(node *TreeNode, path []string) {
	if node == nil {
		return
	}
	newPath := make([]string, len(path))
	copy(newPath, path)
	newPath = append(newPath, strconv.Itoa(node.Val))
	if isLeaf(node) {
		record = append(record, newPath)
		return
	}
	traverse(node.Left, newPath)
	traverse(node.Right, newPath)
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
