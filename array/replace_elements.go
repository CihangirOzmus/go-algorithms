package array

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
