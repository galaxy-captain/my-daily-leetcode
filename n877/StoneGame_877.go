package n877

const IntMin = int(^int32(0) << 31)

func stoneGame2(piles []int) bool {

	length := len(piles)
	record := make([][]int, length)
	for i := range record {
		record[i] = make([]int, length)
		for j := range record[i] {
			record[i][j] = IntMin
		}
		record[i][i] = piles[i]
	}

	for right := 1; right < length; right++ {
		for left := right - 1; left >= 0; left-- {
			record[left][right] = max(piles[left]-record[left+1][right], piles[right]-record[left][right-1])
		}
	}

	return record[0][length-1] > 0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func stoneGame1(piles []int) bool {

	length := len(piles)
	record := make([][]int, length)
	for i := range record {
		record[i] = make([]int, length)
		for j := range record[i] {
			record[i][j] = IntMin
		}
		record[i][i] = piles[i]
	}

	return doGame(piles, record, 0, length-1) > 0
}

func doGame(piles []int, record [][]int, left, right int) int {

	if left == right {
		return record[left][right]
	}

	if record[left][right] != IntMin {
		return record[left][right]
	}

	choseLeft := piles[left] - doGame(piles, record, left+1, right)
	choseRight := piles[right] - doGame(piles, record, left, right-1)
	if choseLeft > choseRight {
		record[left][right] = choseLeft
	} else {
		record[left][right] = choseRight
	}
	return record[left][right]
}
