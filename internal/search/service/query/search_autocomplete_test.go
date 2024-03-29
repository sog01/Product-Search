package query_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
	"github.com/sog01/productdiscovery/internal/search/service/query"
)

func TestSearchSuggestionResult(t *testing.T) {
	ec := createIndices()
	type args struct {
		ctx  context.Context
		req  model.AutocompleteReq
		repo repository.SearchRepository
	}
	tests := []struct {
		name    string
		args    args
		want    model.AutocompleteResp
		wantErr bool
	}{
		{
			name: "search suggestion",
			args: args{
				ctx: context.TODO(),
				req: model.AutocompleteReq{
					Q: "v",
				},
				repo: repository.NewSearchRepository(ec),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := query.SearchAutocomplete(tt.args.ctx, tt.args.req, tt.args.repo)
			if err != nil {
				panic(err)
			}
			fmt.Println(got)
		})
	}
}
