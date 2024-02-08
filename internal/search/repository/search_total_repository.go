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
			count := cli.Count("product_discovery")
			if args.Q != "" {
				count.Query(
					elastic.NewMatchQuery("title", args.Q),
				)
			}
			total, err := count.Do(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed count total product: %v", err)
			}

			return total, nil
		},
	}
}
