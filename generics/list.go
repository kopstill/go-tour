package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	next *List[T]
	val  T
}

func (l List[T]) String() string {
	return fmt.Sprintf("next:: %v | val:: %v", l.next, l.val)
}

// if nothing next, returns 'zero' value of T.
func (l List[T]) get() T {
	return l.val
}

func (l *List[T]) insert(newVal T) {
	copyList := new(List[T])
	copyNext := l.next
	copyVal := l.val
	copyList.next = copyNext
	copyList.val = copyVal
	l.next = copyList
	l.val = newVal
}

func (l *List[T]) remove() {
	// 5 lines to do a deep copy.
	copyList := new(List[T])
	copyNext := l.next
	copyVal := l.val
	copyList.next = copyNext
	copyList.val = copyVal
	l.val = copyList.next.val
	l.next = copyList.next.next
}

func main() {
	var myLinkedList = new(List[string])
	fmt.Println(myLinkedList)
	fmt.Println(myLinkedList.get())
	fmt.Println("-------")

	myLinkedList.insert("hyerin")
	fmt.Println(myLinkedList)
	fmt.Println(myLinkedList.get())
	fmt.Println("-------")

	myLinkedList.remove()
	fmt.Println(myLinkedList)
	fmt.Println(myLinkedList.get())
	fmt.Println("-------")

	myLinkedList.insert("hyerin")
	fmt.Println(myLinkedList)
	fmt.Println(myLinkedList.get())
	fmt.Println("-------")

	myLinkedList.insert("emile")
	fmt.Println(myLinkedList)
	fmt.Println(myLinkedList.get())
	fmt.Println("-------")

	myLinkedList.remove()
	fmt.Println(myLinkedList)
	fmt.Println(myLinkedList.get())
}
