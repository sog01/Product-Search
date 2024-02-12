package model

import "gopkg.in/guregu/null.v4"

type SearchTotalReq struct {
	Q       string      `json:"q"`
	Catalog null.String `json:"catalog"`
}

type SearchTotalResp struct {
	Total int64 `json:"total"`
}
