package main

import "fmt"

func maxArea(height []int) int {
	start := 0
	end := (len(height) - 1)
	area := 0
	for start < end {
		a := 0
		if height[start] > height[end] {
			a = height[end] * (end - start)
			end--
		} else {
			a = height[start] * (end - start)
			start++
		}
		if a > area {
			area = a
		}

	}
	return area
}

func main() {
	height := []int{1, 8, 6, 100, 50, 4, 8, 3, 7}

	fmt.Println(maxArea(height))
}
