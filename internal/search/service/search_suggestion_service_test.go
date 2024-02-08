package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
	"github.com/sog01/productdiscovery/internal/search/service"
)

func TestSearchSuggestionResult(t *testing.T) {
	ec := createIndices()
	type args struct {
		ctx  context.Context
		req  model.SuggestionReq
		repo repository.SearchRepository
	}
	tests := []struct {
		name    string
		args    args
		want    model.SuggestionResp
		wantErr bool
	}{
		{
			name: "search suggestion",
			args: args{
				ctx: context.TODO(),
				req: model.SuggestionReq{
					Text: "v",
				},
				repo: repository.NewSearchRepository(ec),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.SearchSuggestion(tt.args.ctx, tt.args.req, tt.args.repo)
			if err != nil {
				panic(err)
			}
			fmt.Println(got)
		})
	}
}
