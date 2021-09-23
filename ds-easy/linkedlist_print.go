package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func printLinkedList(head *SinglyLinkedListNode) {
	cursor := head
	for cursor != nil {
		d := *cursor
		fmt.Println(d.data)
		cursor = d.next
	}
}
