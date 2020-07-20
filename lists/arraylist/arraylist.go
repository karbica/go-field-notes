package arraylist

import (
	"errors"
	"sync"
)

// List holds ordered elements in a slice.
type List struct {
	elements []interface{}
	length   int
	lock     sync.Mutex
}

// Len returns the length of the list.
func (l *List) Len() int {
	l.lock.Lock()
	defer l.lock.Unlock()

	return len(l.elements)
}

// Empty indicates whether or not the list is empty.
func (l *List) Empty() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	return len(l.elements) == 0
}

// Push adds elements to the end of the list and returns the new length.
func (l *List) Push(elements ...interface{}) int {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.elements = append(l.elements, elements...)
	l.length += len(elements)

	return l.length
}

// Pop removes the last element from the list and returns it.
// If there are no elements, an error is returned.
func (l *List) Pop() (interface{}, error) {
	if l.Empty() {
		return nil, errors.New("arraylist error: cannot pop from empty list")
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	end := l.length - 1
	element := l.elements[end]
	l.elements = l.elements[:end]
	l.length--

	return element, nil
}

// Shift removes the first element from the list and returns it.
// If there are no elements, an error is returned.
func (l *List) Shift() (interface{}, error) {
	if l.Empty() {
		return nil, errors.New("arraylist error: cannot shift from empty list")
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	element := l.elements[0]
	l.elements = l.elements[1:]
	l.length--

	return element, nil
}

// Unshift adds elements to the front of the list and returns the new length.
func (l *List) Unshift(elements ...interface{}) int {
	l.lock.Lock()
	defer l.lock.Unlock()

	front := append(make([]interface{}, len(elements)), elements...)
	l.elements = append(front, l.elements...)
	l.length += len(elements)

	return l.length
}

// New returns a new instance of a list.
func New() *List {
	return &List{elements: make([]interface{}, 0), length: 0}
}
