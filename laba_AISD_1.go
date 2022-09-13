package main

import (
	"fmt"
)

//1.	добавление в конец списка DONE
//2.	добавление в начало списка DONE
//3.	удаление последнего элемента DONE
//4.	удаление первого элемента DONE
//5.	добавление элемента по индексу (вставка перед элементом, который был ранее доступен по этому индексу) DONE
//6.	получение элемента по индексу DONE
//7.	удаление элемента по индексу DONE
//8.	получение размера списка DONE
//9.	удаление всех элементов списка DONE wait
//10.	замена элемента по индексу на передаваемый элемент DONE
//11.	проверка на пустоту списка DONE
//14.  	вставка другого списка в конец DONE

type element struct { // element of list
	title string
	next  *element
}

type singlyLinkedList struct { // list which contains a head element, and lenß
	len  int
	head *element
}

func main() {
	list := new()

	list.AddToTop("Hello")
	list.printAllList() // OUTPUT : "Hello"

	list.AddToTop("Bye")
	list.printAllList() // OUTPUT : "Bye" "Hello"

	list.AddToEnd("Ok")
	list.printAllList() // OUTPUT : "Bye" "Hello" "Ok"

	list.RemoveLastElement()
	list.printAllList() // OUTPUT : "Bye" "Hello"

	list.RemoveFirstElement()
	list.printAllList() // OUTPUT : "Hello"

	list.insertElementByIndex(1, element{
		title: "Insert",
	})
	list.printAllList() // OUTPUT : "Hello" "Insert"

	list.insertElementByIndex(1, element{
		title: "New Insert",
	})
	list.printAllList() // OUTPUT : "Hello" "New Insert" "Insert"

	elem := list.getElementByIndex(1)
	fmt.Println(elem.title) // OUTPUT :  "New Insert"

	fmt.Println(list.Size()) // OUTPUT : 3

	list.deleteElementByIndex(1)
	list.printAllList() // OUTPUT : "Hello" "Insert"

	list.changeElementByIndex(1, element{title: "Changes insert"})
	list.printAllList() // OUTPUT : "Hello" "Changes insert"

	fmt.Print("IsEmpty: ", list.isEmpty(), "\n") //OUTPUT : IsEmpty: false

	newList := new()
	newList.AddToTop("New List 1")
	newList.AddToEnd("New List 2") // new list contains two elements

	list.mergeList(*newList)
	list.printAllList() // OUTPUT : "Hello" "Changes insert" "New List 1" "New List 2"

	list.deleteAll()
	fmt.Print("IsEmpty: ", list.isEmpty(), "\n") //OUTPUT : IsEmpty: true

}

func new() *singlyLinkedList { // constructor of list
	return &singlyLinkedList{}
}

func (list *singlyLinkedList) AddToTop(title string) { // func of adding element with title to top of list
	element := &element{ // create element with input title
		title: title,
	}

	if list.isEmpty() { // list isn't contain any elements
		list.head = element
	} else { // list contains something
		element.next = list.head
		list.head = element
	}
	list.len++ // we just added a new element at list, so we need to increase len

}

func (list *singlyLinkedList) AddToEnd(title string) { // func of adding element with title to the end of list
	element := &element{ // create element with input title
		title: title,
	}

	if list.isEmpty() { // list isn't contain any elements
		list.head = element
	} else { // list contains something
		current := list.head
		for current.next != nil { // while we have a next element, go to next
			current = current.next
		}
		current.next = element // find an end of list, so paste an element
	}

	list.len++

	return

}

func (list *singlyLinkedList) RemoveFirstElement() error {
	if list.isEmpty() { // if list is empty, we can't delete anything
		return fmt.Errorf("Can't remove first element. List is empty") // send error
	}

	list.head = list.head.next // find a first(head)element and pass this
	list.len--                 // we delete an element, so we need to decrease len

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

	list.len-- // we removed element, so we need to decrease len

	return nil
}

func (list *singlyLinkedList) Size() int {
	return list.len
}

func (list *singlyLinkedList) printAllList() error { // for beautiful output
	if list.isEmpty() { // if list is empty, it isn't contains anything
		return fmt.Errorf("printAllList: List is empty")
	}
	current := list.head // create a start position of list
	for current != nil {
		fmt.Print(current.title + " ") // print each element
		current = current.next
	}
	fmt.Println("")
	return nil
}

func (list *singlyLinkedList) getElementByIndex(index int) *element {
	head := list.head // create head element
	for i := 0; i < index; i++ {
		head = head.next // go to the next position
	}
	return head
}

func (list *singlyLinkedList) deleteElementByIndex(index int) error {
	var current *element // create element type

	if index < 0 || index > list.Size() {
		return fmt.Errorf("deleteElementByIndex : index number is invalid")
	}
	if index == 0 { // if we need to delete a first element of list
		current = list.getElementByIndex(index) // get current element by index
		list.head = current.next                // delete first element
	} else {
		current = list.getElementByIndex(index - 1) // get previous element by index
		current.next = current.next.next            // delete current element
	}
	list.len--
	return nil
}

func (list *singlyLinkedList) deleteAll() {
	list.head = nil
	list.len = 0
}

func (list *singlyLinkedList) changeElementByIndex(index int, newElement element) error {
	if index < 0 || index >= list.Size() {
		return fmt.Errorf("changeElementByIndex: index number is invalid.")
	}

	current := list.getElementByIndex(index) // getting element by index
	current.title = newElement.title         // change title of element\
	return nil
}

func (list *singlyLinkedList) isEmpty() bool {
	return list.Size() == 0
}

func (list *singlyLinkedList) mergeList(secondList singlyLinkedList) {
	tail := list.getElementByIndex(list.Size() - 1) // get 'tail' element of current list
	head := secondList.head                         // get 'head' element in new list
	tail.next = head                                //connect current list with new list
}

func (list *singlyLinkedList) insertElementByIndex(index int, element element) { // inserting before an element that was previously available by input index numberß
	if index < 0 || index > list.Size() { // check if index is valid for usß
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
		list.len++
	}

}
