package model

type SearchCategoryReq struct {
	Q string `json:"q"`
}

type ProductSearchCategory struct {
	Category string `json:"category"`
	Count    int64  `json:"count"`
}

type SearchCategoryResp struct {
	Categories []ProductSearchCategory `json:"categories"`
}
