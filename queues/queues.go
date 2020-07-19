package queues

// Queue interface implements queue behaviour.
type Queue interface {
	Len() int
	Empty() bool
	Peek() (interface{}, error)
	Enqueue(...interface{}) error
	Dequeue() (interface{}, error)
}
