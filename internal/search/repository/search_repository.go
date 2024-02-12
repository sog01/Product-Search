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
			search := cli.Search("product_search")

			queries := []elastic.Query{}
			if args.Q != "" {
				queries = append(queries, elastic.NewMatchQuery("title", args.Q))
				search.Highlight(elastic.NewHighlight().
					Field("title").
					PreTags("<strong>").
					PostTags("</strong>"),
				)
			}
			if args.Catalog.String != "" {
				queries = append(queries, elastic.NewTermsQuery("catalog", args.Catalog.String))
			}
			if len(queries) > 0 {
				search.Query(elastic.NewBoolQuery().Must(queries...))
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
			case model.Title:
				search.Sort("title.keyword", true)
			default:
				search.SortBy(elastic.NewFieldSort("_score").Desc()).TrackScores(true)
			}
			search.Sort("id", false)

			res, err := search.Do(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to search product: %v", err)
			}
			result := make(map[string]any)
			sources := []map[string]any{}
			highlights := []map[string][]string{}
			for _, hit := range res.Hits.Hits {
				source := make(map[string]any)
				err := json.Unmarshal(hit.Source, &source)
				if err != nil {
					return nil, fmt.Errorf("failed unmarshal on search product: %v", err)
				}
				sources = append(sources, source)
				highlights = append(highlights, hit.Highlight)
				result["sort"] = hit.Sort
			}
			result["data"] = sources
			result["highlights"] = highlights

			return result, nil
		},
	}
}
