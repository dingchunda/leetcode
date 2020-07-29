package lc

import (
	"sort"
	"strconv"
)

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var queue []int
	max := 0
	for _, meet := range intervals {
		i := 0
		for ; i < len(queue) && meet[0] >= queue[i]; i++ {
		}
		queue = queue[i:]
		queue = append(queue, meet[1])
		for i := len(queue) - 1; i >= 1 && queue[i] < queue[i-1]; i-- {
			queue[i], queue[i-1] = queue[i-1], queue[i]
		}
		if len(queue) > max {
			max = len(queue)
		}
	}
	return max
}

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	ans := make([][3]int, len(costs))
	ans[0][0] = costs[0][0]
	ans[0][1] = costs[0][1]
	ans[0][2] = costs[0][2]
	for i := 1; i < len(costs); i++ {
		ans[i][0] = min(ans[i-1][1], ans[i-1][2]) + costs[i][0]
		ans[i][1] = min(ans[i-1][0], ans[i-1][2]) + costs[i][1]
		ans[i][2] = min(ans[i-1][0], ans[i-1][1]) + costs[i][2]
	}
	return min(min(ans[len(costs)-1][0], ans[len(costs)-1][1]), ans[len(costs)-1][2])
}

func singleNumber3(nums []int) []int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}
	r := xor & (-xor)
	var a, b int
	for _, n := range nums {
		if n&r != 0 {
			a ^= n
		} else {
			b ^= n
		}
	}
	return []int{a, b}
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var ans []string
	var dfs func(r *TreeNode)
	var path string
	dfs = func(r *TreeNode) {
		if r.Left == nil && r.Right == nil {
			ans = append(ans, path)
			return
		}
		if r.Left != nil {
			ln := len(path)
			path += "->" + strconv.Itoa(r.Left.Val)
			dfs(r.Left)
			path = path[:ln]
		}
		if r.Right != nil {
			ln := len(path)
			path += "->" + strconv.Itoa(r.Right.Val)
			dfs(r.Left)
			path = path[:ln]
		}
	}
	path = strconv.Itoa(root.Val)
	dfs(root)
	return ans
}

func countUnivalSubtrees(root *TreeNode) int {
	ans := 0
	var dfs func(r *TreeNode) bool
	dfs = func(r *TreeNode) bool {
		if r == nil {
			return true
		}
		isLeft := dfs(r.Left)
		isRight := dfs(r.Right)
		v := isLeft && isRight && (r.Left == nil || r.Left.Val == r.Val) && (r.Right == nil || r.Right.Val == r.Val)
		if v {
			ans++
		}
		return v
	}
	dfs(root)
	return ans
}
