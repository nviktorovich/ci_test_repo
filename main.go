package main

import "fmt"

func main() {
	a := NewNumList(1, 2, 3, 4, 5)
	fmt.Println(a.Sum())

}

type NumList struct {
	args []int
}

func NewNumList(args ...int) *NumList {
	return &NumList{args: args}
}

func (n *NumList) Sum() int {
	res := 0
	for _, v := range n.args {
		res += v
	}
	return res
	sfsdfsdf
}
