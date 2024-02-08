package search

import "context"

type SuggestionReq struct {
	Text string `json:"text"`
	ctx  context.Context
}

type SuggestionResp struct {
	Suggestions []string `json:"suggestions"`
}
