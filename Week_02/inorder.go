package Week_02

// 二叉树中序遍历
var r []int
func inorderTraversal(root *TreeNode) []int {
	// r := []int{}
	traversal(root)
	return r
}

// 递归 左根右
func traversal(root *TreeNode) {
	if root != nil {
		if root.Left != nil {
			traversal(root.Left)
		}
		r = append(r, root.Val)
		if root.Right != nil {
			traversal(root.Right)
		}
	}
}
