package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var p *ListNode
	c := head

	for c != nil {
		n := c.Next
		c.Next = p
		p = c
		c = n
	}

	return p
}

func genList(arr []int) *ListNode {
	var head, prev *ListNode
	for _, v := range arr {
		tmp := ListNode{
			Val: v,
		}

		if head == nil {
			head = &tmp
			prev = head
			continue
		}

		prev.Next = &tmp
		prev = &tmp
	}

	return head
}

func main() {
	// 输入：head = [1,2,3,4,5]
	// 输出：[5,4,3,2,1]
	r := genList([]int{1, 2, 3, 4, 5})

	t := reverseList(r)
	fmt.Println(t, "t")

	for t != nil {
		fmt.Println(t.Val, t.Next, "t.Next")
		t = t.Next
	}
}
