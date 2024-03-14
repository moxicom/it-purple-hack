package price

import "github.com/moxicom/it-purple-hack/models"

func (service *PriceService) UpdatePrice(r models.UpdateRequest, s models.Storage) error {
	return service.repo.UpdatePrice(
		r.IsDiscount,
		r.MatrixId,
		r.LocationId,
		r.MicrocategoryId,
		r.NewPrice,
	)
}
