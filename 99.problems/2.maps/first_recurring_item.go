// Return first recurrent value
// Given a string, return the first recurring character in it, or -1 if there is no recurring character.
// Given the string [2,5,1,2,3,5,1,2,4], should return 2
// Given the string [2,5,5,2,3,5,1,2,4], should return 5
// Given the string [2,1,1,2,3,5,1,2,4], should return 1
// Given the string [2,3,4,5], should return -1

// Return first recurrent value - Improoved
// Given a string, return the first recurring character in it, or -1 if there is no recurring character.
// The first recurring character is the character that appears twice or more in the string and its index is the smallest
// Given the string [2,5,1,2,3,5,1,2,4], should return 2
// Given the string [2,5,5,2,3,5,1,2,4], should return 2
// Given the string [2,1,1,2,3,5,1,2,4], should return 2
// Given the string [2,3,4,5], should return -1

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(firstRecurringItem([]int{2, 5, 1, 2, 3, 5, 1, 2, 4})) // 2
	fmt.Println(firstRecurringItem([]int{2, 5, 5, 2, 3, 5, 1, 2, 4})) // 5
	fmt.Println(firstRecurringItem([]int{2, 1, 1, 2, 3, 5, 1, 2, 4})) // 1
	fmt.Println(firstRecurringItem([]int{2, 3, 4, 5}))                // -1

	fmt.Println(firstRecurringItemImprooved([]int{2, 5, 1, 2, 3, 5, 1, 2, 4})) // 2
	fmt.Println(firstRecurringItemImprooved([]int{2, 5, 5, 2, 3, 5, 1, 2, 4})) // 2
	fmt.Println(firstRecurringItemImprooved([]int{2, 1, 1, 2, 3, 5, 1, 2, 4})) // 2
	fmt.Println(firstRecurringItemImprooved([]int{2, 3, 4, 5}))                // -1
}

func firstRecurringItem(nums []int) int {
	m := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if m[num] {
			return num
		}
		m[num] = true
	}
	return -1
}

func firstRecurringItemImprooved(nums []int) int {
	m := make(map[int]string, len(nums))

	smallestRepeatedIndex := len(nums)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if m[num] != "" {
			repeatedIndex, _ := strconv.Atoi(m[num])
			if repeatedIndex < smallestRepeatedIndex {
				smallestRepeatedIndex = repeatedIndex
			}
		}
		m[num] = strconv.Itoa(i)
	}

	if smallestRepeatedIndex == len(nums) {
		return -1
	}

	return nums[smallestRepeatedIndex]
}
