package model

import (
	"strings"

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

func NewSort(s string) SortBy {
	switch strings.ToLower(s) {
	case "newest":
		return Newest
	case "highest_price":
		return HighestPrice
	case "lowest_price":
		return LowestPrice
	}
	return 0
}

type ProductSearchResponse struct {
	Id       string  `json:"id"`
	Title    string  `json:"title"`
	CTAURL   string  `json:"cta_url"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
	Catalog  string  `json:"catalog"`
}

type SearchResponse struct {
	Products   []ProductSearchResponse `json:"products"`
	NextCursor string                  `json:"next_cursor"`
}
