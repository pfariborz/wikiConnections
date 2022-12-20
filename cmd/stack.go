package cmd

// Implentation of Stack Object
type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

// Pushes a new element onto the stack
func (s *Stack) push(elem string) {
	*s = append(*s, elem)
}

// Pops an element off of the stack
func (s *Stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		elem := (*s)[index]
		*s = (*s)[:index]
		return elem, true
	}
}
