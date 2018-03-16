package p538

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	solve(root, 0)
	return root
}

func solve(root *TreeNode, mod int) int {
	if root == nil {
		return mod
	}
	tmp := mod
	if root.Right != nil {
		tmp = solve(root.Right, mod)
	}
	root.Val += tmp
	if root.Left != nil {
		return solve(root.Left, root.Val)
	}

	return root.Val
}
