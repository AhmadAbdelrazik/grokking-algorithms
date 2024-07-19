package main

import (
	"slices"
)

func main() {
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallest_index := 0
	for i := range arr {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}

	return smallest_index
}

func selectionSort(arr []int) []int {
	newArr := []int{}

	for range arr {
		idx := findSmallest(arr)
		newArr = append(newArr, arr[idx])
		arr = slices.Delete(arr, idx, idx+1)
	}
	
	return newArr
}
