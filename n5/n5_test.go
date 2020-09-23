package n5

import (
	"my-daily-leetcode/helper"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	helper.AssertString(t, "", longestPalindrome(""))
	helper.AssertString(t, "a", longestPalindrome("a"))
	helper.AssertString(t, "aabaa", longestPalindrome("aabaab"))
}
