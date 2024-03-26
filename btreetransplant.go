package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}
	parent := findParent(root, node)

	if parent == nil {
		root = rplc
	} else if parent.Left == node {
		parent.Left = rplc
		rplc.Parent = parent
	} else {
		parent.Right = rplc
		rplc.Parent = parent
	}
	return root
}

func findParent(root, node *TreeNode) *TreeNode {
	if root == nil || root == node {
		return nil
	}

	if (root.Left == node) || (root.Right == node) {
		return root
	}

	if root.Data > node.Data {
		return findParent(root.Left, node)
	} else {
		return findParent(root.Right, node)
	}
}
