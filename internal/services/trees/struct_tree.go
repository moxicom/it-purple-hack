package trees

import (
	"fmt"
	"sort"
)

type Node struct {
	ID         int64
	Name       string
	ChildrenId []int64
	ParentId   int64
}

type Tree struct {
	Nodes map[int64]Node
}

type RawData struct {
	Name        string
	SubElements []RawData
}

func NewTree(data []RawData, idGlobal *int64) *Tree {
	tree := &Tree{
		Nodes: make(map[int64]Node, 0),
	}
	for _, el := range data {
		tree.insertNodeRecursively(el, idGlobal, -1)
	}
	return tree
}

func (tree *Tree) insertNodeRecursively(data RawData, idGlobal *int64, parentId int64) int64 {
	*idGlobal++
	newNode := Node{
		ID:         *idGlobal,
		Name:       data.Name,
		ChildrenId: []int64{},
		ParentId:   parentId,
	}
	tree.Nodes[newNode.ID] = newNode
	for _, sublocationData := range data.SubElements {
		childId := tree.insertNodeRecursively(sublocationData, idGlobal, newNode.ID)
		newNode.ChildrenId = append(newNode.ChildrenId, childId)
	}
	tree.Nodes[newNode.ID] = newNode
	return newNode.ID
}

func (tree *Tree) PrintTree() {
	nodes := []int64{}
	for _, node := range tree.Nodes {
		if node.ParentId == -1 {
			nodes = append(nodes, node.ID)
		}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i] < nodes[j]
	})
	for _, nodeId := range nodes {
		tree.printNode(nodeId, 0)
	}
}

func (tree *Tree) printNode(nodeId int64, indent int) {
	fmt.Printf("%s%d - %s\n", generateIndent(indent), nodeId, tree.Nodes[nodeId].Name)
	for _, child := range tree.Nodes[nodeId].ChildrenId {
		tree.printNode(child, indent+2)
	}
}

func generateIndent(indent int) string {
	result := ""
	for i := 0; i < indent; i++ {
		result += " "
	}
	return result
}
