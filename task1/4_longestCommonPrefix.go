package main

import "strings"

func longestCommonPrefix(strs []string) string {
	minLen := getMinLen(strs)
	var sb strings.Builder
outer:
	for i := 0; i < minLen; i++ {
		compare := []rune(strs[0])[i]
		for _, str := range strs {
			if []rune(str)[i] != compare {
				break outer
			}
		}
		sb.WriteRune(compare)
	}

	return sb.String()
}

func getMinLen(strs []string) int {
	result := getStrLen(strs[0])
	for _, val := range strs {
		if getStrLen(val) < result {
			result = getStrLen(val)
		}
	}

	return result
}

func getStrLen(str string) int {
	return len([]rune(str))
}
