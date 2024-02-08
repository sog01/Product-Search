package query

import (
	"context"

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

	data, _ := result["data"].([]map[string]any)
	for _, d := range data {
		title, _ := d["title"].(string)
		suggestions = append(suggestions, title)
	}

	return model.SuggestionResp{
		Suggestions: suggestions,
	}, nil
}
