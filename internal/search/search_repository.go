package search

import (
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/pipe"
)

func NewSearchRepository(cli *elastic.Client) SearchRepository {
	return SearchRepository{
		MatchQuery: func(args SearchReq, responses pipe.Responses) (response any, err error) {
			search := cli.Search("product_discovery").
				Query(
					elastic.NewMatchQuery("title", args.Q),
				).
				Sort("id", false)

			if args.NextCursor.String != "" {
				search.SearchAfter(args.NextCursor.String)
			}
			if args.Size > 0 {
				search.Size(args.Size)
			}

			res, err := search.Do(args.ctx)
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
