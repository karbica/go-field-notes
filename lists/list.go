package lists

// List implements the behavior of a list.
type List interface {
	Len() (length int)
	Empty() (empty bool)
	Get(int) (element interface{}, ok error)
	Push(...interface{}) (length int)
	Pop() (element interface{}, ok error)
	Shift() (element interface{}, ok error)
	Unshift(...interface{}) (length int)
}
