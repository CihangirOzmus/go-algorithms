package stacks

import "strconv"

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

func CalPoints(ops []string) int {
	var stack []int
	for _, op := range ops {
		switch op {
		case "+":
			pop := stack[len(stack)-1]
			pop2 := stack[len(stack)-2]
			stack = append(stack, pop+pop2)
		case "C":
			stack = stack[:len(stack)-1]
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		default:
			val, _ := strconv.Atoi(op)
			stack = append(stack, val)
		}
	}

	sum := 0
	for _, score := range stack {
		sum += score
	}
	return sum
}
