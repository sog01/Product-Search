package query

import (
	"context"

	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
)

func SearchTotal(ctx context.Context, args model.SearchTotalReq, repo repository.SearchTotalRepository) (model.SearchTotalResp, error) {
	exec := pipe.PCtx(
		repo.CountTotal,
		composeSearchTotalResponse,
	)

	resp, err := exec(ctx, args)
	if err != nil {
		return model.SearchTotalResp{}, err
	}
	searchTotalResp := pipe.Get[model.SearchTotalResp](resp)
	return searchTotalResp, nil
}

func composeSearchTotalResponse(ctx context.Context, args model.SearchTotalReq, responses pipe.Responses) (response any, err error) {
	res := pipe.Get[map[string]any](responses)
	total, _ := res["total"].(int64)
	return model.SearchTotalResp{
		Total: total,
	}, nil
}
