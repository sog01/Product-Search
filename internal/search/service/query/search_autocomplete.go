package query

import (
	"context"

	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
)

func SearchAutocomplete(ctx context.Context, req model.AutocompleteReq, repo repository.SearchRepository) (model.AutocompleteResp, error) {
	exec := pipe.PCtx(
		repo.MatchQuery,
		composeSuggestionResponse,
	)

	resp, err := exec(ctx, model.SearchReq{
		Q:    req.Q,
		Size: 5,
	})
	if err != nil {
		return model.AutocompleteResp{}, err
	}

	searchResp := pipe.Get[model.AutocompleteResp](resp)
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

	return model.AutocompleteResp{
		Autocompletes: suggestions,
	}, nil
}
