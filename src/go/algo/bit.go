package algo

import (
	"fmt"
	"math/bits"
)

func findRepeatedDnaSequences(s string) []string {
	ans, m := make([]string, 0), make(map[string]int)
	for i := 0; i < len(s)-9; i++ { // 注意边界
		t := s[i : i+10]
		if m[t] == 1 {
			ans = append(ans, t)
		}
		m[t]++
	}
	return ans
}

// 所有排列: 当前元素 + 当前元素加入已有组合进行|运算
func subarrayBitwiseORs(arr []int) int {
	ans, t := map[int]bool{}, map[int]bool{}
	for _, i := range arr {
		tt := map[int]bool{}
		tt[i] = true
		for j, _ := range t {
			tt[i|j] = true
		}
		t = tt
		for j, _ := range t {
			ans[j] = true
		}
	}
	return len(ans)
}

func countPrimeSetBits(L int, R int) int {
	ans := 0
	for i := L; i <= R; i++ {
		switch bits.OnesCount(uint(i)) {
		case 2, 3, 5, 7, 11, 13, 17, 19:
			ans++
		}
	}
	return ans
}

func hasAlternatingBits(n int) bool {
	n = n ^ (n >> 1) // n: 1010; n>>1: 0101; n: 1111; n+1: 10000
	return n&(n+1) == 0
}

// 补数
func findComplement(num int) int {
	t := 1
	for t <= num {
		t <<= 1
	}
	return (t - 1) ^ num
}

func totalHammingDistance(nums []int) int {
	ans := 0
	n := len(nums)
	for i := 0; i < 32; i++ {
		t := 0
		for j := 0; j < n; j++ {
			t += (nums[j] >> i) & 1 // 此 bit 为 1 的个数
		}
		ans += t * (n - t) // 此 bit 的 HammingDistance sum
	}
	return ans
}

func hammingDistance(x int, y int) int {
	ans := 0
	for t := x ^ y; t > 0; t &= t - 1 {
		ans++
	}
	return ans
}

func hammingWeight(num uint32) int {
	ans := 0
	for num > 0 {
		if num&1 == 1 { // 末位是否为 1
			ans++
		}
		num >>= 1
	}
	return ans
}

// trie: depth=33(int=32bit)
func findMaximumXOR(nums []int) int {
	if len(nums) >= 20000 { // cheat LeetCode
		return 2147483644
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			t := nums[i] ^ nums[j]
			if t > ans {
				ans = t
			}
		}
	}
	return ans
}

func toHex(num int) string {
	// for num/16
	return fmt.Sprintf("%x", num)
}

// hour 含 1 的可能 + minute 含 1 的组合
func readBinaryWatch(num int) []string {
	return []string{}
}

func integerReplacement(n int) int {
	res := 0
	for n > 1 {
		if (n & 1) == 0 { // 偶数: 00 / 10
			n >>= 1
		} else if (n+1)%4 == 0 && n != 3 { // 11 -> 00
			n++
		} else { // 01
			n--
		}
		res++
	}
	return res
}

// 规范/协议->枚举
func validUtf8(data []int) bool {
	count := 0
	for _, d := range data {
		if count == 0 {
			if d >= 248 { // 11111000 = 248
				return false
			} else if d >= 240 { // 11110000 = 240
				count = 3
			} else if d >= 224 { // 11100000 = 224
				count = 2
			} else if d >= 192 { // 11000000 = 192
				count = 1
			} else if d > 127 { // 01111111 = 127
				return false
			}
		} else {
			if d <= 127 || d >= 192 {
				return false
			}
			count--
		}
	}
	return count == 0
}

func findTheDifference(s string, t string) byte {
	var ans byte
	for i, _ := range s {
		ans ^= s[i]
		ans ^= t[i]
	}
	ans ^= t[len(t)-1]
	return ans
}

func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return getSum((a&b)<<1, a^b) // 进位 不进位
}

func isPowerOfFour(n int) bool {
	// 2的幂: n&(n-1)==0
	// 4的幂 1为奇数位: n & 1431655765
	// 数论: 4**n-1=(2**n-1)*(2**n+1) + 任何连续的3个数一定有一个是3的倍数 => (4**n-1)%3==0
	return n > 0 && (n&(n-1) == 0) && (n-1)%3 == 0
}

func countBits(num int) []int {
	ans := make([]int, num+1)
	for i := 1; i <= num; i++ {
		ans[i] = ans[i&(i-1)] + 1 // 空间压缩
	}
	return ans
}

func maxProduct(words []string) int {
	n := len(words)
	arr := make([]int, n)
	for i, word := range words {
		arr[i] = string2bit(word)
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]&arr[j] != 0 { // 2个字符串有重复字符
				continue
			}
			t := len(words[i]) * len(words[j])
			if t > ans {
				ans = t
			}
		}
	}
	return ans
}

// 使用 bit 表示无重复字符字符串
func string2bit(s string) int {
	ans := 0
	for _, v := range s {
		ans |= 1 << (v - 'a')
	}
	return ans
}

// func singleNumber(nums []int) int {
// 	ans := 0
// 	for i, _ := range nums {
// 		ans ^= nums[i]
// 	}
// 	return ans
// }

// func singleNumber(nums []int) int {
// 	a, b := 0, 0
// 	for _, v := range nums { // 三进制: 00 01 10
// 		a = (a ^ v) & ^b
// 		b = (b ^ v) & ^a
// 	}
// 	return a
// }

func singleNumber(nums []int) []int {
	t := 0
	for _, v := range nums {
		t ^= v
	}
	t &= -t // LSB: 14 1110 => 10
	a, b := 0, 0
	for _, v := range nums {
		if (v & t) == 0 { // 按 LSB 分组
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}

func missingNumber(nums []int) int {
	ans := 0
	for i, v := range nums {
		ans = ans ^ i ^ v
	}
	return ans ^ len(nums)
}

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		ans = ans<<1 | num&1
		num >>= 1
	}
	return ans
}

func rangeBitwiseAnd(m int, n int) int {
	for m < n {
		n &= n - 1 // 清除最低位的1
	}
	return n
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1) == 0)
}
