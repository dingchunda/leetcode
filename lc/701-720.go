package lc

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{
				Val: val,
			}
		} else {
			insertIntoBST(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{
				Val: val,
			}
		} else {
			insertIntoBST(root.Right, val)
		}
	}
	return root
}
