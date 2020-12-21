package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{{8, 1, 2, 2, 3}, {6, 5, 4, 8}, {7, 7, 7, 7}}
	for _, data := range nums {
		fmt.Println(smallerNumbersThanCurrent(data))
	}
}

type pair struct {
	Pos   int
	Value int
}

func smallerNumbersThanCurrent(nums []int) []int {
	p := make([]pair, len(nums))
	for i, num := range nums {
		p[i] = pair{i, num}
	}
	sort.Slice(p, func(i, j int) bool { return p[i].Value < p[j].Value })

	res := make([]int, len(nums))
	var small int
	for i := range p {
		if i > 0 && p[i].Value != p[i-1].Value {
			small = i
		}
		res[p[i].Pos] = small
	}

	return res
}
