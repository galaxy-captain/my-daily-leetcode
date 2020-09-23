package o47

import (
	"my-daily-leetcode/helper"
	"testing"
)

func TestO47(t *testing.T) {
	helper.AssertInt(t, 12, maxValue([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}
