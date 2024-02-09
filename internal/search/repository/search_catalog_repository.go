package repository

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
)

type SearchCategoryRepository struct {
	GetCategory pipe.FuncCtx[model.SearchCatalogsReq]
}

func NewSearchCatalogsRepository(cli *elastic.Client) SearchCategoryRepository {
	return SearchCategoryRepository{
		GetCategory: func(ctx context.Context, args model.SearchCatalogsReq, responses pipe.Responses) (response any, err error) {
			aggQ := elastic.NewTermsAggregation().Field("catalog")
			search := cli.Search("product_discovery")
			if args.Q != "" {
				search.Query(elastic.NewMatchQuery("title", args.Q))
			}

			resp, err := search.Aggregation("catalog", aggQ).
				Do(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed get product category: %v", err)
			}
			res, _ := resp.Aggregations.Terms("catalog")
			buckets := []map[string]any{}
			for _, bucket := range res.Buckets {
				buckets = append(buckets, map[string]any{
					"catalog": bucket.Key,
					"count":   bucket.DocCount,
				})
			}
			return map[string]any{
				"data": buckets,
			}, nil
		},
	}
}
