package lc

import (
	"math"
	"math/rand"
)

func rand10() int {
	var row, col, idx int
	for {
		row = rand7()
		col = rand7()
		idx = col + (row-1)*7
		if idx <= 40 {
			break
		}
	}
	return 1 + (idx-1)%10
}

func rand7() int {
	return 0
}

type SolutionRandomCycle struct {
	radius float64
	x, y   float64
}

func ConstructorSolutionRandomCycle(radius float64, x_center float64, y_center float64) SolutionRandomCycle {
	return SolutionRandomCycle{
		radius: radius,
		x:      x_center,
		y:      y_center,
	}
}

func (this *SolutionRandomCycle) RandPoint() []float64 {
	r := math.Sqrt(rand.Float64()) * this.radius

	arc := rand.Float64() * 2 * math.Pi
	return []float64{r*math.Cos(arc) + this.x, r*math.Sin(arc) + this.y}
}

type SolutionRects struct {
	rects  [][]int
	sample []int
	sum    int
}

func Constructor(rects [][]int) SolutionRects {
	sample := make([]int, len(rects))
	sum := 0
	for i, rect := range rects {
		sample[i] = (rect[2] - rect[0] + 1) * (rect[3] - rect[1] + 1)
		sum += sample[i]
	}
	return SolutionRects{
		rects:  rects,
		sample: sample,
		sum:    sum,
	}
}

func (this *SolutionRects) Pick() []int {
	if this.sum == 0 {
		return nil
	}
	sample := rand.Intn(this.sum)
	sum := 0
	at := -1
	for i, s := range this.sample {
		sum += s
		if sample < sum {
			at = i
			break
		}
	}
	rect := this.rects[at]
	return []int{
		rand.Intn(rect[2]-rect[0]+1) + rect[0],
		rand.Intn(rect[3]-rect[1]+1) + rect[1],
	}
}
