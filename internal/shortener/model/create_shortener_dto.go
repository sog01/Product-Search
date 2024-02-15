package model

type CreateShortenerReq struct {
	RealURL string `json:"real_url"`
}

type CreateShortenerResp struct {
	Slug string `json:"slug"`
}
