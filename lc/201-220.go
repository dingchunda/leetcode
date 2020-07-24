package lc

import "math"

func isHappy(n int) bool {
	m := map[[9]byte]bool{}
	ans := [9]byte{1, 0, 0, 0, 0, 0, 0, 0, 0}
	var ishappy func(n int) bool
	ishappy = func(n int) bool {
		key := [9]byte{}
		sum := 0
		for n > 0 {
			c := n % 10
			if c > 0 {
				key[c-1]++
			}
			sum += c * c
			n /= 10
		}
		if key == ans {
			return true
		}
		if m[key] {
			return false
		}
		m[key] = true
		return ishappy(sum)
	}
	return ishappy(n)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	var pre *ListNode
	p := head
	for p != nil {
		if p.Val == val {
			if p == head {
				head = p.Next
			}
			if pre != nil {
				pre.Next = p.Next
			}
			p = p.Next
		} else {
			pre = p
			p = p.Next
		}
	}
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head.Next
	pre := head
	pre.Next = nil
	for p != nil {
		pre, p, p.Next = p, p.Next, pre
	}
	return pre
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	finished := make([]bool, numCourses)
	total := numCourses
	for {
		ok := false
		for i := 0; i < numCourses; i++ {
			if finished[i] {
				continue
			}
			restrict := false
			for _, v := range prerequisites {
				if v[0] == i && !finished[v[1]] {
					restrict = true
					break
				}
			}
			if restrict {
				continue
			}
			finished[i] = true
			total--
			ok = true
		}
		if !ok {
			return false
		}
		if total == 0 {
			break
		}
	}
	return true
}

type Trie struct {
	children [26]*Trie
	ok       bool
}

/** Initialize your data structure here. */
func ConstructorTrie() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.ok = true
		return
	}
	root := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if root.children[c] == nil {
			child := Constructor()
			root.children[c] = &child
		}
		root = root.children[c]
		if i == len(word)-1 {
			root.ok = true
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.ok
	}
	root := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if root.children[c] == nil {
			return false
		}
		root = root.children[c]
		if i == len(word)-1 {
			return root.ok
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for i := 0; i < len(prefix); i++ {
		c := prefix[i] - 'a'
		if root.children[c] == nil {
			return false
		}
		root = root.children[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func minSubArrayLen(s int, nums []int) int {
	p, q := 0, 0
	sum := 0
	min := math.MaxInt32
	for q < len(nums) {
		for ; sum < s && q < len(nums); q++ {
			sum += nums[q]
		}
		if q < len(nums) && q-p < min {
			min = q - p
		}

		for ; p < q && sum >= s; p++ {
			sum -= nums[p]
			if q-p < min {
				min = q - p
			}
		}
	}
	if min == math.MaxInt32 {
		return 0
	}
	return min
}

func findOrder(numCourses int, prerequisites [][]int) (rst []int) {
	finished := make([]bool, numCourses)
	total := numCourses
	rst = make([]int, 0, numCourses)
	pre := map[int][]int{}
	for _, v := range prerequisites {
		pre[v[0]] = append(pre[v[0]], v[1])
	}
	ongoing := map[int]bool{}
	for i := 0; i < numCourses; i++ {
		ongoing[i] = true
	}
	for {
		ok := false
		for i := range ongoing {
			if finished[i] {
				continue
			}
			restrict := false

			for _, v := range pre[i] {
				if !finished[v] {
					restrict = true
					break
				}
			}
			if restrict {
				continue
			}
			finished[i] = true
			delete(ongoing, i)
			rst = append(rst, i)
			total--
			ok = true
		}
		if !ok {
			return nil
		}
		if total == 0 {
			break
		}
	}
	return rst
}

func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}
	rst := make([]int, n)
	rst[n-1] = nums[n-1]
	rst[n-2] = max(nums[n-1], nums[n-2])
	for i := n - 3; i >= 1; i-- {
		rst[i] = max(rst[i+1], nums[i]+rst[i+2])
	}
	a1 := rst[1]
	for i := range rst {
		rst[i] = 0
	}
	rst[n-2] = nums[n-2]
	rst[n-3] = max(nums[n-2], nums[n-3])
	for i := n - 4; i >= 2; i-- {
		rst[i] = max(rst[i+1], nums[i]+rst[i+2])
	}
	a2 := rst[2]
	return max(a1, a2+nums[0])
}

func findKthLargest(nums []int, k int) int {
	if k <= 0 || k > len(nums) {
		return -1
	}
	start := 0
	end := len(nums) - 1
	flag := 0
	for start < end {
		v := nums[flag]
		for start < end && nums[end] <= v {
			end--
		}
		if end > start {
			nums[start] = nums[end]
		}
		for start < end && nums[start] >= v {
			start++
		}
		flag = start
		nums[end] = nums[start]
		nums[start] = v
	}
	if flag == k-1 {
		return nums[flag]
	} else if flag >= k {
		return findKthLargest(nums[0:flag], k)
	} else {
		return findKthLargest(nums[flag+1:], k-1-flag)
	}
}

func getSkyline(buildings [][]int) [][]int {
	var ans [][]int
	var forward [][]int
	var tmp [][]int
	for _, b := range buildings {
		first := []int{b[0], b[2]}
		second := []int{b[1], b[2]}
		if len(ans) == 0 {
			ans = append(ans, first)
			forward = append(forward, second)
			continue
		}
		ok := false
		if len(ans) > 0 && ans[len(ans)-1][0] == first[0] {
			if ans[len(ans)-1][1] < first[1] {
				ans[len(ans)-1][1] = first[1]
			}
			ok = true
		}
		if !ok {
			// popup forward line
			i := 0
			for ; i < len(forward) && first[0] > forward[i][0]; i++ {
				if i+1 < len(forward) {
					ans = append(ans, []int{forward[i][0], forward[i+1][1]})
				} else {
					ans = append(ans, []int{forward[i][0], 0})
				}
			}
			forward = forward[i:]
			if len(forward) > 0 && forward[0][0] == first[0] && first[1] != forward[0][1] {
				forward = forward[1:]
			}
			if len(forward) == 0 || forward[0][1] < first[1] {
				ans = append(ans, first)
			}
		}
		tmp = tmp[:0]
		throw := false
		for i := 0; i < len(forward); i++ {
			if second[0] > forward[i][0] {
				if forward[i][1] > second[1] {
					tmp = append(tmp, forward[i])
				}
				continue
			}

			if second[0] == forward[i][0] {
				if second[1] > forward[i][1] {
					forward[i][1] = second[1]
				}
				throw = true
				tmp = append(tmp, forward[i:]...)
				break
			}
			if second[0] < forward[i][0] {
				if second[1] > forward[i][1] {
					tmp = append(tmp, second)

				}
				tmp = append(tmp, forward[i:]...)
				throw = true
				break
			}
		}
		if !throw {
			tmp = append(tmp, second)
		}
		//fmt.Println(i, tmp)
		forward, tmp = tmp, forward
	}
	for i := 0; i < len(forward); i++ {
		if i+1 < len(forward) {
			ans = append(ans, []int{forward[i][0], forward[i+1][1]})
		} else {
			ans = append(ans, []int{forward[i][0], 0})
		}
	}
	return ans
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			diff := nums[i] - nums[j]
			if diff < 0 {
				diff = -diff
			}
			if diff <= t {
				return true
			}
		}
	}
	return false
}

func containsDuplicate(nums []int) bool {
	set := map[int]bool{}
	for _, n := range nums {
		if _, ok := set[n]; ok {
			return true
		}
		set[n] = true
	}
	return false
}
