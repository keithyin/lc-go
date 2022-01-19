package algos

func detectCycle(head *ListNode) *ListNode {

	if head.Next == nil {
		return nil
	}
	slow := head.Next
	faster := head.Next.Next

	for slow != faster {
		if faster.Next == nil {
			return nil
		}
		slow = slow.Next
		faster = faster.Next.Next
	}

	cursor := head
	for slow != cursor {
		slow = slow.Next
		cursor = cursor.Next
	}

	return slow

}
