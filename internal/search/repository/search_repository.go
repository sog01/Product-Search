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
			search := cli.Search("product_discovery").
				Query(
					elastic.NewMatchQuery("title", args.Q),
				)

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
			sources := []json.RawMessage{}
			for _, hit := range res.Hits.Hits {
				sources = append(sources, hit.Source)
				result["sort"] = hit.Sort
			}
			result["sources"] = sources

			return result, nil
		},
	}
}
