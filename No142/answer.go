package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 环形链表II
// 所有遍历过的节点都指向头节点（破环了链表）
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	now := head
	for now != nil {
		if now.Next == head {
			return now
		}
		tmp := now.Next
		now.Next = head
		now = tmp
	}
	return nil
}

func main() {
	head := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}
	head.Next.Next.Next.Next = head.Next
	fmt.Println(detectCycle(head))
}
