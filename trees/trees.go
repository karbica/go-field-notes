package trees

// Tree implements the behaviour of a tree.
type Tree interface {
	Size() (size int)
	Empty() (empty bool)
	Clear()
	Values() (values []interface{})
}
