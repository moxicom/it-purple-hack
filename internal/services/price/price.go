package price

import (
	"github.com/moxicom/it-purple-hack/internal/repository"
	"github.com/moxicom/it-purple-hack/internal/services/trees"
)

type PriceService struct {
	repo           *repository.Repository
	categoriesTree *trees.Tree
	locationsTree  *trees.Tree
}

func NewPriceService(repo *repository.Repository, categoriesTree *trees.Tree, locationsTree *trees.Tree) *PriceService {
	return &PriceService{repo, categoriesTree, locationsTree}
}
