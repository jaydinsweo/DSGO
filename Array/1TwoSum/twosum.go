package main

//TwoSum function here
func TwoSum(array []int, target int) []int {
	// initialise a map of int to int
	m := make(map[int]int)

	for i, num := range array {
		// if (target-num) is eq to the index number - found
		if v, found := m[target-num]; found {
			return []int{v, i}
		}
		m[num] = i
	}
	return nil

}
