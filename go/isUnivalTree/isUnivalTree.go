package isUnivalTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {

	if root.Left != nil && root.Right != nil {
		return root.Left.Val == root.Val && root.Val == root.Right.Val && isUnivalTree(root.Left) && isUnivalTree(root.Right)
	} else if root.Left != nil {
		return root.Val == root.Left.Val && isUnivalTree(root.Left)
	} else if root.Right != nil {
		return root.Val == root.Right.Val && isUnivalTree(root.Right)
	} else {
		return true
	}
}
