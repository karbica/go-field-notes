package binarytree

import (
	"errors"

	queue "github.com/karbica/go-field-notes/queues/arrayqueue"
)

// Tree holds a key, value, and pointers to other nodes.
type Tree struct {
	Key   int
	Value interface{}
	Left  *Tree
	Right *Tree
}

// Insert adds a node to the left or right in level order.
func (t *Tree) Insert(key int, value interface{}) (node *Tree, err error) {
	queue := queue.New()
	queue.Enqueue(t)

	for !queue.Empty() {
		_node, _ := queue.Dequeue()
		node := _node.(*Tree)

		if node.Left == nil {
			node.Left = New(key, value)
			return node.Left, nil
		}

		if node.Right == nil {
			node.Right = New(key, value)
			return node.Right, nil
		}

		queue.Enqueue(node.Left, node.Right)
	}

	return nil, errors.New("btree error: could not insert new node")
}

// Remove finds and removes the node given the key. The remove strategy finds
// the deepest node (rightmost) and replaces the node to remove with it.
func (t *Tree) Remove(key int) (node *Tree) {
	queue := queue.New()
	queue.Enqueue(t)
	var marked *Tree
	var deep *Tree

	if key == t.Key && t.Left == nil && t.Right == nil {
		return nil
	}

	for !queue.Empty() {
		_node, _ := queue.Dequeue()
		node := _node.(*Tree)

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

	return t
}

// Levelorder traverses those nodes from the top down, left to right.
func (*Tree) Levelorder(t *Tree, fn func(*Tree)) {
	queue := queue.New()
	queue.Enqueue(t)

	for !queue.Empty() {
		_node, _ := queue.Dequeue()
		node := _node.(*Tree)

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
func (*Tree) Inorder(node *Tree, fn func(*Tree)) {
	if node != nil {
		node.Inorder(node.Left, fn)
		fn(node)
		node.Inorder(node.Right, fn)
	}
}

// Preorder traverses the nodes from root, left, right. The provided
// callback is invoked at every node.
func (*Tree) Preorder(node *Tree, fn func(*Tree)) {
	if node != nil {
		fn(node)
		node.Preorder(node.Left, fn)
		node.Preorder(node.Right, fn)
	}
}

// Postorder traverses the nodes from left, right, root. The provided
// callback is invoked at every node.
func (*Tree) Postorder(node *Tree, fn func(*Tree)) {
	if node != nil {
		node.Postorder(node.Left, fn)
		node.Postorder(node.Right, fn)
		fn(node)
	}
}

// Remove the deepest (rightmost) node in the binary tree.
func (t *Tree) removeDeep(deep *Tree) {
	queue := queue.New()
	queue.Enqueue(t)

	for !queue.Empty() {
		_node, _ := queue.Dequeue()
		node := _node.(*Tree)

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

// New returns a new instance of the root node in a tree.
func New(key int, value interface{}) *Tree {
	return &Tree{Key: key, Value: value, Left: nil, Right: nil}
}
