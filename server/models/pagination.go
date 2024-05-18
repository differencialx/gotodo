package models

type OffsetPaginationParams struct {
	Limit int `json:"limit" validate:"gte=1"`
	Page  int `json:"page" validate:"gte=1"`
}

type PaginationBody struct {
	Total    int `json:"total"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
