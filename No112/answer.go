package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 路径总和：树中是否存在根节点到叶子节点的路径，使得路径和与目标值相等
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	sum := root.Val
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		if stack[len(stack)-1].Left != nil && stack[len(stack)-1].Left.Val >= -1000 {
			sum += stack[len(stack)-1].Left.Val
			stack = append(stack, stack[len(stack)-1].Left)
		} else if stack[len(stack)-1].Right != nil && stack[len(stack)-1].Right.Val >= -1000 {
			sum += stack[len(stack)-1].Right.Val
			stack = append(stack, stack[len(stack)-1].Right)
		} else {
			if stack[len(stack)-1].Left == nil && stack[len(stack)-1].Right == nil && sum == targetSum { // 找到叶子节点表明找到一条完整路径
				return true
			}
			sum -= stack[len(stack)-1].Val
			stack[len(stack)-1].Val = -2000 // 标记处理过的节点
			stack = stack[:len(stack)-1]    // 弹出栈顶节点
		}
	}
	return false
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	root1 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root2 := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(hasPathSum(nil, 0))
	fmt.Println(hasPathSum(root, 4))
	fmt.Println(hasPathSum(root1, 3))
	fmt.Println(hasPathSum(root2, 37))
}
