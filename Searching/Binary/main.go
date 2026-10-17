package main

import "fmt"

var nums = []int{12, 5, 8, 21, 3, 17, 9, 14, 6, 11}

func main() {
	var num int
	fmt.Print("Enter number you want to search: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	binarySearch(num)
}

func binarySearch(num int) {
	bubbleSort()
	end := (len(nums) - 1)
	for start := 0; start <= end; {
		mid := int((start + end) / 2)
		if nums[mid] == num {
			fmt.Printf("%d found at position %d", num, mid+1)
			return
		} else if nums[mid] < num {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	fmt.Printf("%d not found.", num)
}

func bubbleSort() {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < (len(nums) - i); j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
}
