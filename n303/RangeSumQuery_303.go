package n303

type NumArray struct {
	record []int
}

func Constructor(nums []int) NumArray {

	if nums == nil || len(nums) == 0 {
		return NumArray{}
	}

	na := NumArray{
		record: make([]int, len(nums)),
	}

	na.record[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		na.record[i] = na.record[i-1] + nums[i]
	}

	return na
}

func (this *NumArray) SumRange(i int, j int) int {

	if this.record == nil {
		return 0
	}

	if i == 0 {
		return this.record[j]
	}

	return this.record[j] - this.record[i-1]
}
