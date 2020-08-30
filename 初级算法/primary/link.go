package primary

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

//请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。
//节点并非是最后一个节点
func deleteNode(node *ListNode) {
	if node.Next == nil {
		return
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func NewListNodeFromSlice(data []int) *ListNode {
	var head *ListNode
	var current *ListNode

	for i := 0; i < len(data); i++ {
		node := new(ListNode)
		node.Val = data[i]
		if head == nil {
			head = node
		}

		if current != nil {
			current.Next = node
		}
		current = node
	}
	return head
}

// 反转一个单链表。
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode

	current := head
	for current != nil {
		//保存将要移动的节点
		node := current
		//移动到下个节点
		current = current.Next
		//移动新的头
		node.Next = newHead
		newHead = node
	}
	return newHead
}

func PrintListNode(list *ListNode) {
	current := list
	for current != nil {
		log.Println(current.Val)
		if current.Next != nil {
			current = current.Next
		} else {
			break
		}

	}
}

//GetNodeByIndex  根据索引获取某个节点
func GetNodeByIndex(list *ListNode, index int) *ListNode {
	current := list
	i := 1

	for current != nil {
		if i == index {
			return current
		}
		i++
		current = current.Next
	}
	return nil
}
