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
