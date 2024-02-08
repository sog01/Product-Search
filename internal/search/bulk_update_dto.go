package search

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
)

type ProductSearchUpdate struct {
	Id       uuid.UUID   `json:"id"`
	Title    null.String `json:"title"`
	CTAURL   null.String `json:"cta_url"`
	ImageURL null.String `json:"image_url"`
	Price    null.Float  `json:"price"`
}

type ProductSearchUpdateResponse struct {
	Id uuid.UUID `json:"id"`
	ProductSearchUpdate
}

type BulkUpdateReq struct {
	ProductSearchUpdate []ProductSearchUpdate `json:"products"`
	ctx                 context.Context
}

type BulkUpdateResp struct {
}
