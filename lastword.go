package main

func reverseWord(w string) string {
	r := []rune(w)

	for i := 0; i < len(r)-1; i++ {
		for j := len(r)-1; j > i; j-- {
			r[j], r[j-1] = r[j-1], r[j]
		}
	}

	return string(r)
}

func LastWord(s string) string {
	if len(s) <= 1 {
		return  s + "\n"
	}
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if i == len(s)-1 {
				continue
			} else {
				break
			}
		}

		result += string(s[i])
	}

	return reverseWord(result) + "\n"
}
