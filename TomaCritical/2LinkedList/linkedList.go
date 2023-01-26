package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 203 Remove Linked List Elements
func removeElements(head *ListNode, val int) *ListNode {
	// PS:这里给头节点设置一个空节点会比较好
	// 找到第一个节点
	for head != nil {
		if head.Val != val {
			break
		}
		head = head.Next
	}
	// 0个节点
	if head == nil {
		return nil
	}
	// 1个节点
	if head.Next == nil {
		if head.Val == val {
			return nil
		} else {
			return head
		}
	}
	// 2个节点及以上
	p := head
	q := p.Next
	for q.Next != nil {
		if q.Val == val {
			q = q.Next
			p.Next = q
		} else {
			p = p.Next
			q = q.Next
		}
	}
	if q.Val == val {
		p.Next = nil
	}
	return head
}

// 206 Reverse Linked List
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	q := head.Next
	r := q.Next
	p.Next = nil

	for r != nil {
		q.Next = p
		p = q
		q = r
		r = r.Next
	}
	q.Next = p
	return q
}

// 24 Swap Nodes in Pairs
func swapPairs(head *ListNode) *ListNode {
	// 0 or 1
	if head == nil || head.Next == nil {
		return head
	}
	p, q, r := head, head.Next, head.Next.Next
	head = q
	for r != nil && r.Next != nil {
		q.Next = p
		p.Next = r.Next

		p = r
		q = r.Next
		r = r.Next.Next
	}
	p.Next = r
	q.Next = p
	return head
}

// 19 Remove Nth Node From End of List
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 若为一个节点
	if head.Next == nil {
		return nil
	}
	// virtual head
	virtualHead := &ListNode{Next: head}
	r := &ListNode{Next: virtualHead}
	p, q := virtualHead, virtualHead
	realHead := p
	for i := 0; i < n; i++ {
		q = q.Next
	}
	for q != nil {
		r = r.Next
		p, q = p.Next, q.Next
	}
	r.Next = p.Next
	return realHead.Next
}

// 面试题 02.07. Intersection of Two Linked Lists LCCI
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	p := headA
	for p != nil {
		p = p.Next
		lenA++
	}
	p = headB
	for p != nil {
		p = p.Next
		lenB++
	}
	var q *ListNode
	var diff int
	if lenA > lenB {
		diff = lenA - lenB
		p = headB
		q = headA
	} else {
		diff = lenB - lenA
		p = headA
		q = headB
	}
	for diff > 0 {
		diff--
		q = q.Next
	}
	for p != nil {
		if p == q {
			return p
		}
		p, q = p.Next, q.Next
	}
	return nil
}


