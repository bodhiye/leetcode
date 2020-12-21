package main

import (
	"fmt"
)

func main() {
	len := uniquePaths(2,3)
	fmt.Println(len)
}

func uniquePaths(m, n int) int {
	f := make([]int, m)
	f[0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j-1 >= 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[len(f)-1]
}
