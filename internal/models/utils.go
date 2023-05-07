package models

type PaginateRequest struct {
	Limit int64 `form:"limit,default=10" binding:"omitempty"`
	Page  int64 `form:"page,default=1" binding:"omitempty,min=1"`
}

type FilterRequest struct {
	Search  string `form:"search" binding:"omitempty,ascii"`
	OrderBy string `form:"order_by,default=created_at"`
	Sort    string `form:"sort,default=desc"`
}

type MetaResponse struct {
	Limit    *int64 `json:"limit"`
	Page     int64  `json:"page"`
	LastPage int64  `json:"last_page"`
	Total    int64  `json:"total"`
}
