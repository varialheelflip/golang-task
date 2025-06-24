package main

func twoSum(nums []int, target int) []int {
	var judgeMap = make(map[int]int)
	for i, val := range nums {
		index, ok := judgeMap[target-val]
		if ok {
			return []int{index, i}
		} else {
			judgeMap[val] = i
		}
	}

	return nil
}
