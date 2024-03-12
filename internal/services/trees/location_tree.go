package trees

var locationID int64

// func GetLocationsTree() *Node {
// 	// Create root location
// 	locationID = 0
// 	allRegions := NewNode(rawLocations.Name, locationID)
// 	for _, sublocation := range rawLocations.SubElements {
// 		allRegions.insertNodeRecursively(sublocation, &locationID)
// 	}

// 	return allRegions
// }

func GetLocationsTree() *Tree {
	// Create root location
	locationID = 0
	tree := NewTree(rawLocations, &locationID)
	return tree
}
