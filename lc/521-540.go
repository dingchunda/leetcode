package lc

import (
	"math/rand"
)

func checkSubarraySum(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 && nums[i+1] == 0 {
			return true
		}
	}
	if k == 0 {
		return false
	}
	if k < 0 {
		k = -k
	}
	sum := 0
	find := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if i > 1 {
			find[(sum-nums[i-1])%k] = true
		}
		sum += nums[i]
		if i > 0 && sum > 0 && sum%k == 0 {
			return true
		}
		if sum > k && find[sum%k] {
			return true
		}
	}
	return false
}

type SolutionWeightPIck struct {
	sum int
	w   []int
}

func ConstructorSolutionWeightPIck(w []int) SolutionWeightPIck {
	sum := 0
	for _, v := range w {
		sum += v

	}
	return SolutionWeightPIck{
		sum: sum,
		w:   w,
	}
}

func (this *SolutionWeightPIck) PickIndex() int {
	sample := rand.Intn(this.sum)
	sum := 0
	for i := 0; i < len(this.w); i++ {
		sum += this.w[i]
		if sample < sum {
			return i
		}
	}
	return -1
}
