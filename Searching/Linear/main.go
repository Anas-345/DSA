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
	search(num)
}

func search(num int) {
	for i := 0; i < len(nums); i++ {
		if num == nums[i] {
			fmt.Printf("%d found at position %d", num, i+1)
			return
		}
	}
	fmt.Printf("%d not found.", num)
}
