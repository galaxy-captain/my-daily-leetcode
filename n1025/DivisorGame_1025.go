package n1025

func divisorGame(N int) bool {

	dp := make([]bool, N+1)
	dp[1] = false

	for i := 2; i <= N; i++ {
		for j := 1; j < i; j++ {
			if i%j != 0 {
				continue
			}
			if !dp[i] {
				dp[i] = !dp[i-j]
			}
		}
	}

	return dp[N]
}
