package main

import "fmt"

func main() {
	nums := [4]int{1, 2, 3, 4}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	keys := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range keys {
		fmt.Println("key:", k, ", value:", v)
	}

	for k := range keys {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
