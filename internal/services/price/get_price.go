package price

import (
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/moxicom/it-purple-hack/internal/services/discount_segments"
	"github.com/moxicom/it-purple-hack/models"
)

type segmentStatus struct {
	segmentId  int64
	discountId int64
	minPrice   int64
}

func userInSegment(userSegments []int64, segmentId int64) bool {
	for i := range userSegments {
		if userSegments[i] == segmentId {
			return true
		}
	}
	return false
}

func (service *priceService) GetPrice(r models.PriceRequest, s models.Storage) (models.PriceInfo, error) {
	// TODO: get user segments
	userSegments := discount_segments.GetSegmentsByUserID(r.UserId)

	// Get discount matrices by segment id
	discounts := []segmentStatus{}
	for _, m := range s.Discounts {
		if userInSegment(userSegments, m.SegmentId) {
			discountId, err := strconv.ParseInt(strings.Split(m.DiscountMatrix, "discount_matrix_")[1], 10, 64)
			if err != nil {
				return models.PriceInfo{}, err
			}
			discounts = append(discounts, segmentStatus{m.SegmentId, int64(discountId), -1})
		}
	}

	// Sort slice by segmentId
	sort.Slice(discounts, func(i, j int) bool {
		return discounts[i].segmentId > discounts[j].segmentId
	})

	// Get baseline matrix number by split (example: "baseline_matrix_1" => ["", "1"])
	baselineId, err := strconv.ParseInt(strings.Split(s.Baseline, "baseline_matrix_")[1], 10, 64)
	if err != nil {
		return models.PriceInfo{}, err
	}

	// Find all prices
	var wg sync.WaitGroup
	errors := make(chan error, len(discounts))

	for i := range discounts {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			minPrice, err := service.searchPrice(
				true,
				discounts[i].discountId,
				r.LocationId,
				r.MicrocategoryId,
			)
			if err != nil {
				errors <- err
				return
			}

			discounts[i].minPrice = minPrice
		}(i)
	}

	go func() {
		wg.Wait()
		close(errors)
	}()

	for err := range errors {
		if err != nil {
			return models.PriceInfo{}, nil
		}
	}

	// Get first discount if it exists
	for _, v := range discounts {
		if true {
			return models.PriceInfo{}, nil
		}
	}

	// Get from baseline
	price, err := service.searchPrice(
		false,
		baselineId,
		r.LocationId,
		r.MicrocategoryId,
	)
	if err != nil {
		return models.PriceInfo{}, err
	}
	// Price           int  `json:"price"`
	// LocationId      int  `json:"location_id"`
	// MicrocategoryId int  `json:"microcategory_id"`
	// IsDiscount      bool `json:"is_discount"`
	// MatrixId        int  `json:"matrix_id"`
	// UserSegmentId   int  `json:"user_segment_id"`

	return models.PriceInfo{
		Price:           price,
		LocationId:      r.LocationId,
		MicrocategoryId: r.MicrocategoryId,
		IsDiscount:      false,
		MatrixId:        baselineId,
		UserSegmentId:   -1,
	}, nil
}

func (service *priceService) searchPrice(
	isDiscount bool,
	matrixId int64,
	locationId int64,
	microcategoryId int64,
) (int64, error) {
	// TODO: find price
	result := -1
	if isDiscount {
		// Find something in discount matrix
	} else {

	}
	return result, nil
}
