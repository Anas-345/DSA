package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := int((start + end) / 2)
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[0] {
			if nums[mid] < target || nums[start] > target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if nums[mid] > target || nums[end] < target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return -1
}

func main() {
	nums := []int{6, 7, 1, 2, 3, 4, 5}

	fmt.Println(search(nums, 4))
}
