package program

import (
	"testing"
	"fmt"
)

func createDoubleLinkedListFromInts(values []int) *DoublyLinkedList {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.setHead(&Node{
		value: values[0],
	})
	for _, newValue := range values[1:] {
		doublyLinkedList.setTail(&Node{
			value: newValue,
		})
	}
	return doublyLinkedList
}
func TestDoublyLinkedList(t *testing.T) {
	list := createDoubleLinkedListFromInts([]int{1,2,3,4,5})
	testNode1 := Node{
		value: 9999,
	}
	testNode2 := Node{
		value: 1111,
	}

	// testNode3 := Node{
	// 	value: 1111,
	// }
	list.setHead(&testNode1)
	list.setTail(&testNode2)
	// list.prettyPrint()
	// list.setHead(list.head.next.next.next)
	returnedList := list.prettyPrint()
	fmt.Printf("this is the returned list, %v", returnedList)
	newList := []int{9999, 1, 2, 3,4,5,1111}
	fmt.Printf("this is the returned list, %v", newList)
	fmt.Printf("\nDo they equal: %v", returnedList == newList)
	// list.insertAtPosition(1, &testNode3)
	// list.prettyPrint()
	// list.setTail(&testNode1)
	// list.removeNodesWithValue(9999)
	list.removeNodesWithValue(1111)
	list.prettyPrint()
}