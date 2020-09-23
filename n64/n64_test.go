package n64

import (
	"my-daily-leetcode/helper"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	helper.AssertInt(t, 7, minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
