package lc

import (
	"math"
	"sort"
)

func firstMissingPositive(nums []int) int {
	find := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			find = true
		}
		if nums[i] <= 0 || nums[i] > len(nums) {
			nums[i] = 1
		}
	}
	if !find {
		return 1
	}
	if len(nums) == 1 {
		return 2
	}
	for i := 0; i < len(nums); i++ {
		a := int(math.Abs(float64(nums[i])))
		nums[a-1] = -int(math.Abs(float64(nums[a-1])))
	}
	//fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func trap(height []int) int {
	l := len(height)
	maxLeft := make([]int, l)
	maxRight := make([]int, l)
	for i := 0; i < l; i++ {
		if i == 0 {
			maxLeft[0] = height[0]
		} else {
			if height[i] > maxLeft[i-1] {
				maxLeft[i] = height[i]
			} else {
				maxLeft[i] = maxLeft[i-1]
			}
		}
	}
	for i := 0; i < l; i++ {
		if i == 0 {
			maxRight[i] = height[l-1]
		} else {
			if height[l-i-1] > maxRight[i-1] {
				maxRight[i] = height[l-i-1]
			} else {
				maxRight[i] = maxRight[i-1]
			}
		}
	}

	sum := 0
	for i := 1; i < l-1; i++ {
		left := maxLeft[i-1]
		right := maxRight[l-i-2]
		min := left
		if right < min {
			min = right
		}
		min -= height[i]
		if min > 0 {
			sum += min
		}
	}
	return sum
}

func multiply(num1 string, num2 string) string {
	var rst = make([]byte, len(num1)+len(num2)+2)
	for i := len(num1) - 1; i >= 0; i-- {
		ii := byte(num1[i]) - byte('0')
		for j := len(num2) - 1; j >= 0; j-- {
			jj := byte(num2[j]) - byte('0')
			start := len(num1) - 1 - i + len(num2) - 1 - j
			rst[start] += ii * jj

			for k := start; ; k++ {
				if rst[k] < 10 {
					break
				}
				rst[k+1] += rst[k] / 10
				rst[k] = rst[k] % 10
			}
		}
	}

	ln := 0
	for i := len(rst) - 1; i >= 0; i-- {
		if rst[i] > 0 {
			ln = i + 1
			break
		}
	}

	if ln == 0 {
		return "0"
	}

	var s = make([]byte, ln)
	for i := ln - 1; i >= 0; i-- {
		s[ln-i-1] = rst[i] + byte('0')
	}
	return string(s)
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	cans := make([]int, len(nums))
	cans[len(nums)-1] = 0
	for i := 0; i < len(nums)-1; i++ {
		cans[i] = math.MaxInt32
	}
	for at := len(nums) - 2; at >= 0; at-- {
		for to := at + 1; to <= at+nums[at] && to < len(nums); to++ {
			if cans[to] >= 0 && cans[at] > cans[to]+1 {
				cans[at] = cans[to] + 1
			}
		}
	}
	return cans[0]
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	rst := [][]int{{nums[0]}}
	for i := 1; i < len(nums); i++ {
		var tmp = make([][]int, 0, len(rst)*(len(rst[0])+1))
		for _, v := range rst {
			for l := 0; l <= len(v); l++ {
				bf := make([]int, len(v)+1)
				bf[l] = nums[i]
				copy(bf[:l], v[:l])
				copy(bf[l+1:], v[l:])
				tmp = append(tmp, bf)
			}
		}
		rst = tmp
	}
	return rst
}

func permuteUnique(nums []int) [][]int {
	var travel func(data []bool)
	var buffer []int
	var rst [][]int
	travel = func(data []bool) {
		uni := map[int]int{}
		for i, h := range data {
			if !h {
				uni[nums[i]] = i
			}
		}
		if len(uni) == 0 {
			r := make([]int, len(buffer))
			copy(r, buffer)
			rst = append(rst, r)
			return
		}
		for v, index := range uni {
			data[index] = true
			buffer = append(buffer, v)
			travel(data)
			buffer = buffer[:len(buffer)-1]
			data[index] = false
		}
	}
	hit := make([]bool, len(nums))
	travel(hit)
	return rst
}

func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	l := len(matrix[0])
	for x := 0; x < l/2; x++ {
		for y := x; y < l-x-1; y++ {
			matrix[x][y], matrix[y][l-x-1], matrix[l-x-1][l-y-1], matrix[l-y-1][x] =
				matrix[l-y-1][x], matrix[x][y], matrix[y][l-x-1], matrix[l-x-1][l-y-1]
		}
	}
}

func groupAnagrams(strs []string) [][]string {
	group := map[[26]int][]string{}
	for _, str := range strs {
		dist := [26]int{}
		for _, v := range str {
			dist[v-'a']++
		}
		group[dist] = append(group[dist], str)
	}
	var rst = make([][]string, 0, len(group))
	for _, g := range group {
		rst = append(rst, g)
	}
	return rst
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	neg := false
	if n < 0 {
		n = -n
		neg = true
	}
	neg2 := false
	if x < 0 && n%2 == 1 {
		x = -x
		neg2 = true
	}
	//var max float64
	//max = math.MaxFloat64

	m := map[int]float64{}
	var my func(x float64, n int) float64
	my = func(x float64, n int) float64 {
		if v, ok := m[n]; ok {
			return v
		}
		if n == 1 {
			return x
		}
		var rst float64
		if n%2 == 0 {
			rst = my(x, n/2) * my(x, n/2)
		} else {
			rst = my(x, n/2) * my(x, n/2) * x
		}
		m[n] = rst
		return rst
	}
	rst := my(x, n)
	if neg {
		rst = 1 / rst
	}
	if neg2 {
		rst = -rst
	}
	return rst
}

