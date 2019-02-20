package main

import "fmt"

func primarySum(input int) int {

	sum := 0

	for i := 2; i <= input; i++ {
		flag := false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = true
				break
			}
		}
		if !flag {
			sum += i
		}
	}

	return sum
}

func main() {

	var input int

	fmt.Scanln(&input)

	result := primarySum(input)

	fmt.Println(result)
}
