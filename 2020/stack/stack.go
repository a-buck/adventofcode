package stack

// Stack type
type Stack []interface{}

// Pop element
func (s *Stack) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	i := len(*s) - 1
	item := (*s)[i]
	(*s)[i] = nil
	*s = (*s)[:i]
	return item
}

// Push element
func (s *Stack) Push(val interface{}) {
	*s = append(*s, val)
}

// Peek element
func (s *Stack) Peek() interface{} {
	if len(*s) == 0 {
		return nil
	}
	i := len(*s) - 1
	item := (*s)[i]
	return item
}
