package lc

func findTargetSumWays(nums []int, s int) int {
	rst := map[int]int{}
	rst[nums[0]] += 1
	rst[-nums[0]] += 1
	for i := 1; i < len(nums); i++ {
		tmp := map[int]int{}
		for k, v := range rst {
			tmp[k+nums[i]] += v
			tmp[k-nums[i]] += v
		}
		rst = tmp
	}
	return rst[s]
}
