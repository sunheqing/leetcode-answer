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
		if root.Left != nil { // 左子树非空时
			tmpRight := root.Right // 暂存当前节点的右子树的根节点
			root.Right = root.Left // 将当前节点的左子树移动至树的右节点
			root.Left = nil
			tmpNode := root.Right
			for tmpNode.Right != nil { // 找到当前树的最右节点
				tmpNode = tmpNode.Right
			}
			tmpNode.Right = tmpRight // 将暂存的右子树的根节点接入到当前树的最右节点
		}
		root = root.Right // 移动当前节点
	}
}

func main() {

}
