package lc

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	n := &TreeNode{}
	if t1 != nil {
		n.Val += t1.Val
	}
	if t2 != nil {
		n.Val += t2.Val
	}
	var t1Left, t2Left, t1Right, t2Right *TreeNode
	if t1 != nil {
		t1Left = t1.Left
		t1Right = t1.Right
	}
	if t2 != nil {
		t2Left = t2.Left
		t2Right = t2.Right
	}
	n.Left = mergeTrees(t1Left, t2Left)
	n.Right = mergeTrees(t1Right, t2Right)
	return n
}
