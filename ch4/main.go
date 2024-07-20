package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 2, 5, 7, 3, 7, 8, 5, 3, 2, 6}
	fmt.Printf("Sum(arr): %v\n", Sum(arr))
	fmt.Printf("QuickSort(arr): %v\n", QuickSort(arr))
}

func Sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	} else {
		return arr[0] + Sum(arr[1:])
	}
}


func QuickSort(arr []int) []int {
	fmt.Printf("arr: %v\n", arr)
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less, greater, ans []int
	for _,val := range arr[1:] {
		if val < pivot {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}
	fmt.Printf("less: %v\n", less)
	fmt.Printf("greater: %v\n", greater)
	
	ans = append(ans, QuickSort(less)...)
	ans = append(ans, pivot)
	ans = append(ans, QuickSort(greater)...)

	return ans
}