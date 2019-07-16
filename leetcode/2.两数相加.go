//__author__ = "YaoYao"
//Date: 2019-07-15
package leetcode

import "fmt"

//import (
//	"fmt"
//)

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
//

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head = &ListNode{0, nil}
	p, q, curr := l1, l2, head
	carry := 0
	for p != nil || q != nil {
		x := 0
		y := 0
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		sum := x + y + carry
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
		if p != nil {
			p = p.Next
		} else {
			p = nil
		}
		if q != nil {
			q = q.Next
		} else {
			q = nil
		}
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
	}
	if carry == 1 {
		curr.Next = &ListNode{carry, nil}
	}
	return head.Next
}

func Test() {
	l1 := &ListNode{2, nil}
	l11 := l1
	l11.Next = &ListNode{4, nil}
	l11 = l11.Next
	l11.Next = &ListNode{8, nil}
	//
	l2 := &ListNode{5, nil}
	l22 := l2
	l22.Next = &ListNode{6, nil}
	l22 = l22.Next
	l22.Next = &ListNode{4, nil}
	//
	res := addTwoNumbers(l1, l2)
	fmt.Println(res.Val)
}
