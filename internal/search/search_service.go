package search

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sog01/pipe"
)

type SearchRepository struct {
	MatchQuery pipe.Func[SearchReq]
}

func Search(ctx context.Context, req SearchReq, repo SearchRepository) (SearchResponse, error) {
	exec := pipe.P(
		repo.MatchQuery,
		composeResponse,
	)

	req.ctx = ctx
	resp, err := exec(req)
	if err != nil {
		return SearchResponse{}, err
	}

	searchResp := pipe.Get[SearchResponse](resp)
	return searchResp, nil
}

func composeResponse(args SearchReq, responses pipe.Responses) (response any, err error) {
	result := pipe.Get[map[string]any](responses)
	productListResp := []ProductSearchResponse{}

	sources, _ := result["sources"].([]json.RawMessage)
	sort, _ := result["sort"].([]any)
	for _, source := range sources {
		var pres ProductSearchResponse
		if err := json.Unmarshal(source, &pres); err != nil {
			return nil, fmt.Errorf("failed unmarshal json from product search: %v", err)
		}
		productListResp = append(productListResp, pres)
	}

	sortStrings := []string{}
	for _, s := range sort {
		sortStrings = append(sortStrings, fmt.Sprintf("%v", s))
	}

	return SearchResponse{
		Products:   productListResp,
		NextCursor: strings.Join(sortStrings, ","),
	}, nil
}
