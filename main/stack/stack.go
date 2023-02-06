package stack

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str interface{}) {
	*s = append(*s, str)
}

func (s *Stack) Top() interface{} {
	if len(*s) == 0 {
		return nil
	}
	index := len(*s) - 1
	element := (*s)[index]
	return element
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
