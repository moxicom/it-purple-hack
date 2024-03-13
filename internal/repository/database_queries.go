package repository

import (
	"fmt"
)

func (rep *Repository) getPairPrice(
	matrixName string,
	locationID int64,
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

func (rep *Repository) updatePairPrice(
	matrixName string,
	locationID int64,
	microcategoryId int64,
	newPrice int64,
) error {
	_, err := rep.db.Exec(
		fmt.Sprintf(
			"UPDATE %s SET price = $1 WHERE location_id = $2 AND microcategory_id = $3", matrixName,
		), newPrice, locationID, microcategoryId,
	)
	return err
}

func (rep *Repository) deletePairPrice(
	matrixName string,
	locationID int64,
	microcategoryId int64,
	newPrice int64,
) error {
	_, err := rep.db.Exec(
		fmt.Sprintf(
			"DELETE FROM %s WHERE location_id = $1 AND microcategory_id = $2", matrixName,
		), locationID, microcategoryId,
	)
	return err
}
