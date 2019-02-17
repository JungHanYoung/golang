package main

import "fmt"
import "strconv"

func fibo(num int64) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fibo(num-1) + fibo(num-2)
	}
}

func main() {

	var num string
	fmt.Print("number: ")
	fmt.Scanln(&num)

	parse, err := strconv.ParseInt(num, 10, 0)
	if err == nil {
		result := fibo(parse)
		fmt.Println(result)
	}
}
