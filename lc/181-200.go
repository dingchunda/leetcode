package lc

import "container/list"

func reverseWords2(s []byte) {
	var reverse = func(i, j int) {
		for i < j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
	reverse(0, len(s)-1)
	p := 0
	q := 0
	for p < len(s) && q < len(s) {
		for ; p < len(s) && s[p] == ' '; p++ {
		}
		q = p + 1
		for ; q < len(s) && s[q] != ' '; q++ {

		}
		reverse(p, q-1)
		p = q + 1
	}
	if p < len(s) && q == len(s) {
		reverse(p, q-1)
	}
}

func rotate2(nums []int, k int) {
	k = k % len(nums)
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	i, j = 0, k-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	i, j = k, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}
	rst := make([]int, n)
	rst[n-1] = nums[n-1]
	rst[n-2] = max(nums[n-1], nums[n-2])
	for i := n - 3; i >= 0; i-- {
		rst[i] = max(rst[i+1], nums[i]+rst[i+2])
	}
	return rst[0]
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var rst []int
	var buf, tmp []*TreeNode
	if root != nil {
		buf = append(buf, root)
	}
	for len(buf) > 0 {
		for _, b := range buf {
			if b.Right != nil {
				tmp = append(tmp, b.Right)
			}
			if b.Left != nil {
				tmp = append(tmp, b.Left)
			}
		}
		rst = append(rst, buf[0].Val)
		buf, tmp = tmp, buf
		tmp = tmp[:0]
	}
	return rst
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	tiles := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		tiles[i] = make([]bool, len(grid[0]))
	}
	curLand := 0

	type tile struct {
		i, j int
		land int
	}
	var high = list.New()
	high.PushBack(tile{0, 0, -1})
	for high.Len() > 0 {
		var next tile
		next = high.Front().Value.(tile)
		high.Remove(high.Front())
		i, j := next.i, next.j
		land := next.land
		if tiles[i][j] {
			continue
		}
		tiles[i][j] = true
		var nextLand int
		ty := grid[i][j]
		if ty == '1' {
			if land == -1 {
				curLand++
				nextLand = curLand
			} else {
				nextLand = land
			}
		} else {
			nextLand = -1
		}

		if i-1 >= 0 && !tiles[i-1][j] {
			if grid[i-1][j] == '1' {
				high.PushFront(tile{i - 1, j, nextLand})
			} else {
				high.PushBack(tile{i - 1, j, nextLand})
			}
		}
		if i+1 < len(grid) && !tiles[i+1][j] {
			if grid[i+1][j] == '1' {
				high.PushFront(tile{i + 1, j, nextLand})
			} else {
				high.PushBack(tile{i + 1, j, nextLand})
			}
		}
		if j-1 >= 0 && !tiles[i][j-1] {
			if grid[i][j-1] == '1' {
				high.PushFront(tile{i, j - 1, nextLand})
			} else {
				high.PushBack(tile{i, j - 1, nextLand})
			}
		}
		if j+1 < len(grid[0]) && !tiles[i][j+1] {
			if grid[i][j+1] == '1' {
				high.PushFront(tile{i, j + 1, nextLand})
			} else {
				high.PushBack(tile{i, j + 1, nextLand})
			}
		}
	}
	return curLand
}
