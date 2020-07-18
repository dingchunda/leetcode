package lc

import (
	"math"
	"strings"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	rst := 0
	for i := 1; i < len(prices); i++ {
		if v := prices[i] - min; v > rst {
			rst = v
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return rst
}

func maxProfit2(prices []int) int {
	sum := 0
	base := 0
	hold := false
	for i := 0; i < len(prices); i++ {
		if !hold && i < len(prices)-1 && prices[i] < prices[i+1] {
			hold = true
			base = prices[i]
			continue
		}
		if hold && (i == len(prices)-1 || prices[i] > prices[i+1]) {
			hold = false
			sum += prices[i] - base
		}
	}
	return sum
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m := map[*TreeNode]int{}
	var maxChild func(r *TreeNode) int
	rst := math.MinInt32
	maxChild = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		v, ok := m[r]
		if !ok {
			v = max(0, max(maxChild(r.Left), maxChild(r.Right))) + r.Val
			m[r] = v
		}
		if v := max(0, m[r.Left]) + max(0, m[r.Right]) + r.Val; v > rst {
			rst = v
		}
		return v
	}
	maxChild(root)
	return rst
}

func isPalindrome2(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if !valid(s[i]) {
			i++
		} else if !valid(s[j]) {
			j--
		} else {
			c1 := s[i]
			c2 := s[j]
			if c1 >= 'A' && c1 <= 'Z' {
				c1 += 'a' - 'A'
			}
			if c2 >= 'A' && c2 <= 'Z' {
				c2 += 'a' - 'A'
			}
			//fmt.Println(c1,c2)
			if c1 != c2 {
				return false
			} else {
				i++
				j--

			}
		}
	}
	return true
}

func valid(b byte) bool {
	return b >= '0' && b <= '9' || b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

type border struct {
	Val int
	S   int // -1 +1
}

func longestConsecutive(nums []int) int {
	max := 0
	table := map[border]border{}
	dist := map[int]bool{}
	for _, num := range nums {
		if _, ok := dist[num]; ok {
			continue
		}
		dist[num] = true

		left := border{
			Val: num,
			S:   -1,
		}
		right := border{
			Val: num,
			S:   1,
		}
		rightMax, ok1 := table[left]
		leftMin, ok2 := table[right]
		var newMax int
		if ok1 && ok2 {
			table[leftMin] = rightMax
			table[rightMax] = leftMin
			delete(table, left)
			delete(table, right)
			newMax = rightMax.Val - leftMin.Val - 1
		} else if ok1 && !ok2 {
			delete(table, left)
			newLeft := border{
				Val: num - 1,
				S:   -1,
			}
			table[newLeft] = rightMax
			table[rightMax] = newLeft
			newMax = rightMax.Val - num
		} else if ok2 && !ok1 {
			delete(table, right)
			newRight := border{
				Val: num + 1,
				S:   1,
			}
			table[newRight] = leftMin
			table[leftMin] = newRight
			newMax = num - leftMin.Val
		} else {
			newLeft := border{
				Val: num - 1,
				S:   -1,
			}
			newRight := border{
				Val: num + 1,
				S:   1,
			}
			table[newLeft] = newRight
			table[newRight] = newLeft
			newMax = 1
		}
		if newMax > max {
			max = newMax
		}
	}
	return max
}

func partition2(s string) (rst [][]string) {
	index := make([][]bool, len(s))
	for i := 0; i < len(index); i++ {
		index[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		index[i][i] = true
		left := i
		right := i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			index[left][right] = true
			left--
			right++
		}
		left = i - 1
		right = i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			index[left][right] = true
			left--
			right++
		}
	}

	var buffer []string
	var dfs func(at int)
	dfs = func(at int) {
		if at >= len(s) {
			dst := make([]string, len(buffer))
			copy(dst, buffer)
			rst = append(rst, dst)
			return
		}
		rights := index[at]
		for right := at; right < len(s); right++ {
			if rights[right] {
				buffer = append(buffer, s[at:right+1])
				dfs(right + 1)
				buffer = buffer[:len(buffer)-1]
			}
		}
	}
	dfs(0)
	return
}

func canCompleteCircuit(gas []int, cost []int) int {
	for start := 0; start < len(gas); start++ {
		amount := 0
		first := false
		for k := start; k != start || !first; k = (k + 1) % len(gas) {
			first = true
			amount += gas[k]
			amount -= cost[k]
			if amount < 0 {
				goto next
			}
		}
		return start
	next:
	}
	return -1
}

func singleNumber(nums []int) int {
	xor := 0
	for _, n := range nums {
		xor ^= n

	}
	return xor
}

func singleNumber2(nums []int) int {
	rst := 0
	for i := 0; i < 64; i++ {
		total := 0
		for _, num := range nums {
			if num>>uint(i)&1 == 1 {
				total++
			}
		}
		if total%3 != 0 {
			rst |= 1 << uint(i)
		}
	}
	return rst
}

type RNode struct {
	Val    int
	Next   *RNode
	Random *RNode
}

func copyRandomList(head *RNode) *RNode {
	if head == nil {
		return nil
	}
	p := head
	ch := &RNode{
		Val:    p.Val,
		Random: p,
	}
	p2 := ch
	for p.Next != nil {
		p2.Next = &RNode{
			Val:    p.Next.Val,
			Random: p.Next,
		}
		p = p.Next
		p2 = p2.Next
	}

	p2 = ch
	var after []func()
	for p2 != nil {
		if p2.Random.Random != nil {
			from := ch
			for from != nil {
				if from.Random == p2.Random.Random {
					c := p2
					after = append(after, func() {
						c.Random = from
					})
					break
				}
				from = from.Next
			}
		} else {
			c := p2
			after = append(after, func() {
				c.Random = nil
			})
		}
		p2 = p2.Next
	}
	for _, f := range after {
		f()
	}
	print(ch)
	return ch
}

func wordBreak(s string, wordDict []string) bool {
	data := []byte(s)
	rst := make([]bool, len(s))
	for i := 0; i < len(data); i++ {
		p := string(data[len(data)-1-i:])
		for _, w := range wordDict {
			if len(p) == len(w) {
				if p == w {
					rst[i] = true
					break
				}
			} else if len(p) > len(w) && strings.HasPrefix(p, w) {
				if rst[i-len(w)] {
					rst[i] = true
					break
				}
			}
		}
	}
	return rst[len(rst)-1]
}
