package main

import "fmt"

func splitArraySum(nums []int, s int) int {
	start := 0
	end := arrSum(nums)
	var ans int
	for start <= end {
		mid := int((start + end) / 2)
		if isValid(nums, s, mid) {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}

func isValid(num []int, s int, mid int) bool {
	st := 1
	pages := 0
	for i := 0; i < len(num); i++ {
		if num[i] > mid {
			return false
		}
		if pages+num[i] <= mid {
			pages += num[i]
		} else {
			st++
			pages = num[i]
		}
	}
	if st > s {
		return false
	}
	return true
}

func arrSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func main() {
	nums := []int{7, 2, 5, 10, 8}
	fmt.Println(splitArraySum(nums, 2))
}
