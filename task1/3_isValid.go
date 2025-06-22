package main

func isValid(s string) bool {
	var stack []rune
	for _, val := range []rune(s) {
		stackSize := len(stack)
		// 入栈情况
		if stackSize == 0 || !isPair(stack[stackSize-1], val) {
			stack = append(stack, val)
		} else {
			// 出栈情况
			stack = stack[:stackSize-1]
		}
	}

	return len(stack) == 0
}

func isPair(left rune, right rune) bool {
	return (left == '(' && right == ')') || (left == '[' && right == ']') || (left == '{' && right == '}')
}
