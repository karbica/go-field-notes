package arraystack

import (
	"errors"

	list "github.com/karbica/go-field-notes/lists/arraylist"
)

// Stack controls the state of the stack instance.
type Stack struct {
	items *list.List
}

// New returns a new instance of a stack.
func New() *Stack {
	return &Stack{items: list.New()}
}

// Len returns the number of items in the stack.
func (s *Stack) Len() (length int) {
	return s.items.Len()
}

// Empty indicates whether or not the stack is empty.
func (s *Stack) Empty() (empty bool) {
	return s.items.Empty()
}

// Peek returns the item at the top without removing it.
func (s *Stack) Peek() (top interface{}, ok error) {
	return s.items.Get(s.Len() - 1)
}

// Push adds items to the top of the stack.
func (s *Stack) Push(items ...interface{}) (length int) {
	return s.items.Push(items...)
}

// Pop removes an item from the top of the stack.
func (s *Stack) Pop() (item interface{}, ok error) {
	if s.items.Empty() {
		return nil, errors.New("stack error: cannot pop from empty stack")
	}

	return s.items.Pop()
}
