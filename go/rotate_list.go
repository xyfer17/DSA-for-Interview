/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {

	// return head if below condition fulfilled
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	curr := head
	listLen := 0

	for curr != nil {
		curr = curr.Next
		listLen++
	}

	breakPoint := listLen - k%listLen
	if breakPoint == listLen {
		return head
	}
	curr = head

	for i := 0; i < breakPoint-1; i++ {
		curr = curr.Next
	}

	newHead := curr.Next
	curr.Next = nil

	curr = newHead
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = head

	return newHead
}
