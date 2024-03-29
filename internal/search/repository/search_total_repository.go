package repository

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
)

type SearchTotalRepository struct {
	CountTotal pipe.FuncCtx[model.SearchTotalReq]
}

func NewSearchTotalRepository(cli *elastic.Client) SearchTotalRepository {
	return SearchTotalRepository{
		CountTotal: func(ctx context.Context, args model.SearchTotalReq, responses pipe.Responses) (response any, err error) {
			count := cli.Count("product_search")
			queries := []elastic.Query{}
			if args.Q != "" {
				queries = append(queries, elastic.NewMatchQuery("title", args.Q))
			}
			if args.Catalog.String != "" {
				queries = append(queries, elastic.NewTermQuery("catalog", args.Catalog.String))
			}

			if len(queries) > 0 {
				count.Query(elastic.NewBoolQuery().Filter(queries...))
			}

			total, err := count.Do(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed count total product: %v", err)
			}

			return map[string]any{
				"total": total,
			}, nil
		},
	}
}
