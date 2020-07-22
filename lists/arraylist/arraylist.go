package arraylist

import (
	"errors"
)

// List holds the state of the list instance.
type List struct {
	elements []interface{}
	length   int
}

// New returns a new instance of a list.
func New() *List {
	return &List{elements: make([]interface{}, 0), length: 0}
}

// Len returns the length of the list.
func (l *List) Len() (length int) {
	return len(l.elements)
}

// Empty indicates whether or not the list is empty.
func (l *List) Empty() (empty bool) {
	return len(l.elements) == 0
}

// Get returns an element at a given index within range.
func (l *List) Get(index int) (element interface{}, ok error) {
	if !l.withinRange(index) {
		return nil, errors.New("arraylist error: given index out of list range")
	}

	return l.elements[index], nil
}

// Push adds elements to the end of the list and returns the new length.
func (l *List) Push(elements ...interface{}) (length int) {
	l.elements = append(l.elements, elements...)
	l.length += len(elements)

	return l.length
}

// Pop removes the last element from the list and returns it.
// If there are no elements, an error is returned.
func (l *List) Pop() (element interface{}, ok error) {
	if l.Empty() {
		return nil, errors.New("arraylist error: cannot pop from empty list")
	}

	end := l.length - 1
	element = l.elements[end]
	l.elements = l.elements[:end]
	l.length--

	return element, nil
}

// Shift removes the first element from the list and returns it.
// If there are no elements, an error is returned.
func (l *List) Shift() (element interface{}, ok error) {
	if l.Empty() {
		return nil, errors.New("arraylist error: cannot shift from empty list")
	}

	element = l.elements[0]
	l.elements = l.elements[1:]
	l.length--

	return element, nil
}

// Unshift adds elements to the front of the list and returns the new length.
func (l *List) Unshift(elements ...interface{}) (length int) {
	front := append(make([]interface{}, len(elements)), elements...)
	l.elements = append(front, l.elements...)
	l.length += len(elements)

	return l.length
}

// Check if the given index is within range of the list.
func (l *List) withinRange(index int) (ok bool) {
	if index < 0 || index > l.Len()-1 {
		return false
	}

	return true
}
