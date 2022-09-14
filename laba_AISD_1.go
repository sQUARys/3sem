package main

import (
	"fmt"
)

type element struct { // element of list
	title string
	next  *element
}

type singlyLinkedList struct { // list which contains a head element, and len
	head *element
}

func main() {
	list := new()

	errorHandler(list.printAllList()) // OUTPUT : Handled error : printAllList: List is empty.

	fmt.Println("\nAddToTop func.")
	list.AddToTop("Hello")
	errorHandler(list.printAllList()) // OUTPUT : "Hello";

	fmt.Println("\nAddToTop func.")
	list.AddToTop("Bye")
	errorHandler(list.printAllList()) // OUTPUT : "Bye"; "Hello";

	fmt.Println("\nAddToEnd func.")
	list.AddToEnd("Ok")
	errorHandler(list.printAllList()) // OUTPUT : "Bye"; "Hello"; "Ok";

	fmt.Println("\nRemoveLastElement func.")
	errorHandler(list.RemoveLastElement())
	errorHandler(list.printAllList()) // OUTPUT : "Bye"; "Hello";

	fmt.Println("\nRemoveFirstElement func.")
	errorHandler(list.RemoveFirstElement())
	errorHandler(list.printAllList()) // OUTPUT : "Hello";

	fmt.Println("\ngetSize func.")
	list.insertElementByIndex(1, element{
		title: "Insert",
	})
	errorHandler(list.printAllList()) // OUTPUT : "Hello"; "Insert";

	fmt.Println("\ninsertElementByIndex func.")
	list.insertElementByIndex(1, element{
		title: "New Insert",
	})
	errorHandler(list.printAllList()) // OUTPUT : "Hello"; "New Insert"; "Insert";

	fmt.Println("\ngetElementByIndex func.")
	elem := list.getElementByIndex(1)
	fmt.Println(elem.title) // OUTPUT :  "New Insert"

	fmt.Println("\ngetSize func.")
	fmt.Println(list.getSize()) // OUTPUT : 3

	fmt.Println("\ndeleteElementByIndex func.")
	errorHandler(list.deleteElementByIndex(1))
	errorHandler(list.printAllList()) // OUTPUT : "Hello"; "Insert";

	fmt.Println("\nchangeElementByIndex func.")
	elementToUpdate := element{
		title: "Changes insert",
	}
	errorHandler(list.changeElementByIndex(1, elementToUpdate))
	errorHandler(list.printAllList()) // OUTPUT : "Hello"; "Changes insert"

	fmt.Print("\nIsEmpty: ", list.isEmpty(), "\n") //OUTPUT : IsEmpty: false

	newList := new()
	newList.AddToTop("New List 1")
	newList.AddToEnd("New List 2") // new list contains two elements

	fmt.Println("\nmergeList func.")
	list.mergeList(*newList)
	errorHandler(list.printAllList()) // OUTPUT : "Hello"; "Changes insert"; "New List 1"; "New List 2";

	fmt.Println("\ndeleteAll func.")
	list.deleteAll()
	fmt.Print("\nIsEmpty: ", list.isEmpty(), "\n") //OUTPUT : IsEmpty: true

}

func new() *singlyLinkedList { // constructor of list
	fmt.Println("Create new list.")
	return &singlyLinkedList{}
}

func (list *singlyLinkedList) AddToTop(title string) { // func of adding element with title to top of list
	elementToPaste := &element{ // create element with input title
		title: title,
	}

	if list.isEmpty() { // list isn't contain any elements
		list.head = elementToPaste
	} else { // list contains something
		elementToPaste.next = list.head
		list.head = elementToPaste
	}
}

func (list *singlyLinkedList) AddToEnd(title string) { // func of adding element with title to the end of list
	elementToPaste := &element{ // create element with input title
		title: title,
	}

	if list.isEmpty() { // list isn't contain any elements
		list.head = elementToPaste
	} else { // list contains something
		current := list.head
		for current.next != nil { // while we have a next element, go to next
			current = current.next
		}
		current.next = elementToPaste // find an end of list, so paste an element
	}
	return
}

