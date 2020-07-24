package lc

import "math"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	m := map[int]bool{}
	for _, c := range coins {
		m[c] = true
	}
	rst := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		if m[i] {
			rst[i] = 1
			continue
		}
		if i < 2 {
			continue
		}
		min := math.MaxInt32
		for left := 1; left <= i/2; left++ {
			right := i - left
			if rst[left] > 0 && rst[right] > 0 {
				sum := rst[left] + rst[right]
				if sum < min {
					min = sum
				}
			}
		}
		if min < math.MaxInt32 {
			rst[i] = min
		}
	}
	if rst[amount] > 0 {
		return rst[amount]
	}
	return -1
}

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	white := head
	black := head.Next
	if black == nil {
		return white
	}
	firstBlack := black
	for white.Next != nil && white.Next.Next != nil {
		white.Next = white.Next.Next
		white = white.Next
		if black.Next != nil && black.Next.Next != nil {
			black.Next = black.Next.Next
			black = black.Next
		}
	}
	black.Next = nil
	white.Next = firstBlack
	return head
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var dfs func(i, j int) int
	ilen, jlen := len(matrix), len(matrix[0])
	var path []int
	result := make([]int, ilen*jlen)
	for i := 0; i < len(result); i++ {
		result[i] = -1
	}
	dfs = func(i, j int) int {
		if result[jlen*i+j] != -1 {
			return result[jlen*i+j]
		}
		forward := false
		max := 0
		if i+1 < ilen && matrix[i+1][j] > path[len(path)-1] {
			path = append(path, matrix[i+1][j])
			v := dfs(i+1, j)
			if v > max {
				max = v
			}
			path = path[:len(path)-1]
			forward = true
		}
		if i-1 >= 0 && matrix[i-1][j] > path[len(path)-1] {
			path = append(path, matrix[i-1][j])
			v := dfs(i-1, j)
			if v > max {
				max = v
			}
			path = path[:len(path)-1]
			forward = true
		}
		if j+1 < jlen && matrix[i][j+1] > path[len(path)-1] {
			path = append(path, matrix[i][j+1])
			v := dfs(i, j+1)
			if v > max {
				max = v
			}
			path = path[:len(path)-1]
			forward = true
		}
		if j-1 >= 0 && matrix[i][j-1] > path[len(path)-1] {
			path = append(path, matrix[i][j-1])
			v := dfs(i, j-1)
			if v > max {
				max = v
			}
			path = path[:len(path)-1]
			forward = true
		}
		if !forward {
			result[jlen*i+j] = 1
			return 1
		}
		result[jlen*i+j] = 1 + max
		return 1 + max
	}
	max := 0
	for i := 0; i < ilen; i++ {
		for j := 0; j < jlen; j++ {
			path = append(path, matrix[i][j])
			if v := dfs(i, j); v > max {
				max = v
			}
			path = path[:len(path)-1]
		}

	}
	return max
}

func increasingTriplet(nums []int) bool {
	var left int
	var middle int
	var found bool
	for i := 1; i < len(nums); i++ {
		if found && nums[i] > middle {
			return true
		}
		if !found && nums[i] > nums[i-1] {
			left = nums[i-1]
			middle = nums[i]
			found = true
			continue
		}
		if found && nums[i] > left && nums[i] < middle {
			middle = nums[i]
			continue
		}
		if found && nums[i] > nums[i-1] && nums[i-1] < middle {
			middle = nums[i]
			left = nums[i-1]
		}
	}
	return false
}

func rob3(root *TreeNode) int {
	m := map[*TreeNode]int{}
	var travel func(r *TreeNode) int
	travel = func(r *TreeNode) int {
		if v, ok := m[r]; ok {
			return v
		}
		if r == nil {
			return 0
		}
		sum := r.Val
		sum2 := 0
		if r.Left != nil {
			if r.Left.Left != nil {
				sum += travel(r.Left.Left)
			}
			if r.Left.Right != nil {
				sum += travel(r.Left.Right)
			}
			sum2 += travel(r.Left)
		}
		if r.Right != nil {
			if r.Right.Left != nil {
				sum += travel(r.Right.Left)
			}
			if r.Right.Right != nil {
				sum += travel(r.Right.Right)
			}
			sum2 += travel(r.Right)
		}
		var max int
		if sum > sum2 {
			max = sum
		} else {
			max = sum2
		}
		m[r] = max
		return max
	}
	return travel(root)
}

func countBits(num int) []int {
	rst := make([]int, num+1)
	start := 1
	for i := 0; i <= 31 && start <= num; i++ {
		pair := 1 << uint(i)
		count := 0
		for at := start; at <= num; {
			rst[at]++
			count++
			if count == pair {
				count = 0
				at += pair + 1
			} else {
				at++
			}
		}
		start += pair
	}
	return rst
}

func lengthOfLongestSubstringKDistinct(s string, K int) int {
	if K == 0 {
		return 0
	}
	ans := 0
	i, j := 0, len(s)
	for i <= j {
		mid := (i + j) >> 1
		table := map[byte]int{}
		for k := 0; k < mid; k++ {
			table[s[k]]++
		}
		if len(table) <= K {
			ans = mid
			i = mid + 1
			continue
		}
		find := false
		for k := mid; k < len(s); k++ {
			table[s[k-mid]]--
			if table[s[k-mid]] == 0 {
				delete(table, s[k-mid])
			}
			table[s[k]]++
			if len(table) <= K {
				find = true
				break
			}
		}
		if find {
			ans = mid
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return ans
}

func countComponents(n int, edges [][]int) int {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = -1
	}
	var find func(at int) int
	find = func(at int) int {
		if parent[at] == -1 {
			return at
		}
		return find(parent[at])
	}
	var merge = func(i, j int) {
		seti := find(i)
		setj := find(j)
		if seti != setj {
			parent[seti] = setj
		}
	}
	for _, e := range edges {
		merge(e[0], e[1])
	}
	ans := 0
	for _, p := range parent {
		if p == -1 {
			ans++
		}
	}
	return ans
}
