package model

type SearchTotalReq struct {
	Q string `json:"q"`
}

type SearchTotalResp struct {
	Total int64 `json:"total"`
}
