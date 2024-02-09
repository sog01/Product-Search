package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
)

type SearchRepository struct {
	MatchQuery pipe.FuncCtx[model.SearchReq]
}

func NewSearchRepository(cli *elastic.Client) SearchRepository {
	return SearchRepository{
		MatchQuery: func(ctx context.Context, args model.SearchReq, responses pipe.Responses) (response any, err error) {
			search := cli.Search("product_discovery")

			queries := []elastic.Query{}
			if args.Q != "" {
				queries = append(queries, elastic.NewMatchQuery("title", args.Q))
			}
			if args.Catalog.String != "" {
				queries = append(queries, elastic.NewTermsQuery("catalog", args.Catalog.String))
			}
			if len(queries) > 0 {
				search.Query(elastic.NewBoolQuery().Filter(queries...))
			}

			if args.NextCursor.String != "" {
				searchAfter := []any{}
				for _, c := range strings.Split(args.NextCursor.String, ",") {
					searchAfter = append(searchAfter, c)
				}
				search.SearchAfter(searchAfter...)
			}
			if args.Size > 0 {
				search.Size(args.Size)
			}

			switch args.SortBy {
			case model.Newest:
				search.Sort("created_at", false)
			case model.HighestPrice:
				search.Sort("price", false)
			case model.LowestPrice:
				search.Sort("price", true)
			}
			search.Sort("id", false)

			res, err := search.Do(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to search product: %v", err)
			}
			result := make(map[string]any)
			sources := []map[string]any{}
			for _, hit := range res.Hits.Hits {
				source := make(map[string]any)
				err := json.Unmarshal(hit.Source, &source)
				if err != nil {
					return nil, fmt.Errorf("failed unmarshal on search product: %v", err)
				}
				sources = append(sources, source)
				result["sort"] = hit.Sort
			}
			result["data"] = sources

			return result, nil
		},
	}
}
