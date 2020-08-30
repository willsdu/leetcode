package primary

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	list := NewListNodeFromSlice([]int{9, 5, 4, 1})

	deleteNode(GetNodeByIndex(list, 1))

	PrintListNode(list)

}

func TestReverseLink(t *testing.T) {
	list := NewListNodeFromSlice([]int{9, 5, 4, 1})

	PrintListNode(reverseList(list))

}
