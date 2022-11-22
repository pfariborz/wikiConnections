package cmd

type Queue []string

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) enqueue(elem string) {
	*q = append(*q, elem)
}

func (q *Queue) dequeue() (string, bool) {
	if q.isEmpty() {
		return "", false
	} else {
		elem := (*q)[0]
		*q = (*q)[1:]
		return elem, true
	}
}
