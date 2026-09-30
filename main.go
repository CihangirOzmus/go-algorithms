package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 4, 5, 3, 1, 2}
	fmt.Println(replaceElements(nums))
}

func replaceElements(arr []int) []int {
	j := len(arr) - 1
	rMax := arr[j]
	for j >= 0 {
		if arr[j] > rMax {
			temp := rMax
			rMax = arr[j]
			arr[j] = temp
		} else {
			arr[j] = rMax
		}
		j--
	}
	arr[len(arr)-1] = -1
	return arr
}

func removeElement(nums []int, val int) int {
	j := 0
	for _, num := range nums {
		if num != val {
			nums[j] = num
			j++
		}
	}
	return j
}

func findMaxConsecutiveOnes(nums []int) int {
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
