package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/shortener/model"
)

type GetShortenerRepository struct {
	GetShortener pipe.FuncCtx[model.GetShortenerReq]
}

func NewGetShortenerRepository(ec *elastic.Client) GetShortenerRepository {
	return GetShortenerRepository{
		GetShortener: func(ctx context.Context, args model.GetShortenerReq, responses pipe.Responses) (response any, err error) {
			resp, err := ec.Search("product_search_shortener").Query(elastic.NewTermQuery("slug", args.Slug)).Do(ctx)
			if err != nil {
				return nil, err
			}
			if len(resp.Hits.Hits) == 0 {
				return nil, nil
			}

			r := make(map[string]any)
			err = json.Unmarshal(resp.Hits.Hits[0].Source, &r)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal shortener response: %v", err)
			}
			return r, nil
		},
	}
}
