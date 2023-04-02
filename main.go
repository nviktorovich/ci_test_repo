package main

import "fmt"

func main() {
	fmt.Println(GetSum(1, 2, 3, 5, 7))
}

func GetSum(nums ...int) int {
	var res int = 0
	for _, num := range nums {
		res += num
	}
	return res
	sfsdfsdf
}
