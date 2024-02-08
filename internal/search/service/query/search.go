package query

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
)

func Search(ctx context.Context, req model.SearchReq, repo repository.SearchRepository) (model.SearchResponse, error) {
	exec := pipe.PCtx(
		repo.MatchQuery,
		composeResponse,
	)

	resp, err := exec(ctx, req)
	if err != nil {
		return model.SearchResponse{}, err
	}

	searchResp := pipe.Get[model.SearchResponse](resp)
	return searchResp, nil
}

func composeResponse(ctx context.Context, args model.SearchReq, responses pipe.Responses) (response any, err error) {
	result := pipe.Get[map[string]any](responses)
	productListResp := []model.ProductSearchResponse{}

	sources, _ := result["sources"].([]json.RawMessage)
	sort, _ := result["sort"].([]any)
	for _, source := range sources {
		var pres model.ProductSearchResponse
		if err := json.Unmarshal(source, &pres); err != nil {
			return nil, fmt.Errorf("failed unmarshal json from product search: %v", err)
		}
		productListResp = append(productListResp, pres)
	}

	sortStrings := []string{}
	for _, s := range sort {
		switch v := s.(type) {
		case float64:
			sortStrings = append(sortStrings, strconv.FormatFloat(v, 'f', -1, 64))
		default:
			sortStrings = append(sortStrings, fmt.Sprintf("%v", s))
		}
	}

	return model.SearchResponse{
		Products:   productListResp,
		NextCursor: strings.Join(sortStrings, ","),
	}, nil
}
