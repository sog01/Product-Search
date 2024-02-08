package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
)

type BulkUpdateRepository struct {
	BulkUpdate pipe.FuncCtx[model.BulkUpdateReq]
}

func NewBulkUpdateRepository(cli *elastic.Client) BulkUpdateRepository {
	return BulkUpdateRepository{
		BulkUpdate: func(ctx context.Context, args model.BulkUpdateReq, responses pipe.Responses) (response any, err error) {
			reqs := []elastic.BulkableRequest{}
			for _, product := range args.ProductSearchUpdate {
				id := product.Id.String()
				data := map[string]interface{}{}
				if product.Title.String != "" {
					data["title"] = product.Title.String
				}
				if product.CTAURL.String != "" {
					data["cta_url"] = product.CTAURL.String
				}
				if product.ImageURL.String != "" {
					data["image_url"] = product.ImageURL.String
				}
				if product.Category.String != "" {
					data["category"] = strings.ToLower(product.Category.String)
				}
				data["updated_at"] = time.Now().UTC()
				reqs = append(reqs, elastic.NewBulkUpdateRequest().
					Index("product_discovery").
					Doc(data).
					Id(id),
				)
			}

			_, err = cli.Bulk().
				Add(reqs...).
				Do(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed bulk update product: %v", err)
			}

			return map[string]any{}, nil
		},
	}
}
