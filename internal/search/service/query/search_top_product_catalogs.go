package query

import (
	"context"
	"log"

	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
	"gopkg.in/guregu/null.v4"
)

func SearchTopProductCatalog(ctx context.Context, args model.SearchTopProductCatalogReq, repo repository.SearchRepository) (model.SearchTopProductCatalogResp, error) {
	exec := pipe.PCtx(
		resolveTopProductCatalogs(ctx, repo),
		composeSearchTopProductCatalogs,
	)

	resp, err := exec(ctx, args)
	if err != nil {
		log.Printf("failed search top product catalog: %v\n", err)
		return model.SearchTopProductCatalogResp{}, err
	}

	searchTopProductCatalogResp := pipe.Get[model.SearchTopProductCatalogResp](resp)
	return searchTopProductCatalogResp, nil
}

func resolveTopProductCatalogs(ctx context.Context, repo repository.SearchRepository) pipe.FuncCtx[model.SearchTopProductCatalogReq] {
	return func(ctx context.Context, args model.SearchTopProductCatalogReq, responses pipe.Responses) (response any, err error) {
		exec := pipe.PCtx(
			repo.MatchQuery,
		)
		resp := make(map[string]any)
		for _, catalog := range args.Catalogs {
			response, err := exec(ctx, model.SearchReq{
				Catalog: null.StringFrom(catalog),
				Size:    3,
				SortBy:  model.Newest,
			})
			if err != nil {
				return nil, err
			}
			data := pipe.Get[map[string]any](response)
			resp[catalog] = data["data"]
		}

		return resp, nil
	}
}

func composeSearchTopProductCatalogs(ctx context.Context, args model.SearchTopProductCatalogReq, responses pipe.Responses) (response any, err error) {
	resp := pipe.Get[map[string]any](responses)
	topProductCatalogsData := make(map[string][]model.TopProductCatalogData)
	for catalog, r := range resp {
		data, _ := r.([]map[string]any)
		for _, d := range data {
			imageURL, _ := d["image_url"].(string)
			topProductCatalogsData[catalog] = append(topProductCatalogsData[catalog], model.TopProductCatalogData{
				ImageURL: imageURL,
			})
		}
	}

	topProductCatalogs := []model.TopProductCatalog{}
	for catalog, data := range topProductCatalogsData {
		topProductCatalogs = append(topProductCatalogs, model.TopProductCatalog{
			Catalog: catalog,
			Data:    data,
		})
	}

	return model.SearchTopProductCatalogResp{TopProductCatalogs: topProductCatalogs}, nil
}
