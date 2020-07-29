package lc

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return rec1[0] < rec2[2] && rec1[1] < rec2[3] && rec2[0] < rec1[2] && rec2[1] < rec1[3]
}

func numSimilarGroups(A []string) int {
	var sim = func(a, b string) bool {
		if len(a) != len(b) {
			return false
		}
		left, right := -1, -1
		for i := 0; i < len(a); i++ {
			if a[i] == b[i] {
				continue
			}
			if left == -1 {
				left = i
				continue
			}
			if right == -1 {
				right = i
				continue
			}
			return false
		}
		return left == -1 && right == -1 || left != -1 && right != -1 && a[left] == b[right] && a[right] == b[left]
	}
	parent := make([]int, len(A))
	for i := range parent {
		parent[i] = -1
	}
	var find func(at int) int
	find = func(at int) int {
		if parent[at] == -1 {
			return at
		}
		return find(parent[at])
	}
	ans := 0
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			sa, sb := find(i), find(j)
			if sa == sb {
				continue
			}
			if sim(A[j], A[i]) {
				parent[sb] = sa
			}
		}
	}
	class := map[int]bool{}
	for p := range parent {
		class[find(p)] = true
	}
	ans += len(class)

	return ans
}
