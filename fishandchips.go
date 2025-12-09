package main

func FishAndChips(n int) string {
	result := ""

	switch {
	case n < 0:
		result = "error: number is negative"
	case n%2 == 0 && n%3 == 0:
		result = "fish and chips"
	case n%2 == 0:
		result = "fish"
	case n%3 == 0:
		result = "chips"
	default:
		result = "error: non divisible"
	}

	return result
}
