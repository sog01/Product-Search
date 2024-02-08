package mutation_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/productdiscovery/indices"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
	"github.com/sog01/productdiscovery/internal/search/service/mutation"
	"gopkg.in/guregu/null.v4"
)

func TestBulkInsertResult(t *testing.T) {
	ec := createIndices()
	type args struct {
		ctx  context.Context
		req  model.BulkInsertReq
		repo repository.BulkInsertRepository
	}
	tests := []struct {
		name    string
		args    args
		want    model.BulkInsertResp
		wantErr bool
	}{
		{
			name: "bulk insert",
			args: args{
				ctx: context.TODO(),
				req: model.BulkInsertReq{
					ProductSearchInput: []model.ProductSearchInsert{
						{
							Title:    "VGA Graphic Card",
							CTAURL:   "https://cta_url",
							ImageURL: "https://image_url",
							Category: null.StringFrom("computer"),
							Price:    100,
						},
						{
							Title:    "White Vein",
							CTAURL:   "https://cta_url",
							ImageURL: "https://image_url",
							Category: null.StringFrom("clothes"),
							Price:    900000,
						},
						{
							Title:    "Unknown",
							CTAURL:   "https://cta_url",
							ImageURL: "https://image_url",
							Price:    900000,
						},
					},
				},
				repo: repository.NewBulkInsertRepository(ec),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mutation.BulkInsert(tt.args.ctx, tt.args.req, tt.args.repo)
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
