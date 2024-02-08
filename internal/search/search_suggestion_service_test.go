package search_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sog01/productdiscovery/internal/search"
)

func TestSearchSuggestionResult(t *testing.T) {
	ec := createIndices()
	type args struct {
		ctx  context.Context
		req  search.SuggestionReq
		repo search.SearchRepository
	}
	tests := []struct {
		name    string
		args    args
		want    search.SearchResponse
		wantErr bool
	}{
		{
			name: "search suggestion",
			args: args{
				ctx: context.TODO(),
				req: search.SuggestionReq{
					Text: "v",
				},
				repo: search.NewSearchRepository(ec),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := search.SearchSuggestion(tt.args.ctx, tt.args.req, tt.args.repo)
			if err != nil {
				panic(err)
			}
			fmt.Println(got)
		})
	}
}
