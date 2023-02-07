package answer1145

import "leetcode/utils/structures"

type TreeNode = structures.TreeNode

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	node := findNode(root, x)

	leftCount := nodeCount(node.Left)
	if leftCount > n/2 {
		return true
	}

	rightCount := nodeCount(node.Right)
	if rightCount > n/2 {
		return true
	}

	parentCount := n - leftCount - rightCount - 1
	return parentCount > n/2
}

func findNode(root *TreeNode, x int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == x {
		return root
	}
	if node := findNode(root.Left, x); node != nil {
		return node
	}
	return findNode(root.Right, x)
}

func nodeCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + nodeCount(root.Left) + nodeCount(root.Right)
}
