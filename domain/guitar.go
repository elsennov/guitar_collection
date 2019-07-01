package domain

type (
	Guitar struct {
		Base
		Brand string `json:"brand"`
		Type  string `json:"type"`
		Price int    `json:"price"`
	}

	ViewGuitar struct {
		Id    uint   `json:"id"`
		Brand string `json:"brand"`
		Type  string `json:"type"`
		Price int    `json:"price"`
	}
)
