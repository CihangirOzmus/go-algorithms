// Package Insertionsort solves https://neetcode.io/problems/insertion-sort
package insertionsort

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		for j >= 0 && arr[j+1] < arr[j] {
			arr[j+1], arr[j] = arr[j], arr[j+1]
			j--
		}
	}
	return arr
}
