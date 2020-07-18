package lc

import (
	"container/heap"
	"container/list"
	"math"
	"strconv"
	"unsafe"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	slow, fast := head, head
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	m := map[uintptr]bool{}
	for head != nil {
		p := uintptr(unsafe.Pointer(head))
		if m[p] {
			return head
		}
		m[p] = true
		head = head.Next
	}
	return nil
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if slow.Next == nil {
		return
	}
	tar := head
	p := reverseList(slow.Next)
	slow.Next = nil
	q := head.Next
	isSlow := true
	for p != nil || q != nil {
		if isSlow {
			tar.Next = p
			p = p.Next
			tar = tar.Next
		} else if q != nil {
			tar.Next = q
			q = q.Next
			tar = tar.Next
		}
		isSlow = !isSlow
	}
	print(head)
}

func preorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	if root != nil {
		stack = append(stack, root)
	}
	var rst []int
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rst = append(rst, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return rst
}

type LRUCache struct {
	l *list.List
	m map[int]*list.Element
	c int
}

type kv struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		l: list.New(),
		m: map[int]*list.Element{},
		c: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.m[key]
	if !ok {
		return -1
	}
	this.l.MoveToFront(e)
	return e.Value.(kv).value
}

func (this *LRUCache) Put(key int, value int) {
	e, ok := this.m[key]
	if ok {
		e.Value = kv{key, value}
		this.l.MoveToFront(e)
	} else {
		this.l.PushFront(kv{key, value})
		this.m[key] = this.l.Front()
	}
	if this.l.Len() > this.c {
		e := this.l.Back()
		delete(this.m, e.Value.(kv).key)
		this.l.Remove(e)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right := split(head)
	return mergeList(sortList(left), sortList(right))
}

func split(head *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}
	pre.Next = nil
	return head, slow
}

func evalRPN(tokens []string) int {
	var stack []int
	for i := 0; i < len(tokens); i++ {
		c := tokens[i]
		if c == "+" || c == "-" || c == "*" || c == "/" {
			if len(stack) < 2 {
				return 0
			}
			left, right := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			var v int
			if c == "+" {
				v = left + right
			} else if c == "-" {
				v = left - right
			} else if c == "*" {
				v = left * right
			} else {
				v = left / right
			}
			stack = append(stack, v)
			continue
		}
		v, _ := strconv.ParseInt(c, 10, 31)
		stack = append(stack, int(v))
	}
	return stack[0]
}

func mergeList(fast, slow *ListNode) *ListNode {
	var head *ListNode
	if fast.Val < slow.Val {
		head = fast
		fast = fast.Next
	} else {
		head = slow
		slow = slow.Next
	}

	p := head
	for fast != nil && slow != nil {
		if fast.Val < slow.Val {
			p.Next = fast
			fast = fast.Next
		} else {
			p.Next = slow
			slow = slow.Next
		}
		p = p.Next
	}
	if fast != nil {
		p.Next = fast
	} else if slow != nil {
		p.Next = slow
	}
	return head
}

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	data := []byte(s)
	var reverse = func(i, j int) {
		for i < j {
			data[i], data[j] = data[j], data[i]
			i++
			j--
		}
	}
	reverse(0, len(s)-1)

	for i := 0; i < len(s); i++ {
		if data[i] == ' ' {
			continue
		}
		var j int
		for j = i; j < len(s); j++ {
			if data[j] == ' ' {
				break
			}
		}
		j--
		reverse(i, j)
		i = j
	}

	var trim = func(i, j, k int) {
		for l := j; l < k; l++ {
			data[i+l-j] = data[l]
		}
	}
	blank := 0
	tail := 0
	from := blank
	for blank < len(s) {
		var start int
		for start = from; start < len(s); start++ {
			if data[start] != ' ' {
				break
			}
		}
		var end int
		for end = start; end < len(s); end++ {
			if data[end] == ' ' {
				break
			}
		}
		if start == end {
			break
		}
		if start != blank {
			trim(blank, start, end)
		}
		tail = blank + end - start
		if tail < len(s) {
			data[tail] = ' '
		}
		blank = tail + 1
		if end >= len(s) {
			break
		}
		from = end + 1
	}
	return string(data[0:tail])
}

