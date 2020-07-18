package lc

import (
	"math"
	"strings"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	cnt := 1
	for p.Next != nil {
		p = p.Next
		cnt++
	}
	if cnt == 1 {
		return head
	}
	k = k % cnt
	for i := 0; i < k; i++ {
		var pre *ListNode
		p = head
		for p.Next != nil {
			pre, p = p, p.Next
		}
		pre.Next = nil
		p.Next = head
		head = p
	}
	return head
}

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	rst := make([][]int, m)
	for i := range rst {
		rst[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		rst[0][i] = 1
	}
	for i := 0; i < m; i++ {
		rst[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			rst[i][j] = rst[i-1][j] + rst[i][j-1]
		}
	}
	return rst[m-1][n-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	rst := make([][]int, m)
	for i := range rst {
		rst[i] = make([]int, n)
	}
	if obstacleGrid[m-1][n-1] == 0 {
		rst[m-1][n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		if obstacleGrid[m-1][i] == 1 {
			rst[m-1][i] = 0
		} else {
			rst[m-1][i] = rst[m-1][i+1]
		}
	}
	for i := m - 2; i >= 0; i-- {
		if obstacleGrid[i][n-1] == 1 {
			rst[i][n-1] = 0
		} else {
			rst[i][n-1] = rst[i+1][n-1]
		}
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				rst[i][j] = 0
			} else {
				rst[i][j] = rst[i+1][j] + rst[i][j+1]
			}
		}
	}
	return rst[0][0]
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rst := make([][]int, len(grid))
	for i := 0; i < len(rst); i++ {
		rst[i] = make([]int, len(grid[0]))
	}
	for i := len(grid) - 1; i >= 0; i-- {
		arr := grid[i]
		for j := len(arr) - 1; j >= 0; j-- {
			if i == len(grid)-1 && j == len(grid[0])-1 {
				rst[i][j] = grid[i][j]
				continue
			}
			min := math.MaxInt32
			if i+1 < len(rst) {
				min = grid[i][j] + rst[i+1][j]
			}
			if j+1 < len(grid[0]) {
				if v := grid[i][j] + rst[i][j+1]; v < min {
					min = v
				}
			}
			rst[i][j] = min
		}
	}
	return rst[0][0]
}

func plusOne(digits []int) []int {
	var ex int
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 10 {
			break
		}
		if i > 0 {
			digits[i-1] += digits[i] / 10
			digits[i] %= 10
		} else {
			ex = digits[i] / 10
			digits[i] %= 10
		}
	}
	if ex > 0 {
		var v = []int{ex}
		v = append(v, digits...)
		return v
	}
	return digits
}

func addBinary(a string, b string) string {
	ml := max(len(a), len(b))
	var rst = make([]byte, ml+1)
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		var v byte
		if i >= 0 {
			v += byte(a[i]) - byte('0')
		}
		if j >= 0 {
			v += byte(b[j]) - byte('0')
		}
		start := ml - max(i, j) - 1
		rst[start] += v
		for k := start; ; k++ {
			if rst[k] < 2 {
				break
			}
			rst[k+1] += rst[k] / 2
			rst[k] %= 2
		}
		i--
		j--
	}

	end := 0
	for i := len(rst) - 1; i >= 0; i-- {
		if rst[i] > 0 {
			end = i
			break
		}
	}
	r := make([]byte, end+1)
	for i := end; i >= 0; i-- {
		r[end-i] = rst[i] + byte('0')
	}
	return string(r)
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	rst := make([]int, n)
	rst[0] = 1
	rst[1] = 2
	for i := 2; i < n; i++ {
		rst[i] = rst[i-1] + rst[i-2]
	}
	return rst[n-1]
}

func simplifyPath(path string) string {
	var validPath []string
	for _, path := range strings.Split(path, "/") {
		if path == "." {
			continue
		}
		if path == "" {
			if len(validPath) == 0 {
				validPath = append(validPath, "$")
			}
			continue
		}
		if path == ".." {
			if len(validPath) > 0 {
				if validPath[len(validPath)-1] == "$" {
				} else if validPath[len(validPath)-1] == ".." {
					validPath = append(validPath, "..")
				} else {
					validPath = validPath[:len(validPath)-1]
				}
			} else {
				validPath = append(validPath, "..")
			}
			continue
		}
		validPath = append(validPath, path)
	}
	if len(validPath) == 1 && validPath[0] == "$" {
		return "/"
	}
	rst := strings.Join(validPath, "/")
	if len(rst) > 0 && rst[0] == '$' {
		rst = rst[1:]
	}
	return rst
}

