package n70

import (
	"my-daily-leetcode/helper"
	"testing"
)

func TestN70(t *testing.T) {
	helper.AssertInt(t, 1, climbStairs(0))
	helper.AssertInt(t, 1, climbStairs(1))
	helper.AssertInt(t, 3, climbStairs(3))
}
