package string

import (
	"strconv"
	"strings"
)

func CountDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	aStart, aEnd := toDay(arriveAlice), toDay(leaveAlice)
	bStart, bEnd := toDay(arriveBob), toDay(leaveBob)
	start := max(aStart, bStart)
	end := min(aEnd, bEnd)
	if start <= end {
		return end - start + 1
	}
	return 0
}

func toDay(str string) int {
	daysInMonths := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	parts := strings.Split(str, "-")
	month, _ := strconv.Atoi(parts[0])
	day, _ := strconv.Atoi(parts[1])
	for i := month; i <= 12; i++ {
		day += daysInMonths[i]
	}
	return day
}
