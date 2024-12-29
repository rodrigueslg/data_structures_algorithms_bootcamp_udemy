package main

import "slices"

// Problem: Given an array of integers, return true if the array contains a pair that sums up to a specific target number.
// [1, 2, 3, 9] sum = 8 | result: false
// [1, 2, 4, 4] sum = 8 | result: true

func main() {
	println(hasPairWithSum([]int{}, 8))           // false
	println(hasPairWithSum([]int{1, 2, 3, 9}, 8)) // false
	println(hasPairWithSum([]int{1, 2, 4, 4}, 8)) // true
}

func hasPairWithSum(array []int, sum int) bool {
	var complements []int
	for _, value := range array {
		if slices.Contains(complements, value) {
			return true
		}
		complements = append(complements, sum-value)
	}
	return false
}
