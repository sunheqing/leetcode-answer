package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		if head.Val == val {
			return nil
		}
	} else {
		var preNode, node *ListNode
		preNode = nil
		node = head
		for {
			if node == nil {
				break
			}

			if node.Val == val {
				// 需要删除node节点
				if preNode == nil {
					// 删除头节点，设置新的头节点
					head = node.Next
				} else {
					// 删除非头节点
					preNode.Next = node.Next
				}
			} else {
				// 不需要删除node节点，移动preNode
				preNode = node
			}

			node = node.Next
		}
	}
	return head
}

func main() {
	head := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 6, Next: &ListNode{
					Val: 3, Next: &ListNode{
						Val: 4, Next: &ListNode{
							Val: 5, Next: &ListNode{
								Val: 6, Next: nil,
							},
						},
					},
				},
			},
		},
	}
	head = removeElements(head, 6)
	for {
		if head == nil {
			break
		}
		fmt.Println(head.Val)
		head = head.Next
	}
}
