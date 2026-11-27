package main

import "fmt"

func aggressiveCows(stall []int, cow int) int {
	end := stall[(len(stall)-1)] - stall[0]
	start := end
	for i := 1; i < len(stall); i++ {
		gap := stall[i] - stall[i-1]
		if gap < start {
			start = gap
		}
	}
	ans := -1

	for start <= end {
		mid := int((start + end) / 2)
		if isValid(stall, cow, mid) {
			ans = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return ans
}

func isValid(stall []int, cow int, mid int) bool {
	N := len(stall)
	cows := 1
	cowPosition := 0
	for i := 1; i < N; i++ {
		if mid <= stall[i]-stall[cowPosition] {
			cows++
			cowPosition = i
		}
	}
	if cow <= cows {
		return true
	}
	return false
}

func main() {
	stall := []int{1, 2, 8, 12, 17, 25, 30}
	cows := 4

	fmt.Println(aggressiveCows(stall, cows))
}
