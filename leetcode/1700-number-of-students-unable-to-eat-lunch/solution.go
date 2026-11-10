package numberofstudentsunabletoeatlunch

func CountStudents(students []int, sandwiches []int) int {
	res := len(students)
	count := [2]int{}
	for _, std := range students {
		count[std]++
	}

	for _, sand := range sandwiches {
		if count[sand] > 0 {
			res--
			count[sand]--
		} else {
			break
		}
	}

	return res
}
