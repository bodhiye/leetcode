package main

import "fmt"

func main() {
	fmt.Print(combine(4, 2))
}

func combine(n int, k int) (ans [][]int) {
	tmp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝
		if len(tmp)+(n-cur+1) < k {
			return
		}
		if len(tmp) == k {
			comb := make([]int, k)
			copy(comb, tmp)
			ans = append(ans, comb)
			return
		}
		// 选择当前位置
		tmp = append(tmp, cur)
		dfs(cur + 1)
		// 不选择当前位置
		tmp = tmp[:len(tmp)-1]
		dfs(cur + 1)
	}
	dfs(1)
	return
}
