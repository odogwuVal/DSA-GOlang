package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	left  *node
	right *node
}

type bst struct {
	root *node
	len  int
}

// prints the value of the node when print is called
func (n node) String() string {
	return strconv.Itoa(n.value)
}

func (b bst) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()
}

func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root)
}

func (b bst) inOrderTraversalByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s ", root))
	b.inOrderTraversalByNode(sb, root.right)
}

func (b *bst) add(value int) {
	b.root = b.addByNode(b.root, value)
	b.len++
}
func (b *bst) addByNode(root *node, value int) *node {
	if root == nil {
		return &node{value: value}
	}
	if root.value > value {
		root.left = b.addByNode(root.left, value)
	} else {
		root.right = b.addByNode(root.right, value)
	}
	// this ensures that the structure of the tree is maintained
	return root
}

func (b bst) search(value int) (*node, bool) {
	return b.searchByNode(b.root, value)
}

func (b bst) searchByNode(root *node, value int) (*node, bool) {
	if root == nil {
		return nil, false
	}
	if value == root.value {
		return root, true
	} else if value < root.value {
		return b.searchByNode(root.left, value)
	} else {
		return b.searchByNode(root.right, value)
	}
}

func (b *bst) remove(value int) {
	b.removeByNode(b.root, value)
}

func (b *bst) removeByNode(root *node, value int) *node {
	if root == nil {
		return root
	}
	if value < root.value {
		root.left = b.removeByNode(root.left, value)
	} else if value > root.value {
		root.right = b.removeByNode(root.right, value)
	} else {
		// if the node has one or no child nodes
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			// Node with two children: Get the in-order predecessor (max in left subtree)
			temp := root.left
			// checking the right most node of the left sub tree (max)
			for temp.right != nil {
				temp = temp.right
			}
			// Copy the in-order predecessor's value to this node
			root.value = temp.value
			// Delete the in-order predecessor
			root.left = b.removeByNode(root.left, temp.value)
		}
	}
	return root
}

func main() {
	n := &node{value: 1}
	n.left = &node{value: 0}
	n.right = &node{value: 2}

	b := bst{
		root: n,
		len:  3,
	}

	b.add(5)
	b.add(7)
	b.add(3)
	// fmt.Println(b.search(7))
	b.remove(9)
	fmt.Println(b)

}
