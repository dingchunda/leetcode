package lc

import "sort"

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	opened := 0
	var dfs func(i int)
	dfs = func(i int) {
		visited[i] = true
		opened++
		for _, next := range rooms[i] {
			if !visited[next] {
				dfs(next)
			}
		}
	}
	dfs(0)
	return opened == len(rooms)
}

type ExamRoom struct {
	seat map[int]bool
	n    int
}

func ConstructorExamRoom(N int) ExamRoom {
	return ExamRoom{
		seat: map[int]bool{},
		n:    N,
	}
}

func (this *ExamRoom) Seat() int {
	if len(this.seat) == 0 {
		this.seat[0] = true
		return 0
	}
	var key []int
	for k := range this.seat {
		key = append(key, k)
	}
	sort.Ints(key)
	seatAt := -1
	bestDis := -1
	for i := 0; i < len(key)-1; i++ {
		at := (key[i+1] + key[i]) / 2
		dis := at - key[i]
		if dis > bestDis {
			bestDis = dis
			seatAt = at
		}
	}
	//fmt.Println(seatAt, bestDis, key)
	if len(key) > 0 && key[0] > 0 && key[0] >= bestDis {
		seatAt = 0
		bestDis = key[0]
	}
	if len(key) > 0 && key[len(key)-1] < this.n-1 && this.n-key[len(key)-1]-1 > bestDis {
		seatAt = this.n - 1
	}

	this.seat[seatAt] = true
	return seatAt
}

func (this *ExamRoom) Leave(p int) {
	delete(this.seat, p)
}

func maxDistToClosest(seats []int) int {
	preSit := -1
	bestDis := -1
	lastSit := -1
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			if preSit != -1 {
				at := (i + preSit) / 2
				dis := at - preSit
				if dis > 0 && dis > bestDis {
					bestDis = dis
				}
			} else if i != 0 {
				bestDis = i
			}
			preSit = i
			lastSit = i
		}
	}
	if lastSit != len(seats)-1 && len(seats)-1-lastSit > bestDis {
		return len(seats) - 1 - lastSit
	}
	return bestDis
}
