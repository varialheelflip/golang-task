package main

func merge(intervals [][]int) [][]int {
	// 展不开一点, 逐个合并了
	if len(intervals) == 1 {
		return intervals
	}
	var result = append([][]int{}, intervals[0])
	for i := 1; i < len(intervals); i++ {
		// 过滤出能合并和不能合并的
		var canNotMerge [][]int
		var canMerge [][]int
		var canMergeFlag = false
		for _, val := range result {
			if judgeCanMerge(val, intervals[i]) {
				canMerge = append(canMerge, val)
				canMergeFlag = true
			} else {
				canNotMerge = append(canNotMerge, val)
			}
		}
		if canMergeFlag {
			canMerge = append(canMerge, intervals[i])
		} else {
			canNotMerge = append(canNotMerge, intervals[i])
		}
		// 把能合并的合并了
		if len(canMerge) > 0 {
			result = append(canNotMerge, mergeBatch(canMerge))
		} else {
			result = canNotMerge
		}
	}
	return result
}

func mergeBatch(canMerge [][]int) []int {
	// 按最小左区间, 最大右区间合并
	var left = canMerge[0][0]
	var right = canMerge[0][1]
	for _, val := range canMerge {
		if val[0] < left {
			left = val[0]
		}
		if val[1] > right {
			right = val[1]
		}
	}
	return []int{left, right}
}

func judgeCanMerge(val1 []int, val2 []int) bool {
	if val1[0] <= val2[1] && val1[1] >= val2[0] {
		return true
	}
	if val2[0] <= val1[1] && val2[1] >= val1[0] {
		return true
	}
	return false
}
