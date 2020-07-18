package lc

import (
	"sort"
	"strconv"
)

func countOfAtoms(formula string) string {
	var stack []map[string]int
	stack = append(stack, map[string]int{})

	var atom []byte
	for i := 0; i < len(formula); {
		r := formula[i]
		if r == '(' {
			if len(atom) > 0 {
				stack[len(stack)-1][string(atom)]++
				atom = atom[:0]
			}
			stack = append(stack, map[string]int{})
			i++
			continue
		}
		if r == ')' {
			if len(atom) > 0 {
				stack[len(stack)-1][string(atom)]++
				atom = atom[:0]
			}
			number := 0
			j := i + 1
			for ; j < len(formula) && formula[j] >= '0' && formula[j] <= '9'; j++ {
				number = number*10 + int(formula[j]-'0')
			}
			mergeTo := stack[len(stack)-2]
			for k, v := range stack[len(stack)-1] {
				mergeTo[k] += v * number
			}
			stack = stack[:len(stack)-1]
			i = j
			continue
		}
		if r >= 'A' && r <= 'Z' {
			if len(atom) > 0 {
				stack[len(stack)-1][string(atom)]++
			}
			atom = atom[:0]
			atom = append(atom, r)
			i++
			continue
		}
		if r >= 'a' && r <= 'z' {
			atom = append(atom, r)
			i++
			continue
		}

		if r >= '1' && r <= '9' && len(atom) > 0 {
			number := int(r - '0')
			j := i + 1
			for ; j < len(formula) && formula[j] >= '0' && formula[j] <= '9'; j++ {
				number = number*10 + int(formula[j]-'0')
			}
			stack[len(stack)-1][string(atom)] += number
			atom = atom[:0]
			i = j
			continue
		}
	}
	if len(atom) > 0 {
		stack[0][string(atom)]++
	}
	var sorted []string
	for key := range stack[0] {
		sorted = append(sorted, key)
	}
	sort.Strings(sorted)
	var buffer []byte
	for _, s := range sorted {
		buffer = append(buffer, []byte(s)...)
		if stack[0][s] > 1 {
			buffer = append(buffer, strconv.Itoa(stack[0][s])...)
		}
	}
	return string(buffer)
}

func dailyTemperatures(T []int) []int {
	// [73, 74, 75, 71, 69, 72, 76, 73]
	rst := make([]int, len(T))
	for i := len(T) - 2; i >= 0; i-- {
		p := i + 1
		for p < len(T) {
			if T[p] > T[i] {
				break
			}
			if rst[p] > 0 {
				p += rst[p]
			} else {
				p++
			}
		}
		if p < len(T) {
			rst[i] = p - i
		}
	}
	return rst
}
