package main

import "fmt"

func selection(arr []int) {
	for last := len(arr) - 1; last > 0; last-- {

		maxIndex := 0
		for i := 1; i <= last; i++ {
			if arr[maxIndex] < arr[i] {
				maxIndex = i
			}
		}
		arr[maxIndex], arr[last] = arr[last], arr[maxIndex]
	}
}

func main() {

	arr := []int{9, 1, 2, 8, 6, 7, 5, 3, 0, 4}
	fmt.Println("Initial array is:", arr)
	fmt.Println("")

	selection(arr)

	fmt.Println("Sorted array is:", arr)
}