func minDistance(word1 string, word2 string) int {
	rst := make([][]int, len(word1)+1)
	for i := range rst {
		rst[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word2); i++ {
		rst[0][i] = i
	}
	for i := 0; i <= len(word1); i++ {
		rst[i][0] = i
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				rst[i][j] = rst[i-1][j-1]
			} else {
				v := rst[i-1][j-1]
				if v1 := rst[i][j-1]; v1 < v {
					v = rst[i][j-1]
				}
				if v1 := rst[i-1][j]; v1 < v {
					v = rst[i-1][j]
				}
				rst[i][j] = v + 1
			}
		}
	}
	return rst[len(word1)][len(word2)]
}

func setZeroes(matrix [][]int) {
	zeroAt := false
	for i := 0; i < len(matrix); i++ {
		zero := false
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					zeroAt = true
				}
				if !zero && i > 0 {
					zero = true
					for k := 0; k <= j; k++ {
						matrix[i][k] = 0
					}
				}
				if i > 0 && matrix[0][j] != 0 {
					for k := 0; k < i; k++ {
						matrix[k][j] = 0
					}
				}
			} else {
				if zero {
					matrix[i][j] = 0
				}
				if i > 0 && matrix[0][j] == 0 {
					matrix[i][j] = 0
				}
			}
		}
	}
	if zeroAt {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	var ser func(mx [][]int, s, t int) bool
	ser = func(mx [][]int, s, t int) bool {
		if len(mx) == 0 || s >= t {
			return false
		}
		mid1 := len(mx) / 2
		mid2 := (s + t) / 2
		v := mx[mid1][mid2]
		if v == target {
			return true
		}
		if target < v {
			return ser(mx[:mid1], s, t) || ser(mx[mid1:], s, mid2)
		} else {
			return ser(mx[mid1+1:], s, t) || ser(mx[:mid1+1], mid2+1, t)
		}
	}
	return ser(matrix, 0, len(matrix[0]))
}

func sortColors(nums []int) {
	pleft := 0
	for ; pleft < len(nums); pleft++ {
		if nums[pleft] != 0 {
			break
		}
	}
	pright := len(nums) - 1
	for ; pright >= pleft; pright-- {
		if nums[pright] != 2 {
			break
		}
	}
	p := pleft
	for p <= pright {
		if nums[p] == 0 {
			nums[pleft], nums[p] = nums[p], nums[pleft]
			pleft++
			p++
		} else if nums[p] == 2 {
			nums[pright], nums[p] = nums[p], nums[pright]
			for ; pright >= p; pright-- {
				if nums[pright] != 2 {
					break
				}
			}
		} else {
			p++
		}
	}
}

func minWindow(s string, t string) string {
	cnt := map[byte]int{}
	for _, c := range t {
		cnt[byte(c)]++
	}
	hit := map[byte]int{}
	i, j := 0, 0
	min := len(s) + 1
	var rst string
	for ; i < len(s); i++ {
		if _, ok := cnt[s[i]]; !ok {
			continue
		}
		hit[s[i]]++

		fail := false
		for k, v := range cnt {
			if v > hit[k] {
				fail = true
				break
			}
		}
		if fail {
			continue
		}
		if i-j+1 < min {
			min = i - j + 1
			rst = s[j : i+1]
			//fmt.Println("i", rst)
		}
		for j <= i {
			if _, ok := cnt[s[i]]; !ok {
				j++
				continue
			}
			fail := false
			for k, v := range cnt {
				if v > hit[k] {
					fail = true
					break
				}
			}
			if fail {
				break
			}
			if i-j+1 < min {
				min = i - j + 1
				rst = s[j : i+1]
				//fmt.Println("j", rst)
			}
			hit[s[j]]--
			if hit[s[j]] == 0 {
				delete(hit, s[j])
			}
			j++
		}
	}
	return rst
}

func combine(n int, k int) [][]int {
	var rst [][]int
	var buffer []int
	var hit = make([]bool, n)
	var travel func()
	travel = func() {
		cur := 0
		for _, h := range hit {
			if h {
				cur++
			}
		}
		if cur == k {
			dst := make([]int, len(buffer))
			copy(dst, buffer)
			rst = append(rst, dst)
			return
		}
		var index int
		var last = -1
		for index = 0; index < n; index++ {
			if hit[index] {
				last = index
			}
		}
		for i := last + 1; i < n; i++ {
			if !hit[i] {
				hit[i] = true
				buffer = append(buffer, i+1)
				travel()
				buffer = buffer[:len(buffer)-1]
				hit[i] = false
			}
		}
	}
	travel()
	return rst
}

