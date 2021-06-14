package algo

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func subtractProductAndSum(n int) int {
	a, b := 1, 0
	for n > 0 {
		t := n % 10
		n /= 10
		a *= t
		b += t
	}
	return a - b
}

func numPrimeArrangements(n int) int {
	// p := getPrimeInN(100)
	p := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	t := len(p) // n 以内质数的个数
	for i := 0; i < len(p); i++ {
		if p[i] > n {
			t = i
			break
		}
	}
	fmt.Println(t)
	return factorial(t) * factorial(n-t) % 1000000007
}

func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1) % 1000000007
}

func getPrimeInN(n int) []int {
	// 埃氏筛法 n以内的质数: O(logN) O(N)
	m := map[int]bool{}
	for i := 2; i*i < n; i++ {
		if m[i] {
			continue
		}
		for j := i * i; j < n; j += i { // i的倍数都不是质数
			m[j] = true // 非质数
		}
	}
	ans := make([]int, 0)
	for i := 2; i < n; i++ {
		if !m[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

func dayOfYear(date string) int {
	begin := date[:4] + "-01-01"
	a, _ := time.Parse("2006-01-02", begin)
	b, _ := time.Parse("2006-01-02", date)
	d := b.Sub(a)
	return int(d.Hours()/24) + 1
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func tribonacci(n int) int { // 泰伯纳妾
	a, b, c := 0, 1, 1
	for i := 0; i < n; i++ {
		a, b, c = b, c, a+b+c
	}
	return a
}

func isBoomerang(points [][]int) bool {
	// (x0-x1)(y0-y2) != (x0-x2)(y0-y1)
	return (points[0][0]-points[1][0])*(points[0][1]-points[2][1]) != (points[0][0]-points[2][0])*(points[0][1]-points[1][1])
}

func baseNeg2(N int) string {
	if N == 0 {
		return "0"
	}
	res := ""
	for N != 0 { // 短除法
		t := N % (-2)
		N /= -2
		if t < 0 { // 处理负数
			t += 2
			N++
		}
		res = strconv.Itoa(t) + res
	}
	return res
}

func powerfulIntegers(x int, y int, bound int) []int {
	if bound < 2 {
		return []int{}
	}
	if x == 1 && y == 1 {
		return []int{2}
	}
	m := make(map[int]int)
	if x > y {
		x, y = y, x
	}
	for i := 0; powInt(x, i) <= bound; i++ {
		for j := 0; powInt(y, j)+powInt(x, i) <= bound; j++ {
			m[powInt(y, j)+powInt(x, i)] = 1
		}
		if x == 1 {
			break
		}
	}
	ans := make([]int, 0, len(m))
	for k, _ := range m {
		ans = append(ans, k)
	}
	return ans
}

func powInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func largestTimeFromDigits(arr []int) string {
	flag, ans := false, 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 4; k++ {
				if i == k || j == k {
					continue
				}
				hour, min := 10*arr[i]+arr[j], 10*arr[k]+arr[6-i-j-k]
				if hour < 24 && min < 60 {
					if hour*60+min >= ans {
						ans = hour*60 + min
						flag = true
					}
				}
			}
		}
	}
	if !flag {
		return ""
	}
	return fmt.Sprintf("%02d:%02d", ans/60, ans%60)
}

func surfaceArea(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			v := grid[i][j]
			if v == 0 {
				continue
			}
			ans += 4*v + 2
			// up: i j+1
			if j < len(grid[0])-1 {
				ans -= min(v, grid[i][j+1])
			}
			// down: i j-1
			if j > 0 {
				ans -= min(v, grid[i][j-1])
			}
			// left: i-1 j
			if i > 0 {
				ans -= min(v, grid[i-1][j])
			}
			// right: i+1 j
			if i < len(grid)-1 {
				ans -= min(v, grid[i+1][j])
			}
		}
	}
	return ans
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if rec1[0] == rec1[2] || rec1[1] == rec1[3] || rec2[0] == rec2[2] || rec2[1] == rec2[3] {
		return false
	}
	return rec1[0] < rec2[2] && rec2[0] < rec1[2] && rec1[1] < rec2[3] && rec2[1] < rec1[3]
}

func largestTriangleArea(points [][]int) float64 {
	var ans float64
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				t := TriangleArea(points[i][0], points[i][1], points[j][0], points[j][1], points[k][0], points[k][1])
				fmt.Println(ans, t)
				if t > ans {
					ans = t
				}
			}
		}
	}
	return ans
}

func TriangleArea(x1, y1, x2, y2, x3, y3 int) float64 {
	t := x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)
	if t < 0 {
		t = -t
	}
	return float64(t) / 2
}

func numRabbits(answers []int) int {
	ans := 0
	m := make(map[int]int)
	for _, answer := range answers {
		if m[answer] == 0 {
			ans += answer + 1 // other rabbit
			m[answer] = answer
		} else {
			m[answer]--
		}
	}
	return ans
}

