package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树展开为链表，按照先序遍历顺序，节点依然为树节点（原地修改）
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	for root != nil {

	}
}
