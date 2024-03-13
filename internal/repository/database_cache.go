package repository

import (
	"context"
	"fmt"
	"time"
)

func (rep *Repository) getMatrixCache(
	ctx context.Context,
	matrixName string,
	locationID,
	microcategoryId int64) (int64, error) {
	key := fmt.Sprintf("%s_%v_%v", matrixName, locationID, microcategoryId)
	return rep.cacheDb.Get(ctx, key).Int64()
}

func (rep *Repository) setMatrixCache(
	ctx context.Context,
	matrixName string,
	locationID, microcategoryId,
	price int64) error {
	key := fmt.Sprintf("%s_%v_%v", matrixName, locationID, microcategoryId)
	err := rep.cacheDb.Set(ctx, key, price, time.Minute*5).Err()
	if err == nil {
		fmt.Printf("Inserted new key=%s, value=%v\n", key, price)
	}
	return err
}

func (rep *Repository) deleteMatrixCache(
	ctx context.Context,
	matrixName string,
	locationID,
	microcategoryId int64,
) error {
	key := fmt.Sprintf("%s_%v_%v", matrixName, locationID, microcategoryId)
	return rep.cacheDb.Del(ctx, key).Err()
}
