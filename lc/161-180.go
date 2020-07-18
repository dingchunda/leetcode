package lc

import (
	"sort"
	"strconv"
	"strings"
)

func findPeakElement(nums []int) int {
	var f func(start, end int) int
	f = func(start, end int) int {
		mid := (start + end) / 2
		if mid > 0 && nums[mid] < nums[mid-1] {
			return f(start, mid)
		}
		if mid+1 < len(nums) && nums[mid] < nums[mid+1] {
			return f(mid+1, end)
		}
		return mid
	}
	return f(0, len(nums))
}

func compareVersion(version1 string, version2 string) int {
	var v1 []int
	for _, s := range strings.Split(version1, ".") {
		a, _ := strconv.ParseInt(s, 10, 32)
		v1 = append(v1, int(a))
	}
	var v2 []int
	for _, s := range strings.Split(version2, ".") {
		a, _ := strconv.ParseInt(s, 10, 32)
		v2 = append(v2, int(a))
	}
	var i int
	for i = 0; i < len(v1) && i < len(v2); i++ {
		if v1[i] > v2[i] {
			return 1
		} else if v1[i] < v2[i] {
			return -1
		}
	}
	if i == len(v1) && i < len(v2) {
		for _, v := range v2[i:] {
			if v > 0 {
				return -1
			}
		}
	} else if i == len(v2) && i < len(v1) {
		for _, v := range v1[i:] {
			if v > 0 {
				return 1
			}
		}
	}
	return 0
}

func fractionToDecimal(numerator int, denominator int) string {
	neg := false
	if numerator < 0 && denominator > 0 || numerator > 0 && denominator < 0 {
		neg = true
	}
	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}
	buffer := make([]byte, 0, 16)
	if neg {
		buffer = append(buffer, '-')
	}

	a := numerator / denominator
	buffer = append(buffer, []byte(strconv.Itoa(a))...)
	b := numerator % denominator
	if b == 0 {
		return string(buffer)
	}
	buffer = append(buffer, '.')
	m := map[int]int{}
	for {
		b *= 10
		at, ok := m[b]
		if ok {
			buffer = append(buffer, ' ', ')')
			for i := len(buffer) - 3; i >= at; i-- {
				buffer[i+1] = buffer[i]
			}
			buffer[at] = '('
			break
		}
		m[b] = len(buffer)
		buffer = append(buffer, byte(b/denominator)+'0')
		c := b % denominator
		if c == 0 {
			break
		}
		b = c
	}
	return string(buffer)
}

func twoSum2(numbers []int, target int) []int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		if v := numbers[low] + numbers[high]; v == target {
			return []int{low + 1, high + 1}
		} else if v < target {
			low++
		} else {
			high--
		}
	}
	return nil
}

func majorityElement(nums []int) int {
	count := 0
	candidate := 0
	for _, n := range nums {
		if count == 0 {
			candidate = n
		}
		if n == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack    []*TreeNode
	inserted map[*TreeNode]bool
}

func ConstructorBST(root *TreeNode) BSTIterator {
	var stack []*TreeNode
	if root != nil {
		stack = append(stack, root)
	}
	return BSTIterator{
		stack: stack,
		inserted: map[*TreeNode]bool{
			root: true,
		},
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	for {
		top := this.stack[len(this.stack)-1]
		insert := false
		if top.Right != nil && !this.inserted[top.Right] {
			insert = true
			this.inserted[top.Right] = true
			this.stack[len(this.stack)-1] = top.Right
			this.stack = append(this.stack, top)
		}
		if top.Left != nil && !this.inserted[top.Left] {
			this.stack = append(this.stack, top.Left)
			insert = true
			this.inserted[top.Left] = true
		}
		if !insert {
			this.stack = this.stack[:len(this.stack)-1]
			return top.Val
		}
	}
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

type numbers [][]byte

func (n numbers) Len() int {
	return len(n)
}
func less(a, b []byte) bool {
	var p int
	for p = 0; p < len(a) && p < len(b); p++ {
		if a[p] != b[p] {
			return a[p] > b[p]
		}
	}
	if p == len(a) && p == len(b) {
		return true
	}
	if p == len(a) {
		return less(a, b[p:])
	} else {
		return less(a[p:], b)
	}
}
func (n numbers) Less(i, j int) bool {
	return less(n[i], n[j])
}
func (n numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func largestNumber(nums []int) string {
	strs := make([][]byte, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = []byte(strconv.Itoa(nums[i]))
	}
	sort.Sort(numbers(strs))
	for i := 1; i < len(strs); i++ {
		strs[0] = append(strs[0], strs[i]...)
	}
	var i int
	for i = 0; i < len(strs[0]) && strs[0][i] == '0'; i++ {
	}
	rst := strs[0][i:]
	if len(rst) == 0 {
		return "0"
	}
	return string(rst)
}
