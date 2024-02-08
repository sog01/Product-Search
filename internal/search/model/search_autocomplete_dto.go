package model

type AutocompleteReq struct {
	Q string `json:"text"`
}

type AutocompleteResp struct {
	Autocompletes []string `json:"autocompletes"`
}
