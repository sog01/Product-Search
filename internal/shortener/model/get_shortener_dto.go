package model

type GetShortenerReq struct {
	Slug string `json:"slug"`
}

type GetShortenerResp struct {
	RealURL string `json:"real_url"`
}
