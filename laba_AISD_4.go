package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) GetMin() int {
	if t.left == nil {
		return t.val
	}
	return t.left.GetMin()
}

func (t *TreeNode) GetMax() int {
	if t.right == nil {
		return t.val
	}
	return t.right.GetMin()
}

func (t *TreeNode) PreOrder() {
	if t != nil {
		fmt.Print(t.val, " ")
		t.left.PreOrder()
		t.right.PreOrder()
	}
}

func (t *TreeNode) PostOrder() {
	if t != nil {
		t.left.PostOrder()
		t.right.PostOrder()
		fmt.Print(t.val, " ")
	}
}

func (t *TreeNode) MidOrder() {
	if t != nil {
		t.left.MidOrder()
		fmt.Print(t.val, " ")
		t.right.MidOrder()
	}
}

func (t *TreeNode) Find(value int) (TreeNode, bool) {
	if t == nil {
		return TreeNode{}, false
	}

	switch {
	case value == t.val:
		return *t, true
	case value < t.val:
		return t.left.Find(value)
	default:
		return t.right.Find(value)
	}
}

func (t *TreeNode) Delete(value int) *TreeNode {

	if t == nil {
		return nil
	}

	if value < t.val {
		t.left = t.left.Delete(value)
		return t
	}
	if value > t.val {
		t.right = t.right.Delete(value)
		return t
	}

	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}

	if t.left == nil {
		t = t.right
		return t
	}
	if t.right == nil {
		t = t.left
		return t
	}

	smallestValOnRight := t.right
	for {
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	t.val = smallestValOnRight.val
	t.right = t.right.Delete(t.val)
	return t
}

func (t *TreeNode) Insert(value int) error {

	if t == nil {

		return errors.New("Tree is nil")
	}

	if t.val == value {

		return errors.New("This node value already exists")
	}

	if t.val > value {

		if t.left == nil {

			t.left = &TreeNode{val: value}
			return nil
		}

		return t.left.Insert(value)
	}

	if t.val < value {

		if t.right == nil {

			t.right = &TreeNode{val: value}
			return nil
		}

		return t.right.Insert(value)
	}

	return nil
}

//func (t *TreeNode) GetNext(value int) int {
//	current := t
//	successor := TreeNode{}
//
//	if current.val > value {
//		successor = *current
//		current = current.left
//	} else {
//		current = current.right
//	}
//
//	return successor.val
//}

//func (t *TreeNode) Bfs(node TreeNode) []int {
//	queue := []*TreeNode{}
//	values := []int{}
//	queue = append(queue, &node)
//
//	for len(queue) > 0 {
//		var tempNode = queue[0]
//		queue = queue[1:]
//
//		values = append(values, tempNode.val)
//
//		if tempNode.left != nil {
//			queue = append(queue, tempNode.left)
//		}
//		if tempNode.right != nil {
//			queue = append(queue, tempNode.right)
//		}
//	}
//
//	return values
//}

func main() {
	t := &TreeNode{val: 8}

	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)

	fmt.Println("Pre")
	t.PreOrder()
	fmt.Println("\nPost")
	t.PostOrder()
	fmt.Println("\nMid")
	t.MidOrder()

}
