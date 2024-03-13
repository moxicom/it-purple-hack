package price

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/moxicom/it-purple-hack/internal/services/discount_segments"
	"github.com/moxicom/it-purple-hack/models"
)

type segmentStatus struct {
	segmentId       int64
	discountId      int64
	locationId      int64
	microcategoryId int64
	price           int64
}

func (service *priceService) GetPrice(r models.PriceRequest, s models.Storage) (models.PriceInfo, error) {
	// TODO: get user segments
	userSegments := discount_segments.GetSegmentsByUserID(r.UserId)

	discounts, err := service.getDiscountMatricesBySegments(userSegments, s.Discounts, r.LocationId, r.MicrocategoryId)
	if err != nil {
		return models.PriceInfo{}, err
	}
	fmt.Println(discounts)

	// Get baseline matrix number by split (example: "baseline_matrix_1" => ["", "1"])
	baselineId, err := strconv.ParseInt(strings.Split(s.Baseline, "baseline_matrix_")[1], 10, 64)
	if err != nil {
		return models.PriceInfo{}, err
	}

	// Find all prices in discounts matrices
	var wg sync.WaitGroup
	errors := make(chan error, len(discounts))
	locationsRoadUp, err := service.locationsTree.RoadUp(r.LocationId)
	if err != nil {
		return models.PriceInfo{}, err
	}
	categoriesRoadUp, err := service.categoriesTree.RoadUp(r.MicrocategoryId)
	if err != nil {
		return models.PriceInfo{}, err
	}

	fmt.Println("locations roadup", locationsRoadUp)
	fmt.Println("categories roadup", categoriesRoadUp)

	for i := range discounts {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			priceData, err := service.repo.GetPrice(
				true,
				discounts[i].discountId,
				locationsRoadUp,
				categoriesRoadUp,
			)
			if err != nil {
				errors <- err
				return
			}

			discounts[i].price = priceData.Price
			discounts[i].locationId = priceData.LocationId
			discounts[i].microcategoryId = priceData.MicrocategoryId
		}(i)
	}

	go func() {
		wg.Wait()
		close(errors)
	}()

	for err := range errors {
		if err == nil {
			close(errors)
			return models.PriceInfo{}, err
		}
	}

	// Get first discount if it exists
	for _, v := range discounts {
		if v.price != -1 {
			fmt.Println("Found in discount ", v.discountId)
			return models.PriceInfo{
				Price:           v.price,
				LocationId:      v.locationId,
				MicrocategoryId: v.microcategoryId,
				IsDiscount:      true,
				MatrixId:        v.discountId,
				UserSegmentId:   v.segmentId,
			}, nil
		}
	}

	// Get from baseline
	priceData, err := service.repo.GetPrice(
		false,
		baselineId,
		locationsRoadUp,
		categoriesRoadUp,
	)
	if err != nil {
		return models.PriceInfo{}, err
	}

	return models.PriceInfo{
		Price:           priceData.Price,
		LocationId:      priceData.LocationId,
		MicrocategoryId: priceData.MicrocategoryId,
		IsDiscount:      false,
		MatrixId:        baselineId,
		UserSegmentId:   -1,
	}, nil
}

// Get discount matrices info by segment id
func (service *priceService) getDiscountMatricesBySegments(
	userSegments []int64,
	allDiscountMatrices []models.Discount,
	locationId,
	microcategoryId int64,
) ([]segmentStatus, error) {
	discounts := []segmentStatus{}

	for _, m := range allDiscountMatrices {
		if userInSegment(userSegments, m.SegmentId) {
			discountId, err := strconv.ParseInt(strings.Split(m.DiscountMatrix, "discount_matrix_")[1], 10, 64)
			if err != nil {
				return []segmentStatus{}, err
			}
			discounts = append(discounts, segmentStatus{
				m.SegmentId,
				int64(discountId),
				locationId,
				microcategoryId,
				-1,
			},
			)
		}
	}

	// Sort slice by segmentId
	sort.Slice(discounts, func(i, j int) bool {
		return discounts[i].segmentId > discounts[j].segmentId
	})

	return discounts, nil
}

func userInSegment(userSegments []int64, segmentId int64) bool {
	for i := range userSegments {
		if userSegments[i] == segmentId {
			return true
		}
	}
	return false
}
