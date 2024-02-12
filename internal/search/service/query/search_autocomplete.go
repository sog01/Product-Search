package query

import (
	"context"
	"log"

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
		log.Printf("failed get autocompletes: %v\n", err)
		return model.AutocompleteResp{}, err
	}

	searchResp := pipe.Get[model.AutocompleteResp](resp)
	return searchResp, nil
}

func composeSuggestionResponse(ctx context.Context, args model.SearchReq, responses pipe.Responses) (response any, err error) {
	result := pipe.Get[map[string]any](responses)

	resp := model.AutocompleteResp{}
	data, _ := result["data"].([]map[string]any)
	highlights, _ := result["highlights"].([]map[string][]string)

	if len(data) == 0 || len(highlights) == 0 {
		return resp, nil
	}
	for i, h := range highlights {
		if len(h) == 0 {
			continue
		}
		title, _ := data[i]["title"].(string)
		autocomplete := model.Autocomplete{
			Title: title,
		}
		if len(h["title"]) > 0 {
			autocomplete.Highlight = h["title"][0]
		}

		resp.Autocompletes = append(resp.Autocompletes, autocomplete)
	}
	return resp, nil
}
