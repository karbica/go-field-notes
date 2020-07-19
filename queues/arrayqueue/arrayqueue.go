package arrayqueue

import (
	"errors"
	"sync"
)

// Queue controls the state of values in a FIFO manner.
type Queue struct {
	items []interface{}
	lock  sync.Mutex
}

// Len returns the number of items in the queue.
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.items)
}

// Empty indicates whether or not the queue is empty.
func (q *Queue) Empty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.items) == 0
}

// Peek returns the item at the front without removing it.
func (q *Queue) Peek() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("queue error: cannot peek empty queue")
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	return q.items[0], nil
}

// Enqueue adds items to the end of the queue.
func (q *Queue) Enqueue(items ...interface{}) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = append(q.items, items...)

	return nil
}

// Dequeue removes an item from the front of the queue.
func (q *Queue) Dequeue() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("queue error: cannot dequeue from empty queue")
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	item := q.items[0]
	q.items = q.items[1:]

	return item, nil
}

// New returns a new instance of a queue.
func New() *Queue {
	return &Queue{items: make([]interface{}, 0)}
}
