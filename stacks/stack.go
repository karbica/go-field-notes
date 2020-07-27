package stacks

// Stack implements the behavior of a stack.
type Stack interface {
	Len() (length int)
	Empty() (empty bool)
	Peek() (top interface{}, err error)
	Push(items ...interface{}) (length int)
	Pop() (item interface{}, err error)
}
