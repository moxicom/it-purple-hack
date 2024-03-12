package trees

// import "fmt"

// type Node struct {
// 	ID       int64
// 	Name     string
// 	Children []*Node
// 	Parent   *Node
// }

// type RawData struct {
// 	Name        string
// 	SubElements []RawData
// }

// func (node *Node) insertNodeRecursively(data RawData, idGlobal *int64) {
// 	*idGlobal++
// 	newNode := NewNode(data.Name, *idGlobal)
// 	newNode.Parent = node
// 	node.Children = append(node.Children, newNode)
// 	for _, sublocationData := range data.SubElements {
// 		newNode.insertNodeRecursively(sublocationData, idGlobal)
// 	}
// }

// func NewNode(name string, id int64) *Node {
// 	return &Node{
// 		ID:       id,
// 		Name:     name,
// 		Children: []*Node{},
// 	}
// }

// // Create a child location
// func (l *Node) AddChild(child *Node) {
// 	l.Children = append(l.Children, child)
// 	child.Parent = l
// }

// // Print tree recursively
// func (l *Node) PrintTree(indent int) {
// 	fmt.Printf("%s%d - %s\n", generateIndent(indent), l.ID, l.Name)
// 	for _, child := range l.Children {
// 		child.PrintTree(indent + 2)
// 	}
// }

// // Генерирует отступ для вывода
// func generateIndent(indent int) string {
// 	result := ""
// 	for i := 0; i < indent; i++ {
// 		result += " "
// 	}
// 	return result
// }
