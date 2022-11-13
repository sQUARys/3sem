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
	t := &TreeNode{}
	tmp := 0
	for i := 0; i < 15; i++ {
		val := rand.Int() % 50
		t.Insert(val)
		if i == 7 {
			tmp = val // to work with one number to getprev , getnext , delete
		}
	}

	fmt.Println("Our tree : ")
	t.PrintTree()
	fmt.Println()

	fmt.Println("Pre")
	t.PreOrder()
	fmt.Println("\n\nPost")
	t.PostOrder()
	fmt.Println("\n\nMid")
	t.MidOrder()
	fmt.Println("\n\nGet Min and Get Max:")
	fmt.Println("Min var in tree :", t.GetMin())
	fmt.Println("Max var in tree :", t.GetMax())
	fmt.Println("\n\nGet Prev and Get Next:")
	t.GetPrev(tmp)
	t.GetNext(tmp)

	fmt.Println("\nDegree of tree:", t.GetTreeDegree())

	visited, _ := t.bfs() // second return value is a hashmap which contains all info about each level of binary tree
	fmt.Println("\nVisiting all variables by bfs algo: ", visited)

	fmt.Println("\nIs exist before deleting", tmp, "value : ", t.Find(tmp))
	t.Delete(tmp)

	fmt.Println("\nIs exist after deleting", tmp, "value : ", t.Find(tmp))
	fmt.Println("\nBinary tree after deleting ", tmp)
	t.PrintTree()
}

func (t *TreeNode) GetNext(value int) {
	if t.Find(value) {
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
	if t.Find(value) {
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

func (t *TreeNode) Find(value int) bool {
	if t == nil {
		return false
	}

	switch {
	case value == t.val:
		return true
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

func (t *TreeNode) bfs() (visited []int, hashMap map[int][]TreeNode) {
	hashMap = make(map[int][]TreeNode)
	visited = []int{}

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
			visited = append(visited, node.val)
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
