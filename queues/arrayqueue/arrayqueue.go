package arrayqueue

import (
	"errors"

	list "github.com/karbica/go-field-notes/lists/arraylist"
)

// Queue controls the state of values in a FIFO manner.
type Queue struct {
	items *list.List
}

// New returns a new instance of a queue.
func New() *Queue {
	return &Queue{items: list.New()}
}

// Len returns the number of items in the queue.
func (q *Queue) Len() int {
	return q.items.Len()
}

// Empty indicates whether or not the queue is empty.
func (q *Queue) Empty() bool {
	return q.items.Empty()
}

// Peek returns the item at the front without removing it.
func (q *Queue) Peek() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("queue error: cannot peek empty queue")
	}

	return q.items.Get(0)
}

// Enqueue adds items to the end of the queue.
func (q *Queue) Enqueue(items ...interface{}) int {
	return q.items.Push(items...)
}

// Dequeue removes an item from the front of the queue.
func (q *Queue) Dequeue() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("queue error: cannot dequeue from empty queue")
	}

	return q.items.Shift()
}
