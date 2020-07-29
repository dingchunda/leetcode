package lc

import (
	"math"
	"math/rand"
	"sort"
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

type SolutionPickSec struct {
	nums  []int
	index []int
}

func (s *SolutionPickSec) Len() int {
	return len(s.nums)
}
func (s *SolutionPickSec) Less(i, j int) bool {
	return s.nums[i] < s.nums[j]
}
func (s *SolutionPickSec) Swap(i, j int) {
	s.nums[i], s.nums[j], s.index[i], s.index[j] = s.nums[j], s.nums[i], s.index[j], s.index[i]
}

func ConstructorSolutionPickSec(nums []int) SolutionPickSec {
	index := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		index[i] = i
	}
	s := SolutionPickSec{
		nums:  nums,
		index: index,
	}
	sort.Sort(&s)
	return s
}

func (this *SolutionPickSec) Pick(target int) int {
	left := sort.Search(len(this.nums)-1, func(i int) bool {
		return this.nums[i] >= target
	})
	right := left
	if left < len(this.nums)-1 && this.nums[left+1] == target {
		right = sort.Search(len(this.nums)-1, func(i int) bool {
			return this.nums[i] > target
		})
	}

	//fmt.Println(left,right)
	return this.index[rand.Intn(right-left+1)+left]
}

type SolutionListPick struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func ConstructorSolutionListPick(head *ListNode) SolutionListPick {
	return SolutionListPick{
		head: head,
	}
}

/** Returns a random node's value. */
func (this *SolutionListPick) GetRandom() int {
	big := rand.Intn(math.MaxInt32)
	p := this.head
	cnt := 0
	for p != nil {
		if cnt == big {
			return p.Val
		}
		p = p.Next
		cnt++
	}
	big = big % cnt
	cnt = 0
	p = this.head
	for p != nil {
		if cnt == big {
			return p.Val
		}
		p = p.Next
		cnt++
	}
	return -1
}

func canConstruct(ransomNote string, magazine string) bool {
	hit := [26]int{}
	for _, c := range magazine {
		hit[c-'a']++
	}
	for _, c := range ransomNote {
		hit[c-'a']--
		if hit[c-'a'] < 0 {
			return false
		}
	}
	return true
}

func findTheDifference(s string, t string) byte {
	dic := [26]int{}
	for _, c := range s {
		dic[c-'a']++
	}
	dic2 := [26]int{}
	for _, c := range t {
		dic2[c-'a']++
	}
	for i := 0; i < 26; i++ {
		if dic2[i] == dic[i]+1 {
			return byte(i) + 'a'
		}
	}
	return 0
}
