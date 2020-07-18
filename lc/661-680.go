package lc

import "sort"

func maximumSwap(num int) int {
	var its []item
	n := num
	index := 0
	for n > 0 {
		its = append(its, item{
			v: n % 10,
		})
		index++
		n /= 10
	}
	old := make([]item, len(its))
	copy(old, its)
	sort.Sort(items(its))
	i := len(its) - 1
	for ; i >= 0 && its[i].v == old[i].v; i-- {
	}
	if i >= 0 {
		j := 0
		for ; j < len(old) && old[j].v != its[i].v; j++ {
		}
		old[j], old[i] = old[i], old[j]
	}
	rst := 0
	for i := len(old) - 1; i >= 0; i-- {
		rst = rst*10 + old[i].v
	}
	return rst
}

type items []item
type item struct {
	v int
	i int
}

func (it items) Len() int {
	return len(it)
}
func (it items) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
func (it items) Less(i, j int) bool {
	if it[i].v != it[j].v {
		return it[i].v < it[j].v
	}
	return it[i].i > it[j].i
}

func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			if i+1 == j {
				return true
			}
			return check(s[i+1:j+1]) || check(s[i:j])
		}
	}
	return true
}

func check(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
