package lc

import (
	"container/heap"
	"container/list"
	"sort"
)

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	tiles := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		tiles[i] = make([]bool, len(grid[0]))
	}
	type tile struct {
		i, j int
		land int
	}
	var high = list.New()
	high.PushBack(tile{0, 0, -1})

	area := 0
	maxArea := 0
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
		ty := grid[i][j]
		if ((ty == 0 || ty == 1) && land != 1) && area > 0 {
			if area > maxArea {
				maxArea = area
			}
			area = 0
		}
		if ty == 1 && land != 1 {
			area = 1
		}
		if ty == 1 && land == 1 {
			area++
		}
		if i-1 >= 0 && !tiles[i-1][j] {
			if grid[i-1][j] == 1 {
				high.PushFront(tile{i - 1, j, grid[i][j]})
			} else {
				high.PushBack(tile{i - 1, j, grid[i][j]})
			}
		}
		if i+1 < len(grid) && !tiles[i+1][j] {
			if grid[i+1][j] == 1 {
				high.PushFront(tile{i + 1, j, grid[i][j]})
			} else {
				high.PushBack(tile{i + 1, j, grid[i][j]})
			}
		}
		if j-1 >= 0 && !tiles[i][j-1] {
			if grid[i][j-1] == 1 {
				high.PushFront(tile{i, j - 1, grid[i][j]})
			} else {
				high.PushBack(tile{i, j - 1, grid[i][j]})
			}
		}
		if j+1 < len(grid[0]) && !tiles[i][j+1] {
			if grid[i][j+1] == 1 {
				high.PushFront(tile{i, j + 1, grid[i][j]})
			} else {
				high.PushBack(tile{i, j + 1, grid[i][j]})
			}
		}
	}
	if area > 0 && area > maxArea {
		maxArea = area
	}
	return maxArea
}

func searchBST(root *TreeNode, val int) *TreeNode {
	var travel func(n *TreeNode, visitor func(node *TreeNode) bool) bool
	travel = func(n *TreeNode, visitor func(node *TreeNode) bool) bool {
		if n == nil {
			return false
		}
		if travel(n.Left, visitor) {
			return true
		}
		if visitor(n) {
			return true
		}
		return travel(n.Right, visitor)
	}
	var ans *TreeNode
	travel(root, func(node *TreeNode) bool {
		if node.Val == val {
			ans = node
			return true
		}
		return false
	})
	return ans
}

func topKFrequentWord(words []string, k int) []string {
	table := map[string]int{}
	for _, w := range words {
		table[w]++
	}
	lp := &lightHeapWord{}
	heap.Init(lp)
	for w, fre := range table {
		if lp.Len() == k {
			if lp.less(pariWord{w, fre}, lp.data[0]) {
				continue
			}
			heap.Pop(lp)
		}
		heap.Push(lp, pariWord{w, fre})
	}
	sort.Sort(lp)
	ans := make([]string, 0, k)
	for i := len(lp.data) - 1; i >= 0; i-- {
		ans = append(ans, lp.data[i].word)
	}
	return ans
}

type pariWord struct {
	word string
	fre  int
}

type lightHeapWord struct {
	data []pariWord
}

func (l *lightHeapWord) Push(i interface{}) {
	l.data = append(l.data, i.(pariWord))
}

func (l *lightHeapWord) Pop() interface{} {
	rst := l.data[len(l.data)-1]
	l.data = l.data[:len(l.data)-1]
	return rst
}

func (l *lightHeapWord) Swap(i, j int) {
	l.data[i], l.data[j] = l.data[j], l.data[i]
}
func (l *lightHeapWord) less(p1, p2 pariWord) bool {
	if p1.fre == p2.fre {
		return p1.word > p2.word
	}
	return p1.fre < p2.fre
}

func (l *lightHeapWord) Less(i, j int) bool {
	return l.less(l.data[i], l.data[j])
}
func (l *lightHeapWord) Len() int {
	return len(l.data)
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var dfs func(r *TreeNode) int
	dfs = func(r *TreeNode) int {
		if r.Left == nil && r.Right == nil {
			return 1
		}
		rst := 1
		sum := 1
		for _, n := range []*TreeNode{r.Left, r.Right} {
			if n != nil {
				v := dfs(n)
				if n.Val == r.Val {
					sum += v
				}
				if n.Val == r.Val && v+1 > rst {
					rst = v + 1
				}
			}
		}
		if sum > ans {
			ans = sum
		}
		return rst
	}
	dfs(root)
	if ans > 1 {
		return ans - 1
	}
	return 0
}
