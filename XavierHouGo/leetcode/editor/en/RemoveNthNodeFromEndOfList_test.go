package cn

import (
	"fmt"
	"testing"
)

func TestRemoveNthNodeFromEndOfList(t *testing.T) {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}

	t.Log(removeNthFromEnd(head, 1))
	fmt.Printf("%+v", removeNthFromEnd(head, 1))
}

//leetcode submit region begin(Prohibit modification and deletion)

//Definition for singly-linked list.
//type ListNode struct {
// Val int
// Next *ListNode
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	quick := head
	slow := &ListNode{0, head}

	for i := 0; i < n; i++ {

		quick = quick.Next
		if quick == nil {
			break
		}
	}

	for quick != nil {
		quick = quick.Next
		slow = slow.Next
	}

	if slow.Next == head {
		return head.Next
	}
	slow.Next = slow.Next.Next

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
