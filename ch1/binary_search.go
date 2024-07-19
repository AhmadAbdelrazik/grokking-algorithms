package main

import "fmt"

func main() {
	arr := []int{
		1, 3, 5, 6, 8, 11, 15, 34, 35, 64, 75, 88, 100, 131, 142,
		144, 145, 162, 163, 172, 177, 190, 210, 215, 217, 221, 223,
	}

	fmt.Printf("binarySearch(arr, 5): %v\n", binarySearch(arr, 5))
	fmt.Printf("simpleSearch(arr, 5): %v\n", simpleSearch(arr, 5))
	fmt.Printf("binarySearch(arr, 53): %v\n", binarySearch(arr, 53))
	fmt.Printf("simpleSearch(arr, 53): %v\n", simpleSearch(arr, 53))
	fmt.Printf("binarySearch(arr, 142): %v\n", binarySearch(arr, 142))
	fmt.Printf("simpleSearch(arr, 142): %v\n", simpleSearch(arr, 142))
	fmt.Printf("binarySearch(arr, 217): %v\n", binarySearch(arr, 217))
	fmt.Printf("simpleSearch(arr, 217): %v\n", simpleSearch(arr, 217))
	fmt.Printf("binarySearch(arr, 414): %v\n", binarySearch(arr, 414))
	fmt.Printf("simpleSearch(arr, 414): %v\n", simpleSearch(arr, 414))
}

func binarySearch(arr []int, goal int) bool {
	l := 0
	h := len(arr) - 1

	m := (l + h) / 2
	iterations := 0
	for l <= h {
		iterations++
		if goal < arr[m] {
			h = m - 1
		} else if goal > arr[m] {
			l = m + 1
		} else if goal == arr[m] {
			fmt.Printf("bs iterations: %v\n", iterations)
			return true
		}
		m = (l + h) / 2
	}

	fmt.Printf("bs iterations: %v\n", iterations)
	return false
}

func simpleSearch(arr []int, goal int) bool {
	iterations := 0
	for _, m := range arr {
		iterations++
		if goal == m {
			fmt.Printf("ss iterations: %v\n", iterations)
			return true
		}
	}
	fmt.Printf("ss iterations: %v\n", iterations)
	return false
}
