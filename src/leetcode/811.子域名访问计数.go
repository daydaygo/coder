package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=811 lang=golang
 *
 * [811] 子域名访问计数
 */

// @lc code=start
func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, v := range cpdomains {
		vv := strings.Split(v, " ")
		cnt, _ := strconv.Atoi(vv[0])
		domain := vv[1]
		if _, ok := m[domain]; ok {
			m[domain] += cnt
		} else {
			m[domain] = cnt
		}
		i := strings.Index(domain, ".")
		for i != -1 {
			domain = domain[i+1:]
			if _, ok := m[domain]; ok {
				m[domain] += cnt
			} else {
				m[domain] = cnt
			}

			i = strings.Index(domain, ".")
		}
	}
	var ans []string
	for s, i := range m {
		ans = append(ans, strconv.Itoa(i)+" "+s)
	}
	return ans
}

// @lc code=end
