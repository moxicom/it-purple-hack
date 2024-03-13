package repository

import (
	"context"
	"fmt"
)

func (rep *Repository) DeletePrice(
	matrixId int64,
	locationId int64,
	microcategoryId int64,
	newPrice int64,
) error {
	// Delete cache
	ctx := context.Background()
	err := rep.deleteMatrixCache(
		ctx,
		fmt.Sprintf("%s_%v", discountMatrixName, matrixId),
		locationId,
		microcategoryId,
	)
	if err != nil {
		return err
	}
	// update database row
	return rep.deleteDatabaseRow(matrixId, locationId, microcategoryId, newPrice)
}

func (rep *Repository) deleteDatabaseRow(
	matrixId int64,
	locationId int64,
	microcategoryId int64,
	newPrice int64,
) error {
	return rep.deletePairPrice(
		fmt.Sprintf("%s_%v", discountMatrixName, matrixId),
		locationId,
		microcategoryId,
		newPrice,
	)
}
