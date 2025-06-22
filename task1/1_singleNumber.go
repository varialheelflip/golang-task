package main

func singleNumber(nums []int) int {
	// 遍历放map, 已有就移除, 没有就放进去, 直到最后一个元素
	judgeMap := make(map[int]int)
	for _, num := range nums {
		if judgeMap[num] == 1 {
			delete(judgeMap, num)
		} else {
			judgeMap[num] = 1
		}
	}

	var result int
	for key := range judgeMap {
		result = key
		break
	}

	return result
}
