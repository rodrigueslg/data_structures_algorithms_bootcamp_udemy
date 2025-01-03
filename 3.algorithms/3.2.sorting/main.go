package main

import "fmt"

func main() {
	fmt.Println("Bubble sorted: ", BubbleSort([]int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}))
	fmt.Println("Selection sorted: ", SelectionSort([]int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}))
}

func BubbleSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	i := 0
	sortedCount := 0
	for {
		if sortedCount == len(numbers) {
			break
		}
		if i == len(numbers)-1 {
			sortedCount++
			i = 0
			continue
		}
		if numbers[i] > numbers[i+1] {
			temp := numbers[i]
			numbers[i] = numbers[i+1]
			numbers[i+1] = temp
		}
		i++
	}

	return numbers
}

func SelectionSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	i := 0
	sortedCount := 0
	smallerIndex := 0
	for {
		if sortedCount == len(numbers) {
			break
		}
		if i == len(numbers) {
			smaller := numbers[smallerIndex]
			numbers[smallerIndex] = numbers[sortedCount]
			numbers[sortedCount] = smaller

			sortedCount++
			i = sortedCount
			smallerIndex = sortedCount
			continue
		}
		if numbers[i] < numbers[smallerIndex] {
			smallerIndex = i
		}
		i++
	}

	return numbers
}
