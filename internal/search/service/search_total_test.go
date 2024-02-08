package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
	"github.com/sog01/productdiscovery/internal/search/service"
)

func TestSearchTotalResult(t *testing.T) {
	ec := createIndices()
	type args struct {
		ctx  context.Context
		args model.SearchTotalReq
		repo repository.SearchTotalRepository
	}
	tests := []struct {
		name    string
		args    args
		want    model.SearchTotalResp
		wantErr bool
	}{
		{
			name: "search total",
			args: args{
				ctx: context.TODO(),
				args: model.SearchTotalReq{
					Q: "v",
				},
				repo: repository.NewSearchTotalRepository(ec),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.SearchTotal(tt.args.ctx, tt.args.args, tt.args.repo)
			if err != nil {
				panic(err)
			}
			fmt.Println(got)
		})
	}
}
