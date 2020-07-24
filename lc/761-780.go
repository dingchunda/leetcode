package lc

import "strings"

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
