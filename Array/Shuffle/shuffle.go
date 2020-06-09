package main

// Given an array of 2n elements [x1,x2,x3,..,y1,y2,y3,..]
// Return the array in the form [x1,y1,x2,y2,x3,y3,...]

func Shuffle(nums []int, n int) []int {

	head := nums[:n]
	tail := nums[n:]
	result := []int{}

	for i := 0; i < n; i++ {
		result = append(result, head[i], tail[i])
	}
	return result
}
