package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
	uuid "github.com/satori/go.uuid"
	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
)

type BulkInsertRepository struct {
	BulkInsert pipe.FuncCtx[model.BulkInsertReq]
}

func NewBulkInsertRepository(cli *elastic.Client) BulkInsertRepository {
	return BulkInsertRepository{
		BulkInsert: func(ctx context.Context, args model.BulkInsertReq, responses pipe.Responses) (response any, err error) {
			reqs := []elastic.BulkableRequest{}
			for _, product := range args.ProductSearchInput {
				id := uuid.NewV4().String()
				data := map[string]interface{}{
					"id":         id,
					"title":      product.Title,
					"image_url":  product.ImageURL,
					"cta_url":    product.CTAURL,
					"price":      product.Price,
					"created_at": time.Now().UTC(),
					"updated_at": time.Now().UTC(),
				}
				reqs = append(reqs, elastic.NewBulkCreateRequest().
					Index("product_discovery").
					Doc(data).
					Id(id),
				)
			}

			_, err = cli.Bulk().
				Add(reqs...).
				Do(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed bulk insert product: %v", err)
			}

			return nil, nil
		},
	}
}
