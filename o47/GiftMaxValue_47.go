package o47

func maxValue(grid [][]int) int {

	xl := len(grid[0]) + 1
	yl := len(grid) + 1
	record := make([][]int, yl)
	for y := 0; y < yl; y++ {
		record[y] = make([]int, xl)
		for x := 0; x < xl; x++ {
			record[y][x] = 0
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for y := 1; y < yl; y++ {
		for x := 1; x < xl; x++ {
			record[y][x] = max(record[y-1][x], record[y][x-1]) + grid[y-1][x-1]
		}
	}

	return record[yl-1][xl-1]

	//for y := 0; y < yl; y++ {
	//	record[y] = make([]int, xl)
	//	for x := 0; x < xl; x++ {
	//		if y == 0 && x == 0 {
	//			record[y][x] = grid[y][x]
	//			continue
	//		}
	//		if y == 0 {
	//			record[y][x] = record[y][x-1] + grid[y][x]
	//			continue
	//		}
	//		if x == 0 {
	//			record[y][x] = record[y-1][x] + grid[y][x]
	//			continue
	//		}
	//		if record[y][x-1]+grid[y][x] > record[y-1][x]+grid[y][x] {
	//			record[y][x] = record[y][x-1] + grid[y][x]
	//		} else {
	//			record[y][x] = record[y-1][x] + grid[y][x]
	//		}
	//	}
	//}
	//
	//return record[yl-1][xl-1]
}
