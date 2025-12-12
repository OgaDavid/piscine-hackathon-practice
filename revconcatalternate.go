package main

// import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {

	result := []int{}
	var max []int
	var min []int

	l1 := len(slice1)
	l2 := len(slice2)

	switch {
	case l1 > l2:
		max = slice1
		min = slice2
	case l1 < l2:
		max = slice2
		min = slice1
	case l1 == l2:
		max = slice1
		min = slice2
	}

	// fmt.Printf("max: %d ", max)
	// fmt.Printf("min: %d Ans: ", min)

	if l1 == 0 || l2 == 0 {
		for i := len(max) - 1; i >= 0; i-- {
			result = append(result, max[i])
		}

		return result
	}

	if l1 != l2 {
		for i := len(max) - 1; i >= len(max)-len(min)-1; i-- {
			result = append(result, max[i])
		}

		for i := len(max) - len(min) - 2; i >= 0; i-- {
			result = append(result, slice1[i])
			result = append(result, slice2[i])
		}
		return result
	}

	if l1 == l2 {
		for i := len(max) - 1; i >= 0; i-- {
			result = append(result, max[i])
			result = append(result, min[i])
		}
		return result
	}

	return result
}
