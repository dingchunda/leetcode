package lc

import "math"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	rst := make([][]byte, len(matrix))
	for i := 0; i < len(rst); i++ {
		rst[i] = make([]byte, len(matrix[0]))
	}
	max := byte(0)
	for i := 0; i < len(matrix) && i < len(matrix[0]); i++ {
		for j := i; j < len(matrix[i]); j++ {
			if i == 0 {
				rst[i][j] = matrix[i][j] - '0'
				if rst[i][j] > max {
					max = rst[i][j]
				}
				continue
			}
			if matrix[i][j] == '0' {
				rst[i][j] = 0
				continue
			}
			min := byte(math.MaxUint8)
			if m := rst[i][j-1]; m < min {
				min = m
			}
			if m := rst[i-1][j]; m < min {
				min = m
			}
			if m := rst[i-1][j-1]; m < min {
				min = m
			}
			rst[i][j] = min + 1
			if rst[i][j] > max {
				max = rst[i][j]
			}
		}

		for j := i + 1; j < len(matrix); j++ {
			if i == 0 {
				rst[j][i] = matrix[j][i] - '0'
				if rst[j][i] > max {
					max = rst[j][i]
				}
				continue
			}
			if matrix[j][i] == '0' {
				rst[j][i] = 0
				continue
			}
			min := byte(128)
			if m := rst[j-1][i]; m < min {
				min = m
			}
			if m := rst[j][i-1]; m < min {
				min = m
			}
			if m := rst[j-1][i-1]; m < min {
				min = m
			}
			rst[j][i] = min + 1
			if rst[j][i] > max {
				max = rst[j][i]
			}
		}
	}
	return int(max * max)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	var stack []*TreeNode
	count := 0
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}
	return count
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	s1 := (C - A) * (D - B)
	s2 := (G - E) * (H - F)
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var overlay int
	if F < D && H > B && A < G && E < C {
		k1 := min(H, D) - max(B, F)
		k2 := min(C, G) - max(A, E)
		overlay = k1 * k2
	}
	return s1 + s2 - overlay
}

func calculate(s string) int {
	return calculate2(s, false)
}

func calculate2(s string, neg bool) int {
	i := 0
	for ; i < len(s) && s[i] == ' '; i++ {
	}
	num := 0
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		} else {
			break
		}
	}
	if neg {
		num = -num
	}
again:
	var op byte
	var nextNeg bool
	for ; i < len(s) && op == 0; i++ {
		if s[i] == '+' || s[i] == '*' || s[i] == '/' {
			op = s[i]
		} else if s[i] == '-' {
			op = '+'
			nextNeg = true
		}
	}
	if op == 0 {
		return num
	}
	if op == '*' || op == '/' {
		for ; i < len(s) && s[i] == ' '; i++ {
		}
		num2 := 0
		for ; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				num2 = num2*10 + int(s[i]-'0')
			} else {
				break
			}
		}
		if op == '*' {
			num *= num2
		} else {
			num /= num2
		}
		goto again
	}
	right := calculate2(s[i:], nextNeg)
	switch op {
	case '+':
		return num + right
	case '-':
		return num - right
	case '*':
		return num * right
	case '/':
		return num / right
	}
	return num
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var travel func(r *TreeNode, visitor func(int) bool) bool
	count := 0
	travel = func(r *TreeNode, visitor func(int) bool) bool {
		if r == nil {
			return true
		}
		if r.Left != nil {
			if travel(r.Left, visitor) {
				return true
			}
		}
		if visitor(r.Val) {
			return true
		}
		if r.Right != nil {
			if travel(r.Right, visitor) {
				return true
			}
		}
		return false
	}
	var rst int
	travel(root, func(i int) bool {
		count++
		if count == k {
			rst = i
			return true
		}
		return false
	})
	return rst
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome3(head *ListNode) bool {
	var data []int
	p := head
	for p != nil {
		data = append(data, p.Val)
		p = p.Next
	}
	var i, j int
	if len(data)&1 == 1 {
		i = len(data)/2 - 1
		j = len(data)/2 + 1
	} else {
		i, j = len(data)/2-1, len(data)/2
	}
	for i >= 0 && j < len(data) {
		if data[i] != data[j] {
			return false
		}
		i--
		j++
	}
	return true
}

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	var possible *TreeNode
	var rst []*TreeNode
	var travel func(node *TreeNode) bool
	travel = func(cur *TreeNode) bool {
		rst = append(rst, cur)
		if cur == p {
			if possible != nil {
				return true
			}
			possible = p
		} else if cur == q {
			if possible != nil {
				return true
			}
			possible = q
		}
		if cur.Left != nil {
			if travel(cur.Left) {
				return true
			}
		}
		if cur.Right != nil {
			if travel(cur.Right) {
				return true
			}
		}
		rst = rst[:len(rst)-1]
		if len(rst) > 0 {
			if cur == possible {
				possible = rst[len(rst)-1]
			}
		} else {
			possible = nil
		}
		return false
	}
	travel(root)
	return possible
}

