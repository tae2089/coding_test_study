package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
    Val int
     Left *TreeNode
    Right *TreeNode
}

func checkTree(root *TreeNode) bool {
    leftNode, rightNode := root.Left, root.Right
	sum := leftNode.Val + rightNode.Val
	return root.Val == sum
}