package lc

import "sort"

type task struct {
	t byte
	c int
}

type sortTask []task

func (st sortTask) Swap(i, j int) {
	st[i], st[j] = st[j], st[i]
}
func (st sortTask) Less(i, j int) bool {
	return st[i].c > st[j].c
}
func (st sortTask) Len() int {
	return len(st)
}

func leastInterval(tasks []byte, n int) int {
	count := map[byte]int{}
	for _, t := range tasks {
		count[t]++
	}
	sorted := make([]task, 0, len(count))
	for t, c := range count {
		sorted = append(sorted, task{t, c})
	}
	sort.Sort(sortTask(sorted))
	var deduct = func(pos int) {
		sorted[pos].c--
		for i := pos; i < len(sorted)-1; i++ {
			if sorted[i].c < sorted[i+1].c {
				sorted[i+1], sorted[i] = sorted[i], sorted[i+1]
			} else {
				break
			}
		}
		if sorted[len(sorted)-1].c == 0 {
			sorted = sorted[:len(sorted)-1]
		}
	}
	var jobs []byte
	for len(sorted) > 0 {
		success := false
		for i := 0; i < len(sorted); i++ {
			candidate := sorted[i].t
			ok := true
			for k := len(jobs) - 1; k >= 0 && k >= len(jobs)-n; k-- {
				if jobs[k] == candidate {
					ok = false
					break
				}
			}
			if ok {
				jobs = append(jobs, candidate)
				deduct(i)
				success = true
				break
			}
		}
		if !success {
			jobs = append(jobs, 0)
		}
	}
	return len(jobs)
}
