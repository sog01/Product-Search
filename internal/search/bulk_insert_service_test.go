package search_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sog01/productdiscovery/internal/search"
)

func TestBulkInsertResult(t *testing.T) {
	ec := createIndices()
	type args struct {
		ctx  context.Context
		req  search.BulkInsertReq
		repo search.BulkInsertRepository
	}
	tests := []struct {
		name    string
		args    args
		want    search.SearchResponse
		wantErr bool
	}{
		{
			name: "bulk insert",
			args: args{
				ctx: context.TODO(),
				req: search.BulkInsertReq{
					ProductSearchInput: []search.ProductSearchInsert{
						{
							Title:    "Valve Car",
							CTAURL:   "https://cta_url",
							ImageURL: "https://image_url",
							Price:    100000,
						},
					},
				},
				repo: search.NewBulkInsertRepository(ec),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := search.BulkInsert(tt.args.ctx, tt.args.req, tt.args.repo)
			if err != nil {
				panic(err)
			}
			fmt.Println(got)
		})
	}
}
