package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/productdiscovery/indices"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
	"github.com/sog01/productdiscovery/internal/search/service"
)

func TestSearchResult(t *testing.T) {
	ec := createIndices()
	type args struct {
		ctx  context.Context
		req  model.SearchReq
		repo repository.SearchRepository
	}

	tests := []struct {
		name    string
		args    args
		want    model.SearchResponse
		wantErr bool
	}{
		{
			name: "search",
			args: args{
				ctx: context.TODO(),
				req: model.SearchReq{
					Size:   10,
					SortBy: model.LowestPrice,
				},
				repo: repository.NewSearchRepository(ec),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Search(tt.args.ctx, tt.args.req, tt.args.repo)
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