func subsets(nums []int) [][]int {
	var rst [][]int
	for _, n := range nums {
		end := len(rst)
		for _, c := range rst[:end] {
			dst := make([]int, len(c)+1)
			copy(dst, c)
			dst[len(dst)-1] = n
			rst = append(rst, dst)
		}
		rst = append(rst, []int{n})
	}
	rst = append(rst, nil)
	return rst
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	cover := make([][]bool, len(board))
	for i := 0; i < len(cover); i++ {
		cover[i] = make([]bool, len(board[0]))
	}

	w := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			at := 0
			cover[i][j] = true
			ok := dfs(i, j, &at, board, cover, w)
			cover[i][j] = false
			if ok {
				return true
			}
		}
	}
	return false
}

func dfs(i, j int, at *int, board [][]byte, cover [][]bool, word []byte) bool {
	if word[*at] != board[i][j] {
		return false
	}
	if *at == len(word)-1 {
		return true
	}
	if i-1 >= 0 && !cover[i-1][j] {
		*at++
		cover[i-1][j] = true
		ok := dfs(i-1, j, at, board, cover, word)
		*at--
		cover[i-1][j] = false
		if ok {
			return true
		}
	}
	if i+1 < len(board) && !cover[i+1][j] {
		*at++
		cover[i+1][j] = true
		ok := dfs(i+1, j, at, board, cover, word)
		*at--
		cover[i+1][j] = false
		if ok {
			return true
		}
	}
	if j+1 < len(board[0]) && !cover[i][j+1] {
		*at++
		cover[i][j+1] = true
		ok := dfs(i, j+1, at, board, cover, word)
		*at--
		cover[i][j+1] = false
		if ok {
			return true
		}
	}
	if j-1 >= 0 && !cover[i][j-1] {
		*at++
		cover[i][j-1] = true
		ok := dfs(i, j-1, at, board, cover, word)
		*at--
		cover[i][j-1] = false
		if ok {
			return true
		}
	}
	return false
}

func removeDuplicates2(nums []int) int {
	var rst = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[rst] || rst == 0 || nums[rst] != nums[rst-1] {
			nums[rst+1] = nums[i]
			rst++
		}
	}
	return rst + 1
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	lastLeft := make([]int, len(heights))
	lastLeft[0] = -1
	for i := 1; i < len(heights); i++ {
		pre := i - 1
		for pre >= 0 && heights[pre] >= heights[i] {
			pre = lastLeft[pre]
		}
		if pre >= 0 {
			lastLeft[i] = pre
		} else {
			lastLeft[i] = -1
		}
	}
	firstRight := make([]int, len(heights))
	firstRight[len(firstRight)-1] = len(firstRight)
	for i := len(firstRight) - 2; i >= 0; i-- {
		post := i + 1
		for post < len(heights) && heights[post] >= heights[i] {
			post = firstRight[post]
		}
		if post < len(heights) {
			firstRight[i] = post
		} else {
			firstRight[i] = len(firstRight)
		}
	}
	ans := 0
	for i := 0; i < len(heights); i++ {
		left := 0
		if lastLeft[i] != -1 {
			left = i - lastLeft[i] - 1
		} else {
			left = i
		}
		right := 0
		if firstRight[i] != len(firstRight) {
			right = firstRight[i] - 1 - i
		} else {
			right = len(firstRight) - 1 - i
		}
		if v := (left + right + 1) * heights[i]; v > ans {
			ans = v
		}
	}
	return ans
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var left, mid *ListNode
	var leftTail, midTail *ListNode

	p := head
	for p != nil {
		if p.Val < x {
			if left == nil {
				left = p
				leftTail = p
			} else {
				leftTail.Next = p
				leftTail = p
			}
		} else {
			if mid == nil {
				mid = p
				midTail = p
			} else {
				midTail.Next = p
				midTail = p
			}
		}
		p = p.Next
	}
	if leftTail != nil {
		leftTail.Next = mid
	}
	if midTail != nil {
		midTail.Next = nil
	}
	if left != nil {
		return left
	}
	return mid
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, m)
	copy(tmp, nums1)
	i := 0
	j := 0
	for i < m || j < n {
		if j == n || i < m && tmp[i] <= nums2[j] {
			nums1[i+j] = tmp[i]
			i++
		} else {
			nums1[i+j] = nums2[j]
			j++
		}
	}
}
