package piscine

func ListPopFirst(l *List) (*NodeL, error) {
	if l.Head == nil {
		return nil, nil
	}
	n := l.Head
	l.Head = n.Next
	if l.Head == nil {
		l.Tail = nil
	}
	return n, nil
}
