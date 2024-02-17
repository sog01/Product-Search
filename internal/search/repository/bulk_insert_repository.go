package repository

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
	uuid "github.com/satori/go.uuid"
	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/pkg"
)

type BulkInsertRepository struct {
	UploadImagesURL pipe.FuncCtx[model.BulkInsertReq]
	BulkInsert      pipe.FuncCtx[model.BulkInsertReq]
}

func NewBulkInsertRepository(cli *elastic.Client) BulkInsertRepository {
	return BulkInsertRepository{
		UploadImagesURL: func(ctx context.Context, args model.BulkInsertReq, responses pipe.Responses) (any, error) {
			response := model.BulkInsertReq{}
			for _, product := range args.ProductSearchInput {
				resp, err := http.Get(product.ImageURL)
				if err != nil {
					log.Printf("failed to get image url %v\n", err)
					continue
				}
				defer resp.Body.Close()

				imageName := time.Now().UnixNano()
				imagePath := fmt.Sprintf(pkg.WebPath()+"/images/%d", imageName)
				dst, err := os.Create(imagePath)
				if err != nil {
					return nil, fmt.Errorf("failed to create image file: %v", err)
				}
				defer dst.Close()

				_, err = io.Copy(dst, resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to copy image file: %v", err)
				}

				product.ImageURL = fmt.Sprintf("%s/images/%d", os.Getenv("SERVER.BASEURL"), imageName)
				response.ProductSearchInput = append(response.ProductSearchInput, product)
			}
			return response, nil
		},
		BulkInsert: func(ctx context.Context, args model.BulkInsertReq, responses pipe.Responses) (response any, err error) {
			searchInput := pipe.Get[model.BulkInsertReq](responses)
			if len(searchInput.ProductSearchInput) == 0 {
				searchInput = args
			}

			reqs := []elastic.BulkableRequest{}
			for _, product := range searchInput.ProductSearchInput {
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
				if product.Catalog.String != "" {
					data["catalog"] = strings.ToLower(product.Catalog.String)
				}
				if product.Description.String != "" {
					data["description"] = product.Description.String
				}
				reqs = append(reqs, elastic.NewBulkCreateRequest().
					Index("product_search").
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

			return map[string]any{}, nil
		},
	}
}
