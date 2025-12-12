package main

func RetainFirstHalf(str string) string {
	n := len(str)

	if n == 0 {
		return ""
	}
	if n == 1 {
		return str
	}

	// Convert to runes so we handle multi-byte characters correctly
	runes := []rune(str)
	half := len(runes) / 2

	result := make([]rune, half)
	for i := 0; i < half; i++ {
		result[i] = runes[i]
	}

	return string(result)
}