func findErrorNums(nums []int) []int {
	ans := make([]int, 2)
	m := make(map[int]int)
	for _, num := range nums {
		if m[num] == 1 {
			ans[0] = num
		} else {
			m[num] = 1
		}
	}
	for i := 1; i <= len(nums); i++ {
		if m[i] == 0 {
			ans[1] = i
			break
		}
	}
	return ans
}

func maxCount(m int, n int, ops [][]int) int {
	maxM, maxN := m, n
	for _, op := range ops {
		if maxM > op[0] {
			maxM = op[0]
		}
		if maxN > op[1] {
			maxN = op[1]
		}
	}
	return maxM + maxN
}

func checkPerfectNumber(num int) bool {
	switch num {
	case 6, 28, 496, 8128, 33550336:
		return true
	}
	return false
}

func rand10() int {
	ans := 40
	for ans >= 40 {
		ans = (rand7()-1)*7 + rand7() - 1 // [0,49)
	}
	return ans%10 + 1 // [0,10) -> [1,10]
}

func rand7() int {
	return rand.Intn(7) + 1 // [0,7) => [1,7]
}

func minMoves(nums []int) int {
	sum, min, l := 0, nums[0], len(nums)
	for _, num := range nums {
		sum += num
		if num < min {
			min = num
		}
	}
	return sum - min*l
}

func superPow(a int, b []int) int {
	ans := 1
	for _, i := range b {
		ans = (quickPow(ans, 10, 1337) * quickPow(a, i, 1337)) % 1337
	}
	return ans
}

func quickPow(a, b, c int) int { // 快速计算 a^b % c
	ans := 1
	a %= c
	for b > 0 {
		if (b & 1) == 1 {
			ans = (ans * a) % c
		}
		a = (a * a) % c
		b >>= 1
	}
	return ans
}

func isPowerOfThree(n int) bool {
	// return n > 0 && (1162261467%n == 0) // 3^19
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	for num%2 == 0 {
		num /= 2
	}
	for num%3 == 0 {
		num /= 3
	}
	for num%5 == 0 {
		num /= 5
	}
	return num == 1
}

func addDigits(num int) int {
	for num >= 10 {
		t := 0
		for num > 0 {
			t += num % 10
			num /= 10
		}
		num = t
	}
	return num
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	// 重合部分
	X0, Y0, X1, Y1 := max(A, E), max(B, F), min(C, G), min(D, H)
	return area(A, B, C, D) + area(E, F, G, H) - area(X0, Y0, X1, Y1)
}

func area(x0, y0, x1, y1 int) int {
	l, h := x1-x0, y1-y0
	if l <= 0 || h <= 0 {
		return 0
	}
	return l * h
}

func countPrimes(n int) int {
	// 埃氏筛法 n以内的质数: O(logN) O(N)
	m := map[int]bool{}
	for i := 2; i*i < n; i++ {
		if m[i] {
			continue
		}
		for j := i * i; j < n; j += i { // i的倍数都不是质数
			m[j] = true // 非质数
		}
	}
	ans := 0
	for i := 2; i < n; i++ {
		if !m[i] {
			ans++
		}
	}
	return ans
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			n /= i
		}
	}
	return n == 1
}

func isHappy(n int) bool {
	m := map[int]bool{}
	for n > 0 {
		t, tmp := 0, n
		for tmp > 0 { // 18 => 1*1+8*8
			t += (tmp % 10) * (tmp % 10)
			tmp /= 10
		}
		if t == 1 {
			return true
		}
		if _, ok := m[t]; ok { // 循环退出
			return false
		}
		m[t] = true
		n = t
	}
	return false
}

func trailingZeroes(n int) int {
	ans := 0
	for n > 0 {
		ans += n / 5 // 5的倍数
		n /= 5
	}
	return ans
}

func titleToNumber(s string) int {
	ans, t := 0, 0
	for _, i := range s {
		t = int(i-'A') + 1
		ans = ans*26 + t
	}
	return ans
}

func convertToTitle(n int) string {
	ans := ""
	for n != 0 { // n>0
		n -= 1                             // [1,26] -> [0,25]
		ans = string(byte(n%26)+'A') + ans // 低位 -> 高位
		n /= 26
	}
	return ans
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	t := reverse(x)
	// s := strconv.Itoa(x)
	return t == x
}

func reverse(x int) int {
	ans := 0
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	if ans > 1<<31-1 || ans < -(1<<31) { // [-2^31, 2^31-1]
		return 0
	}
	return ans
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	n := minutesToTest/minutesToDie + 1 // 状态
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(n))))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}

func max3(a, b, c int) int {
	return max(max(a, b), c)
}

// gcd 最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
