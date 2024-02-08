package search_test

import (
	"context"
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/sog01/productdiscovery/internal/search"
	"gopkg.in/guregu/null.v4"
)

func TestBulkUpdateResult(t *testing.T) {
	ec := createIndices()
	type args struct {
		ctx  context.Context
		req  search.BulkUpdateReq
		repo search.BulkUpdateRepository
	}
	tests := []struct {
		name    string
		args    args
		want    search.BulkUpdateResp
		wantErr bool
	}{
		{
			name: "bulk insert",
			args: args{
				ctx: context.TODO(),
				req: search.BulkUpdateReq{
					ProductSearchUpdate: []search.ProductSearchUpdate{
						{
							Id:     uuid.FromStringOrNil("71388de0-4f38-4321-9ed7-750b9286678a"),
							Title:  null.StringFrom("VGA 13x Update1"),
							CTAURL: null.StringFrom("https://cta_url/update1"),
						},
					},
				},
				repo: search.NewBulkUpdateRepository(ec),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := search.BulkUpdate(tt.args.ctx, tt.args.req, tt.args.repo)
			if err != nil {
				panic(err)
			}
			fmt.Println(got)
		})
	}
}
