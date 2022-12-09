package linkedlist

// 142 Linked List Cycle II
func detectCycle(head *ListNode) *ListNode {
    if head==nil || head.Next==nil || head.Next.Next==nil{
		return nil
	}
	slow,fast:=head,head.Next.Next
	// for slow!=fast{
	// 	if slow==fast {
	// 		return slow
	// 	}
	// 	if fast.Next==nil ||fast.Next.Next==nil {
	// 		return nil
	// 	}
	// 	slow=slow.Next
	// 	fast=fast.Next.Next
	// }
	// return fast
	for fast.Next!=nil && fast.Next.Next!=nil{
		if slow==fast {
			for head!=slow{
				head=head.Next
				slow=slow.Next
			}
			return slow
		}
		slow=slow.Next
		fast=fast.Next.Next
	}
	return nil
}