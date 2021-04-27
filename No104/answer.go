package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的最大深度（递归）
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftChildDepth := maxDepth(root.Left)
	rightChildDepth := maxDepth(root.Right)
	if leftChildDepth > rightChildDepth {
		return 1 + leftChildDepth
	}
	return 1 + rightChildDepth
}

// 二叉树的最大深度（非递归，借助一个队列）
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepthVal := 0
	queue := []*TreeNode{root, nil}
	for len(queue) > 0 {
		if queue[0] == nil { // 遇到层的结束标志位，树深加1
			maxDepthVal = maxDepthVal + 1
			queue = queue[1:]
			continue
		}
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
		if queue[0] == nil && len(queue) > 1 { // 当前层结束了，为下一层添加标志位
			queue = append(queue, nil)
		}
	}
	return maxDepthVal
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
	fmt.Println(maxDepth2(root))
}
