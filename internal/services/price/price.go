package price

import (
	"github.com/moxicom/it-purple-hack/internal/repository"
)

type priceService struct {
	repo *repository.Repository
}

func NewPriceService(repo *repository.Repository) *priceService {
	return &priceService{repo}
}
