package piscine

func ListReverse(l *List) {
	var rev *NodeL = nil
	current := l.Head

	for current != nil {
		next := current.Next
		current.Next = rev
		rev = current
		current = next
	}

	l.Head = rev
}
