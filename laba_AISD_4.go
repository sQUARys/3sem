package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func main() {
	t := &TreeNode{val: 8}
	//
	//t.left = &TreeNode{val: 3}       //left subtree
	//t.right = &TreeNode{val: 3}      //right subtree
	//t.left.right = &TreeNode{val: 6} //right subtree of left subtree
	//t.right.left = &TreeNode{val: 5} //left subtree of the left subtree of the right subtree
	//t.left.left = &TreeNode{val: 1}
	//t.right.right = &TreeNode{val: 7}
	//t.right.right.right = &TreeNode{val: 8}
	//t.right.right.left = &TreeNode{val: 9}
	for i := 0; i < 15; i++ {
		t.Insert(rand.Int() % 50)
	}
	//fmt.Println("Pre")
	//t.PreOrder()
	//fmt.Println("\nPost")
	//t.PostOrder()
	//fmt.Println("\nMid")

	//fmt.Print(arr)
	//fmt.Println("Degree :", t.GetTreeDegree())
	t.MidOrder()
	t.GetPrev(49)
	t.GetNext(48)
	//visited, hashMap := t.bfs()
	//fmt.Println("Visited var : ", visited)
	//t.PrintTree()
}

func (t *TreeNode) GetNext(value int) {
	if _, ok := t.Find(value); ok {
		current := t
		next := &TreeNode{}
		for current != nil {
			if value < current.val {
				next = current
				current = current.left
			} else {
				current = current.right
			}
		}
		if next != nil {
			fmt.Println(fmt.Sprintf("Next element of %d is %d.", value, next.val))
		}
	}
}

func (t *TreeNode) GetPrev(value int) {
	if _, ok := t.Find(value); ok {
		current := t
		prev := &TreeNode{}
		for current != nil {
			if value > current.val {
				prev = current
				current = current.right
			} else {
				current = current.left
			}
		}
		if prev != nil {
			fmt.Println(fmt.Sprintf("Previous element of %d is %d.", value, prev.val))
		}
	}
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

func (t *TreeNode) bfs() (visited []TreeNode, hashMap map[int][]TreeNode) {
	hashMap = make(map[int][]TreeNode)
	visited = []TreeNode{}

	if t == nil {
		return
	}

	hashMap[0] = []TreeNode{*t}

	for i := 0; ; {
		nodeArr, ok := hashMap[i]
		if !ok {
			break
		}

		for _, node := range nodeArr {
			visited = append(visited, node)
			_, ok = hashMap[i+1]
			if !ok {
				hashMap[i+1] = []TreeNode{}
			}
			if node.left != nil {
				hashMap[i+1] = append(hashMap[i+1], *node.left)
			}
			if node.right != nil {
				hashMap[i+1] = append(hashMap[i+1], *node.right)
			}

		}
		i += 1
	}

	return
}

func (t *TreeNode) GetTreeDegree() int {
	maxDegree := 0

	if t == nil {
		return maxDegree
	}

	if t.left.GetTreeDegree() > t.right.GetTreeDegree() {
		maxDegree = t.left.GetTreeDegree()
	} else {
		maxDegree = t.right.GetTreeDegree()
	}

	return maxDegree + 1
}

//Printing part

func (t *TreeNode) PrintTree() {
	lines, _, _, _ := t.PrintTreeAux()
	for _, line := range lines {
		fmt.Println(line)
	}
}

func (t *TreeNode) PrintTreeAux() ([]string, int, int, int) {
	// no child
	if t.right == nil && t.left == nil {
		line := fmt.Sprintf("%d", t.val)
		width := len(line)
		height := 1
		middle := width / 2
		return []string{line}, width, height, middle
	}

	// only left child
	if t.right == nil {
		lines, n, p, x := t.left.PrintTreeAux()
		s := fmt.Sprintf("%d", t.val)
		u := len(s)
		firstLine := strings.Repeat(" ", x+1) + strings.Repeat("_", n-x-1) + s
		secondLine := strings.Repeat(" ", x) + "/" + strings.Repeat(" ", n-x-1+u)
		shiftedLines := []string{}
		for _, line := range lines {
			shiftedLines = append(shiftedLines, line+strings.Repeat(" ", u))
		}
		return append([]string{firstLine, secondLine}, shiftedLines...), n + u, p + 2, n + u/2
	}

	//only right child
	if t.left == nil {
		lines, n, p, x := t.right.PrintTreeAux()
		s := fmt.Sprintf("%d", t.val)
		u := len(s)
		firstLine := s + strings.Repeat("_", x) + strings.Repeat(" ", n-x)
		secondLine := strings.Repeat(" ", u+x) + "\\" + strings.Repeat(" ", n-x-1)
		shiftedLines := []string{}
		for _, line := range lines {
			shiftedLines = append(shiftedLines, strings.Repeat(" ", u)+line)
		}
		return append([]string{firstLine, secondLine}, shiftedLines...), n + u, p + 2, u / 2
	}

	// two child
	left, n, p, x := t.left.PrintTreeAux()
	right, m, q, y := t.right.PrintTreeAux()
	s := fmt.Sprintf("%d", t.val)
	u := len(s)
	firstLine := strings.Repeat(" ", x+1) + strings.Repeat("_", n-x-1) + s + strings.Repeat("_", y) + strings.Repeat(" ", m-y)
	secondLine := strings.Repeat(" ", x) + "/" + strings.Repeat(" ", n-x-1+u+y) + "\\" + strings.Repeat(" ", m-y-1)

	if p < q {
		for i := 0; i < q-p; i++ {
			left = append(left, strings.Repeat(" ", n))
		}
	} else if p > q {
		for i := 0; i < p-q; i++ {
			right = append(right, strings.Repeat(" ", m))
		}
	}

	zipLines := Zip(left, right)
	lines := []string{firstLine, secondLine}
	for _, b := range zipLines {
		lines = append(lines, b.First+strings.Repeat(" ", u)+b.Second)
	}
	return lines, n + m + u, Max(p, q) + 2, n + u/2
}

type Pair[T, U any] struct {
	First  T
	Second U
}

func Zip[T, U any](ts []T, us []U) []Pair[T, U] {
	if len(ts) != len(us) {
		panic("slices have different length")
	}
	pairs := make([]Pair[T, U], len(ts))
	for i := 0; i < len(ts); i++ {
		pairs[i] = Pair[T, U]{ts[i], us[i]}
	}
	return pairs
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
