package search

import (
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/pipe"
)

func NewBulkUpdateRepository(cli *elastic.Client) BulkUpdateRepository {
	return BulkUpdateRepository{
		BulkUpdate: func(args BulkUpdateReq, responses pipe.Responses) (response any, err error) {
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
				reqs = append(reqs, elastic.NewBulkUpdateRequest().
					Index("product_discovery").
					Doc(data).
					Id(id),
				)
			}

			_, err = cli.Bulk().
				Add(reqs...).
				Do(args.ctx)
			if err != nil {
				return nil, fmt.Errorf("failed bulk update product: %v", err)
			}

			return nil, nil
		},
	}
}
