package main

import "fmt"

func main() {
	nums := [4]int{1, 2, 3, 4}

	sum := 0

	// 배열 순회 - (_, 값) -> '_'는 값을 사용하지 않을 때 쓴다.
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 배열 순회 - (인덱스, 값)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	keys := map[string]string{"a": "apple", "b": "banana"}

	// 맵 자료구조 순회 - (키, 값)
	for k, v := range keys {
		fmt.Println("key:", k, ", value:", v)
	}

	// 맵 자료구조의 순회 - (키, 값)
	for k := range keys {
		fmt.Println("key:", k)
	}

	// 문자열 파싱 - (인덱스, 값)
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
