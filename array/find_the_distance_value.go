package array

func FindTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	count := 0
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if subs(arr1[i], arr2[j]) <= d {
				count++
				break
			}
		}
	}
	return len(arr1) - count
}

func subs(num1, num2 int) int {
	subs := num1 - num2
	if subs < 0 {
		subs = -subs
	}
	return subs
}
