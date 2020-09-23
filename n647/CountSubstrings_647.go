package n647

func countSubstrings(s string) int {

	length := len(s)
	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, length)
	}

	count := 0
	for x := 0; x < length; x++ {
		dp[x][x] = true
		count++
		for y := x - 1; y >= 0; y-- {
			if x-y == 1 && s[x] == s[y] {
				dp[y][x] = true
				count++
			} else if dp[y+1][x-1] && s[x] == s[y] {
				dp[y][x] = true
				count++
			}
		}
	}

	return count
}
