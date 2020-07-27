package maps

// Map implements the behavior of a map.
type Map interface {
	Size() (size int)
	Empty() (empty bool)
	Get(key interface{}) (value interface{}, ok bool)
	Set(key interface{}, value interface{})
	Delete(key interface{})
	Keys() (keys interface{})
}
