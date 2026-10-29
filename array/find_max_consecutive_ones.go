package array

func FindMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	temp := 0
	for _, num := range nums {
		if num == 1 {
			temp++
		} else {
			maxCount = findMax(temp, maxCount)
			temp = 0
		}
	}

	return findMax(temp, maxCount)
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
