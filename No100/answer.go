package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 判断两棵树是否相等（递归）
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil || q == nil) || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

type doubleNode struct {
	pNode *TreeNode
	qNode *TreeNode
}

// 判断两棵树是否相等（非递归）
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	queue := []*doubleNode{&doubleNode{pNode: p, qNode: q}}
	for len(queue) > 0 {
		pi, qi := queue[0].pNode, queue[0].qNode
		queue = queue[1:]
		if pi == nil && qi == nil {
			continue
		}
		if (pi == nil || qi == nil) || pi.Val != qi.Val {
			return false
		}
		queue = append(queue, &doubleNode{pNode: pi.Left, qNode: qi.Left})
		queue = append(queue, &doubleNode{pNode: pi.Right, qNode: qi.Right})
	}
	return true
}

func main() {
	root1 := &TreeNode{
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
	root2 := &TreeNode{
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
	fmt.Println(isSameTree(root1, root2))
	fmt.Println(isSameTree2(root1, root2))
}
