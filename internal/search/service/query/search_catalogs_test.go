package query_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
	"github.com/sog01/productdiscovery/internal/search/service/query"
)

func TestSearchCatalogsResult(t *testing.T) {
	ec := createIndices()
	type args struct {
		ctx  context.Context
		args model.SearchCatalogsReq
		repo repository.SearchCategoryRepository
	}
	tests := []struct {
		name    string
		args    args
		want    model.SearchCatalogsResp
		wantErr bool
	}{
		{
			name: "search catalogs",
			args: args{
				ctx: context.Background(),
				args: model.SearchCatalogsReq{
					Q: "vga",
				},
				repo: repository.NewSearchCatalogsRepository(ec),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := query.SearchCatalogs(tt.args.ctx, tt.args.args, tt.args.repo)
			if err != nil {
				panic(err)
			}
			fmt.Println(got)
		})
	}
}
