package main

import "fmt"

// iterative approach to get Factorial of a function
func factorial(num int) int {
	result := 1
	if num == 0 {
		return 1
	}
	for i := 1; i <= num; i++ {
		result = result * i
	}
	return result

}

// Recursive approach to get Factorial of a function
func factRecursive(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	result := num * factRecursive(num-1)
	return result
}

func main() {
	fmt.Println(factorial(6))
	fmt.Println(factorial(0))
	fmt.Println(factorial(1))
	fmt.Println(factRecursive(6))

}
