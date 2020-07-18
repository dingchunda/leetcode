package lc

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2

	var head *ListNode
	var ex int
	var pre *ListNode
	for p1 != nil || p2 != nil {
		sum := ex
		if p1 != nil {
			sum += p1.Val
		}
		if p2 != nil {
			sum += p2.Val
		}
		ex = sum / 10
		sum = sum % 10
		node := &ListNode{
			Val: sum,
		}
		if pre != nil {
			pre.Next = node
		}
		pre = node
		if head == nil {
			head = node
		}
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
	if ex > 0 {
		node := &ListNode{
			Val: ex,
		}
		pre.Next = node
	}
	//fmt.Println(head.Val)
	return head
}

func myAtoi(str string) int {
	isPositive := true
	data := []byte(str)

	i := 0
	for i = 0; i < len(str); i++ {
		if data[i] == ' ' {
			continue
		}
		if data[i] == '+' || data[i] == '-' {
			if i+1 >= len(str) {
				return 0
			}
			if data[i+1] < '0' || data[i+1] > '9' {
				return 0
			}
			if data[i] == '-' {
				isPositive = false
			}
			continue
		}
		break
	}
	data = data[i:]

	rst := uint32(0)
	for i := 0; i < len(data); i++ {
		c := data[i]
		if c >= '0' && c <= '9' {
			cutoff := uint32(1<<31 - 1)
			if !isPositive {
				cutoff = 1 << 31
			}
			overflow := false
			if cutoff/10 < rst || cutoff-rst*10 < uint32(c-'0') {
				overflow = true
			}
			if overflow {
				if isPositive {
					return 1<<31 - 1
				} else {
					return -1 << 31
				}
			}
			rst = rst*10 + uint32(c-'0')
		} else {
			break
		}
	}
	if isPositive {
		return int(rst)
	} else {
		return -int(rst)
	}
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	rst := make([][]bool, len(s))
	for i := range rst {
		rst[i] = make([]bool, len(s))
	}
	maxK := -1
	start := -1
	for k := 0; k < len(s); k++ {
		for i := 0; i < len(s)-k; i++ {
			var ok bool
			if k == 0 {
				ok = true
			} else if k == 1 {
				ok = s[i] == s[i+1]
			} else {
				ok = rst[i+1][i+k-1] && s[i] == s[i+k]
			}
			if ok {
				rst[i][i+k] = true
				if k+1 > maxK {
					maxK = k + 1
					start = i
				}
			}
		}
	}
	return s[start : start+maxK]
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	start, end := 0, 1
	rst := 0
	max := 1
	table := map[byte]int{}
	table[byte(s[0])] = 0
	for end < len(s) {
		c := byte(s[end])
		if index, ok := table[c]; !ok {
			table[c] = end
			max++
			//fmt.Println(c)
		} else {
			if max > rst {
				rst = max
			}
			for i := start; i <= index; i++ {
				delete(table, byte(s[i]))
			}
			max -= index - start
			start = index + 1
			table[c] = end

		}
		end++
	}
	if max > rst {
		rst = max
	}
	return rst
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sum := len(nums1) + len(nums2)
	end := sum / 2
	if sum%2 == 0 {
		end--
	}
	i, j := 0, 0
	find := false
	var p []int
	var s int
	total := 0
	for {
		if i >= len(nums1) {
			p = nums2
			s = j
			j++
		} else if j >= len(nums2) {
			p = nums1
			s = i
			i++
		} else if nums1[i] < nums2[j] {
			p = nums1
			s = i
			i++
		} else {
			p = nums2
			s = j
			j++
		}
		if find {
			if i+j-1 == end+1 {
				total += p[s]
				return float64(total) / 2.0
			}
		} else {
			if i+j-1 == end {
				if sum%2 == 1 {
					//fmt.Println(p[s])
					return math.Floor(float64(p[s]))
				}
				total = p[s]
				find = true
			}
		}
	}
}

