package lc

import "sort"

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

type NestedIterator struct {
	data []int
	ptr  int
}

func ConstructorNestedIterator(nestedList []*NestedInteger) *NestedIterator {
	var d []int
	return &NestedIterator{
		data: unpack(nestedList, d),
		ptr:  0,
	}
}

func unpack(nestedList []*NestedInteger, data []int) []int {
	for _, n := range nestedList {

		if n.IsInteger() {
			data = append(data, n.GetInteger())
		} else {
			data = unpack(n.GetList(), data)
		}
	}
	return data
}

func (this *NestedIterator) Next() int {
	v := this.data[this.ptr]
	this.ptr++
	return v
}

func (this *NestedIterator) HasNext() bool {
	return this.ptr < len(this.data)
}

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}

	type item struct {
		k, v int
	}
	var heap []item
	var up func(int)
	up = func(at int) {
		if at == 0 {
			return
		}
		var pre int
		if at%2 == 0 {
			pre = at/2 - 1
		} else {
			pre = at / 2
		}
		if heap[pre].v > heap[at].v {
			heap[pre], heap[at] = heap[at], heap[pre]
			up(pre)
		}
	}
	var down func(int)
	down = func(at int) {
		left, right := at*2+1, at*2+2
		swapLeft := false
		swapRight := false
		if left < len(heap) && heap[left].v < heap[at].v {
			swapLeft = true
		}
		if right < len(heap) {
			if swapLeft {
				if heap[right].v < heap[left].v {
					swapRight = true
					swapLeft = false
				}
			} else {
				if heap[right].v < heap[at].v {
					swapRight = true
				}
			}
		}
		if swapLeft {
			heap[left], heap[at] = heap[at], heap[left]
			down(left)
			return
		}
		if swapRight {
			heap[right], heap[at] = heap[at], heap[right]
			down(right)
			return
		}
		return
	}
	for key, value := range m {
		if len(heap) == k {
			if value < heap[0].v {
				continue
			}
			heap[0].k = key
			heap[0].v = value
			down(0)
		} else {
			heap = append(heap, item{key, value})
			up(len(heap) - 1)
		}
	}
	rst := make([]int, 0, k)
	for _, t := range heap {
		rst = append(rst, t.k)
	}
	return rst
}

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	for _, v := range nums1 {
		m[v] = true
	}
	var rst []int
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			delete(m, v)
			rst = append(rst, v)
		}
	}
	return rst
}

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, v := range nums1 {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	var rst []int
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			rst = append(rst, v)
			m[v]--
			if m[v] == 0 {
				delete(m, v)
			}
		}
	}
	return rst
}

func numberOfPatterns(m int, n int) int {
	border := [][]int{
		{0, 0, 2, 0, 0, 0, 4, 0, 5},
		{0, 0, 0, 0, 0, 0, 0, 5, 0},
		{2, 0, 0, 0, 0, 0, 5, 0, 6},
		{0, 0, 0, 0, 0, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 5, 0, 0, 0, 0, 0},
		{4, 0, 5, 0, 0, 0, 0, 0, 8},
		{0, 5, 0, 0, 0, 0, 0, 0, 0},
		{5, 0, 6, 0, 0, 0, 8, 0, 0}}
	rst := 0
	hit := make([]bool, 10)
	path := 0

	var travel func(at int)
	var nextTravel = func(next int) {
		if hit[next] {
			return
		}
		if next < 1 || next > 9 {
			return
		}
		path++
		hit[next] = true
		travel(next)
		hit[next] = false
		path--
	}
	travel = func(at int) {
		if path >= m {
			rst++
		}
		if path == n {
			return
		}
		for next := 1; next < 10; next++ {
			if next == at {
				continue
			}
			if v := border[at-1][next-1]; v == 0 || hit[v] {
				nextTravel(next)
			}
		}
	}
	nextTravel(1)
	rst *= 4
	old := rst
	nextTravel(2)
	rst += (rst - old) * 3
	nextTravel(5)
	return rst
}

type SummaryRanges struct {
	hit [][]int
}

/** Initialize your data structure here. */
func ConstructorSummaryRanges() SummaryRanges {
	return SummaryRanges{}
}

func (this *SummaryRanges) AddNum(val int) {
	if len(this.hit) == 0 {
		this.hit = append(this.hit, []int{val, val})
		return
	}
	index := sort.Search(len(this.hit), func(i int) bool {
		return this.hit[i][0] >= val
	})
	if index < len(this.hit) && this.hit[index][0] == val {
		return
	}

	mergePre := index > 0 && val <= this.hit[index-1][1]+1
	mergePost := index < len(this.hit) && val+1 == this.hit[index][0]
	if mergePre && mergePost {
		this.hit[index-1][1] = this.hit[index][1]
		for i := index + 1; i < len(this.hit); i++ {
			this.hit[i-1] = this.hit[i]
		}
		this.hit = this.hit[:len(this.hit)-1]
	} else if mergePre && !mergePost {
		this.hit[index-1][1] = max(this.hit[index-1][1], val)
	} else if !mergePre && mergePost {
		this.hit[index][0] = val
	} else {
		this.hit = append(this.hit, nil)
		for i := len(this.hit) - 1; i > index; i-- {
			this.hit[i] = this.hit[i-1]
		}
		this.hit[index] = []int{val, val}
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.hit
}

func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && num&0xaaaaaaaa == 0
}

func rearrangeString(S string, k int) string {
	if k == 0 || k == 1 {
		return S
	}
	m := map[byte]int{}
	for _, c := range S {
		m[byte(c)]++
	}
	type pair struct {
		c   byte
		cnt int
	}
	pool := make([]pair, 0, len(m))
	for k, v := range m {
		pool = append(pool, pair{k, v})
	}
	sort.Slice(pool, func(i, j int) bool {
		return pool[i].cnt > pool[j].cnt
	})
	var buf []byte
	dup := map[byte]bool{}
	for len(pool) > 0 {
		for i := 0; i < k && len(pool) > 0; i++ {
			if len(buf)-k >= 0 {
				delete(dup, buf[len(buf)-k])
			}

			j := 0
			for ; j < len(pool) && dup[pool[j].c]; j++ {
			}
			if j == len(pool) {
				return ""
			}

			buf = append(buf, pool[j].c)
			pool[j].cnt--
			for l := j; l < len(pool)-1 && pool[l].cnt < pool[l+1].cnt; l++ {
				pool[l], pool[l+1] = pool[l+1], pool[l]
			}
			if pool[len(pool)-1].cnt == 0 {
				pool = pool[:len(pool)-1]
			}
			dup[buf[len(buf)-1]] = true
		}
	}
	return string(buf)
}