func (list *singlyLinkedList) RemoveFirstElement() error {
	if list.isEmpty() { // if list is empty, we can't delete anything
		return fmt.Errorf("Can't remove first element. List is empty") // send error
	}

	list.head = list.head.next // find a first(head)element and pass this

	return nil
}

func (list *singlyLinkedList) RemoveLastElement() error {
	if list.isEmpty() { // if list is empty, we can't delete anything
		return fmt.Errorf("RemoveLastElement : List is empty")
	}

	var previous *element // create new memory element

	current := list.head // create a current element which start from head of list

	for current.next != nil { // go to the end of list
		previous = current
		current = current.next
	}

	if previous != nil { // if we go to the end of non-empty list
		previous.next = nil // remove last element
	} else { // if last element is nil, we stay at head, so we need to delete that
		list.head = nil
	}

	return nil
}

func (list *singlyLinkedList) getSize() int {
	size := 0
	for head := list.head; head != nil; head = head.next {
		size++
	}
	return size
}

func (list *singlyLinkedList) printAllList() error { // for beautiful output
	if list.isEmpty() { // if list is empty, it isn't contains anything
		return fmt.Errorf("printAllList: List is empty")
	}

	current := list.head // create a start position of list

	for current != nil {
		fmt.Print(current.title + "; ") // print each element
		current = current.next
	}

	fmt.Println("")
	return nil
}

func (list *singlyLinkedList) getElementByIndex(index int) *element {
	head := list.head // create head element
	for currentIndex := 0; currentIndex < index; currentIndex++ {
		head = head.next // go to the next position
	}
	return head
}

func (list *singlyLinkedList) deleteElementByIndex(index int) error {
	var current *element // create element type

	if index < 0 || index > list.getSize() {
		return fmt.Errorf("deleteElementByIndex : index number is invalid")
	}
	if index == 0 { // if we need to delete a first element of list
		current = list.getElementByIndex(index) // get current element by index
		list.head = current.next                // delete first element
	} else {
		current = list.getElementByIndex(index - 1) // get previous element by index
		current.next = current.next.next            // delete current element
	}
	return nil
}

func (list *singlyLinkedList) deleteAll() {
	list.head = nil
}

func (list *singlyLinkedList) changeElementByIndex(index int, newElement element) error {
	if index < 0 || index >= list.getSize() {
		return fmt.Errorf("changeElementByIndex: index number is invalid.")
	}

	current := list.getElementByIndex(index) // getting element by index
	current.title = newElement.title         // change title of element\
	return nil
}

func (list *singlyLinkedList) isEmpty() bool {
	return list.getSize() == 0
}

func (list *singlyLinkedList) mergeList(secondList singlyLinkedList) {
	tail := list.getElementByIndex(list.getSize() - 1) // get 'tail' element of current list
	head := secondList.head                            // get 'head' element in new list
	tail.next = head                                   //connect current list with new list
}

func (list *singlyLinkedList) insertElementByIndex(index int, element element) { // inserting before an element that was previously available by input index numberß
	if index < 0 || index > list.getSize() { // check if index is valid for usß
		fmt.Println("Index is invalid.")
	} else {
		if index == 0 { // if need to paste before first list's element
			element.next = list.head
			list.head = &element
		} else { // paste in a middle or end of list
			previousElement := list.getElementByIndex(index - 1)
			memory := previousElement.next
			previousElement.next = &element
			element.next = memory
		}
	}
}

func errorHandler(handledError error) { // for handling each error in the list like shell for error(middleware)
	outputFormat := "\nHandled error : %s. " // format of error
	if handledError != nil {                 // if we have an error
		fmt.Println(fmt.Errorf(outputFormat, handledError)) // out error
	}
}
