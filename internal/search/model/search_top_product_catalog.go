package model

type SearchTopProductCatalogReq struct {
	Catalogs []string `json:"catalogs"`
}

type TopProductCatalogData struct {
	ImageURL string `json:"image_url"`
}

type TopProductCatalog struct {
	Catalog string                  `json:"string"`
	Data    []TopProductCatalogData `json:"data"`
}

type SearchTopProductCatalogResp struct {
	TopProductCatalogs []TopProductCatalog `json:"top_product_catalogs"`
}
