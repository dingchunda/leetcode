package lc

func minDistance2(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}

	l1 := len(word1)
	l2 := len(word2)
	var rst = make([][]int, l1+1)
	for i := range rst {
		rst[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			r := rst[i-1][j-1]
			if word1[i-1] == word2[j-1] {
				r++
			}
			if rst[i-1][j] > r {
				r = rst[i-1][j]
			}
			if rst[i][j-1] > r {
				r = rst[i][j-1]
			}
			rst[i][j] = r
		}
	}

	return l1 - rst[l1][l2] + l2 - rst[l1][l2]
}

type NrTreeNode struct {
	Val      int
	Children []*NrTreeNode
}

func postorder(root *NrTreeNode) []int {
	var ans []int

	var travel func(root *NrTreeNode, ans []int) []int
	travel = func(root *NrTreeNode, ans []int) []int {
		if root == nil {
			return ans
		}
		for _, c := range root.Children {
			ans = travel(c, ans)
		}
		ans = append(ans, root.Val)
		return ans
	}
	return travel(root, ans)
}

func preorder(root *NrTreeNode) []int {
	var ans []int

	var travel func(root *NrTreeNode, ans []int) []int
	travel = func(root *NrTreeNode, ans []int) []int {
		if root == nil {
			return ans
		}
		ans = append(ans, root.Val)
		for _, c := range root.Children {
			ans = travel(c, ans)
		}
		return ans
	}
	return travel(root, ans)
}
