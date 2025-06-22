package main

import (
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	numStr := strconv.Itoa(x)
	var sb strings.Builder
	runes := []rune(numStr)
	for i := len(runes) - 1; i >= 0; i-- {
		sb.WriteRune(runes[i]) // 逆序写入字符
	}
	return numStr == sb.String()
}
