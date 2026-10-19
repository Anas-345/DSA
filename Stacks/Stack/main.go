package main

import "fmt"

const N int = 10

var stack = [N]int{}
var top int = -1

func main() {
	isContinue := true
	var choice int
	for isContinue {
		fmt.Println("1- Push value in stack")
		fmt.Println("2- Pop value from stack")
		fmt.Println("3- Exit")
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		switch choice {
		case 1:
			var value int
			fmt.Print("Enter value you want to push: ")
			_, err := fmt.Scan(&value)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			push(value)
		case 2:
			err, value := pop()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%d popped from stack", value)
		case 3:
			isContinue = false
			fmt.Println("Exiting...")
		default:
			fmt.Println("Invalid input")
		}
	}
}

func push(x int) {
	if top == N-1 {
		fmt.Println("Stack is full")
		return
	}
	top++
	stack[top] = x
}

func pop() (error, int) {
	if top == -1 {
		return fmt.Errorf("Stack is empty"), 0
	}
	value := stack[top]
	top--
	return nil, value
}
