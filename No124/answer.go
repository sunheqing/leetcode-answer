package main

import "fmt"

var tmpMax = -1001

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

// todo 二叉树的最大路径和
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeftPathSum := maxPathSum(root.Left)
	maxRightPathSum := maxPathSum(root.Right)
	if m := max([]int{maxLeftPathSum, maxRightPathSum, root.Val + maxLeftPathSum + maxRightPathSum}); m > tmpMax {
		tmpMax = m
	}
	return max([]int{root.Val, root.Val + maxLeftPathSum, root.Val + maxRightPathSum})
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
	root2 := &TreeNode{
		Val: -10,
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
	fmt.Println(maxPathSum(root2))
	fmt.Println(tmpMax)
	fmt.Println(maxPathSum(root))
}
