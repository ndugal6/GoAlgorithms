//Class for a doubly linked list
package program

import (
	// "fmt"
)
type Node struct {
	value int
	prev, next *Node
}

type DoublyLinkedList struct {
	head, tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) setHead(node *Node) {
	ll.setLRCouple(node.prev, node.next)
	if ll.head == nil {
		ll.setLRCouple(node, ll.tail)
		
	} else {
		ll.setLRCouple(node, ll.head)
	}
	node.prev = nil
	ll.head = node
}

func (ll *DoublyLinkedList) setTail(node *Node) {
	ll.setLRCouple(node.prev, node.next)
	if ll.tail == nil {
		ll.setLRCouple(ll.head, node)
	} else {
		ll.setLRCouple(ll.tail, node)
	}
	node.next = nil
	ll.tail = node
}
func (ll *DoublyLinkedList) insertBefore(node *Node, nodeToInsert *Node) {
	if node.prev == nil {
		ll.setHead(nodeToInsert)
	} else {
		ll.setLRCouple(node.prev, nodeToInsert)
		ll.setLRCouple(nodeToInsert, node)
	}
}
func (ll *DoublyLinkedList) setLRCouple(leftNode *Node, rightNode *Node) {
	if leftNode != nil {
		leftNode.next = rightNode;
	}
	if rightNode != nil {
		rightNode.prev = leftNode;
	}
}
func (ll *DoublyLinkedList) insertAfter(node *Node, nodeToInsert *Node) {
	if node.next == nil {
		ll.setTail(nodeToInsert)
	} else {
		ll.setLRCouple(nodeToInsert, node.next)
		ll.setLRCouple(node, nodeToInsert)
	}
}

func (ll *DoublyLinkedList) insertAtPosition(position int, nodeToInsert *Node) {
	if position < 2 {
		ll.setHead(nodeToInsert)
		return
	}
	for currentNode, count := ll.head, 1; currentNode != nil; currentNode, count = currentNode.next, count + 1 {
		if count == position {
			ll.insertBefore(currentNode, nodeToInsert)
			return
		}
	} 
	ll.setTail(nodeToInsert)
}

func (ll *DoublyLinkedList) removeNodesWithValue(value int) {
	for currentNode := ll.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.value == value {
			ll.remove(currentNode)
		}
	}
}

func (ll *DoublyLinkedList) remove(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		ll.setHead(node.next)
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		ll.setTail(node.prev)
	}	
}

func (ll *DoublyLinkedList) containsNodeWithValue(value int) bool {
	for currentNode := ll.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.value == value {
			return true
		}
	}
	return false
}

func (ll *DoublyLinkedList) getNodeWithValue(value int) *Node {
	for currentNode := ll.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.value == value {
			return currentNode
		}
	}
	return nil
}
func (ll *DoublyLinkedList) prettyPrint() []int{
	// fmt.Printf("\nHead: %d \n", ll.head.value)
	// fmt.Printf("Tail: %d \n", ll.tail.value)
	listed := make([]int, 0)
	for currentNode := ll.head; currentNode != nil; currentNode = currentNode.next {
		listed = append(listed, currentNode.value)
		// fmt.Printf("%d ", currentNode.value)
	}
	// fmt.Println()
	return listed
}