package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}

	sorted := true
	for sorted {
		sorted = false
		current := l
		prev := &NodeI{Data: 0, Next: l}

		for current != nil && current.Next != nil {
			if current.Data > current.Next.Data {
				current.Data, current.Next.Data = current.Next.Data, current.Data
				sorted = true
			}

			current = current.Next
			prev = prev.Next
		}
	}

	return l
}
