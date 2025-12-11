package main

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	// Check the first word
	if isLetter(s[0]) && isLower(s[0]) {
		return false
	}

	// Check subsequent words
	for i := 1; i < len(s); i++ {
		if s[i] == ' ' && i+1 < len(s) {
			next := s[i+1]

			// If next word starts with a letter, it must be uppercase
			if isLetter(next) && isLower(next) {
				return false
			}
		}
	}

	return true
}

func isLetter(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func isLower(c byte) bool {
	return (c >= 'a' && c <= 'z')
}
