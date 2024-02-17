package mutation

import (
	"context"
	"fmt"
	"log"

	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
)

func BulkInsert(ctx context.Context, req model.BulkInsertReq, repo repository.BulkInsertRepository) (model.BulkInsertResp, error) {
	exec := pipe.PCtx(
		validateBulkInsertRequest,
		repo.UploadImagesURL,
		repo.BulkInsert,
	)

	_, err := exec(ctx, req)
	if err != nil {
		log.Printf("failed to bulk insert product : %v\n", err)
		return model.BulkInsertResp{}, err
	}

	return model.BulkInsertResp{}, nil
}

func validateBulkInsertRequest(ctx context.Context, req model.BulkInsertReq, responses pipe.Responses) (response any, err error) {
	for i, r := range req.ProductSearchInput {
		if r.Title == "" {
			return nil, fmt.Errorf("empty title, product[%d]", i)
		}
		if r.CTAURL == "" {
			return nil, fmt.Errorf("empty cta_url, product[%d]", i)
		}
		if r.ImageURL == "" {
			return nil, fmt.Errorf("empty image_url, product[%d]", i)
		}
	}
	return nil, nil
}