func reverse(x int) int {
	neg := false
	if x < 0 {
		neg = true
		x = -x
	}
	rst := 0
	var max int
	if neg {
		max = 1 << 31
	} else {
		max = 1<<31 - 1
	}
	for x > 0 {
		if rst > (max-x%10)/10 {
			return 0
		}
		rst = rst*10 + x%10
		x /= 10
	}
	if neg {
		return -rst
	}
	return rst
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		m[v] = i
	}
	for i, v := range nums {
		if rst, ok := m[target-v]; ok && rst != i {
			return []int{i, rst}
		}
	}
	return nil
}

func isPalindrome(x int) bool {
	data := []byte(strconv.Itoa(x))
	at := len(data) / 2
	var i, j int
	if len(data)%2 == 1 {
		i, j = at-1, at+1
	} else {
		i, j = at-1, at
	}
	for i >= 0 && j < len(data) {
		if data[i] != data[j] {
			return false
		}
		i--
		j++
	}
	return true
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := -1
	for i < j {
		if height[i] <= height[j] {
			v := (j - i) * height[i]
			if v > max {
				max = v
			}
			i++
		} else {
			v := (j - i) * height[j]
			if v > max {
				max = v
			}
			j--
		}
	}
	return max
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for index := 0; ; index++ {
		var s *byte
		for _, str := range strs {
			if index >= len(str) {
				return strs[0][:index]
			}
			if s == nil {
				c := byte(str[index])
				s = &c
			} else if str[index] != *s {
				return strs[0][:index]
			}
		}
	}
	return strs[0]
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var rst [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (nums[i] != nums[i-1]) {
			low := i + 1
			high := len(nums) - 1
			sum := -nums[i]
			for low < high {
				lv := nums[low]
				hv := nums[high]
				if total := lv + hv; total == sum {
					rst = append(rst, []int{nums[i], lv, hv})
					for j := low + 1; j <= high; j++ {
						low = j
						if nums[j] != lv {
							break
						}
					}
					if high > low {
						for j := high - 1; j >= low; j-- {
							high = j
							if nums[j] != hv {

								break
							}
						}
					}
					fmt.Println(low, high)

				} else if total < sum {
					low++
				} else {
					high--
				}
			}
		}
	}

	return rst
}

func letterCombinations(digits string) (rst []string) {
	if len(digits) == 0 {
		return
	}
	index := map[byte][]byte{
		2: {'a', 'b', 'c'},
		3: {'d', 'e', 'f'},
		4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'},
		7: {'r', 'q', 'p', 's'},
		8: {'v', 't', 'u'},
		9: {'w', 'x', 'y', 'z'},
	}
	var visit []byte
	var at = 0
	var travel func()
	travel = func() {
		if at < len(digits) {
			for _, c := range index[digits[at]-'0'] {
				visit = append(visit, c)
				at++
				travel()
				at--
				visit = visit[:len(visit)-1]

			}
		} else {
			rst = append(rst, string(visit))
		}
	}
	for _, c := range index[digits[0]-'0'] {
		visit = append(visit, c)
		at = 1
		travel()
		visit = visit[:0]
	}
	return
}

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	if len(nums) < 4 {
		return res
	}
	var l, r int
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r = j+1, len(nums)-1

			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l = l + 1
					}
					for l < r && nums[r] == nums[r-1] {
						r = r - 1
					}
					l++
					r--
				}

			}
		}
	}
	return res
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p []*ListNode
	q := head
	for q != nil {
		p = append(p, q)
		q = q.Next
	}

	if len(p)-n == 0 {
		if len(p) > 1 {
			return p[1]
		} else {
			return nil
		}
	}
	if len(p)-n+1 < len(p) {
		p[len(p)-n-1].Next = p[len(p)-n+1]
	} else {
		p[len(p)-n-1].Next = nil
	}
	return p[0]
}

func isValid(s string) bool {
	var stack []byte
	for _, c := range s {
		switch c {
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '{':
			stack = append(stack, byte(c))
			break
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
			break
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
			break
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
			break
		default:
			return false
		}
	}
	return len(stack) == 0
}
