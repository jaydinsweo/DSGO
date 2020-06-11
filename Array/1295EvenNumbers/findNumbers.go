package main

import "fmt"

func FindNumbers(nums []int) int {

	// init a zero value int
	result := 0

	// loop through nums array - increment if even number of digits
	for _, value := range nums {
		// sprintf formats the string and return the formatted string
		// - not print to console
		// %d - print the argument as a number
		if len(fmt.Sprintf("%d", value))%2 == 0 {
			result++
		}
	}
	return result
}

func main() {
	FindNumbers([]int{12, 32, 555, 2, 42})
}
