package main

import "fmt"

func main() {
	arr := [10]int{9, 1, 2, 8, 6, 7, 5, 3, 0, 4}
	fmt.Println("Initial array is:", arr)
	fmt.Println("")

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Sorted array is:", arr)
}
