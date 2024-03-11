package main

import (
	"fmt"

	"github.com/moxicom/it-purple-hack/internal/services/trees"
)

func main() {
	// TODO: read .env

	// TODO: init db

	// location tree mock
	root_locations := trees.GetLocationsTree()
	root_locations.PrintTree(0)
	fmt.Println()

	// category tree mock
	root_categories := trees.GetCategoryTree()
	root_categories.PrintTree(0)
}
