package arrays

import "sort"

func TargetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	ans := make([]int, 0)
	for i, num := range nums {
		if num == target {
			ans = append(ans, i)
		}
	}
	return ans
}

func GetConcatenation(nums []int) []int {
	n := len(nums)
	res := make([]int, n*2)
	for i := 0; i < n; i++ {
		res[i] = nums[i]
		res[i+n] = nums[i]
	}
	return res
}

func ReplaceElements(arr []int) []int {
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

func RemoveElement(nums []int, val int) int {
	j := 0
	for _, num := range nums {
		if num != val {
			nums[j] = num
			j++
		}
	}
	return j
}

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
