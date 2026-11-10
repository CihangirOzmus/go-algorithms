// Package twosum solves https://leetcode.com/problems/two-sum/
package twosum

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		if v, ok := seen[target-num]; ok {
			return []int{i, v}
		}
		seen[num] = i
	}
	return nil
}
