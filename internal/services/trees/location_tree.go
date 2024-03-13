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
	tree := NewTree(ConvertLocationsToStruct(), &locationID)
	return tree
}

func ConvertLocationsToStruct() []RawData {
	data := []RawData{}
	for name, childs := range rawLocations {
		subElements := []RawData{}
		for _, childName := range childs {
			subElements = append(subElements, RawData{Name: childName})
		}
		data = append(data, RawData{name, subElements})
	}
	return data
}
