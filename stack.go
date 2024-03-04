package interpreter

type Stack []interface{}

func (s *Stack) Push(val interface{}) {
	*s = append(*s, val)
}

func (s *Stack) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	index := len(*s) - 1
	popped := (*s)[index]
	*s = (*s)[:index]
	return popped
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
