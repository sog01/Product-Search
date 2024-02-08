package search

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sog01/pipe"
)

func SearchSuggestion(ctx context.Context, req SuggestionReq, repo SearchRepository) (SuggestionResp, error) {
	exec := pipe.P(
		repo.MatchQuery,
		composeSuggestionResponse,
	)

	resp, err := exec(SearchReq{
		Q:    req.Text,
		Size: 5,
		ctx:  ctx,
	})
	if err != nil {
		return SuggestionResp{}, err
	}

	searchResp := pipe.Get[SuggestionResp](resp)
	return searchResp, nil
}

func composeSuggestionResponse(args SearchReq, responses pipe.Responses) (response any, err error) {
	result := pipe.Get[map[string]any](responses)
	suggestions := []string{}

	sources, _ := result["sources"].([]json.RawMessage)
	for _, source := range sources {
		var pres ProductSearchResponse
		if err := json.Unmarshal(source, &pres); err != nil {
			return nil, fmt.Errorf("failed unmarshal json from product search: %v", err)
		}
		suggestions = append(suggestions, pres.Title)
	}

	return SuggestionResp{
		Suggestions: suggestions,
	}, nil
}
