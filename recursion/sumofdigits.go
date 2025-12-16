package main

func SumOfDigits(n int) int {
	if n < 0 {
		return -1
	}

	if n == 0 {
		return 0
	}

	return n%10 + SumOfDigits(n/10)
}