package trees

func (node *Node) RoadUp() []int64 {
	var path []int64
	current := node
	for current != nil {
		path = append(path, current.ID)
		current = current.Parent
	}

	// The path is reversed
	return path
}

func (node *Node) FindNode(id int64) *Node {
	if node == nil {
		return nil
	}
	if node.ID == id {
		return node
	}
	for i := range node.Children {
		found := node.Children[i].FindNode(id)
		if found != nil {
			return found
		}
	}
	return nil
}
