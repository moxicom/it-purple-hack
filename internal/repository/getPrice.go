package repository

import "github.com/moxicom/it-purple-hack/models"

func (rep *Repository) GetPrice(req models.PriceRequest, s models.Storage) (models.PriceInfo, error) {
	return models.PriceInfo{}, nil
}

func (rep *Repository) GetPriceCache() {
	
}
