package repository

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
)

type SearchCategoryRepository struct {
	GetCategory pipe.FuncCtx[model.SearchCategoryReq]
}

func NewSearchCategoryRepository(cli *elastic.Client) SearchCategoryRepository {
	return SearchCategoryRepository{
		GetCategory: func(ctx context.Context, args model.SearchCategoryReq, responses pipe.Responses) (response any, err error) {
			aggQ := elastic.NewTermsAggregation().Field("category")
			search := cli.Search("product_discovery")
			if args.Q != "" {
				search.Query(elastic.NewMatchQuery("title", args.Q))
			}

			resp, err := search.Aggregation("category", aggQ).
				Do(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed get product category: %v", err)
			}
			res, _ := resp.Aggregations.Terms("category")
			buckets := []map[string]any{}
			for _, bucket := range res.Buckets {
				buckets = append(buckets, map[string]any{
					"category": bucket.Key,
					"count":    bucket.DocCount,
				})
			}
			return buckets, nil
		},
	}
}
