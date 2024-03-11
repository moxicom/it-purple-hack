package price

import (
	"sort"

	"github.com/moxicom/it-purple-hack/models"
)

func (service *priceService) GetPrice(r models.PriceRequest, s models.Storage) (models.PriceInfo, error) {
	// Check all discount matrix
	discounts := make([]int, len(s.Discounts))
	for key := range s.Discounts {
		discounts = append(discounts, key)
	}
	// if user
	sort.Slice(discounts, func(i, j int) bool {
		return discounts[i] > discounts[j]
	})

	// Get from baseline
	return service.repo.GetPrice(r, s)
}
