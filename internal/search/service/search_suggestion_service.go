package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
)

func SearchSuggestion(ctx context.Context, req model.SuggestionReq, repo repository.SearchRepository) (model.SuggestionResp, error) {
	exec := pipe.PCtx(
		repo.MatchQuery,
		composeSuggestionResponse,
	)

	resp, err := exec(ctx, model.SearchReq{
		Q:    req.Text,
		Size: 5,
	})
	if err != nil {
		return model.SuggestionResp{}, err
	}

	searchResp := pipe.Get[model.SuggestionResp](resp)
	return searchResp, nil
}

func composeSuggestionResponse(ctx context.Context, args model.SearchReq, responses pipe.Responses) (response any, err error) {
	result := pipe.Get[map[string]any](responses)
	suggestions := []string{}

	sources, _ := result["sources"].([]json.RawMessage)
	for _, source := range sources {
		var pres model.ProductSearchResponse
		if err := json.Unmarshal(source, &pres); err != nil {
			return nil, fmt.Errorf("failed unmarshal json from product search: %v", err)
		}
		suggestions = append(suggestions, pres.Title)
	}

	return model.SuggestionResp{
		Suggestions: suggestions,
	}, nil
}