func maxProduct(nums []int) int {
	p := 0
	max := -math.MaxInt32
	for p < len(nums) {
		var at int
		for at = p; at < len(nums); at++ {
			if nums[at] != 0 {
				break
			}
			if max < 0 {
				max = 0
			}
		}
		firstNonZero := at
		if firstNonZero >= len(nums) {
			break
		}
		product := 1
		firstNeg, lastNeg := -1, -1
		for ; at < len(nums); at++ {
			if nums[at] == 0 {
				if max < 0 {
					max = 0
				}
				break
			}
			product *= nums[at]
			if nums[at] < 0 {
				if nums[at] > max {
					max = nums[at]
				}
				if firstNeg == -1 {
					firstNeg = at
				}
				lastNeg = at
			}
		}
		// current at is the first zero
		if product > 0 {
			if product > max {
				max = product
			}
			p = at + 1
			continue
		}

		if at-firstNeg > 1 {
			tp := product
			for i := firstNonZero; i <= firstNeg; i++ {
				tp /= nums[i]
			}
			if tp > max {
				max = tp
			}
		}

		if lastNeg-firstNonZero > 0 {
			tp := product
			for i := lastNeg; i < at; i++ {
				tp /= nums[i]
			}
			if tp > max {
				max = tp
			}
		}

		p = at + 1
	}
	return max
}

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	mid := len(nums) / 2
	if mid > 0 && nums[mid] < nums[mid-1] {
		return nums[mid]
	}
	if mid == 0 {
		return nums[mid]
	}
	if nums[mid] > nums[0] && nums[0] > nums[len(nums)-1] {
		return findMin(nums[mid+1:])
	}
	return findMin(nums[:mid])
}

type MinStack struct {
	stack []int
	data  map[int]int
	hp    *stackItems
}

/** initialize your data structure here. */
func ConstructorStak() MinStack {
	a := &stackItems{}
	heap.Init(a)
	return MinStack{
		data: map[int]int{},
		hp:   a,
	}
}

func (this *MinStack) Push(x int) {
	this.data[x]++
	this.stack = append(this.stack, x)
	heap.Push(this.hp, x)
}

func (this *MinStack) Pop() {
	top := this.stack[len(this.stack)-1]
	this.data[top]--
	if this.data[top] == 0 {
		delete(this.data, top)
	}
	this.stack = this.stack[:len(this.stack)-1]
	_ = this.GetMin()
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	for {
		t, ok := this.hp.Top()
		if ok && this.data[t] == 0 {
			heap.Pop(this.hp)
		} else {
			break
		}
	}
	v, _ := this.hp.Top()
	return v
}

type stackItems []int

func (it stackItems) Len() int {
	return len(it)
}
func (it stackItems) Less(i, j int) bool {
	return it[i] < it[j]
}
func (it stackItems) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
func (it *stackItems) Push(v interface{}) {
	*it = append(*it, v.(int))
}
func (it *stackItems) Pop() interface{} {
	data := (*it)[len(*it)-1]
	*it = (*it)[:len(*it)-1]
	return data
}
func (it *stackItems) Top() (int, bool) {
	h := []int(*it)
	if len(h) == 0 {
		return -1, false
	}
	return h[0], true
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p := headA
	for ; p.Next != nil; p = p.Next {
	}
	p.Next = headB
	if headA.Next == nil || headA.Next.Next == nil {
		p.Next = nil
		return nil
	}
	slow := headA.Next
	fast := headA.Next.Next
	for slow != fast && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow != fast {
		p.Next = nil
		return nil
	}
	slow = headA
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	p.Next = nil
	return slow
}
