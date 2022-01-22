package algos

func SortList(head *ListNode) *ListNode {
	return sortList(head)
}

func sortList(head *ListNode) *ListNode {
	return sortListCore(head)
}

func sortListCore(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	part1, part2 := splitList(head)
	part1 = sortListCore(part1)
	part2 = sortListCore(part2)
	return mergeListNode(part1, part2)
}

func splitList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	slowPre := head
	slow := head
	fast := head
	for slow != nil && fast != nil {
		if fast.Next == nil {
			part2 := slow
			slowPre.Next = nil
			return head, part2
		}
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	part2 := slow
	slowPre.Next = nil
	return head, part2

}

func mergeListNode(head1, head2 *ListNode) *ListNode {
	root := new(ListNode)
	cursor := root
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cursor.Next = head1
			head1 = head1.Next
			cursor = cursor.Next
		} else {
			cursor.Next = head2
			head2 = head2.Next
			cursor = cursor.Next
		}
	}
	for head1 != nil {
		cursor.Next = head1
		head1 = head1.Next
		cursor = cursor.Next
	}

	for head2 != nil {
		cursor.Next = head2
		head2 = head2.Next
		cursor = cursor.Next
	}
	return root.Next
}
