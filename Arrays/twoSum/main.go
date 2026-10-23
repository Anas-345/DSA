package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i-1]+nums[j] == target {
				return []int{i - 1, j}
			}
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 22

	fmt.Println(twoSum(nums, target))
}
