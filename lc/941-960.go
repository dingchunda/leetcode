package lc

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var buf, tmp []*TreeNode
	buf = append(buf, root)
	lvl := 0
	lastbroken := false
	for len(buf) > 0 {
		if len(tmp) > 0 && len(tmp) != (1<<uint(lvl-1)) {
			return false
		}
		tmp = tmp[:0]
		end := false
		broken := false
		for _, n := range buf {
			if n.Left != nil {
				if end {
					broken = true
				}
				tmp = append(tmp, n.Left)
			} else {
				end = true
			}
			if n.Right != nil {
				if end {
					broken = true
				}
				tmp = append(tmp, n.Right)
			} else {
				end = true
			}
		}
		//fmt.Println(lvl,len(tmp),lastbroken)
		if len(tmp) == 0 && lastbroken {
			return false
		}
		lastbroken = broken
		buf, tmp = tmp, buf
		lvl++
	}
	return true
}
