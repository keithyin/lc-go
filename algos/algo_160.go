package algos

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	headACnt := 0
	for cur := headA; cur != nil; cur = cur.Next {
		headACnt += 1
	}

	headBCnt := 0
	for cur := headB; cur != nil; cur = cur.Next {
		headBCnt += 1
	}

	cursorA := headA
	cursorB := headB

	if headACnt > headBCnt {
		for i := 0; i < (headACnt - headBCnt); i++ {
			cursorA = cursorA.Next
		}
	} else {
		for i := 0; i < (headBCnt - headACnt); i++ {
			cursorB = cursorB.Next
		}
	}

	for cursorA != nil && cursorB != nil {
		if cursorA == cursorB {
			return cursorA
		}
		cursorA = cursorA.Next
		cursorB = cursorB.Next
	}
	return nil
}
