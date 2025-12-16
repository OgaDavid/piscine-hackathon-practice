package main

// we didn't have to slice arrays at all we can just swap
// starting and end positions and solve recursively

func BinarySearch(arr []int, target, start, end int) int {
	if start > end {
		return -1
	}

	m := start + (end-start)/2
	if arr[m] == target {
		return m
	}

	if target < arr[m] {
		return BinarySearch(arr, target, start, m-1)
	}

	return BinarySearch(arr, target, m+1, end)
}

// func BinarySearch(arr []int, n int) int {
// 	arrlen := len(arr)
// 	mid := arrlen / 2

// 	var tempArr []int

// 	if arrlen == 0 {
// 		return -1
// 	}

// 	if arr[mid] != n {
// 		if arr[mid] < n {
// 			tempArr = arr[mid+1:]
// 			result := BinarySearch(tempArr, n)
// 			if result == -1 {
// 				return -1
// 			}
// 			return (mid + 1) + BinarySearch(tempArr, n)
// 		} else {
// 			tempArr = arr[:mid]
// 			return BinarySearch(tempArr, n)
// 		}
// 	}

// 	return mid
// }
