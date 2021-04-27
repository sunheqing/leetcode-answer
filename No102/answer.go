package main

import "fmt"

// 二叉树的层序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 借助两个队列
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var node *TreeNode
	result := [][]int{[]int{}}
	queue1 := []*TreeNode{root}
	queue2 := []*TreeNode{}

	for len(queue1) > 0 {
		// 取队列1队列顶节点
		node = queue1[0]
		queue1 = queue1[1:]
		// 保存当前层的结果
		result[len(result)-1] = append(result[len(result)-1], node.Val)
		// 下层节点入队列2
		if node.Left != nil {
			queue2 = append(queue2, node.Left)
		}
		if node.Right != nil {
			queue2 = append(queue2, node.Right)
		}
		// 当前队列空且下层队列非空，交换当前队列和下层队列，为下层队列结果预留位置
		if len(queue1) == 0 && len(queue2) > 0 {
			queue1, queue2 = queue2, queue1
			result = append(result, []int{})
		}
	}

	return result
}

// 借助一个队列
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{[]int{}} // nil为每层结束的标志位
	queue := []*TreeNode{root, nil}
	for len(queue) > 0 {
		result[len(result)-1] = append(result[len(result)-1], queue[0].Val) // 保存当前层的结果
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
		if queue[0] == nil { // 当前层结束了
			if len(queue) > 1 { // 还存在下一层待处理，为下一层添加标志位以及存放结果的数组
				queue = append(queue, nil)
				result = append(result, []int{})
			}
			queue = queue[1:] // 去掉当前层标志位
		}
	}
	return result
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   30,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  6,
			Left: nil,
			Right: &TreeNode{
				Val:   31,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(levelOrder(root))
	fmt.Println(levelOrder2(root))
}