func solveNQueens(n int) [][]string {
	attackRow := make([]bool, n)
	attackDia := map[int]bool{}
	attackDia2 := map[int]bool{}

	var dfs func(at int)
	var stand []int
	var rst [][]string
	dfs = func(at int) {
		if at == n {
			var ss = make([]string, 0, n)
			for _, at := range stand {
				s := make([]byte, n)
				for j := 0; j < n; j++ {
					s[j] = '.'
				}
				s[at] = 'Q'
				ss = append(ss, string(s))
			}
			rst = append(rst, ss)
			return
		}
		for j, ok := range attackRow {
			if !ok && !attackDia[-at+j] && !attackDia2[at+j] {
				attackRow[j] = true
				attackDia[-at+j] = true
				attackDia2[at+j] = true
				stand = append(stand, j)
				dfs(at + 1)
				attackRow[j] = false
				attackDia[-at+j] = false
				attackDia2[at+j] = false
				stand = stand[:len(stand)-1]
			}
		}
	}

	dfs(0)
	return rst
}

func maxSubArray(nums []int) int {
	var maxneg = math.MinInt32
	var ok bool
	for _, n := range nums {
		if n < 0 && n > maxneg {
			maxneg = n
		}
		if n >= 0 {
			ok = true
		}
	}
	if !ok {
		return maxneg
	}
	var rst = math.MinInt32
	var maxsub func(nums []int) int
	maxsub = func(nums []int) int {
		if len(nums) == 0 {
			return 0
		}
		sum := 0
		var index int
		for index = 0; index < len(nums) && nums[index] >= 0; index++ {
			sum += nums[index]
		}
		neg := 0
		for ; index < len(nums) && nums[index] < 0; index++ {
			neg += nums[index]
		}
		left := maxsub(nums[index:])
		v := max(sum, sum+neg+left)
		if v > rst {
			rst = v
		}
		return v
	}
	var index int
	for index = 0; index < len(nums) && nums[index] < 0; index++ {
	}
	maxsub(nums[index:])
	return rst
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var ln = len(matrix) / 2
	if len(matrix)%2 == 1 {
		ln = len(matrix)/2 + 1
	}
	rst := make([]int, 0, len(matrix)*len(matrix[0]))
	for i := 0; i < ln; i++ {
		//start := matrix[i][i]
		if i >= len(matrix[0])-i {
			break
		}
		for j := i; j < len(matrix[0])-i; j++ {
			rst = append(rst, matrix[i][j])
		}
		if i+1 >= len(matrix)-i {
			break
		}
		for j := i + 1; j < len(matrix)-i; j++ {
			rst = append(rst, matrix[j][len(matrix[0])-i-1])
		}
		if len(matrix[0])-i-2 < i {
			break
		}
		for j := len(matrix[0]) - i - 2; j >= i; j-- {
			rst = append(rst, matrix[len(matrix)-i-1][j])
		}
		if len(matrix)-i-2 <= i {
			break
		}
		for j := len(matrix) - i - 2; j > i; j-- {
			rst = append(rst, matrix[j][i])
		}
	}
	return rst
}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	cans := make([]bool, len(nums))
	cans[len(nums)-1] = true
	for at := len(nums) - 2; at >= 0; at-- {
		for to := at + 1; to <= at+nums[at] && to < len(nums); to++ {
			if cans[to] {
				cans[at] = true
				break
			}
		}
	}
	return cans[0]
}

func merge(v [][]int) [][]int {
	sort.Sort(intervals(v))

	var rst [][]int
	mergeAt := -1
	for i := 0; i < len(v); i++ {
		if mergeAt == -1 {
			mergeAt = i
			continue
		}
		if v[mergeAt][1] >= v[i][0] {
			if v[mergeAt][1] < v[i][1] {
				v[mergeAt][1] = v[i][1]
			}
			continue
		}
		rst = append(rst, v[mergeAt])
		mergeAt = i
	}
	if mergeAt != -1 {
		rst = append(rst, v[mergeAt])
	}
	return rst
}

type intervals [][]int

func (v intervals) Less(i, j int) bool {
	return v[i][0] < v[j][0]
}

func (v intervals) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v intervals) Len() int {
	return len(v)
}

func getPermutation(n int, k int) string {
	nj := make([]int, n+1)
	nj[1] = 1
	for i := 2; i < n; i++ {
		nj[i] = nj[i-1] * i
	}
	at := make([]int, n)
	for i := 0; i < n-1; i++ {
		start := 0
		sum := nj[n-i-1]
		for k > sum {
			sum += nj[n-i-1]
			start++
		}
		at[i] = start
		k -= nj[n-i-1] * start
	}
	hit := make([]int, n)
	var s = make([]byte, n)
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if hit[j] == 0 {
				if cnt == at[i] {
					hit[j] = 1
					s[i] = byte(j+1) + '0'
					break
				}
				cnt++
			}
		}
	}
	return string(s)
}
