package lc

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var v1, v2 []int
	for p := l1; p != nil; p = p.Next {
		v1 = append(v1, p.Val)
	}
	for p := l2; p != nil; p = p.Next {
		v2 = append(v2, p.Val)
	}
	i, j := len(v1)-1, len(v2)-1
	over := 0
	var p *ListNode
	for i >= 0 && j >= 0 {
		v := v1[i] + v2[j] + over
		n := &ListNode{Val: v % 10, Next: p}
		p = n
		over = v / 10
		i--
		j--
	}

	if i == j {
		if over > 0 {
			return &ListNode{Val: over, Next: p}
		}
		return p
	}
	var lv []int
	var at int
	if i < 0 {
		at = j
		lv = v2
	} else {
		at = i
		lv = v1
	}
	for k := at; k >= 0; k-- {
		v := lv[k] + over
		n := &ListNode{Val: v % 10, Next: p}
		over = v / 10
		p = n
	}
	if over > 0 {
		n := &ListNode{Val: over, Next: p}
		p = n
	}
	return p
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	t1 := map[int]int{}
	for _, a := range A {
		for _, b := range B {
			t1[a+b]++
		}
	}
	t2 := map[int]int{}
	for _, c := range C {
		for _, d := range D {
			t2[c+d]++
		}
	}
	cnt := 0
	for part, v1 := range t1 {
		if v2, ok := t2[-part]; ok {
			cnt += v1 * v2
		}
	}
	return cnt
}
