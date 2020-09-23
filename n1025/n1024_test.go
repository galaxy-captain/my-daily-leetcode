package n1025

import (
	"my-daily-leetcode/helper"
	"testing"
)

func TestN1025(t *testing.T) {
	helper.AssertBool(t, false, divisorGame(1))
	helper.AssertBool(t, true, divisorGame(2))
	helper.AssertBool(t, false, divisorGame(3))
	helper.AssertBool(t, true, divisorGame(6))
}
