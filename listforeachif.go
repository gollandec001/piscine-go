package piscine

func IsPositive_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	case string, rune:
		return false
	}
	return false
}

func IsNegative_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) < 0
	case string, rune:
		return false
	}
	return false
}

func IsNotNumeric_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	case string, rune:
		return true
	}
	return true
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	tmp := &List{}
	for i := (*l).Head; i != nil; i = i.Next {
		if cond(i) {
			f(i)
			ListPushBack(tmp, i)
		}
	}
	l = tmp
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}
