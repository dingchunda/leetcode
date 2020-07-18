package lc

import (
	"container/heap"
	"strconv"
	"strings"
)

func moveZeroes(nums []int) {
	p := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && p == -1 {
			p = i
		}
		if nums[i] != 0 && p >= 0 {
			nums[i], nums[p] = nums[p], nums[i]
			for ; nums[p] != 0; p++ {
			}
		}
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var travel func(r *TreeNode, visitor func(m *TreeNode) bool) bool
	travel = func(r *TreeNode, visitor func(m *TreeNode) bool) bool {
		if r == nil {
			return false
		}
		if travel(r.Left, visitor) {
			return true
		}
		if visitor(r) {
			return true
		}
		return travel(r.Right, visitor)
	}

	var pre *TreeNode
	var rst *TreeNode
	travel(root, func(m *TreeNode) bool {
		if pre == p {
			rst = m
			return true
		}
		pre = m
		return false
	})
	return rst
}

func wallsAndGates(rooms [][]int) {
	if len(rooms) == 0 || len(rooms[0]) == 0 {
		return
	}
	xc := len(rooms)
	yc := len(rooms[0])
	var bfs = func(i, j int) {
		visit := map[int]bool{}
		var buf, tmp []int
		id := i*yc + j
		buf = append(buf, id)
		dis := 0

		for len(buf) > 0 {
			for _, parent := range buf {
				x, y := parent/yc, parent%yc
				if parent < 0 || parent >= xc*yc || visit[parent] || rooms[x][y] == -1 || rooms[x][y] == 0 && dis > 0 {
					continue
				}
				if dis < rooms[x][y] {
					rooms[x][y] = dis
				}
				visit[parent] = true
				if parent%yc != 0 {
					tmp = append(tmp, parent-1)
				}
				if parent%yc != yc-1 {
					tmp = append(tmp, parent+1)
				}
				tmp = append(tmp, parent+yc, parent-yc)
			}
			dis++
			buf, tmp = tmp, buf
			tmp = tmp[:0]
		}
	}

	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == 0 {
				bfs(i, j)
			}
		}
	}
}

func findDuplicate(nums []int) int {
	fast := nums[0]
	slow := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func gameOfLife(board [][]int) {
	if len(board) == 0 {
		return
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			count := 0
			if i-1 >= 0 && (board[i-1][j] == 1 || board[i-1][j] == 3) {
				count++
			}
			if i+1 < len(board) && (board[i+1][j] == 1 || board[i+1][j] == 3) {
				count++
			}
			if j-1 >= 0 && (board[i][j-1] == 1 || board[i][j-1] == 3) {
				count++
			}
			if j+1 < len(board[0]) && (board[i][j+1] == 1 || board[i][j+1] == 3) {
				count++
			}
			if i-1 >= 0 && j-1 >= 0 && (board[i-1][j-1] == 1 || board[i-1][j-1] == 3) {
				count++
			}
			if i-1 >= 0 && j+1 < len(board[0]) && (board[i-1][j+1] == 1 || board[i-1][j+1] == 3) {
				count++
			}
			if i+1 < len(board) && j-1 >= 0 && (board[i+1][j-1] == 1 || board[i+1][j-1] == 3) {
				count++
			}
			if i+1 < len(board) && j+1 < len(board[0]) && (board[i+1][j+1] == 1 || board[i+1][j+1] == 3) {
				count++
			}
			if count == 3 && board[i][j] == 0 {
				board[i][j] = 2
			} else if count < 2 && board[i][j] == 1 {
				board[i][j] = 3
			} else if count > 3 && board[i][j] == 1 {
				board[i][j] = 3
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			}
			if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}

type MedianFinder struct {
	loh *hs
	hih *hs
}

/** initialize your data structure here. */
func ConstructorMedianFinder() MedianFinder {
	lo := &hs{}
	hi := &hs{}
	heap.Init(lo)
	heap.Init(hi)
	return MedianFinder{
		loh: lo,
		hih: hi,
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.loh, num)
	heap.Push(this.hih, -heap.Pop(this.loh).(int))
	if this.hih.Len() > this.loh.Len() {
		heap.Push(this.loh, -heap.Pop(this.hih).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.loh.Len() == this.hih.Len() {
		return float64((*this.loh)[0]-(*this.hih)[0]) / 2
	}
	return float64((*this.loh)[0])
}

type hs []int

func (h *hs) Len() int {
	return len(*h)
}
func (h *hs) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *hs) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hs) Push(i interface{}) {
	*h = append(*h, i.(int))
}
func (h *hs) Pop() (r interface{}) {
	r, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return
}

type Codec struct {
}

func ConstructorCodec() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	if root == nil {
		return ""
	}

	queue := []*TreeNode{root}
	c := []string{strconv.Itoa(root.Val)}

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			add(&c, queue[i].Left)
			add(&c, queue[i].Right)
		}
		queue = queue[l:]
	}

	res := strings.Join(c, ",")
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	c := strings.Split(data, ",")

	if len(data) == 0 {
		return nil
	}

	t := &TreeNode{Val: myAtoi(c[0])}
	queue := []*TreeNode{t}

	i := 1
	for len(queue) > 0 {
		l := len(queue)
		for j := 0; j < l; j++ {
			if c[i] == "nil" {
				queue[j].Left = nil
			} else {
				queue[j].Left = &TreeNode{Val: myAtoi2(c[i])}
				queue = append(queue, queue[j].Left)
			}
			i++
			if c[i] == "nil" {
				queue[j].Right = nil
			} else {
				queue[j].Right = &TreeNode{Val: myAtoi2(c[i])}
				queue = append(queue, queue[j].Right)
			}
			i++
		}
		queue = queue[l:]
	}
	return t
}

func add(c *[]string, node *TreeNode) {
	if node == nil {
		*c = append(*c, "nil")
	} else {
		*c = append(*c, strconv.Itoa(node.Val))
	}
}

func myAtoi2(s string) int {
	a, err := strconv.Atoi(s)
	if err != nil {
		panic(a)
	}
	return a
}

func lengthOfLIS(nums []int) int {
	endWith := map[int]int{}

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		longer := endWith[n]
		longer = max(longer, 1)
		for k, v2 := range endWith {
			if n > k && v2+1 > longer {
				longer = v2 + 1
			}
		}
		endWith[n] = longer
	}
	maxv := 0
	for _, v := range endWith {
		if v > maxv {
			maxv = v
		}
	}
	return maxv
}
