package model

type SearchCatalogsReq struct {
	Q string `json:"q"`
}

type ProductSearchCatalogs struct {
	Catalog string `json:"catalog"`
	Count   int64  `json:"count"`
}

type SearchCatalogsResp struct {
	Catalogs []ProductSearchCatalogs `json:"catalogs"`
}
