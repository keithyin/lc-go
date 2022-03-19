package algos

func rotateRight(head *ListNode, k int) *ListNode {
	listLen := 0
	cursor := head
	for cursor != nil {
		listLen += 1
		cursor = cursor.Next
	}
	k = k % listLen
	if k == 0 {
		return head
	}

	pre := head
	cur := head
	curIdx := 0
	for cur.Next != nil {
		cur = cur.Next
		curIdx++
		if curIdx > k {
			pre = pre.Next
		}
	}
	fakeHead := new(ListNode)
	fakeHead.Next = pre.Next
	pre.Next = nil
	cur.Next = head
	return fakeHead.Next

}
