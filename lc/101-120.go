package lc

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	var isSym func(r1, r2 *TreeNode) bool
	isSym = func(r1, r2 *TreeNode) bool {
		if r1 == nil && r2 == nil {
			return true
		}
		if r1 == nil && r2 != nil || r2 == nil && r1 != nil || r1.Val != r2.Val {
			return false
		}
		return isSym(r1.Right, r2.Left) && isSym(r1.Left, r2.Right)
	}
	return isSym(root, root)
}

func levelOrder(root *TreeNode) (rst [][]int) {
	if root == nil {
		return rst
	}
	type pair struct {
		node  *TreeNode
		level int
	}
	var queue []pair
	var buffer []int
	queue = append(queue, pair{root, 0})
	curLevel := 0
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top.level > curLevel {
			curLevel = top.level
			var curLevels []int
			for _, b := range buffer {
				curLevels = append(curLevels, b)
			}
			rst = append(rst, curLevels)
			buffer = buffer[:0]
		}
		buffer = append(buffer, top.node.Val)
		if top.node.Left != nil {
			queue = append(queue, pair{top.node.Left, top.level + 1})
		}
		if top.node.Right != nil {
			queue = append(queue, pair{top.node.Right, top.level + 1})
		}
	}
	if len(buffer) > 0 {
		var curLevels []int
		for _, b := range buffer {
			curLevels = append(curLevels, b)
		}
		rst = append(rst, curLevels)
		buffer = buffer[:0]
	}
	return rst
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	reverse := true
	var buf = []*TreeNode{root}
	var tmp []*TreeNode
	var rst [][]int
	for len(buf) > 0 {
		var visit []int
		for i := 0; i < len(buf); i++ {
			visit = append(visit, buf[i].Val)
		}
		rst = append(rst, visit)
		for i := len(buf) - 1; i >= 0; i-- {
			if reverse {
				if buf[i].Right != nil {
					tmp = append(tmp, buf[i].Right)
				}
				if buf[i].Left != nil {
					tmp = append(tmp, buf[i].Left)
				}
			} else {
				if buf[i].Left != nil {
					tmp = append(tmp, buf[i].Left)
				}
				if buf[i].Right != nil {
					tmp = append(tmp, buf[i].Right)
				}
			}
		}

		buf, tmp = tmp, buf
		tmp = tmp[:0]
		reverse = !reverse
	}
	return rst
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}
	root := preorder[0]
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root {
			break
		}
	}
	node := &TreeNode{
		Val: root,
	}
	if i > 0 {
		node.Left = buildTree(preorder[1:i+1], inorder[0:i])
	}
	if i+1 < len(preorder) {
		node.Right = buildTree(preorder[i+1:], inorder[i+1:])
	}
	return node
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := len(nums) / 2
	return &TreeNode{
		Val:   nums[index],
		Left:  sortedArrayToBST(nums[0:index]),
		Right: sortedArrayToBST(nums[index+1:]),
	}
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var height func(n *TreeNode) (int, bool)
	heightMap := map[*TreeNode]int{}
	height = func(n *TreeNode) (int, bool) {
		if n == nil {
			return 0, true
		}
		v, ok := heightMap[n]
		if !ok {
			lh, _ := height(n.Left)
			rh, _ := height(n.Right)
			v = max(lh, rh) + 1
			heightMap[n] = v
		}
		lh, ok1 := height(n.Left)
		rh, ok2 := height(n.Right)
		var b bool
		if lh > rh {
			b = lh-rh <= 1
		} else {
			b = rh-lh <= 1
		}
		return v, b && ok1 && ok2
	}

	_, ok := height(root)
	return ok
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	var minp = math.MaxInt32
	if root.Left != nil {
		minp = min(minDepth(root.Left), minp)
	}
	if root.Right != nil {
		minp = min(minDepth(root.Right), minp)
	}
	return minp + 1
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func flatten(root *TreeNode) {
	var travel func(r *TreeNode) *TreeNode
	travel = func(r *TreeNode) *TreeNode {
		if r == nil {
			return nil
		}
		if r.Left == nil && r.Right == nil {
			return r
		}
		left, right := r.Left, r.Right
		next := r
		if left != nil {
			next = travel(r.Left)
			r.Right = left
		}

		if right != nil {
			next.Right = right
			next = travel(right)
		}
		r.Left = nil
		return next
	}
	travel(root)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var tmp1 = []*Node{root}
	var tmp2 []*Node

	for len(tmp1) > 0 {
		for i := 0; i < len(tmp1)-1; i++ {
			tmp1[i].Next = tmp1[i+1]
		}
		for _, n := range tmp1 {
			if n.Left != nil {
				tmp2 = append(tmp2, n.Left)
			}
			if n.Right != nil {
				tmp2 = append(tmp2, n.Right)
			}
		}
		tmp1, tmp2 = tmp2, tmp1
		tmp2 = tmp2[:0]
	}
	return root
}

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if i < len(triangle)-1 {
				triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
			}
		}
	}
	return triangle[0][0]
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	var dfs func(r *TreeNode)
	pathSum := 0
	var path []int
	dfs = func(r *TreeNode) {
		if r.Left == nil && r.Right == nil {
			if pathSum == sum {
				ans = append(ans, path)
			}
			return
		}
		if r.Left != nil {
			path = append(path, r.Left.Val)
			pathSum += r.Left.Val
			dfs(r.Left)
			pathSum -= r.Left.Val
			path = path[:len(path)-1]
		}
		if r.Right != nil {
			path = append(path, r.Right.Val)
			pathSum += r.Right.Val
			dfs(r.Right)
			pathSum -= r.Right.Val
			path = path[:len(path)-1]
		}
	}
	path = append(path, root.Val)
	pathSum = root.Val
	dfs(root)
	return ans
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	var dfs func(r *TreeNode) bool
	pathSum := 0
	dfs = func(r *TreeNode) bool {
		if r.Left == nil && r.Right == nil {
			return pathSum == sum
		}
		if r.Left != nil {
			pathSum += r.Left.Val
			if dfs(r.Left) {
				return true
			}
			pathSum -= r.Left.Val
		}
		if r.Right != nil {
			pathSum += r.Right.Val
			if dfs(r.Right) {
				return true
			}
			pathSum -= r.Right.Val
		}
		return false
	}
	pathSum = root.Val
	return dfs(root)
}
