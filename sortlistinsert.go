package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}

	if l == nil || data_ref <= l.Data {
		newNode.Next = l
		return newNode
	}

	cur := l
	for cur.Next != nil && data_ref > cur.Next.Data {
		cur = cur.Next
	}

	newNode.Next = cur.Next
	cur.Next = newNode

	return l
}
