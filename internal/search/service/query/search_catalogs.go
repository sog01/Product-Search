package query

import (
	"context"

	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
)

func SearchCatalogs(ctx context.Context, args model.SearchCatalogsReq, repo repository.SearchCategoryRepository) (model.SearchCatalogsResp, error) {
	exec := pipe.PCtx(
		repo.GetCategory,
		composeCategoryResponse,
	)

	resp, err := exec(ctx, args)
	if err != nil {
		return model.SearchCatalogsResp{}, err
	}
	categoryResp := pipe.Get[model.SearchCatalogsResp](resp)
	return categoryResp, nil
}

func composeCategoryResponse(ctx context.Context, args model.SearchCatalogsReq, responses pipe.Responses) (response any, err error) {
	res := pipe.Get[map[string]any](responses)
	data, _ := res["data"].([]map[string]any)
	catalogs := []model.ProductSearchCatalogs{}
	for _, d := range data {
		category := model.ProductSearchCatalogs{}
		category.Catalog, _ = d["catalog"].(string)
		category.Count, _ = d["count"].(int64)
		catalogs = append(catalogs, category)
	}
	return model.SearchCatalogsResp{Catalogs: catalogs}, nil
}
