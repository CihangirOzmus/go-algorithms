package stack

func IsValid(s string) bool {
	prMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []rune
	for _, char := range s {
		if char == ')' && len(stack) > 0 && prMap[char] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else if char == ']' && len(stack) > 0 && prMap[char] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else if char == '}' && len(stack) > 0 && prMap[char] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
