package model

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
)

type ProductSearchInsert struct {
	Title    string      `json:"title"`
	CTAURL   string      `json:"cta_url"`
	ImageURL string      `json:"image_url"`
	Price    float64     `json:"price"`
	Catalog  null.String `json:"catalog"`
}

type ProductSearchInsertResponse struct {
	Id uuid.UUID `json:"id"`
	ProductSearchInsert
}

type BulkInsertReq struct {
	ProductSearchInput []ProductSearchInsert `json:"products"`
}

type BulkInsertResp struct {
}
