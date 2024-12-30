// https://leetcode.com/problems/move-zeroes/description/

package main

import "fmt"

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{1, 0})
	moveZeroes([]int{0, 1})
	moveZeroes([]int{2, 1})
}

func moveZeroes(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			temp := nums[right]
			nums[right] = nums[left]
			nums[left] = temp
			left++
		}
	}
	fmt.Println(nums)
}
