package main

import "fmt"

type MyLinkedList struct {
	index int
	val   int
	next  *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{-1, 0, nil}
}

func (this *MyLinkedList) Get(index int) int {
	result := -1
	if this.index == -1 {
		return result
	}
	if this.index == index {
		return this.val
	}
	if this.next == nil {
		return result
	}
	result = this.next.Get(index)
	return result
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.index == -1 {
		this.index = 0
		this.val = val
		return
	}
	old := *this
	*this = MyLinkedList{0, val, &old}
	for currentNode := this.next; currentNode != nil; currentNode = currentNode.next {
		currentNode.index++
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.index == -1 {
		this.AddAtHead(val)
		return
	}
	for currentNode := this; currentNode != nil; currentNode = currentNode.next {
		if currentNode.next == nil {
			currentNode.next = &MyLinkedList{currentNode.index + 1, val, nil}
			return
		}
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if this.index == -1 {
		return
	}
	for currentNode := this; currentNode != nil; currentNode = currentNode.next {
		if currentNode.next == nil {
			if currentNode.index == index-1 {
				this.AddAtTail(val)
			}
			return
		}
		if currentNode.next.index == index {
			currentNode.next = &MyLinkedList{index, val, currentNode.next}
			for newCurrentNode := currentNode.next.next; newCurrentNode != nil; newCurrentNode = newCurrentNode.next {
				newCurrentNode.index++
			}
			return
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.index == -1 {
		return
	}
	if index == 0 {
		if this.next == nil {
			this.index = -1
			return
		}
		*this = *this.next
		for currentNode := this; currentNode != nil; currentNode = currentNode.next {
			currentNode.index--
		}
		return
	}
	old := this
	for currentNode := this.next; currentNode != nil; currentNode = currentNode.next {
		if currentNode.index == index {
			old.next = currentNode.next
			for newCurrentNode := old.next; newCurrentNode != nil; newCurrentNode = newCurrentNode.next {
				newCurrentNode.index--
			}
			return
		}
		old = old.next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	val := list.Get(1)
	list.DeleteAtIndex(0)
	val = list.Get(0)
	fmt.Printf("%d", val)
}
