package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	discountMatrixName = "discount_matrix"
	baselineMatrixName = "baseline_matrix"
)

func (rep *Repository) GetPrice(
	isDiscount bool,
	matrixId int64,
	locationsRoadUp []int64,
	categoriesRoadUp []int64,
) (int64, error) {
	for _, locationId := range locationsRoadUp {
		for _, microcategoryId := range categoriesRoadUp {
			// Check cache
			price, err := rep.getPriceCache(isDiscount, matrixId, locationId, microcategoryId)
			if err != nil && err != redis.Nil {
				return -1, err
			}

			if err == nil && price != -1 {
				// If price has found
				fmt.Printf("Found in cache: Microcategory_id=%v, Location_id=%v, price=%v\n", microcategoryId, locationId, price)
				return price, nil
			} else if err == nil && price == -1 {
				fmt.Printf("Found in cache: Microcategory_id=%v, Location_id=%v, price=%v\n", microcategoryId, locationId, price)
				continue
			}

			fmt.Printf("No redis cache: matrix_id=%v Location_id=%v Microcategory_id=%v\n", matrixId, locationId, microcategoryId)

			// TODO: fetch databse
			price, err = rep.fetchDatabase(isDiscount, matrixId, locationId, microcategoryId)
			if err != nil && err != sql.ErrNoRows {
				// Catched an error
				return -1, err
			} else if err != nil && err == sql.ErrNoRows {
				// Did not find data
				rep.setPriceCache(isDiscount, matrixId, locationId, microcategoryId, -1)
			} else if err == nil {
				// Found data in database
				rep.setPriceCache(isDiscount, matrixId, locationId, microcategoryId, price)
				return price, nil
			}
		}
	}
	return -1, nil
}

func (rep *Repository) getPriceCache(
	isDiscount bool,
	matrixId int64,
	locationId int64,
	microcategoryId int64,
) (int64, error) {
	ctx := context.Background()
	var err error
	price := int64(-1)
	if isDiscount {
		price, err = rep.getMatrixCache(ctx, fmt.Sprintf("%s_%v", discountMatrixName, matrixId), locationId, microcategoryId)
	} else {
		price, err = rep.getMatrixCache(ctx, fmt.Sprintf("%s_%v", baselineMatrixName, matrixId), locationId, microcategoryId)
	}
	return price, err
}

func (rep *Repository) setPriceCache(isDiscount bool,
	matrixId int64,
	locationId int64,
	microcategoryId int64,
	price int64,
) error {
	ctx := context.Background()
	var err error
	if isDiscount {
		err = rep.setMatrixCache(ctx, fmt.Sprintf("%s_%v", discountMatrixName, matrixId), locationId, microcategoryId, price)
	} else {
		err = rep.setMatrixCache(ctx, fmt.Sprintf("%s_%v", baselineMatrixName, matrixId), locationId, microcategoryId, price)
	}
	return err
}

func (rep *Repository) fetchDatabase(
	isDiscount bool,
	matrixId int64,
	locationId int64,
	microcategoryId int64,
) (int64, error) {
	var matrixPrefix string
	if isDiscount {
		matrixPrefix = discountMatrixName
	} else {
		matrixPrefix = baselineMatrixName
	}

	return rep.getPairPrice(
		fmt.Sprintf("%s_%v", matrixPrefix, matrixId),
		locationId,
		microcategoryId,
	)
}
