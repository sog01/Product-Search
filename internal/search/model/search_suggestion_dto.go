package model

type SuggestionReq struct {
	Text string `json:"text"`
}

type SuggestionResp struct {
	Suggestions []string `json:"suggestions"`
}
