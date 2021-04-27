package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的最小深度，根节点到最近的叶子节点（递归）
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right != nil {
		leftChildDepth := minDepth(root.Left)
		rightChildDepth := minDepth(root.Right)
		if leftChildDepth < rightChildDepth {
			return 1 + leftChildDepth
		}
		return 1 + rightChildDepth
	}
	if root.Left != nil {
		return 1 + minDepth(root.Left)
	}
	if root.Right != nil {
		return 1 + minDepth(root.Right)
	}
	return 1
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  311,
				Left: nil,
				Right: &TreeNode{
					Val: 22,
					Left: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  6,
			Left: nil,
			Right: &TreeNode{
				Val: 31,
				Left: &TreeNode{
					Val:   312,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
	}
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 311,
				Left: &TreeNode{
					Val: 22,
					Left: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   17,
				Left:  nil,
				Right: nil,
			},
		},
		Right: nil,
	}
	fmt.Println(minDepth(root))
	fmt.Println(minDepth(root2))
}
