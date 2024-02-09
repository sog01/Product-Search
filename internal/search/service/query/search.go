package query

import (
	"context"
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

	data, _ := result["data"].([]map[string]any)
	sort, _ := result["sort"].([]any)
	for _, d := range data {
		var pres model.ProductSearchResponse
		pres.Id, _ = d["id"].(string)
		pres.Title, _ = d["title"].(string)
		pres.CTAURL, _ = d["cta_url"].(string)
		pres.ImageURL, _ = d["image_url"].(string)
		pres.Price, _ = d["price"].(float64)
		pres.Catalog, _ = d["catalog"].(string)
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
