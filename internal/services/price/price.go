package price

import (
	"github.com/moxicom/it-purple-hack/internal/repository"
	"github.com/moxicom/it-purple-hack/internal/services/trees"
)

type priceService struct {
	repo           *repository.Repository
	categoriesTree *trees.Node
	locationsTree  *trees.Node
}

func NewPriceService(repo *repository.Repository, categoriesTree *trees.Node, locationsTree *trees.Node) *priceService {
	return &priceService{repo, categoriesTree, locationsTree}
}
