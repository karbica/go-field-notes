package lists

// List implements the behaviour of a list.
type List interface {
	Len() int
	Empty() bool
	Get(int) (interface{}, error)
	Push(...interface{}) int
	Pop() (interface{}, error)
	Shift() (interface{}, error)
	Unshift(...interface{}) int
}
