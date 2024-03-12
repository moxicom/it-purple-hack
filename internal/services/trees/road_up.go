package trees

import (
	"fmt"
)

func (tree *Tree) RoadUp(elId int64) ([]int64, error) {
	var path []int64
	current, ok := tree.Nodes[elId]
	if !ok {
		return []int64{}, fmt.Errorf("element with id %v does not exist", elId)
	}
	for current.ParentId != -1 {
		path = append(path, current.ID)
		current = tree.Nodes[current.ParentId]
	}
	path = append(path, current.ID)
	return path, nil
}

// func (node *Node) FindNode(id int64) *Node {
// 	if node == nil {
// 		return nil
// 	}
// 	if node.ID == id {
// 		return node
// 	}
// 	for i := range node.Children {
// 		found := node.Children[i].FindNode(id)
// 		if found != nil {
// 			return found
// 		}
// 	}
// 	return nil
// }
