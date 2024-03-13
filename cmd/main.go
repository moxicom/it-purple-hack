package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/moxicom/it-purple-hack/config"
	"github.com/moxicom/it-purple-hack/internal/repository"
	"github.com/moxicom/it-purple-hack/internal/services/price"
	"github.com/moxicom/it-purple-hack/internal/services/trees"
	"github.com/moxicom/it-purple-hack/models"
)

func main() {
	// TODO: read .env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// TODO: init db
	dbConfig := config.ReadDbInfo()
	postgres := repository.NewPostgres(dbConfig)
	defer postgres.Close()

	redisConfig := config.ReadRedisInfo()
	redis := repository.NewRedis(redisConfig)
	defer redis.Close()

	repo := repository.NewRepository(postgres, redis)
	locationsTree := trees.GetLocationsTree()
	categories := trees.GetCategoriesTree()
	// locationsTree.PrintTree()
	categories.PrintTree()
	// Mock service request
	priceService := price.NewPriceService(repo, categories, locationsTree)
	fmt.Println()
	price, err := priceService.GetPrice(
		models.PriceRequest{
			LocationId:      11,
			MicrocategoryId: 11,
			UserId:          2200,
		},
		models.Storage{
			Baseline: "baseline_matrix_1",
			Discounts: []models.Discount{
				{
					SegmentId:      100,
					DiscountMatrix: "discount_matrix_1",
				},
				{
					SegmentId:      290,
					DiscountMatrix: "discount_matrix_2",
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("IS DISC", price.IsDiscount)
	fmt.Println("MATR_ID", price.MatrixId)
	fmt.Println("PRICE", price.Price)
	fmt.Println("SEGMENT_ID", price.UserSegmentId)
}
