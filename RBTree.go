package main

import "fmt"

/** 1. A red-black tree is a binary search tree
 *  2. Every Link is either red and black
 *  3. Every child that is an empty tree is a leaf with black link to its parent
 *  4. No path from the root to any leaf has two consective red links
 *  5. All path from each node to all its leaves have the same number of black links
 *  6. All red link is to a left child.
 *  7. A left-leaning red-black tree is a red-black tree.
 */

type RBTree struct {
	root *RBTreeNode
}
type RBTreeNode struct {
	left  *RBTreeNode
	right *RBTreeNode
	key   int
	color bool //the color of link to parent, false is red
}

func (t *RBTree) insert(key int) {
	t.root = t.root.insert(key)
	t.root.color = true //BLACK
}
func (t *RBTree) contains(key int) bool {
	if t == nil {
		return false
	}
	if t.root.contains(key) != nil {
		return true
	}
	return false
}
func (t *RBTree) remove(key int) {
	if t.contains(key) {
		t.root = t.root.remove(key)
		if t.root != nil {
			t.root.color = true //BLACK
		}
	}
}

func (node *RBTreeNode) isRed() bool {
	if node == nil {
		return false
	}
	return node.color == false //RED
}

func (node *RBTreeNode) insert(key int) *RBTreeNode {
	if node == nil {
		// fmt.Println("nil node with key", key)
		return &RBTreeNode{key: key}
	}
	if key > node.key {
		// fmt.Println("insert right with key", key)
		node.right = node.right.insert(key)
	} else if key < node.key {
		// fmt.Println("insert left with key", key)
		node.left = node.left.insert(key)
	}
	return node.fixup()
}

func (node *RBTreeNode) fixup() *RBTreeNode {
	//precondition node is not nil
	//fix the red black tree violatioin 4,6
	if node.right.isRed() {
		node = node.rotateLeft()
	}
	if node.left.isRed() && node.left.left.isRed() {
		//red node will never be nil
		node = node.rotateRight()
	}
	if node.left.isRed() && node.right.isRed() {
		node.flipColor()
	}
	return node
}

func (node *RBTreeNode) rotateLeft() *RBTreeNode {
	//precondition node is not nil
	//precondition node.right is red
	//postcondition p is returned with modified color
	//postcondition p.left is red
	h := node
	p := node.right
	h.right = p.left
	p.left = h
	p.color = h.color // maintain the black height
	h.color = false   //RED
	return p
}

func (node *RBTreeNode) rotateRight() *RBTreeNode {
	h := node
	p := node.left
	h.left = p.right
	p.right = h
	p.color = h.color // maintain the black height
	h.color = false   //RED
	return p
}

func (node *RBTreeNode) flipColor() {
	// node, left and right is non empty
	node.color = !node.color
	node.left.color = !node.left.color
	node.right.color = !node.right.color
}

func (node *RBTreeNode) remove(key int) *RBTreeNode {
	// Precondition: key will be valid
	// Step 1: find the node to remove
	// Step 2: find successor node to remove (min of right subtree)
	// Step 3: replace the node with successor's data
	// Step 4: remove successor node
	if key < node.key { // go left subtree
		//push red link down
		if !node.left.isRed() && !node.left.left.isRed() {
			node = node.leftSidePushDownRed()
		}
		node.left = node.left.remove(key)
	} else { //found key and/or deal with right substree
		if node.left.isRed() {
			node = node.rotateRight() //temporary break the constraint
		}
		if key == node.key && node.right == nil {
			node = nil
			return nil
		}
		// deal with the case where node.right is not nil
		if !node.right.isRed() && !node.right.left.isRed() {
			node = node.rightSidePushDownRed()
		}
		if key == node.key {
			node.key = node.right.minData()
			node.right = node.right.removeMin()
		} else {
			node.right = node.right.remove(key)
		}
	}
	return node.fixup()
}
func (node *RBTreeNode) minData() int {
	//precondition node is not nil
	for node.left != nil {
		node = node.left // return node.left.minData() // recursive
	}
	return node.key
}
func (node *RBTreeNode) removeMin() *RBTreeNode {
	//precondition node is not nil
	if node.left == nil { //reach minimum
		node = nil // llrb do no have right child at min so black height is true
		return nil
	}
	if !node.left.isRed() && !node.left.left.isRed() {
		node = node.leftSidePushDownRed()
	}
	node.left = node.left.removeMin()
	return node.fixup()
}

func (node *RBTreeNode) leftSidePushDownRed() *RBTreeNode {
	node.flipColor()
	if node.right.left.isRed() {
		node.right = node.right.rotateRight()
		node = node.rotateLeft()
		node.flipColor()
	}
	return node
}
func (node *RBTreeNode) rightSidePushDownRed() *RBTreeNode {
	node.flipColor()
	if node.left.left.isRed() {
		node = node.rotateRight()
		node.flipColor()
	}
	return node
}
func (node *RBTreeNode) contains(key int) *RBTreeNode {
	if node == nil {
		return nil
	}
	if node.key == key {
		return node
	} else if key > node.key {
		return node.right.contains(key)
	} else {
		return node.left.contains(key)
	}
}

func (t *RBTreeNode) print() {
	if t == nil {
		return
	}
	t.left.print()
	fmt.Println(t.key)
	t.right.print()
}
func (node *RBTreeNode) pprint(indent int) {
	if node == nil {
		return
	}
	indent += 7
	node.right.pprint(indent)
	fmt.Println()
	for i := 7; i < indent; i++ {
		fmt.Print(" ")
	}
	if node.color {
		fmt.Println("Blk", node.key)
	} else {
		fmt.Println("Red", node.key)
	}
	node.left.pprint(indent)
}
func (t *RBTree) print() {
	if t == nil {
		return
	}
	t.root.pprint(0)
}
func main() {
	tree := new(RBTree) //uninitialized,root is nil
	tree.insert(9)
	tree.insert(8)
	tree.insert(7)
	tree.insert(6)
	tree.insert(3)
	fmt.Println(t.contains(3))
	tree.print()
	tree.remove(9)
	tree.remove(2)
	fmt.Println(t.root.minData())
	tree.print()
}
