package lc

type NodeWithParent struct {
	Val    int
	Left   *NodeWithParent
	Right  *NodeWithParent
	Parent *NodeWithParent
}

func inorderSuccessor2(NodeWithParent *NodeWithParent) *NodeWithParent {
	if NodeWithParent.Parent == nil {
		if NodeWithParent.Right != nil {
			p := NodeWithParent.Right
			for p.Left != nil {
				p = p.Left
			}
			return p
		}
		return nil
	} else if NodeWithParent.Parent.Right == NodeWithParent {
		if NodeWithParent.Right != nil {
			p := NodeWithParent.Right
			for p.Left != nil {
				p = p.Left
			}
			return p
		}
		p := NodeWithParent
		for p.Parent != nil && p.Parent.Right == p {
			p = p.Parent
		}
		//if p.Parent != nil && p.Parent.Left == p {
		return p.Parent
		//}
		//return nil
	} else {
		if NodeWithParent.Right != nil {
			p := NodeWithParent.Right
			for p.Left != nil {
				p = p.Left
			}
			return p
		} else {
			return NodeWithParent.Parent
		}
	}
}

func longestPalindromeSubseq(s string) int {
	rst := make([][]int, len(s)+1)
	for i := range rst {
		rst[i] = make([]int, len(s)+1)
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(s); j++ {
			if s[i-1] == s[len(s)-j] {
				rst[i][j] = rst[i-1][j-1] + 1
			} else {
				rst[i][j] = max(rst[i-1][j], rst[i][j-1])
			}
		}
	}
	return rst[len(s)][len(s)]
}

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	a0, a1 := 0, 1
	for k := 2; k <= N; k++ {
		a0, a1 = a1, a0+a1
	}
	return a1
}
