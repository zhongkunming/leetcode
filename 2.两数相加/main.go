package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
两个链表都是逆序存储的，可以通过直接遍历将两个相同位置的值相加

时间复杂度 O(max(l1, l2))
空间复杂度 O(1)
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 声明头指针和尾指针
	var head, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func main() {
	l13 := &ListNode{Val: 3}
	l12 := &ListNode{Val: 6, Next: l13}
	l11 := &ListNode{Val: 2, Next: l12}

	l23 := &ListNode{Val: 3}
	l22 := &ListNode{Val: 4, Next: l23}
	l21 := &ListNode{Val: 2, Next: l22}
	l20 := &ListNode{Val: 5, Next: l21}
	result := addTwoNumbers(l11, l20)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
