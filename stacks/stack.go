package stacks

// Stack implements the behaviour of a stack.
type Stack interface {
	Len() int
	Empty() bool
	Peek() (interface{}, error)
	Push(...interface{}) error
	Pop() (interface{}, error)
}
