package main

func HashCode(dec string) string {
	result := ""

	for _, v := range dec {
		hash := (int(v) + len(dec)) % 127

		if hash < 32 || hash > 126 {
			hash += 33
		}

		result += string(rune(hash))
	}

	return result
}
