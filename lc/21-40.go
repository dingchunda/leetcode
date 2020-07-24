package lc

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var pre *ListNode
	var head *ListNode

	for l1 != nil || l2 != nil {
		node := &ListNode{}
		if l2 == nil || l1 != nil && l1.Val <= l2.Val {
			node.Val = l1.Val
			l1 = l1.Next
		} else {
			node.Val = l2.Val
			l2 = l2.Next
		}
		if pre != nil {
			pre.Next = node
		}
		if head == nil {
			head = node
		}
		pre = node
	}
	return head
}

func generateParenthesis(n int) []string {
	rst := make([][]string, n+1)
	rst[0] = []string{""}
	rst[1] = []string{"()"}

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			for _, v1 := range rst[j-1] {
				for _, v2 := range rst[i-j] {
					rst[i] = append(rst[i], fmt.Sprintf("(%s)%s", v1, v2))
				}
			}
		}
	}
	return rst[n]
}

func mergeKLists(lists []*ListNode) *ListNode {
	np := &priorityQueue{}
	heap.Init(np)
	for _, list := range lists {
		if list != nil {
			heap.Push(np, list)
		}
	}
	var head, pre *ListNode
	for np.Len() > 0 {
		top := heap.Pop(np).(*ListNode)
		if pre != nil {
			pre.Next = top
		}
		if head == nil {
			head = top
		}
		pre = top
		if top.Next != nil {
			heap.Push(np, top.Next)
		}
	}
	return head
}

type priorityQueue []*ListNode

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	no := x.(*ListNode)
	*pq = append(*pq, no)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	no := old[n-1]
	*pq = old[0 : n-1]
	return no
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	rst := head.Next
	var pre *ListNode
	p := head
	for p != nil && p.Next != nil {
		if pre != nil {
			pre.Next, p.Next, p.Next.Next, pre = p.Next, p.Next.Next, p, p
		} else {
			p.Next, p.Next.Next, pre = p.Next.Next, p, p
		}
		p = p.Next
	}
	return rst
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	p := head
	var h *ListNode
	var pre *ListNode
	for {
		ori := p
		cnt := 0
		for ; p != nil && cnt < k; cnt++ {
			if cnt == k-1 {
				p.Next, p = nil, p.Next
			} else {
				p = p.Next
			}
		}
		if cnt != k {
			break
		}
		last := reverseList(ori)
		if pre != nil {
			pre.Next = last
		}
		if h == nil {
			h = last
		}
		pre = ori
		pre.Next = p
	}
	return h
}
func reverseList(head *ListNode) *ListNode {
	p := head.Next
	pre := head
	pre.Next = nil
	for p != nil {
		pre, p, p.Next = p, p.Next, pre
	}
	return pre
}

func removeDuplicates(nums []int) int {
	var rst = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[rst] {
			nums[rst+1] = nums[i]
			rst++
		}
	}
	return rst + 1
}

func removeElement(nums []int, val int) int {
	var p = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[p] = nums[i]
			p++
		}
	}
	return p
}

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	var next func(data []int) bool
	next = func(data []int) bool {
		index := -1
		v := 999999
		for i := 1; i < len(data); i++ {
			if data[i] > data[0] && data[i] < v {
				index = i
			}
		}
		if index != -1 {
			data[0], data[index] = data[index], data[0]
			sort.Ints(data[1:])
			return true
		}
		return false
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if next(nums[i:]) {
			return
		}
	}
	sort.Ints(nums)
	return
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if nums[0] == target {
		return 0
	}
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}
	return binsearch(0, len(nums)-1, nums, target)
}

func binsearch(s, e int, nums []int, target int) int {
	if s > e {
		return -1
	}
	middle := (s + e) / 2
	v := nums[middle]
	if v == target {
		return middle
	}
	if v <= nums[s] && v <= nums[e] {
		if target > v && target <= nums[e] {
			return binsearch(middle+1, e, nums, target)
		}
		return binsearch(s, middle-1, nums, target)
	} else if v >= nums[s] && v >= nums[e] {
		if target >= nums[s] && target < v {
			return binsearch(s, middle-1, nums, target)
		}
		return binsearch(middle+1, e, nums, target)
	} else {
		if target < v {
			return binsearch(s, middle-1, nums, target)
		} else {
			return binsearch(middle+1, e, nums, target)
		}
	}
}

func searchRange(nums []int, target int) []int {
	start := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target+1
	})
	return []int{start, end - 1}
}

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	subs := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i] = map[byte]bool{}
		cols[i] = map[byte]bool{}
		subs[i] = map[byte]bool{}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c == '.' {
				continue
			}
			c = c - '0'
			if rows[i][c] || cols[j][c] || subs[i/3*3+j/3][c] {
				return false
			}
			rows[i][c] = true
			cols[j][c] = true
			subs[i/3*3+j/3][c] = true

		}
	}
	return true
}

func combinationSum(candidates []int, target int) [][]int {
	rst := make([][][]int, target+1)
	buffer := make([][][]int, target+1)
	for _, c := range candidates {
		for i := 0; i < len(buffer); i++ {
			buffer[i] = buffer[i][:0]
		}

		for i := 1; i*c <= target; i++ {
			var v []int
			for k := 0; k < i; k++ {
				v = append(v, c)
			}
			buffer[i*c] = append(buffer[i*c], v)
			for k := i*c + 1; k <= target; k++ {
				for _, aa := range rst[k-i*c] {
					v2 := make([]int, len(v))
					copy(v2, v)
					v2 = append(v2, aa...)
					buffer[k] = append(buffer[k], v2)
				}
			}
		}
		for i := 0; i < len(buffer); i++ {
			rst[i] = append(rst[i], buffer[i]...)
		}
	}
	return rst[target]
}

func longestValidParentheses(s string) int {
	var stack []string
	ans := 0
	for _, c := range s {
		if c == '(' {
			stack = append(stack, string(c))
		} else {
			if len(stack) == 0 {
				continue
			}
			if stack[len(stack)-1] == "(" {
				stack[len(stack)-1] = "2"
				sum := 2
				if len(stack) >= 2 && stack[len(stack)-2] != "(" {
					num, _ := strconv.ParseInt(stack[len(stack)-2], 10, 32)
					sum += int(num)
					stack[len(stack)-2] = strconv.Itoa(sum)
					stack = stack[:len(stack)-1]
				}
				if sum > ans {
					ans = sum
				}
			} else {
				// a number
				if len(stack) == 1 {
					stack = stack[:0]
				} else {
					//  (,num, ( ,num
					//  (  , (  ,num
					num, _ := strconv.ParseInt(stack[len(stack)-1], 10, 32)
					sum := int(num) + 2
					if len(stack) >= 3 && stack[len(stack)-3] != "(" {
						num2, _ := strconv.ParseInt(stack[len(stack)-3], 10, 32)
						sum += int(num2)
						stack[len(stack)-3] = strconv.Itoa(sum)
						stack = stack[:len(stack)-2]
					} else {
						stack[len(stack)-2] = strconv.Itoa(sum)
						stack = stack[:len(stack)-1]
					}
					if sum > ans {
						ans = sum
					}
				}
			}
		}
	}
	return ans
}