func productExceptSelf(nums []int) []int {
	rst := make([]int, len(nums))
	if len(nums) == 2 {
		rst[0] = nums[1]
		rst[1] = nums[0]
		return rst
	}

	rst[0] = 1
	for i := 1; i < len(nums); i++ {
		rst[i] = rst[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 1; i-- {
		rst[0] *= nums[i+1]
		rst[i] *= rst[0]
	}
	rst[0] *= nums[1]
	return rst
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	var heap []int
	var pos = map[int][]int{}
	var swap = func(i, j int) {
		ivs, jvs := pos[heap[i]], pos[heap[j]]
		for k, v := range ivs {
			if v == i {
				for s, v := range jvs {
					if v == j {
						ivs[k], jvs[s] = jvs[s], ivs[k]
						break
					}
				}
				break
			}
		}
		pos[heap[i]] = ivs
		pos[heap[j]] = jvs
		heap[i], heap[j] = heap[j], heap[i]
	}

	var down func(at int)
	down = func(at int) {
		left := at*2 + 1
		right := at*2 + 2
		if left < len(heap) && heap[left] > heap[at] {
			if right < len(heap) && heap[right] > heap[left] {
				swap(at, right)
				down(right)
			} else {
				swap(at, left)
				down(left)
			}
		}
		if right < len(heap) && heap[right] > heap[at] {
			swap(at, right)
			down(right)
		}
	}
	var up func(at int)
	up = func(at int) {
		var back int
		if at%2 == 1 {
			back = at / 2
		} else {
			back = at/2 - 1
		}
		if back >= 0 && heap[at] > heap[back] {
			swap(at, back)
			up(back)
		}
	}
	var rst []int
	i := 0
	for ; i < len(nums) && i < k; i++ {
		heap = append(heap, nums[i])
		pos[nums[i]] = append(pos[nums[i]], len(heap)-1)
		up(len(heap) - 1)
	}
	rst = append(rst, heap[0])
	//fmt.Println(heap, pos)
	for ; i < len(nums); i++ {
		at := pos[nums[i-k]][0]
		pos[nums[i-k]] = pos[nums[i-k]][1:]
		pos[nums[i]] = append(pos[nums[i]], at)
		heap[at] = nums[i]
		up(at)
		down(at)
		rst = append(rst, heap[0])
		//fmt.Println(i, heap, pos)
	}
	return rst
}

func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	return resecureSearchMatrix(matrix, target, 0, len(matrix[0]))
}
func resecureSearchMatrix(matrix [][]int, target int, start, end int) bool {
	if len(matrix) == 0 {
		return false
	}
	if start == end {
		return false
	}
	at := len(matrix) / 2
	pos := (start + end) / 2
	v := matrix[at][pos]
	if v == target {
		return true
	}
	if v < target {
		if resecureSearchMatrix(matrix, target, pos+1, end) {
			return true
		}
		if resecureSearchMatrix(matrix[at+1:], target, start, pos+1) {
			return true
		}
	} else {
		if resecureSearchMatrix(matrix[at:], target, start, pos) {
			return true
		}
		if resecureSearchMatrix(matrix[:at], target, start, end) {
			return true
		}
	}
	return false
}
