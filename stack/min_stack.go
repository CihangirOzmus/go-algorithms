package stack

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	minTmp := val
	if len(s.minStack) > 0 {
		pop := s.minStack[len(s.minStack)-1]
		if minTmp > pop {
			minTmp = pop
		}
	}
	s.minStack = append(s.minStack, minTmp)
}

func (s *MinStack) Pop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
	}
	if len(s.minStack) > 0 {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.minStack) > 0 {
		return s.minStack[len(s.minStack)-1]
	}
	return -1
}
