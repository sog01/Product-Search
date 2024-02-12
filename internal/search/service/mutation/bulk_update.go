package mutation

import (
	"context"
	"fmt"

	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
)

func BulkUpdate(ctx context.Context, req model.BulkUpdateReq, repo repository.BulkUpdateRepository) (model.BulkUpdateResp, error) {
	exec := pipe.PCtx(
		validateBulkUpdateRequest,
		repo.BulkUpdate,
	)

	_, err := exec(ctx, req)
	if err != nil {
		return model.BulkUpdateResp{}, err
	}

	return model.BulkUpdateResp{}, nil
}

func validateBulkUpdateRequest(ctx context.Context, args model.BulkUpdateReq, responses pipe.Responses) (response any, err error) {
	for i, productUpdate := range args.ProductSearchUpdate {
		var notEmpty bool
		if productUpdate.Title.String != "" {
			notEmpty = true
		}
		if productUpdate.CTAURL.String != "" {
			notEmpty = true
		}
		if productUpdate.ImageURL.String != "" {
			notEmpty = true
		}
		if productUpdate.Description.String != "" {
			notEmpty = true
		}
		if !notEmpty {
			return nil, fmt.Errorf("empty product[%d]", i)
		}
	}
	return nil, nil
}
