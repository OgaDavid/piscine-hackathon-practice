package main

func isPrimeNumber(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func FindPrevPrime(nb int) int {
	num := 0
	for i := nb; i > 0; i-- {
		if isPrimeNumber(i) {
			num = i
			break
		} else {
			continue
		}
	}

	return num
}
