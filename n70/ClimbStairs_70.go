package n70

func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	p := 1
	c := 1
	tmp := 0
	for i := 2; i <= n; i++ {
		tmp = c
		c = p + c
		p = tmp
	}
	return c
}
