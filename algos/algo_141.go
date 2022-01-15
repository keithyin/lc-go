package algos

func hasCycle(head *ListNode) bool {

	slow := head
	fast := head

	for slow != nil && fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next

		if slow == fast {
			return true
		}

	}

	return false

}
