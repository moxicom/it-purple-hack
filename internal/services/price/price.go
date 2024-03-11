package price

import (
	"github.com/moxicom/it-purple-hack/internal/repository"
	"github.com/moxicom/it-purple-hack/models"
)

type priceService struct {
	repo *repository.Postgres
}

func NewPriceService(repo *repository.Postgres) *priceService {
	return &priceService{repo}
}

func (s *priceService) GetPrice(r models.PriceRequest) {
	// Check
}
