package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a interface{}) {
	fmt.Println(a)
}

func main() {
	a := 10
	b := 15

	result := plus(a, b)

	fmt.Println(result)

	plusPlus(plus(a, b))
}
