package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 环形链表
// 快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
}

// 修改链表值为一个特殊值（破环了链表）
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	now := head
	for now != nil {
		if now.Val == 1234567890 {
			return true
		}
		now.Val = 1234567890
		now = now.Next
	}
	return false
}

// 所有遍历过的节点都指向头节点（破环了链表）
func hasCycle3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	now := head
	for now != nil {
		if now.Next == head {
			return true
		}
		tmp := now.Next
		now.Next = head
		now = tmp
	}
	return false
}

func main() {

}
