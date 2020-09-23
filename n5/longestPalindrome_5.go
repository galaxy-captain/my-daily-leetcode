package n5

func longestPalindrome(s string) string {

	length := len(s)

	if length <= 1 {
		return s
	}

	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, length)
	}

	begin := 0
	max := 1
	for x := 0; x < length; x++ {
		dp[x][x] = true
		for y := x - 1; y >= 0; y-- {

			if x-y == 1 && s[x] == s[y] {
				dp[y][x] = true
			} else if dp[y+1][x-1] && s[x] == s[y] {
				dp[y][x] = true
			}

			if dp[y][x] && x-y+1 > max {
				max = x - y + 1
				begin = y
			}
		}
	}

	return s[begin : begin+max]
}
