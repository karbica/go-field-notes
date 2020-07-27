package binarytree

import (
	queue "github.com/karbica/go-field-notes/queues/arrayqueue"
)

// Tree holds the state of a tree.
type Tree struct {
	Root *Node
	size int
}

// Node holds the state of a node.
type Node struct {
	Key   int
	Value interface{}
	Left  *Node
	Right *Node
}

// NewTree returns a new instance of a tree.
func NewTree(key int, value interface{}) (tree *Tree) {
	root := NewNode(key, value)
	return &Tree{Root: root, size: 1}
}

// NewNode returns a new instance of a node.
func NewNode(key int, value interface{}) (node *Node) {
	return &Node{Key: key, Value: value, Left: nil, Right: nil}
}

// Insert adds a node to the left or right in level order.
// The newly inserted node is then returned.
func (t *Tree) Insert(key int, value interface{}) (node *Node) {
	if t.Root == nil {
		t.Root = NewNode(key, value)
		return t.Root
	}

	q := queue.New()
	q.Enqueue(t.Root)

	for !q.Empty() {
		_n, _ := q.Dequeue()
		n := _n.(*Node)

		if n.Left == nil {
			n.Left = NewNode(key, value)
			node = n.Left
			break
		}

		if n.Right == nil {
			n.Right = NewNode(key, value)
			node = n.Right
			break
		}

		q.Enqueue(n.Left, n.Right)
	}

	t.size++
	return node
}

// Remove finds and removes the node given the key. The remove strategy finds
// the deepest node (rightmost) and replaces the node to remove with it.
func (t *Tree) Remove(key int) {
	root := t.Root

	if key == root.Key && root.Left == nil && root.Right == nil {
		t.Root = nil
		return
	}

	var marked *Node
	var deep *Node
	queue := queue.New()
	queue.Enqueue(root)

	for !queue.Empty() {
		_node, _ := queue.Dequeue()
		node := _node.(*Node)

		if node != nil {
			if key == node.Key {
				marked = node
			}
			deep = node
			queue.Enqueue(node.Left, node.Right)
		}
	}

	marked.Key = deep.Key
	marked.Value = deep.Value
	t.removeDeep(deep)
	t.size--
}

// Levelorder traverses those nodes from the top down, left to right.
func (*Tree) Levelorder(t *Tree, fn func(*Node)) {
	queue := queue.New()
	queue.Enqueue(t)

	for !queue.Empty() {
		_node, _ := queue.Dequeue()
		node := _node.(*Node)

		fn(node)

		if node.Left != nil {
			queue.Enqueue(node.Left)
		}

		if node.Right != nil {
			queue.Enqueue((node.Right))
		}
	}
}

// Inorder traverses the nodes from left, root, right. The provided
// callback is invoked at every node.
func (t *Tree) Inorder(node *Node, fn func(*Node)) {
	if node != nil {
		t.Inorder(node.Left, fn)
		fn(node)
		t.Inorder(node.Right, fn)
	}
}

// Preorder traverses the nodes from root, left, right. The provided
// callback is invoked at every node.
func (t *Tree) Preorder(node *Node, fn func(*Node)) {
	if node != nil {
		fn(node)
		t.Preorder(node.Left, fn)
		t.Preorder(node.Right, fn)
	}
}

// Postorder traverses the nodes from left, right, root. The provided
// callback is invoked at every node.
func (t *Tree) Postorder(node *Node, fn func(*Node)) {
	if node != nil {
		t.Postorder(node.Left, fn)
		t.Postorder(node.Right, fn)
		fn(node)
	}
}

// Remove the deepest (rightmost) node in the binary tree.
func (t *Tree) removeDeep(deep *Node) {
	queue := queue.New()
	queue.Enqueue(t.Root)

	for !queue.Empty() {
		_node, _ := queue.Dequeue()
		node := _node.(*Node)

		if node.Left != nil {
			if node.Left == deep {
				node.Left = nil
				return
			}
			queue.Enqueue(node.Left)
		}

		if node.Right != nil {
			if node.Right == deep {
				node.Right = nil
				return
			}
			queue.Enqueue(node.Right)
		}
	}
}
