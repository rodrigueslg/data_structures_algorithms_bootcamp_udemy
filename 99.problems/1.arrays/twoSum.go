// https://leetcode.com/problems/two-sum/description/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	//fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	complMap := make(map[int]int)
	for i, v := range nums {
		if val, err := complMap[v]; err {
			return []int{val, i}
		}
		complMap[target-v] = i
	}
	return nil
}
