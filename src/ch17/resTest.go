package ch17

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	le := right - left + 1
	var nodes = make([]int, le)
	var i, tail, p = 1, le - 1, head
	for p != nil {
		if left <= i && right >= i {
			nodes[tail] = p.Val
			tail--
		} else if i > right {
			break
		}
		i++
		p = p.Next
	}
	p = head
	i = 1
	for p != nil {
		if left <= i && right >= i {
			p.Val = nodes[i-left]
		} else if i > right {
			break
		}
		p = p.Next
		i++
	}
	return head
}
