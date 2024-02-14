package model

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
)

type ProductSearchUpdate struct {
	Id          uuid.UUID   `json:"id" swaggertype:"string"`
	Title       null.String `json:"title" swaggertype:"string"`
	Description null.String `json:"description" swaggertype:"string"`
	CTAURL      null.String `json:"cta_url" swaggertype:"string"`
	ImageURL    null.String `json:"image_url" swaggertype:"string"`
	Price       null.Float  `json:"price" swaggertype:"number"`
	Catalog     null.String `json:"catalog" swaggertype:"string"`
}

type ProductSearchUpdateResponse struct {
	Id uuid.UUID `json:"id"`
	ProductSearchUpdate
}

type BulkUpdateReq struct {
	ProductSearchUpdate []ProductSearchUpdate `json:"products"`
}

type BulkUpdateResp struct {
}
