package lc

import "sort"

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var buf, tmp []*TreeNode
	buf = append(buf, root)
	lvl := 0
	lastbroken := false
	for len(buf) > 0 {
		if len(tmp) > 0 && len(tmp) != (1<<uint(lvl-1)) {
			return false
		}
		tmp = tmp[:0]
		end := false
		broken := false
		for _, n := range buf {
			if n.Left != nil {
				if end {
					broken = true
				}
				tmp = append(tmp, n.Left)
			} else {
				end = true
			}
			if n.Right != nil {
				if end {
					broken = true
				}
				tmp = append(tmp, n.Right)
			} else {
				end = true
			}
		}
		//fmt.Println(lvl,len(tmp),lastbroken)
		if len(tmp) == 0 && lastbroken {
			return false
		}
		lastbroken = broken
		buf, tmp = tmp, buf
		lvl++
	}
	return true
}

func canReorderDoubled(A []int) bool {
	var check = func(data []int, forward bool) bool {
		if len(data)&1 == 1 {
			return false
		}
		visited := make([]bool, len(data))
		for i := 0; i < len(data); i++ {
			if visited[i] {
				continue
			}
			if !forward && data[i]%2 != 0 {
				return false
			}
			var pair int
			if forward {
				pair = data[i] * 2
			} else {
				pair = data[i] / 2
			}
			index := sort.Search(len(data), func(i int) bool {
				return data[i] >= pair
			})
			for ; index < len(data) && data[index] == pair && visited[index]; index++ {
			}
			if index < len(data) && data[index] == pair {
				visited[index] = true
			} else {
				return false
			}
		}
		return true
	}

	sort.Ints(A)
	index1 := sort.Search(len(A), func(i int) bool {
		return A[i] >= 0
	})
	if !check(A[:index1], false) {
		return false
	}
	A = A[index1:]

	index2 := sort.Search(len(A), func(i int) bool {
		return A[i] > 0
	})
	if index2 < len(A) && A[index2] > 0 {
		if !check(A[index2:], true) {
			return false
		}
		A = A[:index2]
	}
	return len(A)&1 == 0
}
