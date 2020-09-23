package n53

func maxSubArray(nums []int) int {
	last := nums[0]
	max := last
	for i := 1; i < len(nums); i++ {
		if last < 0 {
			last = nums[i]
		} else {
			last = last + nums[i]
		}
		if last > max {
			max = last
		}
	}
	return max
}
