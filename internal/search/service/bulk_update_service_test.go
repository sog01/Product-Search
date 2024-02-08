package service_test

import (
	"context"
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
	"github.com/sog01/productdiscovery/internal/search/service"
	"gopkg.in/guregu/null.v4"
)

func TestBulkUpdateResult(t *testing.T) {
	ec := createIndices()
	type args struct {
		ctx  context.Context
		req  model.BulkUpdateReq
		repo repository.BulkUpdateRepository
	}
	tests := []struct {
		name    string
		args    args
		want    model.BulkUpdateResp
		wantErr bool
	}{
		{
			name: "bulk insert",
			args: args{
				ctx: context.TODO(),
				req: model.BulkUpdateReq{
					ProductSearchUpdate: []model.ProductSearchUpdate{
						{
							Id:     uuid.FromStringOrNil("4cbc68b5-e151-4f20-b84e-4abe53e4b210"),
							Title:  null.StringFrom("VGA 13x Update1 Repository"),
							CTAURL: null.StringFrom("https://cta_url/update1"),
						},
					},
				},
				repo: repository.NewBulkUpdateRepository(ec),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.BulkUpdate(tt.args.ctx, tt.args.req, tt.args.repo)
			if err != nil {
				panic(err)
			}
			fmt.Println(got)
		})
	}
}
