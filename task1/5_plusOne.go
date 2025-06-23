package main

func plusOne(digits []int) []int {
	var inFlag int
	for i := len(digits) - 1; i >= 0; i-- {
		switch i {
		case len(digits) - 1:
			{
				if digits[i] == 9 {
					digits[i] = 0
					inFlag = 1
				} else {
					digits[i] += 1
				}
			}
		default:
			{
				if digits[i]+inFlag == 10 {
					digits[i] = 0
					inFlag = 1
				} else {
					digits[i] += inFlag
					inFlag = 0
				}
			}
		}
	}
	if inFlag == 1 {
		return append([]int{1}, digits...)
	}

	return digits
}
