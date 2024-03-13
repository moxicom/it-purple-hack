package repository

import "fmt"

func (rep *Repository) UpdatePrice(
	isDiscount bool,
	matrixId int64,
	locationId int64,
	microcategoryId int64,
	newPrice int64,
) error {
	// update cache
	err := rep.setPriceCache(isDiscount, matrixId, locationId, microcategoryId, newPrice)
	if err != nil {
		return err
	}
	// update database row
	return rep.updateDatabaseRow(isDiscount, matrixId, locationId, microcategoryId, newPrice)
}

func (rep *Repository) updateDatabaseRow(
	isDiscount bool,
	matrixId int64,
	locationId int64,
	microcategoryId int64,
	newPrice int64,
) error {
	var matrixPrefix string
	if isDiscount {
		matrixPrefix = discountMatrixName
	} else {
		matrixPrefix = baselineMatrixName
	}
	return rep.updatePairPrice(
		fmt.Sprintf("%s_%v", matrixPrefix, matrixId),
		locationId,
		microcategoryId,
		newPrice,
	)
}
