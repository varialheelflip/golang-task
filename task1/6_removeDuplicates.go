package main

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	writeIndex := 0
	for readIndex := 1; readIndex < len(nums); readIndex++ {
		if nums[readIndex] == nums[writeIndex] {
			continue
		} else {
			writeIndex++
			nums[writeIndex] = nums[readIndex]
		}
	}
	return writeIndex + 1
}
