package query

import (
	"context"

	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
)

func SearchCategory(ctx context.Context, args model.SearchCategoryReq, repo repository.SearchCategoryRepository) (model.SearchCategoryResp, error) {
	exec := pipe.PCtx(
		repo.GetCategory,
		composeCategoryResponse,
	)

	resp, err := exec(ctx, args)
	if err != nil {
		return model.SearchCategoryResp{}, err
	}
	categoryResp := pipe.Get[model.SearchCategoryResp](resp)
	return categoryResp, nil
}

func composeCategoryResponse(ctx context.Context, args model.SearchCategoryReq, responses pipe.Responses) (response any, err error) {
	res := pipe.Get[[]map[string]any](responses)

	categories := []model.ProductSearchCategory{}
	for _, r := range res {
		category := model.ProductSearchCategory{}
		category.Category, _ = r["category"].(string)
		category.Count, _ = r["count"].(int64)
		categories = append(categories, category)
	}
	return model.SearchCategoryResp{Categories: categories}, nil
}
