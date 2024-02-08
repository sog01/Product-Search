package search_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/productdiscovery/indices"
	"github.com/sog01/productdiscovery/internal/search"
)

func TestSearchRepositoryResult(t *testing.T) {
	ec := createIndices()
	type args struct {
		ctx  context.Context
		req  search.SearchReq
		repo search.SearchRepository
	}

	tests := []struct {
		name    string
		args    args
		want    search.SearchResponse
		wantErr bool
	}{
		{
			name: "search",
			args: args{
				ctx: context.TODO(),
				req: search.SearchReq{
					Q: "vga",
				},
				repo: search.NewSearchRepository(ec),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := search.Search(tt.args.ctx, tt.args.req, tt.args.repo)
			if err != nil {
				panic(err)
			}
			fmt.Println(got)
		})
	}
}

func createIndices() *elastic.Client {
	ec, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	indices.CreateProductDiscovery(ec)
	return ec
}
