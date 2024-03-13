package repository

import (
	"fmt"
)

func (rep *Repository) getPairPrice(
	matrixName string,
	locationID,
	microcategoryId int64,
) (int64, error) {
	var price int64
	err := rep.db.QueryRow(
		fmt.Sprintf(
			"SELECT price FROM %s WHERE location_id = $1 AND microcategory_id = $2", matrixName,
		), locationID, microcategoryId,
	).Scan(&price)
	if err != nil {
		return 0, err
	}
	return price, nil
}
