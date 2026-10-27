package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := int((start + end) / 2)

		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] {
			return mid
		} else if arr[mid] > arr[mid+1] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{0, 10, 5, 2}

	fmt.Println(peakIndexInMountainArray(arr))
}
