package main

import (
	"fmt"
)

func Hash() {
	m := make(map[int]int)
	if v, ok := m[1]; ok { // 查
		fmt.Println(v)
	}
	m[10]++ // 增/改
	fmt.Println(m)

	// 减少一次 for 循环
	var ans, n, k int
	var sum []int
	sumHash := make(map[int]int)
	for i := 1; i <= n; i++ {
		// for j := 0; j < i; j++ { // hash: 减少一次遍历
		// 	if sum[i]-sum[j] == k {
		// 		ans++
		// 	}
		// }
		j := sum[i] - k // sum[j]
		if v, ok := sumHash[j]; ok {
			ans += v
		}
		m[i]++
	}
}
