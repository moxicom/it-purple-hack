package trees

var categoryID int64

// func GetCategoryTree() *Node {
// 	// Create root category
// 	categoryID = 0
// 	allRegions := NewNode(rawCategories.Name, categoryID)
// 	for _, subCategory := range rawCategories.SubElements {
// 		allRegions.insertNodeRecursively(subCategory, &categoryID)
// 	}

// 	return allRegions
// }

func GetCategoriesTree() *Tree {
	// Create root location
	categoryID = 0
	tree := NewTree(rawCategories, &categoryID)
	return tree
}
