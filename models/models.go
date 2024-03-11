package models

type PriceRequest struct {
	Location_id      int `json:"location_id"`
	Microcategory_id int `json:"microcategory_id"`
	User_id          int `json:"user_id"`
}

type PriceInfo struct {
	Price            int `json:"price"`
	Location_id      int `json:"location_id"`
	Microcategory_id int `json:"microcategory_id"`
	Matrix_id        int `json:"matrix_id"`
	User_segment_id  int `json:"user_segment_id"`
}

type Storage struct {
	Baseline  string           `json:"baseline"`
	Discounts []map[int]string `json:"discounts"`
}
