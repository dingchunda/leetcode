package lc

import (
	"container/heap"
	"strings"
)

func numJewelsInStones(J string, S string) int {
	m := map[byte]int{}
	for _, s := range []byte(J) {
		m[s] = 1
	}
	sum := 0
	for _, s := range []byte(S) {
		if _, ok := m[s]; ok {
			sum++
		}
	}
	return sum
}

func canTransform(start string, end string) bool {
	if strings.ReplaceAll(start, "X", "") != strings.ReplaceAll(end, "X", "") {
		return false
	}
	j := 0
	for i := 0; i < len(start); i++ {
		if start[i] == 'R' {
			if j < i {
				j = i
			}
			for ; j < len(end) && end[j] != 'R'; j++ {
			}
			if j == len(end) {
				return false
			}
			j++
		}
	}

	j = 0
	for i := 0; i < len(end); i++ {
		if end[i] == 'L' {
			if j < i {
				j = i
			}
			for ; j < len(start) && start[j] != 'L'; j++ {
			}
			if j == len(start) {
				return false
			}
			j++
		}
	}
	return true
}

func slidingPuzzle(board [][]int) int {
	visited := map[[6]int]bool{}
	var buf, tmp [][6]int
	final := [6]int{1, 2, 3, 4, 5, 0}
	transfer := [6][4]int{
		{1, -2, 0, 0},
		{1, -1, -2, 0},
		{-1, -2, 0, 0},
		{1, 2, 0, 0},
		{-1, 1, 2, 0},
		{-1, 2, 0, 0},
	}
	init := [6]int{board[0][0], board[0][1], board[0][2], board[1][0], board[1][1], board[1][2]}
	ans := 0
	buf = append(buf, init)
	for len(buf) > 0 {
		tmp = tmp[:0]
		for _, cr := range buf {
			visited[cr] = true
			if cr == final {
				return ans
			}
			zeroAt := -1
			for at, v := range cr {
				if v == 0 {
					zeroAt = at
					break
				}
			}
			for _, tr := range transfer[zeroAt] {
				st := cr
				switch tr {
				case 1:
					st[zeroAt], st[zeroAt+1] = st[zeroAt+1], st[zeroAt]
					break
				case -1:
					st[zeroAt], st[zeroAt-1] = st[zeroAt-1], st[zeroAt]
					break
				case 2:
					st[zeroAt], st[zeroAt-3] = st[zeroAt-3], st[zeroAt]
					break
				case -2:
					st[zeroAt], st[zeroAt+3] = st[zeroAt+3], st[zeroAt]
					break
				default:
					continue
				}
				if !visited[st] {
					tmp = append(tmp, st)
				}
			}
		}
		ans++
		buf, tmp = tmp, buf
	}
	return -1
}

func reorganizeString(S string) string {
	m := map[byte]int{}
	for _, c := range S {
		m[byte(c)]++
	}
	lh := &lightHeapOrg{}
	heap.Init(lh)
	for k, v := range m {
		heap.Push(lh, pairOrg{k, v})
	}
	var buf []byte
	for lh.Len() > 0 {
		if len(buf) > 0 && buf[len(buf)-1] == lh.data[0].c {
			index := -1
			if lh.Len() > 1 {
				index = 1
			}
			if lh.Len() > 2 && lh.data[2].cnt > lh.data[1].cnt {
				index = 2
			}
			if index == -1 {
				return ""
			}
			buf = append(buf, lh.data[index].c)
			lh.data[index].cnt--
			if lh.data[index].cnt > 0 {
				heap.Fix(lh, index)
			} else {
				lh.data[index].cnt = 1e9
				heap.Fix(lh, index)
				heap.Pop(lh)
			}
		} else {
			buf = append(buf, lh.data[0].c)
			lh.data[0].cnt--
			if lh.data[0].cnt == 0 {
				heap.Pop(lh)
			} else {
				heap.Fix(lh, 0)
			}
		}
	}
	return string(buf)
}

type pairOrg struct {
	c   byte
	cnt int
}
type lightHeapOrg struct {
	data []pairOrg
}

func (l *lightHeapOrg) Push(i interface{}) {
	l.data = append(l.data, i.(pairOrg))
}
func (l *lightHeapOrg) Pop() interface{} {
	rst := l.data[len(l.data)-1]
	l.data = l.data[:len(l.data)-1]
	return rst
}
func (l *lightHeapOrg) Swap(i, j int) {
	l.data[i], l.data[j] = l.data[j], l.data[i]
}
func (l *lightHeapOrg) Less(i, j int) bool {
	return l.data[i].cnt > l.data[j].cnt
}
func (l *lightHeapOrg) Len() int {
	return len(l.data)
}
