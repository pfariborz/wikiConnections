package cmd

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(elem string) {
	*s = append(*s, elem)
}

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
