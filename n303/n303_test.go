package n303

import (
	"my-daily-leetcode/helper"
	"testing"
)

func TestN303(t *testing.T) {

	var na NumArray

	na = Constructor([]int{-2, 0, 3, -5, 2, -1})
	helper.AssertInt(t, 1, na.SumRange(0, 2))
	helper.AssertInt(t, -1, na.SumRange(2, 5))

	na = Constructor([]int{})
	na.SumRange(0, 0)
	helper.AssertInt(t, 0, na.SumRange(2, 5))

}
