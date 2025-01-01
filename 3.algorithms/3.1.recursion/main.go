package main

func main() {
	// fmt.Println(FindFactorialIterative(5))
	// fmt.Println(FindFactorialRecursive(5))

	// fmt.Println(FindFibonacciIterative(8))
	// fmt.Println(FindFibonacciRecursive(8))
	// fmt.Println(FindFibonacciRecursive(45)) // // O(2^N) - very slow, too many calculations

	//fmt.Println(ReverseStringRecursive("yoyo master"))
	ReverseStringRecursive("yoyo master")
}

func FindFactorialIterative(number int) int {
	result := 1
	for i := number; i > 1; i-- {
		result *= i
	}
	return result
}

func FindFactorialRecursive(number int) int {
	if number == 2 {
		return 2
	}
	return number * FindFactorialRecursive(number-1)
}

func FindFibonacciIterative(index int) int {
	fib := 1
	prev := 0
	for i := 1; i < index; i++ {
		fib += prev
		prev = fib - prev
	}
	return fib
}

func FindFibonacciRecursive(index int) int { // O(2^N)
	if index < 2 {
		return index
	}
	return FindFibonacciRecursive(index-1) + FindFibonacciRecursive(index-2)
}

func ReverseStringRecursive(s string) string {
	if s == "" {
		return ""
	}

	return ReverseStringRecursive(s[1:]) + string(s[0])
}
