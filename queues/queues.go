package queues

// Queue implements the behaviour of a queue.
type Queue interface {
	Len() int
	Empty() bool
	Peek() (interface{}, error)
	Enqueue(...interface{}) error
	Dequeue() (interface{}, error)
}
