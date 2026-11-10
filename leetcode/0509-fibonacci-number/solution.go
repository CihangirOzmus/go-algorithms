// Package fibonaccinumber solves https://leetcode.com/problems/fibonacci-number/
package fibonaccinumber

var memo = make(map[int]int)

func Fib(n int) int {
	if n <= 1 {
		return n
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = Fib(n-1) + Fib(n-2)
	return memo[n]
}
