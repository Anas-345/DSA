package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		mid := int((start + end) / 2)
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			start = mid + 2
		} else {
			end = mid
		}
	}
	return nums[start]
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(nums))
}
