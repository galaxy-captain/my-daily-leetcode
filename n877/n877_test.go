package n877

import (
	"my-daily-leetcode/helper"
	"testing"
)

func TestN877(t *testing.T) {
	helper.AssertBool(t, true, stoneGame2([]int{5, 3, 4, 5}))
	helper.AssertBool(t, true, stoneGame2([]int{7, 7, 12, 16, 41, 48, 41, 48, 11, 9, 34, 2, 44, 30, 27, 12, 11, 39, 31, 8, 23, 11, 47, 25, 15, 23, 4, 17, 11, 50, 16, 50, 38, 34, 48, 27, 16, 24, 22, 48, 50, 10, 26, 27, 9, 43, 13, 42, 46, 24}))
}
