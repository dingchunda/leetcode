package lc

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	rst := make([][]bool, len(s))
	for i := range rst {
		rst[i] = make([]bool, len(s))
	}
	sum := 0
	for k := 0; k < len(s); k++ {
		for i := 0; i < len(s)-k; i++ {
			var ok bool
			if k == 0 {
				ok = true
			} else if k == 1 {
				ok = s[i] == s[i+1]
			} else {
				ok = rst[i+1][i+k-1] && s[i] == s[i+k]
			}
			if ok {
				rst[i][i+k] = true
				sum++
			}
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

func findTarget(root *TreeNode, k int) bool {
	f := &F{
		m: map[int]*TreeNode{},
	}
	return f.find(root, k)
}

type F struct {
	m map[int]*TreeNode
}

func (i *F) find(root *TreeNode, k int) bool {
	if i.m[k-root.Val] != nil && i.m[k-root.Val] != root {
		return true
	}
	i.m[root.Val] = root
	if root.Left != nil {
		if i.find(root.Left, k) {
			return true
		}
	}
	if root.Right != nil {
		if i.find(root.Right, k) {
			return true
		}
	}
	return false
}

func findClosestElements(arr []int, k int, x int) []int {
	firstAt := findFirst(arr, 0, len(arr)-1, x)
	//fmt.Println(firstAt)
	var rst = []int{arr[firstAt]}
	left, right := firstAt-1, firstAt+1
	for i := 1; i < k; i++ {
		if left < 0 {
			if right < len(arr) {
				rst = append(rst, arr[right])
				right++
			} else {
				return rst
			}
		} else if right == len(arr) {
			if left >= 0 {
				t := rst
				rst = []int{arr[left]}
				rst = append(rst, t...)
				left--
			} else {
				return rst
			}
		} else {
			if dist(arr[left], x) <= dist(arr[right], x) {
				t := rst
				rst = []int{arr[left]}
				rst = append(rst, t...)
				left--
			} else {
				rst = append(rst, arr[right])
				right++
			}
		}
	}
	return rst
}

func findFirst(arr []int, start, end int, x int) int {
	if start == end {
		return start
	}
	at := (start + end) / 2
	if arr[at] == x {
		return at
	} else if arr[at] > x {
		if at > 0 {
			if dist(arr[at-1], x) > dist(arr[at], x) {
				return at
			} else {
				return findFirst(arr, start, at-1, x)
			}
		} else {
			return at
		}
	} else {
		if at < len(arr)-1 {
			if arr[at] != arr[at+1] && dist(arr[at], x) <= dist(arr[at+1], x) {
				return at
			} else {
				return findFirst(arr, at+1, end, x)
			}
		} else {
			return at
		}
	}
}

func dist(i, j int) int {
	if j >= i {
		return j - i
	}
	return i - j
}
