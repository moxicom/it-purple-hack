package main

import (
	"fmt"
	"log"
	"net/http"

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

	PriceService := price.NewPriceService(repo, categories, locationsTree)

	//
	// Mock service requests
	//
	fmt.Println()
	testStorage := models.Storage{
		Baseline: "baseline_matrix_1",
		Discounts: []models.Discount{
			{
				SegmentId:      168,
				DiscountMatrix: "discount_matrix_1",
			},
			{
				SegmentId:      290,
				DiscountMatrix: "discount_matrix_2",
			},
		},
	}
	price, err := PriceService.GetPrice(
		models.PriceRequest{
			LocationId:      7,
			MicrocategoryId: 18,
			UserId:          2200,
			// 168, 290,
		},
		testStorage,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("IS DISC", price.IsDiscount)
	fmt.Println("MATR_ID", price.MatrixId)
	fmt.Println("PRICE", price.Price)
	fmt.Println("SEGMENT_ID", price.UserSegmentId)
	fmt.Println("LOCATION_ID", price.LocationId)
	fmt.Println("MICROCATEGORY_ID", price.MicrocategoryId)

	fmt.Println()
	fmt.Println("Update")
	err = PriceService.UpdatePrice(
		models.UpdateRequest{
			LocationId:      7,
			MicrocategoryId: 18,
			MatrixId:        1,
			IsDiscount:      true,
			NewPrice:        12345,
		},
		testStorage,
	)

	if err != nil {
		panic(err)
	}

	price, err = PriceService.GetPrice(
		models.PriceRequest{
			LocationId:      7,
			MicrocategoryId: 18,
			UserId:          2200,
			// 168, 290,
		},
		testStorage,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("IS DISC", price.IsDiscount)
	fmt.Println("MATR_ID", price.MatrixId)
	fmt.Println("PRICE", price.Price)
	fmt.Println("SEGMENT_ID", price.UserSegmentId)
	fmt.Println("LOCATION_ID", price.LocationId)
	fmt.Println("MICROCATEGORY_ID", price.MicrocategoryId)

	fmt.Println()
	fmt.Println("Delete")
	err = PriceService.DeleteDiscount(
		models.UpdateRequest{
			LocationId:      7,
			MicrocategoryId: 18,
			MatrixId:        1,
			IsDiscount:      true,
		},
		testStorage,
	)

	if err != nil {
		panic(err)
	}

	price, err = PriceService.GetPrice(
		models.PriceRequest{
			LocationId:      7,
			MicrocategoryId: 18,
			UserId:          2200,
			// 168, 290,
		},
		testStorage,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("IS DISC", price.IsDiscount)
	fmt.Println("MATR_ID", price.MatrixId)
	fmt.Println("PRICE", price.Price)
	fmt.Println("SEGMENT_ID", price.UserSegmentId)
	fmt.Println("LOCATION_ID", price.LocationId)
	fmt.Println("MICROCATEGORY_ID", price.MicrocategoryId)

	http.HandleFunc("POST /get_price", api_handlers.get_price(&PriceService, testStorage))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
