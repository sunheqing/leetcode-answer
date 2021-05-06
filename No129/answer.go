package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 求根节点到叶节点数字之和
// 非递归，借助栈实现
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		if stack[len(stack)-1].Left != nil && stack[len(stack)-1].Left.Val >= 0 {
			stack = append(stack, stack[len(stack)-1].Left)
		} else if stack[len(stack)-1].Right != nil && stack[len(stack)-1].Right.Val >= 0 {
			stack = append(stack, stack[len(stack)-1].Right)
		} else {
			if stack[len(stack)-1].Left == nil && stack[len(stack)-1].Right == nil { // 找到叶子节点表明找到一条完整路径
				for i := 0; i < len(stack); i++ {
					sum += stack[i].Val * int(math.Pow(10, float64(len(stack)-1-i)))

				}
			}
			stack[len(stack)-1].Val = -1 // 标记处理过的节点
			stack = stack[:len(stack)-1] // 弹出栈顶节点
		}
	}
	return sum
}

// 非递归，借助栈实现，改进了求和方式
func sumNumbers2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	oneSum := root.Val
	sum := 0
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		if stack[len(stack)-1].Left != nil && stack[len(stack)-1].Left.Val >= 0 {
			oneSum = 10*oneSum + stack[len(stack)-1].Left.Val
			stack = append(stack, stack[len(stack)-1].Left)
		} else if stack[len(stack)-1].Right != nil && stack[len(stack)-1].Right.Val >= 0 {
			oneSum = 10*oneSum + stack[len(stack)-1].Right.Val
			stack = append(stack, stack[len(stack)-1].Right)
		} else {
			if stack[len(stack)-1].Left == nil && stack[len(stack)-1].Right == nil { // 找到叶子节点表明找到一条完整路径
				sum += oneSum
			}
			oneSum = (oneSum - stack[len(stack)-1].Val) / 10
			stack[len(stack)-1].Val = -1 // 标记处理过的节点
			stack = stack[:len(stack)-1] // 弹出栈顶节点
		}
	}
	return sum
}

type NewTreeNode struct {
	Val    int
	Father *NewTreeNode
	Left   *NewTreeNode
	Right  *NewTreeNode
}

// 如果每个根节点可以保存父节点指针时：非递归，不借助其他数据结构
func sumNumbers3(root *NewTreeNode) int {
	if root == nil {
		return 0
	}
	oneSum := root.Val
	sum := 0
	for root != nil {
		if root.Left != nil && root.Left.Val >= 0 {
			oneSum = 10*oneSum + root.Left.Val
			root.Left.Father = root
			root = root.Left
		} else if root.Right != nil && root.Right.Val >= 0 {
			oneSum = 10*oneSum + root.Right.Val
			root.Right.Father = root
			root = root.Right
		} else {
			if root.Left == nil && root.Right == nil { // 找到叶子节点表明找到一条完整路径
				sum += oneSum
			}
			oneSum = (oneSum - root.Val) / 10
			root.Val = -1      // 标记处理过的节点
			root = root.Father // 回退到父节点
		}
	}
	return sum
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
	fmt.Println(sumNumbers(nil))
	fmt.Println(sumNumbers(root))
	fmt.Println(sumNumbers(root1))
	fmt.Println(sumNumbers(root2))

	fmt.Println(sumNumbers2(nil))
	fmt.Println(sumNumbers2(root))
	fmt.Println(sumNumbers2(root1))
	fmt.Println(sumNumbers2(root2))

	root11 := &NewTreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root22 := &NewTreeNode{
		Val: 10,
		Left: &NewTreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &NewTreeNode{
			Val: 20,
			Left: &NewTreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &NewTreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(sumNumbers3(root11))
	fmt.Println(sumNumbers3(root22))
}
