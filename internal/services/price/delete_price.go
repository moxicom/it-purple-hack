package price

import (
	"errors"

	"github.com/moxicom/it-purple-hack/models"
)

func (service *priceService) DeleteDiscount(r models.UpdateRequest, s models.Storage) error {
	if r.IsDiscount {
		return service.repo.DeletePrice(
			r.MatrixId,
			r.LocationId,
			r.MicrocategoryId,
		)
	}
	return errors.New("can not delete row from baseline matrix")
}
