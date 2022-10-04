package leetcode

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if targetSum-root.Val == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
