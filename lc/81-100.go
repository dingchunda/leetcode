package lc

import "unsafe"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	h := head
	var pre *ListNode
	for p != nil && p.Next != nil {
		z := p
		for z.Next != nil && z.Next.Val == z.Val {
			z = z.Next
		}
		if p == h && z != p {
			h = z.Next
		}
		if p != z && pre != nil {
			pre.Next = z.Next
		}
		if p == z {
			pre = z
		}
		p = z.Next
	}
	return h
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	for p != nil && p.Next != nil {
		if p.Next.Val == p.Val {
			q := p.Next
			for ; q != nil && q.Val == p.Val; q = q.Next {
			}
			p.Next = q
			p = q
		} else {
			p = p.Next
		}
	}
	return head
}

func subsetsWithDup(nums []int) [][]int {
	cntTable := map[int]int{}
	for _, n := range nums {
		cntTable[n]++
	}
	sets := make([][][]int, 0, len(cntTable))
	for n, cnt := range cntTable {
		var set [][]int
		set = append(set, nil)
		for i := 1; i <= cnt; i++ {
			s := make([]int, i)
			for j := 0; j < i; j++ {
				s[j] = n
			}
			set = append(set, s)
		}
		sets = append(sets, set)
	}
	var rst [][]int
	var buffer []int
	var travel func(at int)
	travel = func(at int) {
		if at == len(sets) {

			dst := make([]int, len(buffer))
			copy(dst, buffer)
			rst = append(rst, dst)

			return
		}
		for _, s := range sets[at] {
			ln := len(buffer)
			buffer = append(buffer, s...)
			travel(at + 1)
			buffer = buffer[:ln]
		}
	}
	travel(0)
	return rst
}

func numDecodings(s string) int {
	rst := make([]int, len(s))
	if s[len(s)-1] > '0' {
		rst[0] = 1
	}
	if len(s) == 1 {
		return rst[0]
	}
	for i := 1; i < len(s); i++ {
		if s[len(s)-i-1] > '0' {
			rst[i] = rst[i-1]
		}
		da := (s[len(s)-i-1]-'0')*10 + (s[len(s)-i] - '0')
		if da >= 10 && da <= 26 {
			if i == 1 {
				rst[i] += 1
			} else {
				rst[i] += rst[i-2]
			}
		}
	}
	return rst[len(s)-1]
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	p := head
	c := 0
	var h *ListNode
	for p != nil {
		c++
		if c == m {
			break
		}
		h = p
		p = p.Next
	}
	if c != m || m == n {
		return head
	}
	var pre *ListNode
	start := p
	p, pre = p.Next, p
	for p != nil {
		p, p.Next, pre = p.Next, pre, p
		c++
		if c == n {
			break
		}
	}
	if h != nil {
		h.Next = pre
	}
	start.Next = p
	if m == 1 {
		return pre
	}
	return head
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack = []*TreeNode{root}
	var rst []int
	var m = map[uintptr]bool{}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		stacked := false
		if top.Right != nil {
			p := uintptr(unsafe.Pointer(top.Right))
			if !m[p] {
				stack = append(stack, top.Right)
				stack = append(stack, top)
				m[p] = true
				stacked = true
			}
		}
		if top.Left != nil {
			p := uintptr(unsafe.Pointer(top.Left))
			if !m[p] {
				if !stacked {
					stack = append(stack, top)
					stacked = true
				}
				stack = append(stack, top.Left)
				m[p] = true
			}
		}
		if !stacked {
			rst = append(rst, top.Val)
		}
	}
	return rst
}

func numTrees(n int) int {
	rst := make([]int, n+1)
	rst[0] = 1
	rst[1] = 1
	for i := 2; i <= n; i++ {
		max := 0
		for j := 1; j <= i; j++ {
			max += rst[i-j] * rst[j-1]
		}
		rst[i] = max
	}
	return rst[n]
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) (ok bool) {
	pre := -1
	ok = true
	visit(root, func(v int) bool {
		defer func() {
			pre = v
		}()

		if pre == -1 {
			return true
		} else {
			ok = v > pre
			return v > pre
		}
	})
	return
}

func visit(root *TreeNode, visitor func(v int) bool) bool {
	if root == nil {
		return false
	}
	if root.Left != nil {
		if !visit(root.Left, visitor) {
			return false
		}
	}
	if !visitor(root.Val) {
		return false
	}
	if root.Right != nil {
		if !visit(root.Right, visitor) {
			return false
		}
	}
	return true
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
