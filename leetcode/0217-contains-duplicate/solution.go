// Package containsduplicate solves https://leetcode.com/problems/contains-duplicate/
package containsduplicate

func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}
