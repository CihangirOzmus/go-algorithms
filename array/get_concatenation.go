package array

func GetConcatenation(nums []int) []int {
	n := len(nums)
	res := make([]int, n*2)
	for i := range n {
		res[i] = nums[i]
		res[i+n] = nums[i]
	}
	return res
}
