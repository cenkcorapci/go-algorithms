package main

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	if position == 0 {
		return &SinglyLinkedListNode{data: data, next: llist}
	}
	cursor := llist
	for position > 1 {
		position -= 1
		cursor = cursor.next
	}
	insert := SinglyLinkedListNode{data: data, next: cursor.next}
	cursor.next = &insert
	return llist
}
