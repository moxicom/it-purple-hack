package models

type PriceRequest struct {
	LocationId      int64 `json:"location_id"`
	MicrocategoryId int64 `json:"microcategory_id"`
	UserId          int64 `json:"user_id"`
}

type UpdateRequest struct {
	LocationId      int64 `json:"location_id"`
	MicrocategoryId int64 `json:"microcategory_id"`
	MatrixId        int64 `json:"matrix_id"`
	IsDiscount      bool  `json:"is_discount"`
	NewPrice        int64 `json:"new_price"`
}

type PriceInfo struct {
	Price           int64 `json:"price"`
	LocationId      int64 `json:"location_id"`
	MicrocategoryId int64 `json:"microcategory_id"`
	IsDiscount      bool  `json:"is_discount"`
	MatrixId        int64 `json:"matrix_id"`
	UserSegmentId   int64 `json:"user_segment_id"`
}

type Discount struct {
	SegmentId      int64
	DiscountMatrix string
}

type Storage struct {
	Baseline  string     `json:"baseline"`
	Discounts []Discount `json:"discounts"`
}
