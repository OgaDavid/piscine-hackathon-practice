package main

func containsOnlyAlphabet(str string) bool {
	for _, v := range str {
		if (v < 'a' || v > 'z') && (v < 'A' || v > 'Z') {
			return false
		}
	}
	return true
}

func isUpperCase(s rune) bool {
	return s >= 'A' && s <= 'Z'
}

func CamelToSnakeCase(str string) string {
	result := ""

	if len(str) == 0 || !containsOnlyAlphabet(str) {
		return str
	}

	for i := 0; i < len(str); i++ {
		if i != 0 && isUpperCase(rune(str[i])) && i+1 < len(str) && !isUpperCase(rune(str[i+1])) {
			result += "_"
			result += string(str[i])
		} else if !isUpperCase(rune(str[i])) || (i == 0 && isUpperCase(rune(str[i]))) {
			result += string(str[i])
		} else {
			return str
		}
	}

	return result
}
