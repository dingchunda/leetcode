package lc

import (
	"math"
	"sort"
)

func minCostII(costs [][]int) int {
	if len(costs) == 0 || len(costs[0]) == 0 {
		return 0
	}
	ans := make([][]int, len(costs))
	color := len(costs[0])
	for i := range ans {
		ans[i] = make([]int, color)
	}
	for i := 0; i < color; i++ {
		ans[0][i] = costs[0][i]
	}
	for i := 1; i < len(costs); i++ {
		mv := math.MaxInt32
		mAt := -1
		for index, v := range ans[i-1] {
			if v < mv {
				mAt = index
				mv = v
			}
		}
		for j := 0; j < color; j++ {
			if j != mAt {
				ans[i][j] = mv + costs[i][j]
			}
		}
		smv := math.MaxInt32
		for index, v := range ans[i-1] {
			if index != mAt && v < smv {
				smv = v
			}
		}
		ans[i][mAt] = smv + costs[i][mAt]
	}
	mv := math.MaxInt32
	for i := 0; i < color; i++ {
		if v := ans[len(costs)-1][i]; v < mv {
			mv = v
		}
	}
	return mv
}

func missingNumber(nums []int) int {
	sort.Ints(nums)
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != k {
			return k
		}
		k++
	}
	return k
}

func alienOrder(words []string) string {
	total := map[byte]bool{}
	m := map[[2]byte]int{}
	for i := 0; i < len(words); i++ {
		if i < len(words)-1 {
			left, right := words[i], words[i+1]
			j := 0
			for ; j < len(left) && j < len(right) && left[j] == right[j]; j++ {
			}
			if j < len(left) && j == len(right) {
				return ""
			}
			if j < len(left) && j < len(right) {
				dir := 1
				a, b := left[j], right[j]
				if a > b {
					a, b = b, a
					dir = -1
				}
				dir2, ok := m[[2]byte{a, b}]
				if ok && dir != dir2 {
					return ""
				}
				if !ok {
					m[[2]byte{a, b}] = dir
				}
			}
		}
		for _, c := range words[i] {
			total[byte(c)] = true
		}
	}

	//fmt.Println(m)
	income := map[byte]map[byte]bool{}
	for pair := range m {
		if _, ok := income[pair[0]]; !ok {
			income[pair[0]] = map[byte]bool{}
		}
		if _, ok := income[pair[1]]; !ok {
			income[pair[1]] = map[byte]bool{}
		}
	}
	for pair, dir := range m {
		if dir > 0 {
			income[pair[1]][pair[0]] = true
		} else {
			income[pair[0]][pair[1]] = true
		}
	}
	//fmt.Println(income)

	var rst []byte
	for len(income) > 0 {
		var only byte
		find := false
		for c, sources := range income {
			if len(sources) == 0 {
				only = c
				find = true
				break
			}
		}
		if !find {
			return ""
		}
		rst = append(rst, only)
		for _, sources := range income {
			delete(sources, only)
		}
		delete(income, only)
		delete(total, only)
	}
	for k := range total {
		rst = append(rst, k)
	}
	//fmt.Println(rst)
	return string(rst)
}

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		ans := make([]int, n)
		for i := 0; i < n; i++ {
			if ans[i] == 2 {
				continue
			}
			for j := 0; j < n; j++ {
				if i != j {
					if knows(j, i) {
						ans[j] = 2
						if knows(i, j) {
							ans[i] = 2
							break
						}
					} else {
						ans[i] = 2
						break
					}
				}
			}
			//fmt.Println(ans[i])
			if ans[i] == 2 {
				continue
			}
			return i
		}
		return -1
	}
}

func numSquares(n int) int {
	rst := make([]int, n+1)
	rst[0] = 0
	rst[1] = 1
	for i := 2; i <= n; i++ {
		for p := 1; ; p++ {
			l := p * p
			if l > i {
				break
			}
			cur := rst[i-l] + 1
			if rst[i] == 0 {
				rst[i] = cur
			} else if cur < rst[i] {
				rst[i] = cur
			}
		}
	}
	return rst[n]
}
