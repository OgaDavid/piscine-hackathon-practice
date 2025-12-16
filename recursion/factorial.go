package main

// 5! -> 5*4! => 4! -> 4*3! ....

// recursive solution
func Factorial(n int) int {
	if n < 0 {
		return -1
	}

	// base case
	if n == 0 || n == 1 {
		return 1
	}

	return n * Factorial(n-1)
}

// iterative solution
func IterativeFactorial(n int) int {
	if n < 0 {
		return -1
	}

	res := 1

	for i := 2; i <= n; i++ {
		res *= i
	}

	return res
}