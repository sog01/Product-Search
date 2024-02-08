package model

import (
	"gopkg.in/guregu/null.v4"
)

type SearchReq struct {
	Q          string      `json:"q"`
	NextCursor null.String `json:"next_cursor"`
	Size       int         `json:"size"`
	SortBy     SortBy      `json:"sortBy"`
}

type SortBy int

const (
	Newest SortBy = iota + 1
	HighestPrice
	LowestPrice
)

type ProductSearchResponse struct {
	Id       string  `json:"id"`
	Title    string  `json:"title"`
	CTAURL   string  `json:"cta_url"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
}

type ProductSearchPagination struct {
	Total      int    `json:"total"`
	NextCursor string `json:"nextCursor"`
}

type SearchResponse struct {
	Products   []ProductSearchResponse `json:"products"`
	NextCursor string                  `json:"nextCursor"`
}
