package lc

func findAnagrams(s string, p string) (rst []int) {
	var m = map[byte]int{}
	for _, v := range []byte(p) {
		m[v]++
	}
	pstart := 0
	for phead := 0; phead < len(s); phead++ {
		c := s[phead]
		v, _ := m[c]
		if v == 0 {
			for ; pstart < phead; pstart++ {
				m[s[pstart]]++
				if m[c] > 0 {
					m[c]--
					break
				}
			}
			pstart++
		} else {
			m[c]--
			if phead-pstart == len(p)-1 {
				rst = append(rst, pstart)
				m[s[pstart]]++
				pstart++
			}
		}
	}
	return
}

func findKthNumber(n int, k int) int {
	var count = func(root int) int {
		rst := 0
		next := root + 1
		for root <= n {
			rst += min(next, n+1) - root
			root *= 10
			next *= 10
		}
		return rst
	}
	var dfs func(root, k int) int
	dfs = func(root, k int) int {
		if k == 1 {
			return root
		}
		c := count(root)
		if c < k {
			k -= c
			return dfs(root+1, k)
		} else {
			k--
			return dfs(root*10, k)
		}
	}
	return dfs(1, k)
}
