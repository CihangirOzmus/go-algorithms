package array

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
