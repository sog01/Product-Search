package query_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
	"github.com/sog01/productdiscovery/internal/search/service/query"
)

func TestSearchCategory(t *testing.T) {
	ec := createIndices()
	type args struct {
		ctx  context.Context
		args model.SearchCategoryReq
		repo repository.SearchCategoryRepository
	}
	tests := []struct {
		name    string
		args    args
		want    model.SearchCategoryResp
		wantErr bool
	}{
		{
			name: "search category",
			args: args{
				ctx: context.Background(),
				args: model.SearchCategoryReq{
					Q: "vga",
				},
				repo: repository.NewSearchCategoryRepository(ec),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := query.SearchCategory(tt.args.ctx, tt.args.args, tt.args.repo)
			if err != nil {
				panic(err)
			}
			fmt.Println(got)
		})
	}
}
