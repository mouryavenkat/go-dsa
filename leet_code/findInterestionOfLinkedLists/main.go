package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var headALen, headBLen int
	tempHeadA := headA
	for tempHeadA != nil {
		headALen++
		tempHeadA = tempHeadA.Next
	}
	tempHeadB := headB
	for tempHeadB != nil {
		headBLen++
		tempHeadB = tempHeadB.Next
	}
	if headALen > headBLen {
		for i := 0; i < headALen-headBLen; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < headBLen-headALen; i++ {
			headB = headB.Next
		}
	}
	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
