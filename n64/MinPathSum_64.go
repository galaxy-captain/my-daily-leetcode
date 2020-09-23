package n64

func minPathSum(grid [][]int) int {

	x := len(grid[0])
	y := len(grid)

	dp := make([][]int, y)
	for i := range dp {
		dp[i] = make([]int, x)
	}

	for i := y - 1; i >= 0; i-- {
		for j := x - 1; j >= 0; j-- {
			if j == x-1 && i == y-1 {
				dp[i][j] = grid[i][j]
				continue
			}
			if j == x-1 {
				dp[i][j] = grid[i][j] + dp[i+1][j]
				continue
			}
			if i == y-1 {
				dp[i][j] = grid[i][j] + dp[i][j+1]
				continue
			}
			if grid[i][j]+dp[i+1][j] < grid[i][j]+dp[i][j+1] {
				dp[i][j] = grid[i][j] + dp[i+1][j]
			} else {
				dp[i][j] = grid[i][j] + dp[i][j+1]
			}
		}
	}

	return dp[0][0]
}
