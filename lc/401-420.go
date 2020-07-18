package lc

import (
	"container/heap"
	"sort"
	"strconv"
)

func removeKdigits(num string, k int) string {
	digit := []byte(num)
	for l := 1; l <= k && len(digit) > 0; l++ {
		i := 0
		for ; i < len(digit) && (i != len(digit)-1 && digit[i] <= digit[i+1]); i++ {
		}
		for j := i + 1; j < len(digit); j++ {
			digit[j-1] = digit[j]
		}
		digit = digit[:len(digit)-1]
	}
	i := 0
	for ; i < len(digit) && digit[i] == '0'; i++ {
	}
	digit = digit[i:]
	if len(digit) > 0 {
		return string(digit)
	}
	return "0"
}

func canCross(stones []int) bool {
	if len(stones) == 2 {
		return stones[1]-stones[0] == 1
	}
	var cross []map[int]bool

	maxJump := stones[len(stones)-1]
	jump := map[int]bool{}
	if j := stones[len(stones)-1] - stones[len(stones)-2]; j < maxJump {
		jump[j] = true
	}
	cross = append(cross, jump)
	for i := len(stones) - 3; i >= 0; i-- {
		s := 0
		jump := map[int]bool{}
		if j := stones[len(stones)-1] - stones[i]; j < maxJump {
			jump[j] = true
		}
		for j := len(stones) - 2; j > i; j-- {
			need := stones[j] - stones[i]
			if (need < maxJump && cross[s][need]) || (need < maxJump && need >= 1 && cross[s][need-1]) || (need+1 < maxJump && cross[s][need+1]) {
				jump[need] = true
			}
			s++
		}
		cross = append(cross, jump)
	}
	//for _, v := range cross {
	//	fmt.Println(v)
	//}
	return cross[len(cross)-1][1]
}

func reconstructQueue(people [][]int) [][]int {
	if len(people) == 0 {
		return nil
	}
	sort.Sort(pair(people))
	var rst = make([][]int, len(people))
	for _, p := range people {
		slot := 0
		i := 0
		for ; slot < p[1]; i++ {
			if rst[i] == nil || rst[i][0] == p[0] {
				slot++
			}
		}
		for ; i < len(rst) && rst[i] != nil; i++ {
		}
		rst[i] = p
	}
	return rst
}

type pair [][]int

func (p pair) Len() int {
	return len(p)
}
func (p pair) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p pair) Less(i, j int) bool {
	if p[i][0] != p[j][0] {
		return p[i][0] < p[j][0]
	}
	return p[i][1] < p[j][1]
}

type node struct {
	i       int
	j       int
	height  int
	visited bool
}

type priorityQueue2 []*node

func (pq priorityQueue2) Len() int {
	return len(pq)
}

func (pq priorityQueue2) Less(i, j int) bool {
	return pq[i].height < pq[j].height
}

func (pq priorityQueue2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue2) Push(x interface{}) {
	no := x.(*node)
	*pq = append(*pq, no)
}

func (pq *priorityQueue2) Pop() interface{} {
	old := *pq
	n := len(old)
	no := old[n-1]
	*pq = old[0 : n-1]
	return no
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) < 3 || len(heightMap[0]) < 3 {
		return 0
	}

	nq := &priorityQueue{}

	heap.Init(nq)

	l0 := len(heightMap)
	l1 := len(heightMap[0])

	visit := make([][]bool, l0)
	for i := range visit {
		visit[i] = make([]bool, l1)
	}

	for i := 0; i < l1; i++ {
		heap.Push(nq, &node{
			i:      0,
			j:      i,
			height: heightMap[0][i],
		})
		visit[0][i] = true
		heap.Push(nq, &node{
			i:      l0 - 1,
			j:      i,
			height: heightMap[l0-1][i],
		})
		visit[l0-1][i] = true
	}
	for i := 1; i < l0-1; i++ {
		heap.Push(nq, &node{
			i:      i,
			height: heightMap[i][0],
		})
		visit[i][0] = true
		heap.Push(nq, &node{
			i:       i,
			j:       l1 - 1,
			height:  heightMap[i][l1-1],
			visited: true,
		})
		visit[i][l1-1] = true
	}
	sum := 0
	var dir = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for nq.Len() > 0 {
		top := heap.Pop(nq).(*node)
		for _, n := range dir {
			x, y := top.i+n[0], top.j+n[1]
			if x >= 0 && x < l0 && y >= 0 && y < l1 && !visit[x][y] {
				if top.height > heightMap[x][y] {
					sum += top.height - heightMap[x][y]
				}
				visit[x][y] = true
				heap.Push(nq, &node{
					i:      x,
					j:      y,
					height: max(heightMap[x][y], top.height),
				})
			}
		}
	}
	return sum
}

func fizzBuzz(n int) []string {
	rst := make([]string, n)
	for i := 3; i < n+1; i += 3 {
		rst[i-1] = "Fizz"
	}
	for i := 5; i < n+1; i += 5 {
		if rst[i-1] != "" {
			rst[i-1] = "FizzBuzz"
		} else {
			rst[i-1] = "Buzz"
		}
	}
	for i := 0; i < n; i++ {
		if rst[i] == "" {
			rst[i] = strconv.Itoa(i + 1)
		}
	}
	return rst
}

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	over := byte(0)
	var rst []byte
	for i >= 0 && j >= 0 {
		sum := num1[i] - '0' + num2[j] - '0' + over
		rst = append(rst, sum%10+'0')
		over = sum / 10
		i--
		j--
	}
	left := num1
	at := i
	if i == -1 {
		left = num2
		at = j
	}
	for at >= 0 {
		sum := left[at] - '0' + over
		rst = append(rst, sum%10+'0')
		over = sum / 10
		at--
	}
	if over != 0 {
		rst = append(rst, over+'0')
	}
	i, j = 0, len(rst)-1
	for i < j {
		rst[i], rst[j] = rst[j], rst[i]
		i++
		j--
	}
	return string(rst)
}

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	rst := make([]bool, sum+1)
	rst[0] = true
	for _, n := range nums {
		for l := sum; l >= n; l-- {
			rst[l] = rst[l] || rst[l-n]
		}
	}
	return rst[sum]
}
