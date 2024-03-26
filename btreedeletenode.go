package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		root.Data = minValue(root.Right)

		root.Right = BTreeDeleteNode(root.Right, &TreeNode{Data: root.Data})
	}
	return root
}

func minValue(node *TreeNode) string {
	min := node
	for min.Left != nil {
		min = min.Left
	}
	return min.Data
}
