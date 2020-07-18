package lc

import (
	"math/rand"
	"strings"
)

type Solution struct {
	origin []int
}

func ConstructorSolution(nums []int) Solution {
	return Solution{
		origin: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	rst := make([]int, len(this.origin))
	copy(rst, this.origin)
	return rst
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	rst := make([]int, len(this.origin))
	copy(rst, this.origin)
	if len(rst) == 0 {
		return rst
	}
	for i := 0; i < len(rst); i++ {
		at := rand.Intn(len(rst))
		rst[i], rst[at] = rst[at], rst[i]
	}
	return rst
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func lexicalOrder(n int) []int {
	var rst = make([]int, 0, n)
	var dfs func(num int)
	dfs = func(num int) {
		if num > n {
			return
		}
		rst = append(rst, num)
		for i := 0; i < 10; i++ {
			num = num*10 + i
			if num <= n {
				dfs(num)
				num = (num - i) / 10
			}
		}
	}
	for i := 1; i < 10; i++ {
		dfs(i)
	}
	return rst
}

func firstUniqChar(s string) int {
	m := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]] = false
		} else {
			m[s[i]] = true
		}

	}
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok && v {
			return i
		}
	}
	return -1
}

func decodeString(s string) string {
	buffer := make([]byte, 0, len(s))
	var poss []int
	var times []int
	var curTime int
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			curTime = curTime*10 + int(c-'0')
			continue
		}
		if c == '[' {
			times = append(times, curTime)
			curTime = 0
			poss = append(poss, len(buffer))
			continue
		}
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			buffer = append(buffer, c)
			continue
		}
		if c == ']' {
			time := times[len(times)-1]
			pos := poss[len(poss)-1]
			src := buffer[pos:]
			for cur := 0; cur < time-1; cur++ {
				buffer = append(buffer, src...)
			}
			times = times[:len(times)-1]
			poss = poss[:len(poss)-1]
		}
	}
	return string(buffer)
}

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	hit := map[rune]int{}
	for _, c := range s {
		hit[c]++
	}
	for key, v := range hit {
		if v < k {
			max := 0
			for _, sub := range strings.Split(s, string(key)) {
				if rst := longestSubstring(sub, k); rst > max {
					max = rst
				}
			}
			return max
		}
	}
	return len(s)
}
