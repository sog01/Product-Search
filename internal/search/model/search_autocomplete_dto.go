package model

type AutocompleteReq struct {
	Q string `json:"text"`
}

type Autocomplete struct {
	Title     string `json:"title"`
	Highlight string `json:"highlight"`
}

type AutocompleteResp struct {
	Autocompletes []Autocomplete `json:"autocompletes"`
}
