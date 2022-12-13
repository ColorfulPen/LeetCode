package linkedlist

import (
	"fmt"
	"testing"
)

func Test142(t *testing.T) {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 1, Next: a}
	c := &ListNode{Val: 1, Next: b}
	fmt.Println(detectCycle(c))
}
